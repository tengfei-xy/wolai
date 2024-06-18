package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/tengfei-xy/go-log"
)

func getUserInfo() (WUserInfo, error) {
	var wui WUserInfo
	var p ReqUserInfoStruct
	h, err := getUserIDHtml()
	if err != nil {
		return WUserInfo{}, err
	}
	err = json.Unmarshal(h, &p)
	if err != nil {
		return WUserInfo{}, err
	}
	if p.Code != 1000 {
		return WUserInfo{}, fmt.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
	}
	wui.userid = p.Data.UserID

	l := len(p.Data.WorkspaceList)
	if l == 0 {
		log.Fatal("未发现工作空间")
	}
	wui.ws = make([]workspaceInfo, l)

	for i := range wui.ws {
		wui.ws[i].id = p.Data.WorkspaceList[i]
	}
	return wui, nil
}
func getUserIDHtml() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/authentication/user/getUserInfo`, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", `https://www.wolai.com/`)
	req.Header.Set("Cookie", config.Cookie)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("host", `api.wolai.com`)
	req.Header.Set("Origin", `https://www.wolai.com`)
	req.Header.Set("Sec-Fetch-Dest", `empty`)
	req.Header.Set("Sec-Fetch-Site", `same-site`)
	req.Header.Set("Accept-Language", `zh-CN,zh;q=0.9`)
	req.Header.Set("Accept", `application/json, text/plain, */*`)
	req.Header.Set("Content-Type", ` application/json`)
	req.Header.Set("Sec-Fetch-Mode", `cors`)
	req.Header.Set("wolai-os-platform", `mac`)
	req.Header.Set("x-client-timezone", `Asia/Shanghai`)
	req.Header.Set("wolai-app-version", `1.2.2-4`)
	req.Header.Set("wolai-client-platform", `web`)
	req.Header.Set("x-client-timeoffset", `-480`)
	req.Header.Set("wolai-client-version", ``)
	req.Header.Set("reqId", "7NYkp7BR1wdJm1oXFRevfJ")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("状态码:%d", resp.StatusCode)
	}

	resp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("内部错误:%v", err)
	}
	return resp_data, nil
}
func (wui *WUserInfo) outputUserID() {
	log.Infof("发现用户: %s", wui.userid)
}

type WUserInfo struct {
	userid string
	ws     []workspaceInfo
}
