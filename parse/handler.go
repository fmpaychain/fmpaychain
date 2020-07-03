package parse

import (
	"log"
	"strings"
)

//func main() {
//	var reqStr string = "0000001400012016112417033000000000000000000162284816987298900791100001509" //卡验证请求定长报文

//	// 从文件中读取字段信息进行解析'
//	var resultMap map[string]string = UnpackLeague(reqStr)
//	var msg string = PackBusiness(resultMap)
//	log.Println("msg=%v", msg)
//	resultMap = UnpackBusiness(msg)
//	log.Println(resultMap)
//	bankReqMsg := "000003920001<?xml version='1.0' encoding='UTF-8'?><Msg><MsgHead><MsgVer>1</MsgVer><MsgID>0001</MsgID><MsgType>0</MsgType><ProcID>20161124170330000000000000000001</ProcID><ProcTime>Jan _2 15:04:05</ProcTime><Sender>A001</Sender><Receiver></Receiver></MsgHead><MsgCont><CardNo>6228481698729890079</CardNo><BankNo>1100001509</BankNo></MsgCont></Msg>"
//	resultBankMap := UnpackBank(bankReqMsg)
//	log.Println("resultBankMap", resultBankMap)

//}

/**
解析联盟积分请求报文
*/
func UnpackLeague(Msg string) (map[string]string, error) {
	var err error
	//解析交易码
	txcode := Substr(Msg, 8, 4)
	//解析定长报文
	resultMap, err := UnpackLeagueMsg(Msg, txcode)
	//设置Sender=A000
	resultMap["Sender"] = "A000"
	return resultMap, err
}

/**
生成联盟积分商城报文
*/
func PackLeague(PlafromData map[string]string) (string, error) {
	return PackLeagueMsg(PlafromData)

}

/**
解析发卡行系统报文
*/
func UnpackBank(Msg string) (map[string]string, error) {
	msgId := getMsgID(Msg)
	var resultMap map[string]string
	var err error
	// 读写xml文件分支
	if strings.EqualFold(msgId, "1001") {
		resultMap, err = ReadXML1001(Msg)
	} else if strings.EqualFold(msgId, "1003") {
		resultMap, err = ReadXML1003(Msg)
	} else if strings.EqualFold(msgId, "1004") {
		resultMap, err = ReadXML1004(Msg)
	} else if strings.EqualFold(msgId, "1007") {
		resultMap, err = ReadXML1007(Msg)
	} else if strings.EqualFold(msgId, "1014") {
		resultMap, err = ReadXML1014(Msg)
	}
	return resultMap, err
}

/**
生成发卡行系统报文
*/
func PackBank(platData map[string]string, BankNo string) (string, error) {
	var msgId string = platData["MsgID"]
	platData["BankNo"] = BankNo
	var msg string
	var err error
	// 读写xml文件分支
	if strings.EqualFold(msgId, "1001") {
		msg, err = WriteXML1001(platData)
	} else if strings.EqualFold(msgId, "1003") {
		msg, err = WriteXML1003(platData)
	} else if strings.EqualFold(msgId, "1004") {
		msg, err = WriteXML1004(platData)
	} else if strings.EqualFold(msgId, "1007") {
		msg, err = WriteXML1007(platData)
	} else if strings.EqualFold(msgId, "1014") {
		msg, err = WriteXML1014(platData)
	}
	return msg, err
}

