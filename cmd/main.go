package main

import (
	"flag"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"sdigo/pkg/data"
)

func main() {
	file := flag.String("file", "", "A file to parse")
	rulesPath := flag.String("rules", "", "A file with rules")

	flag.Parse()

	if *file == "" {
		log.Fatal("PCAP file not found")
		return
	}

	if *rulesPath == "" {
		log.Fatal("Rules file not found")
		return
	}

	var handler *pcap.Handle
	var err error

	// Open the pcap file or device.
	if *file != "" {
		if handler, err = pcap.OpenOffline(*file); err != nil {
			panic(err)
		}
	}

	rules, err := data.ParseRules(*rulesPath)
	if err != nil {
		log.Fatalf("an error occurred while parsing json file: %s", err.Error())
		return
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range source.Packets() {
		data.ProcessPacketInfo(*rules, packet)
	}
}
