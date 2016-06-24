package main

import (
	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"io"
	"net/http"
	"strconv"
)

type FlowTable struct {
	addressMap map[string]string
}

var flowTable *FlowTable

func runNetworkAnalyzer(networkInterface string) {
	if handle, err := pcap.OpenLive(networkInterface, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			tcpLayer := packet.Layer(layers.LayerTypeTCP)

			if ipLayer != nil && tcpLayer != nil {
				ip, _ := ipLayer.(*layers.IPv4)
				tcp, _ := tcpLayer.(*layers.TCP)

				srcPort := strconv.Itoa(int(tcp.SrcPort))
				dstPort := strconv.Itoa(int(tcp.DstPort))

				srcString := ip.SrcIP.String() + ":" + srcPort
				dstString := ip.DstIP.String() + ":" + dstPort
				flowTable.addressMap[srcString] = dstString
			}
		}
	}
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(flowTable.addressMap)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(data))
}

func main() {
	flowTable = &FlowTable{
		map[string]string{},
	}

	go runNetworkAnalyzer("eth0")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":7777", nil)
}
