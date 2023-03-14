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
	"strings"

	"github.com/spf13/cobra"
)

// cntUniqCmd represents the jsonkeys command
var cntUniqCmd = &cobra.Command{
	Use:   "cntUniq",
	Short: "cnt uniq words splitted by commas",
	Run: func(cmd *cobra.Command, args []string) {
		m := map[string]int{}
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Fprintln(os.Stderr, err)
				}
				break
			}
			query := strings.Split(input, ",")
			for _, q := range query {
				m[q]++
			}
		}
		bs, _ := json.Marshal(m)
		fmt.Fprintln(os.Stdout, string(bs))
	},
}

func init() {
	rootCmd.AddCommand(cntUniqCmd)
}
