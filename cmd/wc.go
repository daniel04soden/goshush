package cmd

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "Counts the amount of abritrary words present in a given file",
	Long:  `Passes in a single file and from there, loops through the file and keeps a count of the running words`,
	Run:   wc,
}

func countWords(file *os.File) int{
	scanner:=bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	var wordCount int = 0

	for scanner.Scan(){
		wordCount++
	}

	return wordCount
}

func countLines(file *os.File) int{
	scanner:=bufio.NewScanner(file)

	var lineCount int = 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func countSize(file *os.File) int64{
	info,err:=os.Stat(file.Name())

	if err!=nil{
		log.Fatal((err))
	}

	return info.Size()
}

func wc(cmd *cobra.Command, args []string){
	fileName := strings.Join(args, " ")

	file,err:= os.Open(fileName)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()
	wordCount := countWords(file)
    file.Seek(0, 0) // Reset back
    lineCount := countLines(file)
    file.Seek(0, 0) // Reset back
    size := countSize(file)

    fmt.Println("Words: ", wordCount, " | Lines: ", lineCount, " | Size/Bytes: ", size)
}

func init() {
	RootCmd.AddCommand(wcCmd)
}
