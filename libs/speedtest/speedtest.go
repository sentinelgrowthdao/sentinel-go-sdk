package speedtest

import (
	"errors"

	sdkmath "cosmossdk.io/math"
	"github.com/showwin/speedtest-go/speedtest"
)

// performTests runs the ping, download, and upload tests on the target server.
func performTests(target *speedtest.Server) error {
	// Perform the ping test
	if err := target.PingTest(nil); err != nil {
		return err
	}

	// Perform the download test
	if err := target.DownloadTest(); err != nil {
		return err
	}

	// Perform the upload test
	if err := target.UploadTest(); err != nil {
		return err
	}

	// Wait for the context to be ready after the tests
	target.Context.Wait()
	return nil
}

// RunTest performs a speed test and returns download and upload speeds.
func RunTest() (dlSpeed, upSpeed sdkmath.Int, err error) {
	// Create a new Speedtest client
	st := speedtest.New()

	// Fetch the list of servers from the Speedtest service
	servers, err := st.FetchServers()
	if err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Find the best server from the list
	targets, err := servers.FindServer(nil)
	if err != nil {
		return sdkmath.ZeroInt(), sdkmath.ZeroInt(), err
	}

	// Iterate through the list of target servers to find a valid result
	for _, target := range targets {
		// Perform the tests on the target server
		if err := performTests(target); err != nil {
			target.Context.Reset()
			continue
		}

		// Convert download and upload speeds to sdkmath.Int
		dlSpeed, _ = sdkmath.NewIntFromString(target.DLSpeed.String())
		upSpeed, _ = sdkmath.NewIntFromString(target.ULSpeed.String())

		// Check if the speeds are positive
		if !dlSpeed.IsPositive() || !upSpeed.IsPositive() {
			target.Context.Reset()
			continue
		}

		// A valid result was found, exit the loop
		return dlSpeed, upSpeed, nil
	}

	// Return an error if no valid result was found
	return sdkmath.ZeroInt(), sdkmath.ZeroInt(), errors.New("no server provided valid results")
}
