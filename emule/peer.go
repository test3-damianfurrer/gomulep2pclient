package emule

import (
	"fmt"
	"io"
	"net"
	//util "github.com/AltTechTools/gomule-tst/emule"
	util "github.com/AltTechTools/gomule-SockSrvClient/emule"
	//"time"
	"errors"
	//"github.com/test3-damianfurrer/gomule/tree/sharedtest/emule"
	libdeflate "github.com/4kills/go-libdeflate/v2"
	sam "github.com/eyedeekay/sam3/helper"
)

type Peer struct {
	Host     string
	Port       int
	Username   string
	Uuid	   []byte
	Debug      bool
	I2P	   bool
	SAM      string
	SAMPort  int
	//Ctcpport   int
	///ClientConn net.Conn
	//Comp	   libdeflate.Compressor
	//DeComp	   libdeflate.Decompressor
	listener	net.Listener
	SrvTCPCompression         bool
	SrvTCPNewTags             bool
	SrvTCPUnicode             bool
	SrvTCPRelatedSearch       bool
	SrvTCPTypeTagInterger     bool
	SrvTCPLargeFiles          bool
	SrvTCPObfuscation         bool
}
/*
type PeerClient struct {
	Debug      bool
	Peer	*Peer
	//Ctcpport   int
	PeerConn net.Conn
	Comp	   libdeflate.Compressor
	DeComp	   libdeflate.Decompressor
}*/

