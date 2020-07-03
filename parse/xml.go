package parse

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type StringMsgReq0001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0001 `xml:"MsgCont"`
}

type StringMsgReq0002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0002 `xml:"MsgCont"`
}

type StringMsgReq0003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0003 `xml:"MsgCont"`
}

type StringMsgReq0004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0004 `xml:"MsgCont"`
}

type StringMsgReq0005 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0005 `xml:"MsgCont"`
}

type StringMsgReq0006 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0006 `xml:"MsgCont"`
}

type StringMsgReq0007 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0007 `xml:"MsgCont"`
}

type StringMsgReq0008 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0008 `xml:"MsgCont"`
}

type StringMsgReq0009 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0009 `xml:"MsgCont"`
}

type StringMsgRsp0001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0001 `xml:"MsgCont"`
}

type StringMsgRsp0002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0002 `xml:"MsgCont"`
}

type StringMsgRsp0003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0003 `xml:"MsgCont"`
}

type StringMsgRsp0004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0004 `xml:"MsgCont"`
}

type StringMsgRsp0005 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0005 `xml:"MsgCont"`
}

type StringMsgRsp0006 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0006 `xml:"MsgCont"`
}

type StringMsgRsp0007 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0007 `xml:"MsgCont"`
}

type StringMsgRsp0008 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0008 `xml:"MsgCont"`
}

type StringMsgRsp0009 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0009 `xml:"MsgCont"`
}

type StringMsgReq1001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContReq0001 `xml:"MsgCont"`
}

type StringMsgReq1003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContReq0003 `xml:"MsgCont"`
}

type StringMsgReq1004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContReq1004 `xml:"MsgCont"`
}

type StringMsgReq1007 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContReq0007 `xml:"MsgCont"`
}
type StringMsgReq1014 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContReq1014 `xml:"MsgCont"`
}

type StringMsgRsp1001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1001 `xml:"MsgCont"`
}

type StringMsgRsp1003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1003 `xml:"MsgCont"`
}

type StringMsgRsp1004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1004 `xml:"MsgCont"`
}

type StringMsgRsp1007 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1007 `xml:"MsgCont"`
}

