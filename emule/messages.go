package emule

import (
	"fmt"
	util "github.com/AltTechTools/gomule-SockSrvClient/emule"
	"net"
	libdeflate "github.com/4kills/go-libdeflate/v2" //libdeflate.Compressor
)

func handlePeerMsg(protocol byte,buf []byte, pc *PeerClient){
    	//0xd4
	switch protocol {
		case 0xe3:
			decodeE3(buf[0],buf[1:len(buf)],pc)
		case 0xd4:
			decodeD4(buf[0],buf[1:len(buf)],pc.DeComp,pc)
		default:
			fmt.Println("ERROR: only std 0xE3 protocol supported")
	}
}

func decodeD4(btype byte,buf []byte,dc libdeflate.Decompressor, pc *PeerClient){
	fmt.Printf("DEBUG: 0xd4 type 0x%x\n",btype)
	blen, decompressed, err := dc.Decompress(buf, nil, 1)
	if err != nil {
		fmt.Println("ERROR: failed to decompress buffer",err)
		return
	}
	fmt.Println("DEBUG: decompressed length:",blen)
	fmt.Println("DEBUG: decompressed",decompressed[0:30])
	decodeE3(btype,decompressed,pc)
}

func decodeE3(btype byte,buf []byte, pc *PeerClient){
	switch btype {
			/*case 0x38:
				prcServerTextMsg(buf)
			case 0x40:
				prcIdChange(buf,client)
			case 0x34:
				prcServerStatus(buf)
			case 0x32:
				prcServerList(buf)
			case 0x41:
				prcServerIdentification(buf)
			*/
            default:
            	fmt.Printf("ERROR: Msg type 0x%x not supported\n",btype)
        }
}

func p2phello(buf []byte,protocol byte,conn net.Conn,debug bool){
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
	
}
