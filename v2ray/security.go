package v2ray

// Security is a custom type used to represent different security settings.
type Security byte

// Constants for Security type with automatic incrementation for each security setting.
const (
	SecurityUnspecified Security = iota // Default value for unspecified security
	SecurityNone                        // SecurityNone represents no security
	SecurityTLS                         // SecurityTLS represents TLS security
)

// String returns a string representation of the Security type.
func (s Security) String() string {
	switch s {
	case SecurityNone:
		return "none"
	case SecurityTLS:
		return "tls"
	default:
		return "" // Return empty string for unspecified or unknown security settings
	}
}

// IsValid checks if the Security value is valid.
func (s Security) IsValid() bool {
	return s.String() != ""
}

// NewSecurityFromString converts a string to a Security type.
func NewSecurityFromString(v string) Security {
	switch v {
	case "none":
		return SecurityNone
	case "tls":
		return SecurityTLS
	default:
		return SecurityUnspecified // Returns the default security if no match is found
	}
}
