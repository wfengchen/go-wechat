package wechatapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"strings"
)

type WechatApp struct {
}

type WechatLoginInfo struct {
	SessionKey string `json:"sessionkey"`
	Openid     string `json:"openid"`
	Unionid    string `json:"Unionid"`
	Errcode    int64  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

var (
	URL_CODE2SESSION string = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

func Login(appid, secret, code string) (WechatLoginInfo, error) {
	if code == "" {
		return WechatLoginInfo{}, errors.New("Code Empty.")
	}

	url := getUrlCode2Session(appid, secret, code)
	rsp, err := http.Get(url)

	if err != nil {
		return WechatLoginInfo{}, errors.New("Api return Faild.")
	}

	defer rsp.Body.Close()

	body_bytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		return WechatLoginInfo{}, errors.New("Call ioutil.ReadAll(rsp.Body) return false.")
	}

	rsp_string := string(body_bytes)
	rsp_data := WechatLoginInfo{}
	json.NewDecoder(strings.NewReader(rsp_string)).Decode(&rsp_data)

	return rsp_data, nil
}

func getUrlCode2Session(appid, secret, code string) string {
	return fmt.Sprintf(URL_CODE2SESSION, appid, secret, code)
}
