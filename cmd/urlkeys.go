/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// urlkeysCmd represents the jsonkeys command
var urlkeysCmd = &cobra.Command{
	Use:   "urlkeys",
	Short: "url query keys",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Fprintln(os.Stderr, err)
				}
				break
			}
			query := strings.Split(input, "&")
			kvs := []string{}
			for _, q := range query {
				kv := strings.Split(q, "=")

				//把 abc[0] 归一为 abc
				tmp := strings.Split(kv[0], "[")

				kvs = append(kvs, tmp[0])
			}
			sort.Strings(kvs)
			fmt.Fprintln(os.Stdout, strings.Join(kvs, ","))
		}
	},
}

func init() {
	rootCmd.AddCommand(urlkeysCmd)
}
