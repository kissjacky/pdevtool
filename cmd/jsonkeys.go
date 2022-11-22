/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// jsonkeysCmd represents the jsonkeys command
var jsonkeysCmd = &cobra.Command{
	Use:   "jsonkeys",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// sigs := make(chan os.Signal, 1)
		// signal.Notify(sigs, syscall.SIGPIPE)
		// go func() {
		// 	for {
		// 		sig := <-sigs
		// 		fmt.Println("SIG:")
		// 		fmt.Println(sig)
		// 		os.Stderr.Write([]byte("pi\n"))
		// 		signal.Reset(syscall.SIGPIPE)
		// 	}
		// }()

		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Fprintln(os.Stderr, err)
				}
				break
			}
			var o map[string]interface{}
			err = json.Unmarshal([]byte(input), &o)
			if err != nil {
				fmt.Println(input, err)
				continue
			}
			keys := sort.StringSlice{}
			for k := range o {
				keys = append(keys, k)
			}
			sort.Sort(keys)
			fmt.Fprintln(os.Stdout, strings.Join(keys, ","))
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonkeysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonkeysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonkeysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
