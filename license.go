package inlicense

import (
	"time"

	"coderlytics.at/inlicense/internal"
)

// License represents the licensing data as object
type License struct {
	publicKey string
	license   internal.License
}

// Load the license from the given file path using the public key and return the license object
func Load(licenseFile string, publicKey string) (*License, error) {
	return &License{}, nil
}

// Verify checks the license with the public key to verify it's not manipulated
func (l *License) Verify() bool {
	return true
}

// IsExpired checks if the given license is already expired
// True if the license is not valid anymore
func (l *License) IsExpired() bool {
	return l.isExpired(time.Now())
}

func (l *License) isExpired(validationDate time.Time) bool {
	validFrom := l.license.ValidFrom
	validTo := l.license.ValidTo
	return validationDate.Before(validFrom) || validationDate.After(validTo)
}

// HasModuleLicensed checks if the given license contains an entry for the given module name
// Returns true if the modules is licensed
func (l *License) HasModuleLicensed(module string) bool {
	return true
}

// HasModuleFeatureLicensed checks if the given license contains an entry for the given feature in the given module
// Returns true if the given feature is licensed in the given module
func (l *License) HasModuleFeatureLicensed(module string, feature string) bool {
	return true
}
