package main

import (
  "fmt"
  "flag"
  "os"
  "os/exec"
  "log"
  "bytes"
  "strings"
  "strconv"
)

// TODO this will eventually return a map
// TODO worth noting we don't detect the end of a dictionary
func bencodeUnmarshal(bencodedString string) string {
  parsingBuffer := ""
  remainingValueLength := 0
  parsingDictionary := false

  var words []string

  for _, c := range bencodedString {
    char := string(c)

    if parsingDictionary {
      if remainingValueLength != 0 {
        parsingBuffer += char
        remainingValueLength--
        if remainingValueLength == 0 { // finished parsing element value
          words = append(words, parsingBuffer)
          parsingBuffer = ""
        }
      } else if char == ":" { // finished parsing element length
        remainingValueLength, _ = strconv.Atoi(parsingBuffer)
        parsingBuffer = ""
      } else { // we must be parsing element length
        parsingBuffer += char
      }
    } else {
      if char == "d" { // we are parsing a dictionary
        parsingDictionary = true
        fmt.Println("begin parsing dictionary")
      }
    }
  }

  return strings.Join(words, ",")
}

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
  }

}
