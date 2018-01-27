package govern

import (
  "strconv"
  "testing"
)

func TestBencodeDecodeResponseMessage(t *testing.T) {
  evalResponse := "d2:id4:hurl2:ns9:boot.user7:session36:112830d1-8a29-4f55-88e9-6e7e735044c75:value17:#'boot.user/sevened2:id4:hurl7:session36:112830d1-8a29-4f55-88e9-6e7e735044c76:statusl4:doneee"

  output := bencodeUnmarshal(evalResponse)

  if output != "haha" {
    t.Errorf("output was not haha, instead got %v", output)
  }
}

// TODO this should be moved to it's own bencode(_test).go file
func TestBencodeDecode(t *testing.T) {
  bencodedString := "d2:id7:test-id" +
    "2:ns9:boot.user" +
    "7:session36:a647fb12-54ae-4313-8358-1161810de8f3" +
    "5:value17:#'boot.user/devile"

  output := bencodeUnmarshal(bencodedString)

  if output != "id,test-id,"+
    "ns,boot.user,"+
    "session,a647fb12-54ae-4313-8358-1161810de8f3,"+
    "value,#'boot.user/devil" {
    t.Errorf("the unmarshalled string was not haha, got %v", output)
  }
}

func TestBencodeMultipleDictionaryDecode(t *testing.T) {
  bencodedString := "d2:id7:test-id" +
    "2:ns16:watermarker.core" +
    "7:session36:57fea508-bc66-42af-b167-a3469da03ec" +
    "35:value30:#'watermarker.core/this-method" +
    "e" +
    "d" +
    "2:id7:test-id" +
    "7:session36:57fea508-bc66-42af-b167-a3469da03ec" +
    "36:statusl4:done" +
    "e" // + "e" // is this required? surely extra d at front if so.

  output := bencodeUnmarshal(bencodedString)

  expectedMessage := "id,test-id," +
    "ns,watermarker.core," +
    "session,57fea508-bc66-42af-b167-a3469da03ec3," +
    "value,#'watermarker.core/this-method"

  if output != expectedMessage {
    t.Errorf("the unmarshalled string was not %s, got %v", expectedMessage, output)
  }
}

func TestBencodeEncode(t *testing.T) {
  message := map[string]string{
    "id":      "test-id",
    "ns":      "boot.user",
    "session": "a647fb12-54ae-4313-8358-1161810de8f3",
    "value":   "#'boot.user/devil",
  }

  bencodedString := bencodeMarshall(message)

  expectedMessage := "d2:id7:test-id2:ns9:boot.user7:session36:a647fb12-54ae-4313-8358-1161810de8f35:value17:#'boot.user/devile"

  if bencodedString != expectedMessage {
    t.Errorf("the bencoded string doesn't match your expected string, got %v", bencodedString)
  }
}

func TestStringConcat(t *testing.T) {
  start := "hello"
  start += " there"

  if start != "hello there" {
    t.Errorf("string concat doesn't behave as you expect, got %v", start)
  }
}

func TestStringConversion(t *testing.T) {
  conversion, _ := strconv.Atoi("39")
  if conversion != 39 {
    t.Errorf("string conversion to int doesn't behave as you expect, got %v", conversion)
  }
}
