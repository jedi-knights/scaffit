/*
Package cmd implements the command line interface for scaffit.

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
	"github.com/jedi-knights/scaffit/pkg"
	"github.com/jedi-knights/scaffit/pkg/generators"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Creates a new module component.",
	Long:  `Creates a new module component.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err      error
			location string
		)

		if location, err = newCmd.Flags().GetString("location"); err != nil {
			log.Fatalf("Error reading location: %v\n", err)
		}

		if location, err = filepath.Abs(location); err != nil {
			log.Fatalf("Error resolving absolute path: %v", err)
		}

		log.Printf("location: %s\n", location)

		fsys := pkg.NewFileSystem()

		if err = generators.NewModuleGenerator(*fsys, location).Generate(); err != nil {
			log.Fatalf("Error generating module: %v\n", err)
		}
	},
}

func init() {
	newCmd.AddCommand(moduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	_ = moduleCmd.MarkFlagRequired("location")
}
