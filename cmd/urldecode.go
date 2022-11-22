/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// urldecodeCmd represents the urldecode command
var urldecodeCmd = &cobra.Command{
	Use:   "urldecode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			// Read the keyboad input.
			input, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Fprintln(os.Stderr, err)
				}
				break
			}
			input, err = url.QueryUnescape(input)
			fmt.Fprint(os.Stdout, input)
		}
	},
}

func init() {
	rootCmd.AddCommand(urldecodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urldecodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urldecodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
