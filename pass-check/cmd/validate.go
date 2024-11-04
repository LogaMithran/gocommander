/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var fpassword string

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates the given password is strong or not",
	Long:  `Validate will validates the given password is strong or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(fpassword)
		hasLower := regexp.MustCompile(`[a-z]`).MatchString(fpassword)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(fpassword)
		hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_.,]`).MatchString(fpassword)

		// Return true if all criteria are met
		if hasUpper && hasLower && hasNumber && hasSpecial {
			fmt.Printf("Your password << %v >> is strong", fpassword)
		} else {
			fmt.Printf("Your password << %v >> is not strong", fpassword)
		}

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	validateCmd.Flags().StringVarP(&fpassword, "password", "p", "", "To give password using the flag")
	if err := validateCmd.MarkFlagRequired("password"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(validateCmd)

}
