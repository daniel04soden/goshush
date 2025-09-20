package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Prints the bottom contents of a file",
	Long:  `tail is a command that loops through a file passed and liss only the bottom contents of the file`,
	Run:   tail,
}

func readBottomFile(fileName string) string {
	var store []string
	var res string
	var EOF int = 0

	file,err:= os.Open(fileName)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()


	s:= bufio.NewScanner(file)

	for s.Scan() {
		store = append(store,s.Text())
		EOF++
	}

	for i:=EOF-10; i<EOF;i++ {
		res += store[i] + "\n"
	}

	return res
}

func tail(cmd *cobra.Command, args []string) {
	content := readBottomFile(strings.Join(args, " "))
	fmt.Println(content)
}

func init() {
	RootCmd.AddCommand(tailCmd)
}
