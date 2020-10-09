package utils

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	TemplateCode string
	Url          string
	Method       string
	AppCode      string
	Length       int
}

func NewConfig(appCode string) *Config {
	return &Config{
		TemplateCode: "M09DD535F4",
		Url:          "https://smssend.shumaidata.com/sms/send?receive=%s&tag=%s&templateId=%s",
		Method:       "POST",
		AppCode:      appCode,
		Length:       6,
	}
}

func Rand(lenth int) string {
	var code string
	rand.Seed(time.Now().Unix())
	for i := 0; i < lenth; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

func Send(method, url, appCode string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "APPCODE"+" "+appCode)
	client.Do(req)
	return nil
}
