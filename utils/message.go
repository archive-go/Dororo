package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WxMessage struct {
	Touser string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

//通知相关工具

func WxSendMsg(msg string) (error) {
	resp , err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=wwf0cbb1d540eb320b&corpsecret=qHTR5i2AlEtHESGQbfQySSMzfrl2_CWZFZRSuFroY5c")
	if err != nil {
		return err
	}
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	token := data["access_token"]
	postData := WxMessage{
		Touser:  "@all",
		Msgtype: "text",
		Agentid: "1000004",
		Text: struct {
			Content string `json:"content"`
		}{
			Content:msg,
		},
		Safe: 0,
	}
	jsondata , err := json.Marshal(postData)

	if err != nil {
		return  err
	}


	_ , err = http.Post(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s" , token) , "application/json" ,  bytes.NewReader(jsondata) )
	if err != nil {
		return err
	}
	return nil
}



