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
	//sam "github.com/eyedeekay/sam3/helper"
)

type PeerClient struct {
	Debug      bool
	Peer	*Peer
	//Ctcpport   int
	PeerConn net.Conn
	Comp	   libdeflate.Compressor
	DeComp	   libdeflate.Decompressor
}

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
			this.p2phello(buf,protocol)
		}
		
		
	}
}

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
		/*if err != io.EOF {
			fmt.Println("ERROR:", err.Error())
			}
		*/
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

func (this *PeerClient) p2phello(buf []byte,protocol byte){
	debug:=this.Debug
	conn:=this.PeerConn
	dataindex:=1
	hashsize := int(buf[dataindex])
	dataindex+=1
	if debug {
		fmt.Println("Hash size", hashsize)
		fmt.Println("Hash", buf[dataindex:dataindex+hashsize])
	}
	dataindex+=hashsize
	if debug {
		fmt.Println("clientid", buf[dataindex:dataindex+4])
		fmt.Println("tcpport", buf[dataindex+4:dataindex+4+2])
	}
	dataindex+=4
	dataindex+=2
	tagcount := int(util.ByteToUint32(buf[dataindex:dataindex+4]))
	if debug {
		fmt.Println("tag count", tagcount)
	}
	dataindex+=4
		
	//if debug {
	//	fmt.Println("all else (p2phello)", buf[dataindex:len(buf)])
	//}
	
	//func ReadTags(pos int, buf []byte, tags int,debug bool)(totalread int, ret []*OneTag){
	tagsreadb, tagarr := util.ReadTags(dataindex,buf,tagcount,debug)
	if debug {
		fmt.Println("Tags read bytes / tags count", tagsreadb,len(tagarr))
	}
	for i := 0; i < len(tagarr); i++ {
		switch tagarr[i].NameByte {
			case 0x1:
				if tagarr[i].Type == byte(2) {
					if debug {
						fmt.Printf("Debug Name Tag: %s\n",tagarr[i].Value)
					}
				}
			case 0x11:
				if debug {
					fmt.Printf("Debug Version Tag: %d\n",util.ByteToUint32(tagarr[i].Value))
				}
			case 0x20:
				if debug {
					fmt.Printf("Debug Flags Tag: %b\n",util.ByteToUint32(tagarr[i].Value))
				}
			case 0x0f:
				if debug {
					fmt.Printf("Debug Port Tag: %d\n",util.ByteToUint32(tagarr[i].Value))
				}
			case 0x60:
				if debug {
					fmt.Printf("Debug ipv6 Tag: %d\n",tagarr[i].Value)
				}
			default:
				if debug {
					fmt.Printf("Warning: unknown tag 0x%x\n",tagarr[i].NameByte)
					fmt.Println(" ->Value: ",tagarr[i].Value)
				}
		}
	}
	
	//reply
	//replyHello(uuid_b []byte, clientid uint32, tcpport uint16, tags *[]util.OneTag, conn net.Conn)
	var tagarr []*util.OneTag
	username_b := util.StringToByte("username")
	tag = &util.OneTag{Type: 2, NameString: "", NameByte: 0x1, Value: username_b, ValueLen: uint16(len(username_b))}
	tagarr = append(tagarr,tag)
	tag = &util.OneTag{Type: 3, NameString: "", NameByte: 0x11, Value: util.ByteToUint32(60), ValueLen: uint16(4)}
	tagarr = append(tagarr,tag)
	replyHello(this.Peer.Uuid, 0, uint16(this.Peer.Port), &tagarr, conn net.Conn)
	//test
	
}