/**
解析业务接口适配器请求报文
*/
func UnpackBusiness(Msg string) (map[string]string, error) {

	msgType := getMsgType(Msg)
	log.Println("==UnpackBusiness msgType==%v", msgType)
	msgId := getMsgID(Msg)
	log.Println("==UnpackBusiness msgId==%v", msgId)
	var resultMap map[string]string
	var err error
	// 读写xml文件分支
	if strings.EqualFold(msgType, "0") {
		if strings.EqualFold(msgId, "0001") || strings.EqualFold(msgId, "1001") {
			resultMap, err = ReadXMLReq0001(Msg)
		} else if strings.EqualFold(msgId, "0002") {
			resultMap, err = ReadXMLReq0002(Msg)
		} else if strings.EqualFold(msgId, "0003") || strings.EqualFold(msgId, "1003") {
			resultMap, err = ReadXMLReq0003(Msg)
		} else if strings.EqualFold(msgId, "0004") || strings.EqualFold(msgId, "1004") {
			resultMap, err = ReadXMLReq0004(Msg)
		} else if strings.EqualFold(msgId, "0005") {
			resultMap, err = ReadXMLReq0005(Msg)
		} else if strings.EqualFold(msgId, "0006") {
			resultMap, err = ReadXMLReq0006(Msg)
		} else if strings.EqualFold(msgId, "0007") || strings.EqualFold(msgId, "1007") {
			resultMap, err = ReadXMLReq0007(Msg)
		} else if strings.EqualFold(msgId, "0008") {
			resultMap, err = ReadXMLReq0008(Msg)
		} else if strings.EqualFold(msgId, "0009") {
			resultMap, err = ReadXMLReq0009(Msg)
		}
	} else if strings.EqualFold(msgType, "1") {
		if strings.EqualFold(msgId, "0001") || strings.EqualFold(msgId, "1001") {
			log.Println("==UnpackBusiness msgId==%v", msgId)
			resultMap, err = ReadXMLRsp0001(Msg)
		} else if strings.EqualFold(msgId, "0002") {
			resultMap, err = ReadXMLRsp0002(Msg)
		} else if strings.EqualFold(msgId, "0003") || strings.EqualFold(msgId, "1003") {
			resultMap, err = ReadXMLRsp0003(Msg)
		} else if strings.EqualFold(msgId, "0004") || strings.EqualFold(msgId, "1004") {
			resultMap, err = ReadXMLRsp0004(Msg)
		} else if strings.EqualFold(msgId, "0005") {
			resultMap, err = ReadXMLRsp0005(Msg)
		} else if strings.EqualFold(msgId, "0006") {
			resultMap, err = ReadXMLRsp0006(Msg)
		} else if strings.EqualFold(msgId, "0007") || strings.EqualFold(msgId, "1007") {
			resultMap, err = ReadXMLRsp0007(Msg)
		} else if strings.EqualFold(msgId, "0008") {
			resultMap, err = ReadXMLRsp0008(Msg)
		} else if strings.EqualFold(msgId, "0009") {
			resultMap, err = ReadXMLRsp0009(Msg)
		}
	}
	return resultMap, err
}

/**
生成业务接口适配器请求报文
*/
func PackBusiness(platData map[string]string) (string, error) {
	log.Println("==PackBusiness platData==%v", platData)
	var msgType string = platData["MsgType"]
	var msgId string = platData["MsgID"]
	log.Println("==PackBusiness MsgType==%v", msgType)
	var msg string
	var err error
	// 读写xml文件分支
	if strings.EqualFold(msgType, "0") {
		if strings.EqualFold(msgId, "0001") || strings.EqualFold(msgId, "1001") {
			msg, err = WriteXMLReq0001(platData)
		} else if strings.EqualFold(msgId, "0002") {
			msg, err = WriteXMLReq0002(platData)
		} else if strings.EqualFold(msgId, "0003") || strings.EqualFold(msgId, "1003") {
			msg, err = WriteXMLReq0003(platData)
		} else if strings.EqualFold(msgId, "0004") || strings.EqualFold(msgId, "1004") {
			msg, err = WriteXMLReq0004(platData)
		} else if strings.EqualFold(msgId, "0005") {
			msg, err = WriteXMLReq0005(platData)
		} else if strings.EqualFold(msgId, "0006") {
			msg, err = WriteXMLReq0006(platData)
		} else if strings.EqualFold(msgId, "0007") || strings.EqualFold(msgId, "1007") {
			msg, err = WriteXMLReq0007(platData)
		} else if strings.EqualFold(msgId, "0008") {
			msg, err = WriteXMLReq0008(platData)
		} else if strings.EqualFold(msgId, "0009") {
			msg, err = WriteXMLReq0009(platData)
		}
	} else if strings.EqualFold(msgType, "1") {
		log.Println("==PackBusiness MsgType==%v", msgType)
		if strings.EqualFold(msgId, "0001") || strings.EqualFold(msgId, "1001") {
			msg, err = WriteXMLRsp0001(platData)
		} else if strings.EqualFold(msgId, "0002") {
			msg, err = WriteXMLRsp0002(platData)
		} else if strings.EqualFold(msgId, "0003") || strings.EqualFold(msgId, "1003") {
			msg, err = WriteXMLRsp0003(platData)
		} else if strings.EqualFold(msgId, "0004") || strings.EqualFold(msgId, "1004") {
			msg, err = WriteXMLRsp0004(platData)
		} else if strings.EqualFold(msgId, "0005") {
			msg, err = WriteXMLRsp0005(platData)
		} else if strings.EqualFold(msgId, "0006") {
			msg, err = WriteXMLRsp0006(platData)
		} else if strings.EqualFold(msgId, "0007") || strings.EqualFold(msgId, "1007") {
			msg, err = WriteXMLRsp0007(platData)
		} else if strings.EqualFold(msgId, "0008") {
			msg, err = WriteXMLRsp0008(platData)
		} else if strings.EqualFold(msgId, "0009") {
			msg, err = WriteXMLRsp0009(platData)
		}
	}
	return msg, err
}

func getMsgType(msg string) string {
	begin := strings.Index(msg, "<MsgType>")
	end := strings.Index(msg, "</MsgType>")
	return msg[begin+9 : end]
}
func getMsgID(msg string) string {
	begin := strings.Index(msg, "<MsgID>")
	end := strings.Index(msg, "</MsgID>")
	return msg[begin+7 : end]
}
