package main

import (
  "net"
  "bufio"
  "fmt"
  "time"
)

func main() {
  // expectedMessage := "d2:id7:test-id2:ns9:boot.user7:session36:a647fb12-54ae-4313-8358-1161810de8f35:value17:#'boot.user/devile"

  newSessionMessage := "d2:op5:clonee"

  // message := "d4:code"+
  // "16:(def thing 123)"+
  // "2:id5:an-id"+
  // "2:op4:eval"+
  // "7:session9:a-sessione"

  replAddress, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
  socket, _ := net.DialTCP("tcp4", nil, replAddress)
  socket.Write([]byte(newSessionMessage))

  bufReader := bufio.NewReader(socket)

  timeoutDuration := 5 * time.Second
  socket.SetReadDeadline(time.Now().Add(timeoutDuration))
  bytes, _ := bufReader.ReadBytes('\n')

  fmt.Printf("%s", bytes)

  // socket.Write([]byte(message))
  socket.Close()
}
