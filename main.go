/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/kissjacky/pdevtool/cmd"
)

func main() {
	err := LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
