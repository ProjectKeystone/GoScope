package main

import (
	"strconv"
	"time"
  "net"
)

func PortScan(protocol, hostname string, port int) bool {
  address := hostname + ":" + strconv.Itoa(port)

  connection, err := net.DialTimeout(protocol, address, 60*time.Second)
  if err != nil {
    //returns false if the port is not accepting connections
    return false
  }
  defer connection.Close()

  return true
}