type StringMsgRsp1014 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHeadBank    `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1014 `xml:"MsgCont"`
}

type MsgHeadBank struct {
	XMLName  xml.Name `xml:"MsgHead"`
	MsgVer   MsgVer   `xml:"MsgVer"`
	MsgID    MsgID    `xml:"MsgID"`
	MsgType  MsgType  `xml:"MsgType"`
	ProcID   ProcID   `xml:"ProcID"`
	ProcTime ProcTime `xml:"ProcTime"`
	Sender   Sender   `xml:"Sender"`
	Receiver Receiver `xml:"Receiver"`
}

type MsgHead struct {
	XMLName    xml.Name   `xml:"MsgHead"`
	MsgVer     MsgVer     `xml:"MsgVer"`
	MsgID      MsgID      `xml:"MsgID"`
	MsgType    MsgType    `xml:"MsgType"`
	ProcID     ProcID     `xml:"ProcID"`
	ProcTime   ProcTime   `xml:"ProcTime"`
	Sender     Sender     `xml:"Sender"`
	SenderID   SenderID   `xml:"SenderID"`
	Receiver   Receiver   `xml:"Receiver"`
	ReceiverID ReceiverID `xml:"ReceiverID"`
}

type MsgContReq0001 struct {
	XMLName xml.Name `xml:"MsgCont"`
	CardNo  CardNo   `xml:"CardNo"`
	BankNo  BankNo   `xml:"BankNo"`
}

type MsgContReq0002 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	UserName UserName `xml:"UserName"`
}

type MsgContReq0003 struct {
	XMLName xml.Name `xml:"MsgCont"`
	CardNo  CardNo   `xml:"CardNo"`
	BankNo  BankNo   `xml:"BankNo"`
}

type MsgContReq0004 struct {
	XMLName        xml.Name       `xml:"MsgCont"`
	UserName       UserName       `xml:"UserName"`
	LeagueUserName LeagueUserName `xml:"LeagueUserName"`
	CardNo         CardNo         `xml:"CardNo"`
	BankNo         BankNo         `xml:"BankNo"`
	BankPoint      BankPoint      `xml:"BankPoint"`
	Rate           Rate           `xml:"Rate"`
	BlockPoint     BlockPoint     `xml:"BlockPoint"`
}

type MsgContReq1004 struct {
	XMLName    xml.Name   `xml:"MsgCont"`
	UserName   UserName   `xml:"UserName"`
	CardNo     CardNo     `xml:"CardNo"`
	BankNo     BankNo     `xml:"BankNo"`
	BankPoint  BankPoint  `xml:"BankPoint"`
	Rate       Rate       `xml:"Rate"`
	BlockPoint BlockPoint `xml:"BlockPoint"`
}

type MsgContReq0005 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	UserName UserName `xml:"UserName"`
}

type MsgContReq0006 struct {
	XMLName    xml.Name   `xml:"MsgCont"`
	UserName   UserName   `xml:"UserName"`
	BlockPoint BlockPoint `xml:"BlockPoint"`
	GoodsId    GoodsId    `xml:"GoodsId"`
	MerchantId MerchantId `xml:"MerchantId"`
}

type MsgContReq0007 struct {
	XMLName    xml.Name   `xml:"MsgCont"`
	CardNo     CardNo     `xml:"CardNo"`
	BankNo     BankNo     `xml:"BankNo"`
	CashAmount CashAmount `xml:"CashAmount"`
}

type MsgContReq0008 struct {
	XMLName      xml.Name     `xml:"MsgCont"`
	UserName     UserName     `xml:"UserName"`
	ThemUserName ThemUserName `xml:"ThemUserName"`
	BlockPoint   BlockPoint   `xml:"BlockPoint"`
}

type MsgContReq0009 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	UserName  UserName  `xml:"UserName"`
	OriProcID OriProcID `xml:"OriProcID"`
}

type MsgContRsp0001 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	Response  Response  `xml:"Response"`
	ResMsg    ResMsg    `xml:"ResMsg"`
	BankPoint BankPoint `xml:"BankPoint"`
}

type MsgContRsp0002 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	Response Response `xml:"Response"`
	ResMsg   ResMsg   `xml:"ResMsg"`
}

type MsgContRsp0003 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	Response  Response  `xml:"Response"`
	ResMsg    ResMsg    `xml:"ResMsg"`
	BankPoint BankPoint `xml:"BankPoint"`
}

type MsgContRsp0004 struct {
	XMLName    xml.Name   `xml:"MsgCont"`
	Response   Response   `xml:"Response"`
	ResMsg     ResMsg     `xml:"ResMsg"`
	BlockPoint BlockPoint `xml:"BlockPoint"`
}

type MsgContRsp0005 struct {
	XMLName    xml.Name   `xml:"MsgCont"`
	Response   Response   `xml:"Response"`
	ResMsg     ResMsg     `xml:"ResMsg"`
	BlockPoint BlockPoint `xml:"BlockPoint"`
}

type MsgContRsp0006 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	Response Response `xml:"Response"`
	ResMsg   ResMsg   `xml:"ResMsg"`
}

type MsgContRsp0007 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	Response Response `xml:"Response"`
	ResMsg   ResMsg   `xml:"ResMsg"`
}

type MsgContRsp0008 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	Response Response `xml:"Response"`
	ResMsg   ResMsg   `xml:"ResMsg"`
}

type MsgContRsp0009 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	Response  Response  `xml:"Response"`
	ResMsg    ResMsg    `xml:"ResMsg"`
	OriProcID OriProcID `xml:"OriProcID"`
}

type MsgContReq1014 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	CardNo    CardNo    `xml:"CardNo"`
	BankPoint BankPoint `xml:"BankPoint"`
}

type MsgContRsp1001 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	Response  Response  `xml:"Response"`
	ResMsg    ResMsg    `xml:"ResMsg"`
	BankPoint BankPoint `xml:"BankPoint"`
}

type MsgContRsp1003 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	Response  Response  `xml:"Response"`
	ResMsg    ResMsg    `xml:"ResMsg"`
	BankPoint BankPoint `xml:"BankPoint"`
}

type MsgContRsp1004 struct {
	XMLName    xml.Name   `xml:"MsgCont"`
	Response   Response   `xml:"Response"`
	ResMsg     ResMsg     `xml:"ResMsg"`
	BlockPoint BlockPoint `xml:"BlockPoint"`
}

type MsgContRsp1007 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	Response Response `xml:"Response"`
	ResMsg   ResMsg   `xml:"ResMsg"`
}

type MsgContRsp1014 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	Response Response `xml:"Response"`
	ResMsg   ResMsg   `xml:"ResMsg"`
}

type MsgVer struct {
	XMLName   xml.Name `xml:"MsgVer"`
	InnerText string   `xml:",innerxml"`
}

type MsgID struct {
	XMLName   xml.Name `xml:"MsgID"`
	InnerText string   `xml:",innerxml"`
}

type MsgType struct {
	XMLName   xml.Name `xml:"MsgType"`
	InnerText string   `xml:",innerxml"`
}

type ProcID struct {
	XMLName   xml.Name `xml:"ProcID"`
	InnerText string   `xml:",innerxml"`
}

type ProcTime struct {
	XMLName   xml.Name `xml:"ProcTime"`
	InnerText string   `xml:",innerxml"`
}

type Sender struct {
	XMLName   xml.Name `xml:"Sender"`
	InnerText string   `xml:",innerxml"`
}

type SenderID struct {
	XMLName   xml.Name `xml:"SenderID"`
	InnerText string   `xml:",innerxml"`
}

type Receiver struct {
	XMLName   xml.Name `xml:"Receiver"`
	InnerText string   `xml:",innerxml"`
}

type ReceiverID struct {
	XMLName   xml.Name `xml:"ReceiverID"`
	InnerText string   `xml:",innerxml"`
}

type CardNo struct {
	XMLName   xml.Name `xml:"CardNo"`
	InnerText string   `xml:",innerxml"`
}

type BankNo struct {
	XMLName   xml.Name `xml:"BankNo"`
	InnerText string   `xml:",innerxml"`
}

type UserName struct {
	XMLName   xml.Name `xml:"UserName"`
	InnerText string   `xml:",innerxml"`
}

type BankPoint struct {
	XMLName   xml.Name `xml:"BankPoint"`
	InnerText string   `xml:",innerxml"`
}

type BlockPoint struct {
	XMLName   xml.Name `xml:"BlockPoint"`
	InnerText string   `xml:",innerxml"`
}

type Rate struct {
	XMLName   xml.Name `xml:"Rate"`
	InnerText string   `xml:",innerxml"`
}

type GoodsId struct {
	XMLName   xml.Name `xml:"GoodsId"`
	InnerText string   `xml:",innerxml"`
}

type MerchantId struct {
	XMLName   xml.Name `xml:"MerchantId"`
	InnerText string   `xml:",innerxml"`
}

type CashAmount struct {
	XMLName   xml.Name `xml:"CashAmount"`
	InnerText string   `xml:",innerxml"`
}

type ThemUserName struct {
	XMLName   xml.Name `xml:"ThemUserName"`
	InnerText string   `xml:",innerxml"`
}

type OriProcID struct {
	XMLName   xml.Name `xml:"OriProcID"`
	InnerText string   `xml:",innerxml"`
}

type Response struct {
	XMLName   xml.Name `xml:"Response"`
	InnerText string   `xml:",innerxml"`
}

type ResMsg struct {
	XMLName   xml.Name `xml:"ResMsg"`
	InnerText string   `xml:",innerxml"`
}

type LeagueUserName struct {
	XMLName   xml.Name `xml:"LeagueUserName"`
	InnerText string   `xml:",innerxml"`
}

// 以下是生成内部格式xml“请求”报文对应的方法
func WriteXMLReq0001(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0001.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0001
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0002(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0002.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0002
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.UserName.InnerText = strings.TrimSpace(inputMap["UserName"])

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0003(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0003.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0003
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0004(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0004.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0004
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.UserName.InnerText = inputMap["UserName"]
	result.MsgMsgCont.LeagueUserName.InnerText = inputMap["LeagueUserName"]
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]
	result.MsgMsgCont.BankPoint.InnerText = inputMap["BankPoint"]
	result.MsgMsgCont.Rate.InnerText = inputMap["Rate"]
	result.MsgMsgCont.BlockPoint.InnerText = inputMap["BlockPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0005(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0005.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0005
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.UserName.InnerText = strings.TrimSpace(inputMap["UserName"])

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0006(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0006.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0006
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.UserName.InnerText = inputMap["UserName"]
	result.MsgMsgCont.BlockPoint.InnerText = inputMap["BlockPoint"]
	result.MsgMsgCont.GoodsId.InnerText = inputMap["GoodsId"]
	result.MsgMsgCont.MerchantId.InnerText = inputMap["MerchantId"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0007(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0007.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0007
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]
	result.MsgMsgCont.CashAmount.InnerText = inputMap["CashAmount"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0008(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0008.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0008
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.UserName.InnerText = inputMap["UserName"]
	result.MsgMsgCont.ThemUserName.InnerText = inputMap["ThemUserName"]
	result.MsgMsgCont.BlockPoint.InnerText = inputMap["BlockPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLReq0009(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Req0009.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq0009
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.UserName.InnerText = inputMap["UserName"]
	result.MsgMsgCont.OriProcID.InnerText = inputMap["OriProcID"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0001(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0001.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0001
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]
	result.MsgMsgCont.BankPoint.InnerText = inputMap["BankPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0002(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0002.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0002
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0003(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0003.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0003
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]
	result.MsgMsgCont.BankPoint.InnerText = inputMap["BankPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0004(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0004.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0004
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]

	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]
	result.MsgMsgCont.BlockPoint.InnerText = inputMap["BlockPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0005(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0005.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0005
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]
	result.MsgMsgCont.BlockPoint.InnerText = inputMap["BlockPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0006(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0006.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0006
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")

	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0007(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0007.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0007
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0008(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0008.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0008
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]

	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXMLRsp0009(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	var err error
	content, err := ioutil.ReadFile(GetPath() + "/" + "Rsp0009.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgRsp0009
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"]
	result.MsgMsgCont.Response.InnerText = inputMap["Response"]
	result.MsgMsgCont.ResMsg.InnerText = inputMap["ResMsg"]
	result.MsgMsgHead.SenderID.InnerText = inputMap["SenderID"]
	result.MsgMsgHead.ReceiverID.InnerText = inputMap["ReceiverID"]
	result.MsgMsgCont.OriProcID.InnerText = inputMap["OriProcID"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXMLReq0001(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0001
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["BankNo"] = result.MsgMsgCont.BankNo.InnerText
	resultMap["CardNo"] = result.MsgMsgCont.CardNo.InnerText

	return resultMap, err
}

func ReadXMLReq0002(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0002
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["UserName"] = result.MsgMsgCont.UserName.InnerText

	return resultMap, err
}

func ReadXMLReq0003(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0003
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["BankNo"] = result.MsgMsgCont.BankNo.InnerText
	resultMap["CardNo"] = result.MsgMsgCont.CardNo.InnerText

	return resultMap, err
}

func ReadXMLReq0004(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0004
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["UserName"] = result.MsgMsgCont.UserName.InnerText
	resultMap["LeagueUserName"] = result.MsgMsgCont.LeagueUserName.InnerText
	resultMap["BankNo"] = result.MsgMsgCont.BankNo.InnerText
	resultMap["CardNo"] = result.MsgMsgCont.CardNo.InnerText
	resultMap["BankPoint"] = result.MsgMsgCont.BankPoint.InnerText
	resultMap["Rate"] = result.MsgMsgCont.Rate.InnerText
	resultMap["BlockPoint"] = result.MsgMsgCont.BlockPoint.InnerText

	return resultMap, err
}

func ReadXMLReq0005(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0005
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["UserName"] = result.MsgMsgCont.UserName.InnerText

	return resultMap, err
}

func ReadXMLReq0006(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0006
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["UserName"] = result.MsgMsgCont.UserName.InnerText
	resultMap["BlockPoint"] = result.MsgMsgCont.BlockPoint.InnerText
	resultMap["GoodsId"] = result.MsgMsgCont.GoodsId.InnerText
	resultMap["MerchantId"] = result.MsgMsgCont.MerchantId.InnerText

	return resultMap, err
}

func ReadXMLReq0007(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0007
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["BankNo"] = result.MsgMsgCont.BankNo.InnerText
	resultMap["CardNo"] = result.MsgMsgCont.CardNo.InnerText
	resultMap["CashAmount"] = result.MsgMsgCont.CashAmount.InnerText

	return resultMap, err
}

func ReadXMLReq0008(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0008
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["UserName"] = result.MsgMsgCont.UserName.InnerText
	resultMap["ThemUserName"] = result.MsgMsgCont.ThemUserName.InnerText
	resultMap["BlockPoint"] = result.MsgMsgCont.BlockPoint.InnerText

	return resultMap, err
}

func ReadXMLReq0009(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgReq0009
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["UserName"] = result.MsgMsgCont.UserName.InnerText
	resultMap["OriProcID"] = result.MsgMsgCont.OriProcID.InnerText

	return resultMap, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXMLRsp0001(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0001
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BankPoint"] = result.MsgMsgCont.BankPoint.InnerText

	return resultMap, err
}

func ReadXMLRsp0002(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0002
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText

	return resultMap, err
}

func ReadXMLRsp0003(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0003
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BankPoint"] = result.MsgMsgCont.BankPoint.InnerText

	return resultMap, err
}

func ReadXMLRsp0004(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0004
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BlockPoint"] = result.MsgMsgCont.BlockPoint.InnerText

	return resultMap, err
}

func ReadXMLRsp0005(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0005
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BlockPoint"] = result.MsgMsgCont.BlockPoint.InnerText

	return resultMap, err
}

func ReadXMLRsp0006(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0006
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText

	return resultMap, err
}

func ReadXMLRsp0007(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0007
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText

	return resultMap, err
}

func ReadXMLRsp0008(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0008
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText

	return resultMap, err
}

func ReadXMLRsp0009(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp0009
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["SenderID"] = result.MsgMsgHead.SenderID.InnerText
	resultMap["ReceiverID"] = result.MsgMsgHead.ReceiverID.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["OriProcID"] = result.MsgMsgCont.OriProcID.InnerText

	return resultMap, err
}

// 以下是发卡行系统报文生成方法
func WriteXML1001(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string
	content, err := ioutil.ReadFile(GetPath() + "/" + "1001.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq1001
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = "A002"
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXML1003(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string

	content, err := ioutil.ReadFile(GetPath() + "/" + "1003.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq1003
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = "A002"
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXML1004(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string

	content, err := ioutil.ReadFile(GetPath() + "/" + "1004.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq1004
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = "A002"
	result.MsgMsgCont.UserName.InnerText = inputMap["UserName"]
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]
	result.MsgMsgCont.BankPoint.InnerText = inputMap["BankPoint"]
	result.MsgMsgCont.Rate.InnerText = inputMap["Rate"]
	result.MsgMsgCont.BlockPoint.InnerText = inputMap["BlockPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXML1007(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string

	content, err := ioutil.ReadFile(GetPath() + "/" + "1007.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq1007
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = ""
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankNo.InnerText = inputMap["BankNo"]
	result.MsgMsgCont.CashAmount.InnerText = inputMap["CashAmount"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

func WriteXML1014(inputMap map[string]string) (string, error) {
	var buf bytes.Buffer
	var msg string

	content, err := ioutil.ReadFile(GetPath() + "/" + "1014.xml")
	if err != nil {
		return "", err
	}
	var result StringMsgReq1014
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	result.MsgMsgHead.MsgVer.InnerText = "1"
	result.MsgMsgHead.MsgID.InnerText = inputMap["MsgID"]
	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"]
	result.MsgMsgHead.ProcID.InnerText = inputMap["ProcID"]
	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"]
	result.MsgMsgHead.Receiver.InnerText = ""
	result.MsgMsgCont.CardNo.InnerText = inputMap["CardNo"]
	result.MsgMsgCont.BankPoint.InnerText = inputMap["BankPoint"]

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		fmt.Println("len:=%v", strconv.Itoa(len(xmlOutPutData)))
		strLen := strconv.Itoa(len(xmlOutPutData))
		buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
		buf.WriteString(strLen)
		//buf.WriteString(inputMap["MsgID"])

		xmlOutPutDataNew := append([]byte(buf.String()), xmlOutPutData...)
		msg = string(xmlOutPutDataNew)
		//写入文件
		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
		fmt.Println("OK~")
		//		return string(xmlOutPutDataNew)
	} else {
		fmt.Println(outPutErr)
		//		return ""
	}
	return msg, err
}

// 以下是解析发卡行系统请求报文方法
func ReadXML1001(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp1001
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BankPoint"] = result.MsgMsgCont.BankPoint.InnerText
	log.Println("==1001==", result)
	return resultMap, err
}

func ReadXML1003(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp1003
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BankPoint"] = result.MsgMsgCont.BankPoint.InnerText

	return resultMap, err
}

func ReadXML1004(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp1004
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText
	resultMap["BlockPoint"] = result.MsgMsgCont.BlockPoint.InnerText

	return resultMap, err
}

func ReadXML1007(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp1007
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText

	return resultMap, err
}

func ReadXML1014(msg string) (map[string]string, error) {
	var resultMap map[string]string
	resultMap = make(map[string]string)
	var result StringMsgRsp1014
	var content []byte = []byte(Substr(msg, 8, len(msg)-8))
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(result)
	log.Println(result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgID"] = result.MsgMsgHead.MsgID.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcID"] = result.MsgMsgHead.ProcID.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["Response"] = result.MsgMsgCont.Response.InnerText
	resultMap["ResMsg"] = result.MsgMsgCont.ResMsg.InnerText

	return resultMap, err
}
