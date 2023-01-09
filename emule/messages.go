package emule

import (
	"fmt"
	util "github.com/AltTechTools/gomule-tst/emule"
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
