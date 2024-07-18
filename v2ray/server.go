package v2ray

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v4/process"
	proxymancommand "github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	statscommand "github.com/v2fly/v2ray-core/v5/app/stats/command"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	"github.com/v2fly/v2ray-core/v5/common/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	sentinelsdk "github.com/sentinel-official/sentinel-go-sdk/v1/types"
	"github.com/sentinel-official/sentinel-go-sdk/v1/utils"
	"github.com/sentinel-official/sentinel-go-sdk/v1/v2ray/types"
)

const (
	// RequestLen is the expected length of request used for peer operations (1 byte for type + 16 bytes for UUID).
	RequestLen = 1 + 16

	// InfoLen represents the length of the server information (2 bytes for version + 1 byte for type).
	InfoLen = 2 + 1
)

// Ensure Server implements sentinelsdk.ServerService interface.
var _ sentinelsdk.ServerService = (*Server)(nil)

// Server represents the V2Ray server instance.
type Server struct {
	homeDir string             // Home directory of the V2Ray server.
	name    string             // Name of the server instance.
	info    []byte             // Information about the server instance.
	cmd     *exec.Cmd          // Command to run the V2Ray server.
	pm      *types.PeerManager // Peer manager for handling peer information.
}

// NewServer creates a new instance of the V2Ray server with the specified home directory.
func NewServer(homeDir string) *Server {
	return &Server{
		homeDir: homeDir,
		info:    make([]byte, InfoLen),
		cmd:     nil,
		pm:      types.NewPeerManager(),
	}
}

// configFilePath returns the full path of the V2Ray server's configuration file.
func (s *Server) configFilePath() string {
	return filepath.Join(s.homeDir, fmt.Sprintf("%s.json", s.name))
}

// pidFilePath returns the file path of the server's PID file.
func (s *Server) pidFilePath() string {
	return filepath.Join(s.homeDir, fmt.Sprintf("%s.pid", s.name))
}

