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
	var constantRunning bool = true
	if constantRunning{
		scanner := bufio.NewScanner(os.Stdin)
		for {
			currentDir, err := os.Getwd()
			if err != nil {
				print(err)
			}

			fmt.Print(currentDir,"~>")
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

		}
	}else{
		cmd.RootCmd.Execute()
	}
}
