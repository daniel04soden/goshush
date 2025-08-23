package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "Touch creates a file",
	Long:  `Touch takes in a flag for the name from a user to create a file of that name`,
	Run:   touch,
}

func touch(cmd *cobra.Command, args []string) {
	fileName := strings.Join(args, " ")
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name())
}

func init() {
	RootCmd.AddCommand(touchCmd)
}
