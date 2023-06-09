package iec

// APDU 应用协议数据单元
type Apdu struct {
	apci APCI
	asdu ASDU
}