// readPIDFromFile reads the PID from the server's PID file.
func (s *Server) readPIDFromFile() (int32, error) {
	// Read PID from the PID file.
	data, err := os.ReadFile(s.pidFilePath())
	if err != nil {
		return 0, err
	}

	// Convert PID data to integer.
	pid, err := strconv.ParseInt(string(data), 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(pid), nil
}

// writePIDToFile writes the given PID to the server's PID file.
func (s *Server) writePIDToFile(pid int) error {
	// Convert PID to byte slice.
	data := []byte(strconv.Itoa(pid))

	// Write PID to file with appropriate permissions.
	if err := os.WriteFile(s.pidFilePath(), data, 0644); err != nil {
		return err
	}

	return nil
}

// clientConn establishes a gRPC client connection to the V2Ray server.
func (s *Server) clientConn() (*grpc.ClientConn, error) {
	// Define the target address for the gRPC client connection.
	target := "127.0.0.1:23"

	// Establish a gRPC client connection with specified options:
	// - WithBlock: Blocks until the underlying connection is established.
	// - WithTransportCredentials: Configures insecure transport credentials for the connection.
	return grpc.Dial(
		target,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

// handlerServiceClient establishes a gRPC client connection to the V2Ray server's handler service.
func (s *Server) handlerServiceClient() (*grpc.ClientConn, proxymancommand.HandlerServiceClient, error) {
	// Establish a gRPC client connection using the clientConn method.
	conn, err := s.clientConn()
	if err != nil {
		return nil, nil, err
	}

	// Create a new HandlerServiceClient using the established connection.
	client := proxymancommand.NewHandlerServiceClient(conn)

	// Return both the connection and the client.
	return conn, client, nil
}

// statsServiceClient establishes a gRPC client connection to the V2Ray server's stats service.
func (s *Server) statsServiceClient() (*grpc.ClientConn, statscommand.StatsServiceClient, error) {
	// Establish a gRPC client connection using the clientConn method.
	conn, err := s.clientConn()
	if err != nil {
		return nil, nil, err
	}

	// Create a new StatsServiceClient using the established connection.
	client := statscommand.NewStatsServiceClient(conn)

	// Return both the connection and the client.
	return conn, client, nil
}

// Info returns the server's information.
func (s *Server) Info() []byte {
	return s.info
}

// Type returns the service type of the server.
func (s *Server) Type() sentinelsdk.ServiceType {
	return sentinelsdk.ServiceTypeV2Ray
}

// IsUp checks if the V2Ray server process is running.
func (s *Server) IsUp(ctx context.Context) (bool, error) {
	// Read PID from file.
	pid, err := s.readPIDFromFile()
	if err != nil {
		return false, err
	}

	// Retrieve process with the given PID.
	proc, err := process.NewProcessWithContext(ctx, pid)
	if err != nil {
		return false, err
	}

	// Check if the process is running.
	ok, err := proc.IsRunningWithContext(ctx)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	// Retrieve the name of the process.
	name, err := proc.NameWithContext(ctx)
	if err != nil {
		return false, err
	}

	// Check if the process name matches constant v2ray.
	if name != v2ray {
		return false, nil
	}

	return true, nil
}

// PreUp writes the configuration to the config file before starting the server process.
func (s *Server) PreUp(v interface{}) error {
	// Check for valid parameter type.
	cfg, ok := v.(*types.ServerConfig)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	// Write configuration to file.
	return cfg.WriteFile(s.configFilePath())
}

// PostUp performs operations after the server process is started.
func (s *Server) PostUp() error {
	// Check if command or process is nil.
	if s.cmd == nil || s.cmd.Process == nil {
		return fmt.Errorf("nil command or process")
	}

	// Write PID to file.
	if err := s.writePIDToFile(s.cmd.Process.Pid); err != nil {
		return err
	}

	return nil
}

// PreDown performs operations before the server process is terminated.
func (s *Server) PreDown() error {
	return nil
}

// Down terminates the V2Ray server process.
func (s *Server) Down(ctx context.Context) error {
	// Read PID from file.
	pid, err := s.readPIDFromFile()
	if err != nil {
		return err
	}

	// Retrieve process with the given PID.
	proc, err := process.NewProcessWithContext(ctx, pid)
	if err != nil {
		return err
	}

	// Terminate the process.
	if err := proc.TerminateWithContext(ctx); err != nil {
		return err
	}

	// Reset the command.
	s.cmd = nil
	return nil
}

// PostDown performs cleanup operations after the server process is terminated.
func (s *Server) PostDown() error {
	// Remove PID file.
	if err := utils.RemoveFile(s.pidFilePath()); err != nil {
		return err
	}

	return nil
}

// AddPeer adds a new peer to the V2Ray server.
func (s *Server) AddPeer(ctx context.Context, req []byte) ([]byte, error) {
	// Check if the request length is valid.
	if len(req) != RequestLen {
		return nil, fmt.Errorf("invalid request length; expected %d, got %d", RequestLen, len(req))
	}

	// Establish a gRPC client connection to the handler service.
	conn, client, err := s.handlerServiceClient()
	if err != nil {
		return nil, err
	}

	// Ensure the connection is closed when done.
	defer func() {
		if err = conn.Close(); err != nil {
			panic(err)
		}
	}()

	// Encode the request buffer to email using base64 encoding and extract proxy type.
	email := base64.StdEncoding.EncodeToString(req)
	proxy := types.Proxy(req[0])

	// Parse the UUID from the request buffer.
	uid, err := uuid.ParseBytes(req[1:])
	if err != nil {
		return nil, err
	}

	// Prepare gRPC request to add a user to the handler.
	in := &proxymancommand.AlterInboundRequest{
		Tag: proxy.Tag(),
		Operation: serial.ToTypedMessage(
			&proxymancommand.AddUserOperation{
				User: &protocol.User{
					Level:   0,
					Email:   email,
					Account: proxy.Account(uid),
				},
			},
		),
	}

	// Send the request to add a user to the handler.
	_, err = client.AlterInbound(ctx, in)
	if err != nil {
		return nil, err
	}

	// Update the local peer collection with the new peer information.
	s.pm.Put(
		&types.Peer{
			Email: email,
		},
	)

	// Return nil for success (no additional data to return in response).
	return nil, nil
}

// HasPeer checks if a peer exists in the V2Ray server's peer list.
func (s *Server) HasPeer(_ context.Context, req []byte) (bool, error) {
	// Check if the request length is valid.
	if len(req) != RequestLen {
		return false, fmt.Errorf("invalid request length; expected %d, got %d", RequestLen, len(req))
	}

	// Encode the request buffer to email using base64 encoding.
	email := base64.StdEncoding.EncodeToString(req)
	peer := s.pm.Get(email)

	// Return true if the peer exists, otherwise false.
	return peer != nil, nil
}

// RemovePeer removes a peer from the V2Ray server.
func (s *Server) RemovePeer(ctx context.Context, req []byte) error {
	// Check if the request length is valid.
	if len(req) != RequestLen {
		return fmt.Errorf("invalid request length; expected %d, got %d", RequestLen, len(req))
	}

	// Establish a gRPC client connection to the handler service.
	conn, client, err := s.handlerServiceClient()
	if err != nil {
		return err
	}

	// Ensure the connection is closed when done.
	defer func() {
		if err = conn.Close(); err != nil {
			panic(err)
		}
	}()

	// Encode the data buffer to email using base64 encoding and extract proxy type.
	email := base64.StdEncoding.EncodeToString(req)
	proxy := types.Proxy(req[0])

	// Prepare gRPC request to remove a user from the handler.
	in := &proxymancommand.AlterInboundRequest{
		Tag: proxy.Tag(),
		Operation: serial.ToTypedMessage(
			&proxymancommand.RemoveUserOperation{
				Email: email,
			},
		),
	}

	// Send the request to remove a user from the handler.
	_, err = client.AlterInbound(ctx, in)
	if err != nil {
		// If the user is not found, continue without error.
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}

	// Remove the peer information from the local collection.
	s.pm.Delete(email)

	// Return nil for success.
	return nil
}

// PeerCount returns the number of peers connected to the V2Ray server.
func (s *Server) PeerCount() int {
	return s.pm.Len()
}

// PeerStatistics retrieves statistics for each peer connected to the V2Ray server.
func (s *Server) PeerStatistics(ctx context.Context) (items []*sentinelsdk.PeerStatistic, err error) {
	// Establish a gRPC client connection to the stats service.
	conn, client, err := s.statsServiceClient()
	if err != nil {
		return nil, err
	}

	// Ensure the connection is closed when done.
	defer func() {
		if err = conn.Close(); err != nil {
			panic(err)
		}
	}()

	// Define a function to process each peer in the local collection.
	fn := func(key string, _ *types.Peer) (bool, error) {
		// Prepare gRPC request to get uplink traffic stats.
		in := &statscommand.GetStatsRequest{
			Reset_: false,
			Name:   fmt.Sprintf("user>>>%s>>>traffic>>>uplink", key),
		}

		// Send the request to get uplink traffic stats.
		res, err := client.GetStats(ctx, in)
		if err != nil {
			// If the stat is not found, continue to the next peer.
			if !strings.Contains(err.Error(), "not found") {
				return false, err
			}
		}

		// Extract uplink traffic stats or use an empty stat if not found.
		upLink := res.GetStat()
		if upLink == nil {
			upLink = &statscommand.Stat{}
		}

		// Prepare gRPC request to get downlink traffic stats.
		in = &statscommand.GetStatsRequest{
			Reset_: false,
			Name:   fmt.Sprintf("user>>>%s>>>traffic>>>downlink", key),
		}

		// Send the request to get downlink traffic stats.
		res, err = client.GetStats(ctx, in)
		if err != nil {
			// If the stat is not found, continue to the next peer.
			if !strings.Contains(err.Error(), "not found") {
				return false, err
			}
		}

		// Extract downlink traffic stats or use an empty stat if not found.
		downLink := res.GetStat()
		if downLink == nil {
			downLink = &statscommand.Stat{}
		}

		// Append peer statistics to the result collection.
		items = append(
			items,
			&sentinelsdk.PeerStatistic{
				Key:      key,
				Download: downLink.GetValue(),
				Upload:   upLink.GetValue(),
			},
		)

		return false, nil
	}

	// Iterate over each peer and retrieve statistics.
	if err := s.pm.Iterate(fn); err != nil {
		return nil, err
	}

	// Return the constructed collection of peer statistics.
	return items, nil
}
