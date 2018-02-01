package main

import (
  "fmt"
  "flag"
  "os"
  "io/ioutil"
  "os/exec"
  "log"
  "bytes"
  "strings"
  "github.com/russross/blackfriday"
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
    markdownFile, _ := ioutil.ReadFile(filePath)
    output := blackfriday.Run(markdownFile)
    f, _ := os.Create("thing.html")
    f.Write(output)
    f.Sync()
    f.Close()
  }

}
