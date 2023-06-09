package iec

const (
	// StartFrame  启动帧
	StartFrame byte  = 0x68
	TypeIFrame uint8 = 1
	TypeSFrame uint8 = 2
	TypeUFrame uint8 = 3
	// TestFRAct 测试激活 一般由子站发送
	TestFRAct byte = 0x43
	// TestFRCon 测试确认
	TestFRCon byte = 0x83
	// StopDTAct 停止激活
	StopDTAct byte = 0x13
	// StopDTCon 停止确认
	StopDTCon byte = 0x23
	// StartDTAct 启动激活
	StartDTAct byte = 0x07
	// StartDTCon 启动确定
	StartDTCon byte = 0x0b
)

// APCI APCI应用协议控制单元
type APCI struct {
	start                  byte
	ctr1, ctr2, ctr3, ctr4 byte
	//apduc长度
	apduLen byte
	frame   FRAME
}

type FRAME struct {
	FrameType uint8
}

// IFRAME I帧
type IFRAME struct {
	FRAME
	//发送序号
	TXCounter uint16
	//接收序号
	RXCounter uint16
}

// SFRAME S帧
type SFRAME struct {
	FRAME
	//接收序号
	RXCounter uint16
}

// UFRAME U帧
type UFRAME struct {
	FRAME
	// IType 指令类型
	IType byte
}
