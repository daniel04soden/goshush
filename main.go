/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/daniel04soden/goshush/cmd"
)

func main() {
	var constantRunning bool = false
	//TODO Fix inconsistencies with repeated flag commands like ls and ps
	if constantRunning{
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("goshush> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break
		}
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd.RootCmd.SetArgs(args)
		cmd.RootCmd.Execute()

		//TODO fix issues with cd where directory isn't fully changed - works in constant running
	}}else{
		cmd.RootCmd.Execute()
	}
}
