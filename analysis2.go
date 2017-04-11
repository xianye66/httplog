package main

import (
	"encoding/csv"
	"os"
	"io"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"strings"
	"time"
)


func main(){
	file,err := os.Open("logdata/access_log2")
	mysqlutil.GlobalMysqlUtil = mysqlutil.MysqlUtil{}
	mysqlutil.GlobalMysqlUtil.InitMySqlUtil("115.159.71.248", 3306,
		"remote", "Iknowthat", "webanalytics")
	if err != nil{
		print(err.Error())
	}
	defer file.Close()
	reader :=csv.NewReader(file)
	for   {
		record,err := reader.Read()
		//analysisSingleRecord(record)
		if err == io.EOF{
			break
		}
		if len(record) != 11 {
			println(strings.Join(record," "))
			continue
		}
		println(record[3][1:len(record[3])-1])
		tm2, _ := time.Parse("01/Jan/2006:15:04:05 -0700", record[3][1:len(record[3])-1])
		record[3] = tm2.Format("2006-01-02 15:04:05")
		println(record[3])
		new := make([]interface{}, len(record))
		for i, v := range record {
			new[i] = v
		}
		err = mysqlutil.GlobalMysqlUtil.Insert("INSERT INTO http_log " +
			"(`ip`,`remote_login_name`,`remote_login_user`,`request_date`,`domain`,`path`,`request_first_line`,`status`,`byte_size`,`refer_url`,`user_agent`)" +
			" VALUES (?,?,?,?,?,?,?,?,?,?,?)",new...)
		if err != nil {
			print(err.Error())
		}

	}

}
