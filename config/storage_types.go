// Copyright © 2026 Ping Identity Corporation

package config

// StorageType represents the type of storage backend for tokens.
type StorageType string

const (
	// StorageTypeSecureLocal uses a secure local storage (e.g. the OS keychain).
	StorageTypeSecureLocal StorageType = "secure_local"

	// StorageTypeNone disables token storage.
	StorageTypeNone StorageType = "none"
)

// String returns the string representation of the StorageType.
func (s StorageType) String() string {
	return string(s)
}

// IsValid reports whether the StorageType is a value supported by this package.
func (s StorageType) IsValid() bool {
	switch s {
	case StorageTypeSecureLocal, StorageTypeNone:
		return true
	default:
		return false
	}
}
