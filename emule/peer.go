package emule

import (
	"fmt"
	// "io"
	"net"
	util "github.com/AltTechTools/gomule-tst/emule"
	"time"
	//"github.com/test3-damianfurrer/gomule/tree/sharedtest/emule"
	libdeflate "github.com/4kills/go-libdeflate/v2"
)

type Peer struct {
	Server     string
	Port       int
	Username   string
	Uuid	   []byte
	Debug      bool
	//Ctcpport   int
	///ClientConn net.Conn
	Comp	   libdeflate.Compressor
	DeComp	   libdeflate.Decompressor
	SrvTCPCompression         bool
	SrvTCPNewTags             bool
	SrvTCPUnicode             bool
	SrvTCPRelatedSearch       bool
	SrvTCPTypeTagInterger     bool
	SrvTCPLargeFiles          bool
	SrvTCPObfuscation         bool
}

func NewPeerInstance(server string, port int, debug bool) *Client {
	return &Client{
		Server:   server,
		Port:     port,
		Username: "gomuleclientuser",
		Debug:   debug}
}
func (this *Peer) SetTCPFlags(tcpmap uint32){
	this.SrvTCPCompression		= false
	this.SrvTCPNewTags		= false
	this.SrvTCPUnicode		= false
	this.SrvTCPRelatedSearch 	= false
	this.SrvTCPTypeTagInterger 	= false
	this.SrvTCPLargeFiles 		= false
	this.SrvTCPObfuscation 		= false
	
	if tcpmap & uint32(0x00000001) != 0 {
		this.SrvTCPCompression = true
		fmt.Println("this.SrvTCPCompression")
	}
	if tcpmap & uint32(0x00000008) != 0 {
		this.SrvTCPNewTags = true
		fmt.Println("this.SrvTCPNewTags")
	}
	if tcpmap & uint32(0x00000010) != 0 {
		this.SrvTCPUnicode = true
		fmt.Println("this.SrvTCPUnicode")
	}
	if tcpmap & uint32(0x00000040) != 0{
		this.SrvTCPRelatedSearch = true
		fmt.Println("this.SrvTCPRelatedSearch")
	}
	if tcpmap & uint32(0x00000080) != 0 {
		this.SrvTCPTypeTagInterger = true
		fmt.Println("this.SrvTCPTypeTagInterger")
	}
	if tcpmap & uint32(0x00000100) != 0 {
		this.SrvTCPLargeFiles = true
		fmt.Println("this.SrvTCPLargeFiles")
	}
	if tcpmap & uint32(0x00000400) != 0 {
		this.SrvTCPObfuscation = true
		fmt.Println("this.SrvTCPObfuscation")
	}
/*
		// Server TCP flags
#define SRV_TCPFLG_COMPRESSION          0x00000001
#define SRV_TCPFLG_NEWTAGS                      0x00000008
#define SRV_TCPFLG_UNICODE                      0x00000010
#define SRV_TCPFLG_RELATEDSEARCH        0x00000040
#define SRV_TCPFLG_TYPETAGINTEGER       0x00000080
#define SRV_TCPFLG_LARGEFILES           0x00000100
#define SRV_TCPFLG_TCPOBFUSCATION	0x00000400
		*/

}
