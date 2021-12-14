package main

import (
	"fmt"
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
	var licCmd = &cobra.Command{Use: "lic [private key] [license config]",
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
	privateKey, err := internal.NewPrivateKey()

	if err != nil {
		println("Error generating private key")
		println(err)
		os.Exit(1)
	}

	b64Key, err := privateKey.ToBase64String()

	if err != nil {
		println("Error converting private key to base64 string")
		println(err)
		os.Exit(1)
	}
	println(fmt.Sprintf("Private key (keep in a secure place): %s", b64Key))

	publicKey := privateKey.GetPublicKey()
	b64pKey := publicKey.ToBase64String()
	println(fmt.Sprintf("Public key (put into your code to verify the license): %s", b64pKey))
}

// generateLicense generates the license file using the private key and the license configuration file
func generateLicense(cmd *cobra.Command, args []string) {
	// TODO implement
}
