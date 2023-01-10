package emule

import (
	"fmt"
	util "github.com/AltTechTools/gomule-SockSrvClient/emule"
	"net"
	libdeflate "github.com/4kills/go-libdeflate/v2" //libdeflate.Compressor
)

/*
body := make([]byte,0)
	//body = append(body,0x6a,0xff,0x9d,0x13,0xba,0x4f,0x4b,0x67,0xaf,0x0c,0xf6,0xa5,0x14,0xc4,0xd4,0x99) //client uuid this.Uuid
	body = append(body,this.Uuid...) //client uuid
	abuf := util.UInt32ToByte(uint32(0))
	body = append(body,abuf...) //client id 0 default
	body = append(body,util.UInt16ToByte(uint16(this.Ctcpport))...) //tcp port default
	body = append(body,util.UInt32ToByte(uint32(4))...) //tag count
	body = append(body,util.EncodeByteTagString(util.EncodeByteTagNameInt(0x1),this.Username)...)
	body = append(body,util.EncodeByteTagInt(util.EncodeByteTagNameInt(0x11),uint32(0x3C))...)
	body = append(body,util.EncodeByteTagInt(util.EncodeByteTagNameInt(0x20),uint32(0b1100011101))...)
	body = append(body,util.EncodeByteTagInt(util.EncodeByteTagNameInt(0xfb),util.ByteToUint32([]byte{128, 13, 4, 3}))...)
	//body = append(body,util.EncodeByteTagInt(util.EncodeByteTagNameInt(0x20),)...)
	
	fmt.Println("Size body", len(body))
	data := util.EncodeByteMsg(0xE3,0x01,body)
	this.ClientConn.Write(data)
  */
//tag types:
//2 string
//3 int
//4 float

func replyHello(uuid_b []byte, clientid uint32, tcpport uint16, tags *[]util.OneTag, conn net.Conn){
	sendHello(0x4C,uuid_b,clientid,tcpport,tags,conn)
}

func greetHello(uuid_b []byte, clientid uint32, tcpport uint16, tags *[]util.OneTag, conn net.Conn){
	sendHello(0x01,uuid_b,clientid,tcpport,tags,conn)
}

/*
//TODO: check buf len on all and prevent read > len(buf)
type OneTag struct {
	Type byte
	NameByte byte
	NameString string
	Value []byte
	ValueLen uint16
}
*/

func sendHello(msgtype byte, uuid_b []byte, clientid uint32, tcpport uint16, tags *[]util.OneTag, conn net.Conn){
	body := make([]byte,0)
	body = append(body,byte(len(uuid_b))) //check , but should be 16 always
	body = append(body,uuid_b...)
	body = append(body,util.UInt32ToByte(clientid)...)
	body = append(body,util.UInt16ToByte(tcpport)...)
	taglen:=len(*tags)
	body = append(body,util.UInt32ToByte(taglen)...)
	for i:=0;i<taglen;i++ {
		var tagbytes_b []byte
		var tagname_b []byte
		if util.OneTag.NameByte != 0 {
			tagname_b = util.EncodeByteTagNameInt(0x11)
		} else if util.OneTag.NameString != "" {
			tagname_b = util.EncodeByteTagNameStr(util.OneTag.NameString)
		} else {
			panic("no tag name")
		}
		if util.OneTag.Type == 3 {
			strvalbuf := make([]byte,0) //strvalbuf := make([]byte,util.OneTag.ValueLen+2)
			strvalbuf = append(strvalbuf,util.UInt16ToByte(util.OneTag.ValueLen)...)
			strvalbuf = append(strvalbuf,util.OneTag.Value...) 
			tagbytes_b = util.EncodeByteTag(util.OneTag.Type,tagname_b,strvalbuf)
		} else {
			tagbytes_b = util.EncodeByteTag(util.OneTag.Type,tagname_b,util.OneTag.Value)
		}
		body = append(body,tagbytes_b...)
	}
	//data := util.EncodeByteMsg(0xE3,0x01,body)
	data := util.EncodeByteMsg(0xE3,msgtype,body)
	//conn.Write(data)
	fmt.Println("DEBUG HELLO:", data) //test
}
