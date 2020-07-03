package parse

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Requestbody struct {
	req string
}

type StringResources struct {
	XMLName      xml.Name     `xml:"resources"`
	Resource0001 Resource0001 `xml:"Resource0001"`
	Resource0002 Resource0002 `xml:"Resource0002"`
	Resource0003 Resource0003 `xml:"Resource0003"`
	Resource0004 Resource0004 `xml:"Resource0004"`
	Resource0005 Resource0005 `xml:"Resource0005"`
	Resource0006 Resource0006 `xml:"Resource0006"`
	Resource0007 Resource0007 `xml:"Resource0007"`
	Resource0008 Resource0008 `xml:"Resource0008"`
	Resource0009 Resource0009 `xml:"Resource0009"`
	Resource1014 Resource1014 `xml:"Resource1014"`
}
type Resource0001 struct {
	XMLName        xml.Name         `xml:"Resource0001"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0002 struct {
	XMLName        xml.Name         `xml:"Resource0002"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0003 struct {
	XMLName        xml.Name         `xml:"Resource0003"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0004 struct {
	XMLName        xml.Name         `xml:"Resource0004"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0005 struct {
	XMLName        xml.Name         `xml:"Resource0005"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0006 struct {
	XMLName        xml.Name         `xml:"Resource0006"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0007 struct {
	XMLName        xml.Name         `xml:"Resource0007"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0008 struct {
	XMLName        xml.Name         `xml:"Resource0008"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource0009 struct {
	XMLName        xml.Name         `xml:"Resource0009"`
	ResourceString []ResourceString `xml:"string"`
}

type Resource1014 struct {
	XMLName        xml.Name         `xml:"Resource1014"`
	ResourceString []ResourceString `xml:"string"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}

/**
解析联盟积分请求报文
*/
func UnpackLeagueMsg(str string, msgId string) (map[string]string, error) {
	log.Println("==UnpackLeagueMsg msg==" + str)

	//定义返回值Map
	var offset int = 0
	var ResourceString []ResourceString
	resultMap := make(map[string]string)
	content, err := ioutil.ReadFile(GetPath() + "/" + "reqConfiguration.xml")
	if err != nil {
		return nil, err
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(msgId, "0001") {
		ResourceString = result.Resource0001.ResourceString
	} else if strings.EqualFold(msgId, "0002") {
		ResourceString = result.Resource0002.ResourceString
	} else if strings.EqualFold(msgId, "0003") {
		ResourceString = result.Resource0003.ResourceString
	} else if strings.EqualFold(msgId, "0004") {
		ResourceString = result.Resource0004.ResourceString
	} else if strings.EqualFold(msgId, "0005") {
		ResourceString = result.Resource0005.ResourceString
	} else if strings.EqualFold(msgId, "0006") {
		ResourceString = result.Resource0006.ResourceString
	} else if strings.EqualFold(msgId, "0007") {
		ResourceString = result.Resource0007.ResourceString
	} else if strings.EqualFold(msgId, "0008") {
		ResourceString = result.Resource0008.ResourceString
	} else if strings.EqualFold(msgId, "0009") {
		ResourceString = result.Resource0009.ResourceString
	}
	log.Println(ResourceString)

	for _, o := range ResourceString {
		//		log.Println(ResourceString)
		//		log.Println(o.StringName + "===" + o.InnerText)
		len, err1 := strconv.Atoi(o.InnerText)
		if err1 != nil {
			return nil, err
		}
		resultMap[o.StringName] = strings.TrimSpace(Substr(str, offset, len))
		//		log.Println("len=%v, offset=%v\n", len, offset)
		//		log.Println("substr=%v\n", resultMap[o.StringName])
		offset = offset + len
	}
	log.Println("==UnpackLeagueMsg RESULTMAP==")

	log.Println(resultMap)

	return resultMap, err

}

/**
生成联盟积分定长返回报文
*/
func PackLeagueMsg(inputMap map[string]string) (string, error) {
	log.Println("==PackLeagueMsg inputMap==")
	log.Println(inputMap)
	//拼接用
	var buf bytes.Buffer
	//获取交易码
	msgId := inputMap["MsgID"]
	var ResourceString []ResourceString
	content, err := ioutil.ReadFile(GetPath() + "/" + "resConfiguration.xml")
	if err != nil {
		return "", err
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}
	if strings.EqualFold(msgId, "0001") {
		ResourceString = result.Resource0001.ResourceString
	} else if strings.EqualFold(msgId, "0002") {
		ResourceString = result.Resource0002.ResourceString
	} else if strings.EqualFold(msgId, "0003") {
		ResourceString = result.Resource0003.ResourceString
	} else if strings.EqualFold(msgId, "0004") {
		ResourceString = result.Resource0004.ResourceString
	} else if strings.EqualFold(msgId, "0005") {
		ResourceString = result.Resource0005.ResourceString
	} else if strings.EqualFold(msgId, "0006") {
		ResourceString = result.Resource0006.ResourceString
	} else if strings.EqualFold(msgId, "0007") {
		ResourceString = result.Resource0007.ResourceString
	} else if strings.EqualFold(msgId, "0008") {
		ResourceString = result.Resource0008.ResourceString
	} else if strings.EqualFold(msgId, "0009") {
		ResourceString = result.Resource0009.ResourceString
	}
	var msgcont bytes.Buffer
	for _, o := range ResourceString {
		//		log.Println(ResourceString)
		//		log.Println(o.StringName + "===" + o.InnerText) //定长报文是否要补齐
		temStr := inputMap[o.StringName]
		//		log.Println("---------temStr-------", temStr)

		innerText := o.InnerText
		//		log.Println("---------o.InnerText-------", innerText)
		textLen, ok := strconv.Atoi(innerText)
		//		log.Println("---------textLen-------", textLen)
		if ok == nil {
			for i := len(temStr); i < textLen; i++ {
				temStr = temStr + " "
				//				log.Println("temStr:=%v", temStr)

			}
			msgcont.WriteString(temStr)
			log.Println("==msgcont==%v", msgcont.String())
		}

	}
	strLen := strconv.Itoa(len(msgcont.String()))
	buf.WriteString(Substr("00000000", 0, 8-len(strLen)))
	buf.WriteString(strLen)
	buf.WriteString(msgcont.String())
	log.Println("==PackLeagueMsg msg==", buf.String())
	return buf.String(), err

}

/**

字符串子串截取
*/
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func GetPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}
