package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/tengfei-xy/go-log"
)

func getUserID() (string, error) {
	var p UserInfoStruct
	h, err := getUserIDHtml()
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(h, &p)
	if err != nil {
		return "", err
	}
	if p.Code != 1000 {
		return "", fmt.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
	}
	log.Infof("发现用户:%s", p.Data.UserID)
	return p.Data.UserID, nil
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
	req.Header.Set("wolai-app-version", `1.2.0-18`)
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

type UserInfoStruct struct {
	Code    int          `json:"code"`
	Data    UserInfoData `json:"data"`
	Message string       `json:"message"`
}
type UserInfoData struct {
	UserID                          string   `json:"userId"`
	Mobile                          []string `json:"mobile"`
	CountryCode                     string   `json:"countryCode"`
	Email                           string   `json:"email"`
	UserName                        string   `json:"userName"`
	Avatar                          string   `json:"avatar"`
	EmailVerified                   bool     `json:"emailVerified"`
	Password                        bool     `json:"password"`
	Pin                             bool     `json:"pin"`
	UserUnshareable                 bool     `json:"userUnshareable"`
	UserUnUploadable                bool     `json:"userUnUploadable"`
	CreditUnavailable               bool     `json:"creditUnavailable"`
	UserHash                        string   `json:"userHash"`
	RecommendCode                   string   `json:"recommendCode"`
	RegisterTime                    int64    `json:"registerTime"`
	IsEligibleForEducationDiscount  bool     `json:"isEligibleForEducationDiscount"`
	EducationDiscountExpirationDate int      `json:"educationDiscountExpirationDate"`
	IsNewUser                       bool     `json:"isNewUser"`
	RegisterMethod                  string   `json:"registerMethod"`
	DisableKefu                     bool     `json:"disableKefu"`
	InvitedUserCount                int      `json:"invitedUserCount"`
	WorkspaceList                   []string `json:"workspaceList"`
	WechatOpenID                    string   `json:"wechatOpenId"`
	WechatWebsiteOpenID             string   `json:"wechatWebsiteOpenId"`
	WechatUnionID                   string   `json:"wechatUnionId"`
}