func NewPeerInstance(server string, port int, debug bool) *Peer {
	return &Peer{
		Host:   server,
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

func (this *Peer) yoursam() string {
	return fmt.Sprintf("%s:%d", this.SAM, this.SAMPort)
}

func (this *Peer) Start() {
	var ln net.Listener
	var err error
	if this.I2P {
		ln, err = sam.I2PListener("go-imule-servr", this.yoursam(), "go-imule-server")
	} else {
		ln, err = net.Listen("tcp", fmt.Sprintf("%s:%d", this.Host, this.Port))
	}
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	this.listener = ln
	fmt.Printf("Starting peer %s:%d\n", this.Host, this.Port)

	for {
		conn, err := this.listener.Accept()
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			continue
		}
		go this.respConn(conn)
	}
}
/*
func (this *PeerClient) Start() {
	for {
		buf, protocol, err, buflen := this.read(this.PeerConn) //maybe get this from this instead
		//tst
		fmt.Println("Protocol",protocol)
		fmt.Println("Buf len",buflen)
		if err != nil {
			if err == io.EOF {
				if this.Debug {
				    fmt.Printf("DEBUG: %v disconnected\n", this.PeerConn.RemoteAddr())
				}
				//logout(uhash, this.Debug, this.db) //logout(chigh_id, cport, this.Debug, this.db)
			} else if errors.Is(err, net.ErrClosed) {
				if this.Debug {
					fmt.Println("DEBUG: conn closed due to invalid client data")
				}
			}else {
				fmt.Println("ERROR: from read:", err.Error())
			}
			this.DeComp.Close()
			this.Comp.Close()
			this.PeerConn.Close()
			return
		}
		if this.Debug {
			fmt.Printf("DEBUG: type 0x%02x\n", buf[0])
		}
		if buf[0] == 0x01 { //p2p hello
			//uhash = login(buf, protocol, conn, this.Debug, this.db,HighId(this.Host),uint16(this.Port), this.Ssname, this.Ssdesc, this.Ssmsg, this.getTCPFlags())//chigh_id, cport, uhash = login(buf, protocol, conn, this.Debug, this.db)
			p2phello(buf,protocol,this.PeerConn,this.Debug)
		}
		
		
	}
}*/

func (this *Peer) respConn(conn net.Conn) {
	//var chigh_id uint32
	//var cport int16
	
	//test
	var err error
	//uhash := make([]byte, 16)
	//client := SockSrvClient{Conn: conn}
	pc := PeerClient{PeerConn: conn, Peer: this, Debug: this.Debug}
	
	
	pc.DeComp, err = libdeflate.NewDecompressor()
	if err != nil {
		fmt.Println("ERROR libdeflate Decompressor:", err.Error())
		return
	}
	pc.Comp, err = libdeflate.NewCompressor()
	if err != nil {
		fmt.Println("ERROR libdeflate Compressor:", err.Error())
		return
	}
	
	pc.Start()
	
	if this.Debug {
		fmt.Printf("DEBUG: %v connected\n", conn.RemoteAddr())
	}
	/*
	for {
		buf, protocol, err, buflen := this.read(conn)
		//tst
		fmt.Println("Protocol",protocol)
		fmt.Println("Buf len",buflen)
		if err != nil {
			if err == io.EOF {
				if this.Debug {
				    fmt.Printf("DEBUG: %v disconnected\n", conn.RemoteAddr())
				}
				//logout(uhash, this.Debug, this.db) //logout(chigh_id, cport, this.Debug, this.db)
			} else if errors.Is(err, net.ErrClosed) {
				if this.Debug {
					fmt.Println("DEBUG: conn closed due to invalid client data")
				}
			}else {
				fmt.Println("ERROR: from read:", err.Error())
			}
			pc.DeComp.Close()
			pc.Comp.Close()
			pc.PeerConn.Close()
			return
		}
		if this.Debug {
			fmt.Printf("DEBUG: type 0x%02x\n", buf[0])
		}
		if buf[0] == 0x01 { //p2p hello
			//uhash = login(buf, protocol, conn, this.Debug, this.db,HighId(this.Host),uint16(this.Port), this.Ssname, this.Ssdesc, this.Ssmsg, this.getTCPFlags())//chigh_id, cport, uhash = login(buf, protocol, conn, this.Debug, this.db)
			p2phello(buf,protocol,conn,this.Debug)
		}
		
		
	}
	*/
}
/*
func (this *PeerClient) read(conn net.Conn) (buf []byte, protocol byte, err error, buflen int) {
	//possible protocols:
	//0xe3 - ed2k
	//0xc5 - emule
	//0xd4 -zlib compressed
	protocol = 0xE3
	buf = make([]byte, 5)
	err = nil
	var n int = 0

	n, err = conn.Read(buf)
	if err != nil {
		/if err != io.EOF {
			fmt.Println("ERROR:", err.Error())
			}
		//
		return
	}
	if buf[0] == 0xE3 {
		protocol = 0xE3
	} else if buf[0] == 0xD4 {
		protocol = 0xD4
	} else if buf[0] == 0xC5 {
		protocol = 0xC5
	} else {
		fmt.Printf("ERROR: unsuported protocol 0x%02x\n", protocol)
		err = errors.New("unsuported protocol")
		return
	}
	if this.Debug {
		fmt.Printf("DEBUG: selected protocol 0x%02x(by byte 0x%02x)\n", protocol, buf[0])
	}
	size := util.ByteToUint32(buf[1:n])
	//if this.Debug {
	//	fmt.Printf("DEBUG: size %v -> %d\n", buf[1:n], size)
	//}
	buf = make([]byte, 0)
	toread := size
	var tmpbuf []byte
	for{
		if toread > 1024  {
			tmpbuf = make([]byte, 1024)
		} else {
			tmpbuf = make([]byte, toread)
		}
		n, err = conn.Read(tmpbuf)
		if err != nil {
			fmt.Println("ERROR: on read to buf", err.Error())
			//return
		}
		buf = append(buf, tmpbuf[0:n]...)
		if n < 0 {
			fmt.Println("WARNING: n (conn.Read) < 0, some problem")
			n = 0
		}
		toread -= uint32(n)
		if toread <= 0 {
			break;
		}
	}
	//buf = make([]byte, size)
	//n, err = conn.Read(buf)
	//if err != nil {
	//	fmt.Println("ERROR: on read to buf", err.Error())
	//	//return
	//}
	n = int(size-toread)
	if this.Debug {
		fmt.Printf("DEBUG: size %d, n %d\n", size, n)
	}
	buflen = n
	return
}
*/
