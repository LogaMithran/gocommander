package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand/v2"
)

var (
	lowerCase string = "abcdefghijklmnopqrstuvxyz"
	upperCase string = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	numbers   string = "0123456789"
	special   string = "£$&()*+[]@#^-_!?"
	iteration int    = 4
	password  string = ""
)

func giveStrongPassword() {
	for range 15 {
		step := rand.IntN(iteration)
		switch step {
		case 0:
			lowerCaseIndex := int(rand.Float64() * float64(len(lowerCase)-1))
			password += string(lowerCase[lowerCaseIndex])
		case 1:
			upperCaseindex := int(rand.Float64() * float64(len(upperCase)-1))
			password += string(upperCase[upperCaseindex])
		case 2:
			numberIndex := int(rand.Float64() * float64(len(numbers)-1))
			password += string(numbers[numberIndex])
		case 3:
			specialIndex := int(rand.Float64() * float64(len(special)-1))
			password += string(special[specialIndex])
		}
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "To get some strong password",
	Long:  `A command that will give strong passwords that are unable to crack.`,
	Run: func(cmd *cobra.Command, args []string) {
		giveStrongPassword()
		fmt.Printf("\n\nYour strong passoword is %v.\nKeep it secret", password)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
