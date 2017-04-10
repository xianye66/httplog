package main

import (
	"encoding/csv"
	"os"
	"io"
	"strings"
	"github.com/kevin-zx/go-util/mysqlUtil"
)

type spiderinfo struct {
	spiderName string
	crawlerTimes int
	EndDateTime string
	startDateTime string
	date string

}



type statistics struct {

}

var spiderMap = map[string]*spiderinfo{}

func main(){
	file,err := os.Open("logdata/access_log2")
	mysqlutil.GlobalMysqlUtil = mysqlutil.MysqlUtil{}
	mysqlutil.GlobalMysqlUtil.InitMySqlUtil("localhost", 3306,
		"root", "", "test")
	if err != nil{
		print(err.Error())
	}
	defer file.Close()
	reader :=csv.NewReader(file)
	reader.Comma = ' '
	for   {
		record,err := reader.Read()
		analysisSingleRecord(record)
		if err == io.EOF{
			break
		}
	}
	println(spiderMap["Baiduspider"].spiderName, spiderMap["Baiduspider"].crawlerTimes,spiderMap["Baiduspider"].date,spiderMap["Baiduspider"].startDateTime,spiderMap["Baiduspider"].EndDateTime)
	println(spiderMap["Googlebot"].spiderName, spiderMap["Googlebot"].crawlerTimes,spiderMap["Googlebot"].date,spiderMap["Googlebot"].startDateTime,spiderMap["Googlebot"].EndDateTime)
}

func analysisSingleRecord(record []string)  {

	if len(record) == 12{
		judgeSpider(record[11],strings.Replace(record[3] ,"[","",-1))

	}else {

	}
}

func judgeSpider(userAgent string,date string)  {
	dateString := strings.Split(date,":")[0]
	if strings.Contains(userAgent,"Baiduspider") {
		test("Baiduspider", dateString, date)
	}
	if strings.Contains(userAgent,"Googlebot") {
		test("Googlebot", dateString, date)
	}
	
}

func test(spider string, dateString string, date string)  {
	if spiderMap[spider] == nil {
		spiderMap[spider] = &spiderinfo{spiderName:spider,startDateTime:date, date:dateString, }
	}else{
		if spiderMap[spider].date != dateString {
			println(spiderMap[spider].spiderName, spiderMap[spider].crawlerTimes,spiderMap[spider].date,spiderMap[spider].startDateTime,spiderMap[spider].EndDateTime)
			spiderMap[spider] = &spiderinfo{spiderName:spider,startDateTime:date, date:dateString }
		}else{
			(spiderMap[spider]).EndDateTime = date
			(spiderMap[spider]).crawlerTimes += 1
		}
	}
}