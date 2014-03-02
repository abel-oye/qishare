package controllers

import (
	"bytes"
	"fmt"
	"github.com/robfig/config"
	"github.com/robfig/revel"
	"github.com/robfig/revel/mail"
)

/*
* 公共方法
 */

//发送邮件
func sendMail(subject, content string, tos []string) {
	fmt.Println("发送邮件")
	mailer := &mail.Mailer{}
	from := ""
	if mailer.UserName == "" {
		//basePath := revel.BasePath
		//c, _ := config.ReadDefault(basePath + "/conf/qishare.conf")
		c, _ := config.ReadDefault(revel.BasePath + "/conf/qishare.conf")
		/*mailer.UserName, _ = c.String("smtp", "smtp.username")
		mailer.Password, _ = c.String("smtp", "smtp.password")
		from, _ = c.String("smtp", "smtp.from")
		mailer.Server, _ = c.String("smtp", "smtp.server")
		mailer.Port, _ = c.Int("smtp", "smtp.port")*/
		fmt.Println("BasePath" + revel.BasePath)
		fmt.Println(c)
		mailer.UserName = "e23jiang@sina.com"
		mailer.Password = "7442008"
		from = "e23jiang@sina.com"
		mailer.Server = "smtp.sina.com"
		mailer.Port = 25
	}
	//mailer := &mail.Mailer{Server: "smtp.sina.com", Port: 25, UserName: "e23jiang@sina.com", Password: "7442008"}
	message := &mail.Message{From: from, To: tos, Subject: subject, PlainBody: bytes.NewBufferString(content)}

	err := mailer.SendMessage(message)
	if err != nil {
		fmt.Println(err)
	}
}
