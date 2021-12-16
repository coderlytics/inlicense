package internal

import "time"

// License holds the license information with all details
type License struct {
	ValidFrom time.Time `mapstructure:"validFrom" validate:"datetime" default:"01.01.1970 00:00"`
	ValidTo   time.Time `mapstructure:"validTo" validate:"datetime,required"`
	Product   Product   `mapstructure:"product" validate:"required"`
}

// Product holds the licensed products with the licensed modules
type Product struct {
	Name    string `mapstructure:"name" validate:"alphanum,required"`
	Modules []Module
}

// Module holds the licensed modules with the licensed features
type Module struct {
	Features []Feature `mapstructure:"features" validate:"required"`
}

// Feature holds the licensed features with custom configuration parameters
type Feature struct {
	Name string `mapstructure:"name" validate:"alphanum,required"`
}

// GenerateLicense generates a new license with the given data using the private key for signing
func GenerateLicense(data string, privateKey string) (string, error) {
	println(data)
	println(privateKey)
	return "", nil
}
