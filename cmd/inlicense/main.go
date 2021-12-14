package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"coderlytics.at/inlicense/internal"
	"github.com/spf13/cobra"
)

func main() {
	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate the private and public key pair",
		Long:  `Generates the private and public key pair which get used to create and validate a license.`,
		Run:   generateKeyPair,
	}
	var licCmd = &cobra.Command{Use: "lic [license config file] [private key]",
		Short: "Generate a license file",
		Long:  `Geneates the license file using the private key with the information from the license configuration.`,
		Args:  cobra.MinimumNArgs(2),
		Run:   generateLicense,
	}

	var rootCmd = &cobra.Command{Use: "inlicense"}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(genCmd, licCmd)
	rootCmd.Execute()
}

// generateKeyPair generates the private and public key pair for creating and validating a license
func generateKeyPair(cmd *cobra.Command, args []string) {
	priv, pub, err := internal.GenerateBase64Keys()

	if err != nil {
		println("Error generating key pair")
		println(err.Error())
		os.Exit(1)
	}

	println(fmt.Sprintf("Private key (keep in a secure place): %s", priv))
	println(fmt.Sprintf("Public key (put into your code to verify the license): %s", pub))
}

// generateLicense generates the license file using the private key and the license configuration file
func generateLicense(cmd *cobra.Command, args []string) {
	cfgFile := args[0]
	privKey := args[1]

	if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
		println(fmt.Sprintf("Configuration file %s does not exist", cfgFile))
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		println(fmt.Sprintf("Error reading file %s", cfgFile))
		println(err.Error())
	}

	lic, err := internal.GenerateLicense(data, privKey)

	if err != nil {
		println("Unable to create license key")
		println(err.Error())
		os.Exit(1)
	}

	println(lic)
}
