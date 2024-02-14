/*
Copyright © 2024 jspc <james@zero-internet.org.uk>

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
	"os"

	"github.com/jspc/synthdown"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const errInvalid = 1

var failer failerFunc = os.Exit

type failerFunc = func(int)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a synthdown configuration",
	Long: `Validate a synthdown configuration

This command can be used to read files, and validate them as synthdown`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := synthdown.New(viper.GetString("file"))
		if err != nil {
			fmt.Printf("%v\n", err)
			failer(errInvalid)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
