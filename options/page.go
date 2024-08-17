package options

import (
	"encoding/base64"
	"errors"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/flags"
)

// Page represents page-related options.
type Page struct {
	CountTotal bool   `json:"count_total" toml:"count_total"` // CountTotal indicates whether to include total count in paged queries.
	Key        string `json:"key" toml:"key"`                 // Key is the base64-encoded key for page.
	Limit      uint64 `json:"limit" toml:"limit"`             // Limit is the maximum number of results per page.
	Offset     uint64 `json:"offset" toml:"offset"`           // Offset is the offset for page.
	Reverse    bool   `json:"reverse" toml:"reverse"`         // Reverse indicates whether to reverse the order of results in page.
}

// NewPage creates a new Page instance with default values.
func NewPage() *Page {
	return &Page{
		CountTotal: flags.DefaultPageCountTotal,
		Key:        flags.DefaultPageKey,
		Limit:      flags.DefaultPageLimit,
		Offset:     flags.DefaultPageOffset,
		Reverse:    flags.DefaultPageReverse,
	}
}

// WithCountTotal sets the CountTotal field and returns the updated Page instance.
func (p *Page) WithCountTotal(v bool) *Page {
	p.CountTotal = v
	return p
}

// WithKey sets the Key field with a base64-encoded value and returns the updated Page instance.
func (p *Page) WithKey(v []byte) *Page {
	p.Key = base64.StdEncoding.EncodeToString(v)
	return p
}

// WithLimit sets the Limit field and returns the updated Page instance.
func (p *Page) WithLimit(v uint64) *Page {
	p.Limit = v
	return p
}

// WithOffset sets the Offset field and returns the updated Page instance.
func (p *Page) WithOffset(v uint64) *Page {
	p.Offset = v
	return p
}

// WithReverse sets the Reverse field and returns the updated Page instance.
func (p *Page) WithReverse(v bool) *Page {
	p.Reverse = v
	return p
}

// GetCountTotal returns the CountTotal field.
func (p *Page) GetCountTotal() bool {
	return p.CountTotal
}

// GetKey returns the decoded Key field.
func (p *Page) GetKey() []byte {
	key, err := base64.StdEncoding.DecodeString(p.Key)
	if err != nil {
		panic(err)
	}

	return key
}

// GetLimit returns the Limit field.
func (p *Page) GetLimit() uint64 {
	return p.Limit
}

// GetOffset returns the Offset field.
func (p *Page) GetOffset() uint64 {
	return p.Offset
}

// GetReverse returns the Reverse field.
func (p *Page) GetReverse() bool {
	return p.Reverse
}

// ValidatePageKey checks if the provided key is a valid base64-encoded string.
func ValidatePageKey(key string) error {
	if key == "" {
		return nil
	}
	if _, err := base64.StdEncoding.DecodeString(key); err != nil {
		return errors.New("key must be a valid base64-encoded string")
	}

	return nil
}

// ValidatePageLimit validates the Limit field.
func ValidatePageLimit(limit uint64) error {
	if limit == 0 {
		return errors.New("limit must be greater than zero")
	}

	return nil
}

// ValidatePageOffset validates the Offset field.
func ValidatePageOffset(offset uint64) error {
	if offset < 0 {
		return errors.New("offset must be non-negative")
	}

	return nil
}

// Validate validates all the fields of the Page struct.
func (p *Page) Validate() error {
	if err := ValidatePageKey(p.Key); err != nil {
		return err
	}
	if err := ValidatePageLimit(p.Limit); err != nil {
		return err
	}
	if err := ValidatePageOffset(p.Offset); err != nil {
		return err
	}

	return nil
}

// PageRequest creates a new PageRequest with the configured options.
func (p *Page) PageRequest() *query.PageRequest {
	return &query.PageRequest{
		Key:        p.GetKey(),
		Offset:     p.GetOffset(),
		Limit:      p.GetLimit(),
		CountTotal: p.GetCountTotal(),
		Reverse:    p.GetReverse(),
	}
}

// NewPageFromCmd creates and returns a Page from the given cobra command's flags.
func NewPageFromCmd(cmd *cobra.Command) (*Page, error) {
	// Retrieve the value of the "page.count-total" flag.
	countTotal, err := flags.GetPageCountTotal(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.key" flag.
	key, err := flags.GetPageKey(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.limit" flag.
	limit, err := flags.GetPageLimit(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.offset" flag.
	offset, err := flags.GetPageOffset(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.reverse" flag.
	reverse, err := flags.GetPageReverse(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Page instance populated with the retrieved flag values.
	return &Page{
		CountTotal: countTotal,
		Key:        key,
		Limit:      limit,
		Offset:     offset,
		Reverse:    reverse,
	}, nil
}
