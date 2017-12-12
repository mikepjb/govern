package main

import (
  "strconv"
  "testing"
)

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

func TestBencodeEncode(t *testing.T) {
  message := map[string]string{
    "id": "test-id",
    "ns": "boot.user",
    "session": "a647fb12-54ae-4313-8358-1161810de8f3",
    "value": "#'boot.user/devil",
  }

  bencodedString := bencodeMarshall(message)

  if bencodedString != "something special" {
    t.Errorf("the bencoded string was not something special, got %v", bencodedString)
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
