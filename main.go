package main

import (
	"flag"
	"fmt"
	"strings"
	// "net"
)

var (
	outputFlag string
)

type PortList []string
type HostList []string

func (hl *HostList) String() string {
	return fmt.Sprintln(*hl)
}

func (pl *PortList) String() string {
	return fmt.Sprintln(*pl)
}

func (hl *HostList) Set(value string) error {
	*hl = strings.Split(value, ",")
	return nil
}

func (pl *PortList) Set(value string) error {
	*pl = strings.Split(value, ",")
	return nil
}

func main() {
	var hosts HostList
	var ports PortList
	flag.Var(&hosts, "h", "Host to scan")
	flag.Var(&ports, "p", "Port to scan")
	outputFlag := flag.String("o", "outfile.txt", "Output format")
	flag.Parse()
	fmt.Println("Hosts:", hosts)
	fmt.Println("Ports:", ports)
	fmt.Println("Output:", *outputFlag)

}
