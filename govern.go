package govern

import (
	"bytes"
	"flag"
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runCommand(command string) {
	args := strings.Fields(command)

	var cmd *exec.Cmd

	if len(args) < 2 {
		cmd = exec.Command(args[0])
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out.String())
}

// Converts markdown to html and opens the resulting file in chrome.
func governMarkdown(filePath string) {
	markdownFile, _ := ioutil.ReadFile(filePath)
	output := blackfriday.Run(markdownFile)
	dir, _ := ioutil.TempDir("", "markdown")
	outputPath := filepath.Join(dir, "converted_markdown.html")
	f, _ := os.Create(outputPath)
	runCommand(strings.Join([]string{"open", outputPath}, " "))
	f.Write(output)
	f.Sync()
	f.Close()
}

func main() {
	filePtr := flag.String("file", "", "File to determine a task for")
	flag.Parse()

	if *filePtr == "" {
		fmt.Println("no file supplied to govern, please use the -file flag")
		os.Exit(0)
	}

	// TODO find the full filepath by combining current dir with parameter
	filePath := *filePtr

	if strings.HasSuffix(filePath, ".clj") {
		fmt.Println("clojure file found, executing load-file")
		// TODO inspect file for deftest OR clojure.test for more stable address.
	} else if strings.HasSuffix(filePath, ".go") {
		fmt.Println("go file found, executing go test")
		runCommand("go test")
	} else if strings.HasSuffix(filePath, ".md") {
		governMarkdown(filePath)
	}

}
