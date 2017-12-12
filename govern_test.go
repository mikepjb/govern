package main

import (
  "testing"
)

// TODO this should be moved to it's own bencode(_test).go file
func TestBencodeDecode(t *testing.T) {
  bencodedString := "d2:id7:test-id" +
  "2:ns9:boot.user" +
  "7:session36:a647fb12-54ae-4313-8358-1161810de8f3" +
  "5:value17:#'boot.user/devile"

  output := bencodeUnmarshal(bencodedString)

  if output != "haha" {
    t.Errorf("the unmarshalled string was not haha, got %v", output)
  }
}
