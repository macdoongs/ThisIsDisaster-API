package controllers

import (
	"net/smtp"

	"github.com/revel/revel"
)

type SMTP struct {
	*revel.Controller
}

func (c SMTP) Index() revel.Result {
	var admin string = "rladudals02@gmail.com"
	var password string = "password"
	var server string = "smtp.gmail.com"

	// 메일서버 로그인 정보 설정
	auth := smtp.PlainAuth("", admin, password, server)

	from := admin
	to := []string{"rladudals02@ajou.ac.kr"} // 복수 수신자 가능

	// 메시지 작성
	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다\r\n"
	message := []byte(headerSubject + headerBlank + body)

	var code int = RESULT_CODE_SUCCESS
	var msg string = "Success"
	var rType string = "SMTP"

	// 메일 보내기
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		code = RESULT_CODE_FAILURE
		msg = "Fail"
	}

	// JSON response
	response := make(map[string]interface{})
	response["result_code"] = code
	response["result_msg"] = msg
	response["response_type"] = rType

	return c.RenderJSON(response)
}
