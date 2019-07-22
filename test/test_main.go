package main

import (
	"fmt"
	"io/ioutil"
	"klook.libs/utils"
	"os"
	"strings"
	"time"
)

func testSplit() {
	bufStr := "1. 国家药监局决定从严查处吉林长春长生疫苗案件，对长春长生所有疫苗生产、销售 全流程、全链条进行彻查，对全国疫苗生产企业全面开展（ ），研究完善我国疫苗管理 体制。 A. 定期检查 B. 预约检查 C. 通知检查 D. 飞行检查 【答案】D 【解析】7 月 23 日下午，国家药监局召开党组扩大会议，传达学习习近平总书记对吉 林长春长生疫苗案件重要指示精神，研究贯彻落实措施。会议决定，一是在前期工作基础上， 进一步增加人员，充实案件查处工作领导小组力量，全力配合国务院调查组工作。二是对长 春长生所有疫苗生产、销售全流程、全链条进行彻查，尽快查清事实真相，锁定证据线索。 三是坚持重拳出击，对不法分子严惩不贷、以儆效尤；对失职渎职的，从严处理、严肃问责。 四是针对人民群众关切的热点问题，做好解疑释惑工作。五是举一反三，对全国疫苗生产企 业全面开展飞行检查，严查严控风险隐患。六是对疫苗全生命周期监管制度进行系统分析， 逐一解剖问题症结，研究完善我国疫苗管理体制。 故本题正确选项为 D 项。"
	fmt.Println(bufStr[4:6])

}
func GetNextDayBegin() int64 {
	t := time.Now().AddDate(0, 0, 1)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tm1.Unix()
}
func testInt() {
	a := 100000
	fmt.Println(string(a))
}
func testTransInt() {
	/*a := 10
	str := strconv.Itoa(a)*/
}
func testMapRem() {
	allTopicsMap := make(map[int]struct{})
	allTopicsMap[100] = struct{}{}
	//delete(allTopicsMap, 100)
	fmt.Println(len(allTopicsMap))

}
func testTime() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))
}

//xxx
func testMap() {
	var dat = map[string]interface{}{
		"data": "hongyuqin",
	}

	fmt.Println(dat)
}

//把interface 序列化为json字节
func testJson() {
	//b, err := json.Marshal(u)
	//json.Marshal()
}
func testCh() {
	done := make(chan bool)

	for i := 0; i < 4; i++ {
		go func() {
			time.Sleep(time.Second)
			go func() {
				time.Sleep(time.Second)
				done <- true
			}()
			done <- true
		}()
	}
	for i := 0; i < 8; i++ {
		<-done
	}
}
func GetFileContentAsStringLines(filePath string) ([]string, error) {
	fmt.Printf("get file content as lines: %v", filePath)
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file: %v error: %v", filePath, err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	fmt.Printf("get file content as lines: %v, size: %v", filePath, len(result))
	return result, nil
}
func genInsertMqSql() {
	dat, err := ioutil.ReadFile("./file/test.txt")
	if err != nil {
		fmt.Printf("gen sql error")
		return
	}
	//fmt.Print(string(dat))
	arr := strings.Split(string(dat), "\n")
	fmt.Println("arr len is :", len(arr))
	//0.创建文件 可以写入
	f, err := os.Create("./file/dbsql.txt")
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer f.Close()
	//1.以换行分割
	for _, str := range arr {
		//2.0 以 "{ 开头才处理
		if strings.HasPrefix(str, `--`) {
			continue
		}
		insertMqConsumerLogPattern := `insert into erp_kl_mq_consumer_log(unique_reference,type,topic,channel,body,status) values ('%s','%s','%s','%s',%s,2);`
		insertHttpRetryPattern := `insert into erp_http_retry (retry_type,max_retry,current_retry,retry_status,source_platform,unique_reference) values ('%s',%d,%d,'%s','%s','%s');`
		insertUniqueReferencePattern := `insert into erp_unique_reference (unique_reference) values ('%s');`
		if !strings.HasPrefix(str, `topic`) {
			//uuidString := utils.GetUUIDString()
			//分成 unique_reference,type,topic,channel,`body`
			firstEnd := strings.Index(str, `, topic`)
			secondBegin := strings.Index(str, `, topic`)
			secondEnd := strings.Index(str, `, channel`)
			thirdBegin := strings.Index(str, `, channel`)
			fmt.Printf("json is :%s \n", str[0:firstEnd])
			fmt.Printf(",topic is :%s\n", str[secondBegin+8:secondEnd])
			fmt.Printf(",channel is :%s\n", str[thirdBegin+10:len(str)-1])
			fmt.Printf("=============\n")
			json := str[0:firstEnd]
			topic := str[secondBegin+8 : secondEnd]
			channel := str[thirdBegin+10 : len(str)-1]
			dataType := getDataType(topic)
			//break
			uuid := utils.GetUUIDString()
			//3.组装sql
			//3.1 mqConsumerLog
			insertMqLogSql := fmt.Sprintf(insertMqConsumerLogPattern, uuid, dataType, topic, channel, fmt.Sprintf(`AES_ENCRYPT('%s','nRajDntseUtIIJqs')`, json))
			n3, err := f.WriteString(insertMqLogSql + "\n")
			if err != nil {
				fmt.Println("生成sql失败")
				return
			}
			f.Sync()
			//3.2 httpRetry
			insertRetrySql := fmt.Sprintf(insertHttpRetryPattern, dataType, 10, 0, "Waiting", "klk", uuid)
			n3, err = f.WriteString(insertRetrySql + "\n")
			if err != nil {
				fmt.Println("生成sql失败")
				return
			}
			f.Sync()
			//3.3 erpUniqueReference
			insertUUIDSql := fmt.Sprintf(insertUniqueReferencePattern, uuid)
			n3, err = f.WriteString(insertUUIDSql + "\n")
			if err != nil {
				fmt.Println("生成sql失败")
				return
			}
			f.Sync()
			fmt.Printf("wrote %d bytes\n", n3)
		}
	}

}
func getDataType(topic string) string {
	switch topic {
	case "ActivitySyncERPTopic":
		return "Item"
	case "OrderPayCaptureCompleteTopic":
		return "OrderCreate"
	case "TicketConfirmTopic":
		return "OrderConfirm"
	case "OrderRedeemTopic":
		return "OrderRedeem"
	case "OrderRefundApplyTopic":
		return "OrderRefund"
	case "OrderRefundCompleteTopic":
		return "OrderRefundSuccess"
	}
	return ""
}

//类型推断
func infer() {
	m := make(map[string]string)
	m["name"] = "nick"
	var i interface{} = m
	switch v := i.(type) {
	case map[string]string:
		fmt.Println("map is :", v["name"])
	}
}
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	infer()
	//genInsertMqSql()
	//testMap()
	//testCh()
	//testMap()
	//testTime()
	//bytes := []byte("dfsdfs")
	//fmt.Println("bytes is :",string(bytes))
	//testMapRem()
	//testInt()
	/*log.SetOutput(os.Stdout)
	log.Info("main ")
	xlsx, err := xlsx.OpenFile("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, sheet := range xlsx.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}*/
	//testSplit()
	/*a, err := strconv.Atoi("")
	if err != nil {
		fmt.Println("a is :", a)
	}*/
	//fmt.Println(GetNextDayBegin())
}
