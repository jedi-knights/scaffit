/*
Copyright © 2023 Omar Crosby <omar.crosby@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project in a specified directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err    error
			result string
		)

		prompt := promptui.Select{
			Label: "Select Module Path Type",
			Items: []string{"github.com", "gitlab.com", "bitbucket.org", "other"},
		}

		if _, result, err = prompt.Run(); err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		switch result {
		case "github.com":
			fmt.Println("github.com")
		case "gitlab.com":
			fmt.Println("gitlab.com")
		case "bitbucket.org":
			fmt.Println("bitbucket.org")
		case "other":
			fmt.Println("other")
		}

		prompt := promptui.Prompt{
			Label: "What is your module's name?",
			Validate: func(input string) error {
				// Module names should be in lowercase letters and avoid hyphens (-) or special characters.
				// Use underscores or camelCase for multi-word module names. For example, mylibrary or
				// myLibrary.
				if len(input) < 1 {
					return fmt.Errorf("Module name must be at least 1 character")
				}

				// Check to ensure the input is all lowercase
				for _, c := range input {
					if c < 'a' || c > 'z' {
						return fmt.Errorf("Module name must be all lowercase")
					}
				}

				// Check to ensure the input does not contain a hyphen
				if strings.Contains(input, "-") {
					return fmt.Errorf("Module name must not contain a hyphen")
				}

				// Check to ensure the input does not contain a special character
				if strings.ContainsAny(input, "~!@#$%^&*()+`-={}|[]\\:\";'<>?,./") {
					return fmt.Errorf("Module name must not contain a special character")
				}

				// Check to ensure the input does not contain a space
				if strings.Contains(input, " ") {
					return fmt.Errorf("Module name must not contain a space")
				}

				return nil
			},
		}

		if _, result, err = prompt.Run(); err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		fmt.Printf("Your module will be named %q\n", result)

		fmt.Println("new called")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
