package main

import (
	"fmt"
)

func testSplit() {
	bufStr := "1. 国家药监局决定从严查处吉林长春长生疫苗案件，对长春长生所有疫苗生产、销售 全流程、全链条进行彻查，对全国疫苗生产企业全面开展（ ），研究完善我国疫苗管理 体制。 A. 定期检查 B. 预约检查 C. 通知检查 D. 飞行检查 【答案】D 【解析】7 月 23 日下午，国家药监局召开党组扩大会议，传达学习习近平总书记对吉 林长春长生疫苗案件重要指示精神，研究贯彻落实措施。会议决定，一是在前期工作基础上， 进一步增加人员，充实案件查处工作领导小组力量，全力配合国务院调查组工作。二是对长 春长生所有疫苗生产、销售全流程、全链条进行彻查，尽快查清事实真相，锁定证据线索。 三是坚持重拳出击，对不法分子严惩不贷、以儆效尤；对失职渎职的，从严处理、严肃问责。 四是针对人民群众关切的热点问题，做好解疑释惑工作。五是举一反三，对全国疫苗生产企 业全面开展飞行检查，严查严控风险隐患。六是对疫苗全生命周期监管制度进行系统分析， 逐一解剖问题症结，研究完善我国疫苗管理体制。 故本题正确选项为 D 项。"
	fmt.Println(bufStr[4:6])

}
func main() {
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
	testSplit()
}
