package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "Searches a file for pattern",
	Long:  `Using a search pattern, a file is opened, read line by line and if the pattern is found print it`,
	Run:   grep,
}

func grepFile(fileName string, pattern string) string {
	content := ""
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), pattern) {
			content = scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}

func grep(cmd *cobra.Command, args []string) {
	patternReturn := grepFile(args[0], args[1])
	fmt.Println(patternReturn)
}

func init() {
	RootCmd.AddCommand(grepCmd)
}
