package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var headCmd = &cobra.Command{
	Use:   "head",
	Short: "Prints the top contents of a file",
	Long:  `head is a command that loops through a file passed and liss only the top contents of the file`,
	Run:   head,
}

func readTopFile(fileName string) string {
	var res string
	file,err:= os.Open(fileName)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()


	var linesRead = 0
	s:= bufio.NewScanner(file)

	for s.Scan() && linesRead < 10{
		res += s.Text() + "\n"
		linesRead++
	}

	return res
}

func head(cmd *cobra.Command, args []string) {
	content := readTopFile(strings.Join(args, " "))
	fmt.Println(content)
}

func init() {
	RootCmd.AddCommand(headCmd)
}
