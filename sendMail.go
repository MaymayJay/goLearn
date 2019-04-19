package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"time"
)

const (
	mailServerIp 	= "smtp server ip"
	mailServerUser  = "user name"
	mailServerPwd   = "password"
)



func main()  {

	addr := "xxxxxxx:25"
	from := "xxxxxx"
	to   := []string{"xxxxxxx"}
	sendMail(addr,from,to,"html","这里是主题","这里是内容")
}

func sendMail(addr string,from string,to []string,contentType string,subject string,body string)  {

	if contentType == "html" {
		contentType = " text/html; charset=UTF-8"
	}else {
		contentType = " text/plain; charset=UTF-8"
	}
	buffer := bytes.NewBuffer(nil)
	header := fmt.Sprintf("To:%s\r\n"+
		"From:%s\r\n"+
		"Subject:%s\r\n"+
		"Content-Type:"+ contentType + "\r\n"+
		"Date:%s\r\n\r\n"+
		"%s", to[0],from, subject,time.Now().String(),body)

	fmt.Println(header)
	buffer.WriteString(header)

	auth := smtp.PlainAuth("",mailServerUser,mailServerPwd,mailServerIp)
	err := smtp.SendMail(addr,auth,from,to,buffer.Bytes())

	if err != nil {
		fmt.Println("发送失败",err)
	}else {
		fmt.Println("发送成功")
	}
}

