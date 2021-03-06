package email

import (
	"fmt"
	"go_api_frame/web/common/config"
	"go_api_frame/web/common/util"
	"gopkg.in/gomail.v2"
	"runtime/debug"
	"strings"
	"time"
)

func Email(str, url, ua, ip string) {
	if !config.MyConfig.SendEmail {
		//开启发送邮件功能
		return
	}
	now := time.Now().Local().Format(util.TimeFormat)
	DebugStack := ""
	for _, v := range strings.Split(string(debug.Stack()), "\n") {
		DebugStack += v + "<br>"
	}

	subject := fmt.Sprintf("【系统告警】%s 服务出错了！", config.MyConfig.ServiceName)

	body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", str))
	body = strings.ReplaceAll(body, "{RequestTime}", now)
	body = strings.ReplaceAll(body, "{RequestURL}", url)
	body = strings.ReplaceAll(body, "{RequestUA}", ua)
	body = strings.ReplaceAll(body, "{RequestIP}", ip)
	body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

	// 执行发邮件
	m := gomail.NewMessage()
	//设置发件人
	m.SetHeader("From", config.MyConfig.FromEmailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(config.MyConfig.ToEmailUSer, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", subject)

	//设置邮件正文
	m.SetBody("text/html", body)

	d := gomail.NewDialer(config.MyConfig.Host, config.MyConfig.Port, config.MyConfig.FromEmailUser, config.MyConfig.EmailPass)

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
}
