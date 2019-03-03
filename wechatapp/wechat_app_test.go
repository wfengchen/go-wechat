package wechatapp

import (
	"testing"
	"wechat/wechatapp"
)

func TestCode2Session(t *testing.T) {
	code := "021WokCN1JzTF81KAoAN17GdCN1WokCd"

	appid := "wx371628456b599a40"
	secret := "1dd8aed1b39474758ede01eba23d3f40"

	_, err := wechatapp.Login(appid, secret, code)
	if err != nil {
		t.Errorf("Login Faild. ERR:" + err.Error())
	}
}
