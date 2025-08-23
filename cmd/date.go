package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Displays the date and time",
	Long:  "In a long form, depending on flags, displays the date time and day of the week",
	Run:   date,
}

func date(cmd *cobra.Command, args []string) {
	timeCurrent := time.Now()
	fmt.Println(timeCurrent.Day(), "/", timeCurrent.Month(), "/", timeCurrent.Year(), "|", timeCurrent.Hour(), ":", timeCurrent.Minute(), ":", timeCurrent.Second())
}

func init() {
	RootCmd.AddCommand(dateCmd)
}
