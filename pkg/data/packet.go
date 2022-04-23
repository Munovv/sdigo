package data

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"strings"
)

func ProcessPacketInfo(rules []Rule, packet gopacket.Packet) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		glog.Info("TCP layer is detected.")

		tcphandler, _ := tcpLayer.(*layers.TCP)
		srcport := tcphandler.SrcPort
		destport := tcphandler.DstPort

		iplayer := packet.Layer(layers.LayerTypeIPv4)
		httphandler, _ := iplayer.(*layers.IPv4)
		srcip := httphandler.SrcIP
		destip := httphandler.DstIP

		dstMsg := fmt.Sprintf("%s:%s", destip, destport)
		srcMsg := fmt.Sprintf("%s:%s", srcip, srcport)

		msgRes := ""
		for _, rule := range rules {
			if msg, _ := DetectRule(rule, packet); msg != "" {
				msgRes += msg + " "
			}
		}

		if msgRes != "" {
			fmt.Printf("%s --> %s . DETECTED: { %s}\n", dstMsg, srcMsg, msgRes)
		}
	} else {
		fmt.Println("TCP layer is not detected.")
	}
}

func DetectRule(rule Rule, packet gopacket.Packet) (string, []byte) {
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		if strings.Contains(string(applicationLayer.Payload()), rule.Value) {

			if glog.V(1) {
				glog.Infof("%s", rule.Value)
			}
			return fmt.Sprintf("%s", rule.Value), applicationLayer.LayerContents()
		} else {
			return "", nil
		}
	} else {
		return "", nil
	}
}
