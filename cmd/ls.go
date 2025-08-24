package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	long   bool
	hidden bool
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all files in the current directory",
	Long:  "using the io/ioutil package, we list all files found in our current directory",
	Run:   ls,
}

func ls(cmd *cobra.Command, args []string) {
	files, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !hidden && strings.HasPrefix(file.Name(), ".") {
			continue
		} else {
			if long {
				info, err := file.Info()
				if err != nil {
					log.Fatal(err)
					continue
				}
				if file.IsDir() {
					fmt.Printf("📁 %s\t%d b \t%s\n", info.Mode(), info.Size(), file.Name())
				} else {
					fmt.Printf("%s\t%d b \t%s\n", info.Mode(), info.Size(), file.Name())
				}
			} else {
				if file.IsDir() {
					fmt.Println("📁", file.Name())
				} else {
					fmt.Println(file.Name())
				}
			}
		}
	}
}

func init() {
	lsCmd.Flags().BoolVarP(&long, "long", "l", false, "Shows more info on list")
	lsCmd.Flags().BoolVarP(&hidden, "all", "a", false, "Includes hidden files")
	RootCmd.AddCommand(lsCmd)
}
