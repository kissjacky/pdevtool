/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kissjacky/pdevtool/util"
	"github.com/spf13/cobra"
)

// import flag "github.com/spf13/pflag"

var matchCls = regexp.MustCompile(`(\w+)\([^\(\)]+\)`)
var regStr = regexp.MustCompile(`"[^"]+"`)

// pyObj2jsonCmd represents the jsonkeys command
var pyObj2jsonCmd = &cobra.Command{
	Use:   "py2json",
	Short: "convert py log string to json",
	Run: func(cmd *cobra.Command, args []string) {
		omitNull, _ := cmd.Flags().GetString("omit_null")
		file, _ := cmd.Flags().GetString("file")
		absFile, err := filepath.Abs(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		dir := filepath.Dir(absFile)
		targetPyPath := filepath.Join(dir, "py2json.py")
		bs, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		orig := string(bs)
		orig = strings.ReplaceAll(orig, " ", "")
		orig = strings.ReplaceAll(orig, "\n", "")

		s := orig
		s = strings.ReplaceAll(s, "\\\"", "ESCAPE_DOUBLE_QUOTE")
		s = regStr.ReplaceAllString(s, `"STR_PLACEHOLDER"`)

		classes := map[string]bool{}
		for {
			mm := matchCls.FindAllStringSubmatch(s, -1)
			if len(mm) <= 0 {
				break
			}
			for _, v := range mm {
				if len(v) < 1 {
					continue
				}
				classes[v[1]] = true
			}

			s = matchCls.ReplaceAllString(s, "None")
		}
		fs, err := os.OpenFile(targetPyPath, os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		util.Templates["py2json.tmpl"].Execute(fs, map[string]interface{}{
			"Classes": classes,
			"ObjStr":  orig,
		})

		bs, _ = exec.Command(targetPyPath).Output()
		outputStr := string(bs)
		if omitNull == "1" {
			var tmp interface{}
			json.Unmarshal(bs, &tmp)
			removeNulls(tmp)
			bs, _ := json.Marshal(tmp)
			outputStr = string(bs)

		}
		fmt.Fprintln(os.Stdout, outputStr)
	},
}

func removeNulls(m interface{}) {
	if m == nil {
		return
	}
	switch v := m.(type) {
	case int:
	case float64:
	case string:
	case []interface{}:
		for _, u := range v {
			removeNulls(u)
		}
	case map[string]interface{}:
		for i, u := range v {
			if u == nil {
				delete(v, i)
			}
			removeNulls(u)
		}
	default:
		// i isn't one of the types above
	}
}

func init() {
	fs := pyObj2jsonCmd.PersistentFlags()
	fs.String("file", "./tmp/python_obj.log", "file contains python object string log")
	fs.String("omit_null", "0", "if omit null field, default 0,set 1 to omit")
	rootCmd.AddCommand(pyObj2jsonCmd)
}
