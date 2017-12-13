package govern

import (
  "strconv"
  "strings"
  "sort"
  "fmt"
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
      }
    }
  }

  return strings.Join(words, ",")
}

func bencodeMarshall(message map[string]string) string {
  bencodedMessage := "d"
  var keys []string

  for k := range message {
    keys = append(keys, k)
  }

  sort.Strings(keys)

  for _, key := range keys {
    value := message[key]
    bencodedMessage += fmt.Sprintf("%v:%s", len(key), key)
    bencodedMessage += fmt.Sprintf("%v:%s", len(value), value)
  }

  bencodedMessage += "e"

  return bencodedMessage
}

