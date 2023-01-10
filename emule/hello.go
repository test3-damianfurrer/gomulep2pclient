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
