package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getTeam(spacceID string) (teamStruct, error) {

	h, err := getTeamHtml(spacceID)
	if err != nil {
		return teamStruct{}, err
	}
	ts, err := getTeamStruct(h)
	if err != nil {
		return teamStruct{}, err
	}
	if ts.Code != 1000 {
		return teamStruct{}, fmt.Errorf("错误的数据,%s", ts.Message)
	}
	if len(ts.Data) == 0 {
		return teamStruct{}, fmt.Errorf("空数据,可能您没有导出的权限")
	}
	return ts, nil
}
func getTeamHtml(spaceID string) (strings.Builder, error) {
	var sb strings.Builder

	client := &http.Client{}
	data := fmt.Sprintf("{\"spaceId\":\"%s\"}", spaceID)
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/workspace/getTeams`, strings.NewReader(data))
	if err != nil {
		return sb, err
	}
	req.Header.Set("Cookie", config.Cookie)
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("Host", `api.wolai.com`)
	req.Header.Set("Origin", `https://www.wolai.com`)
	req.Header.Set("Accept-Language", `en-US,en;q=0.9`)
	req.Header.Set("Accept", ` application/json, text/plain, */*`)
	req.Header.Set("Content-Type", ` application/json`)
	req.Header.Set("wolai-os-platform", `mac`)
	req.Header.Set("x-client-timezone", `Asia/Shanghai`)
	req.Header.Set("wolai-app-version", `1.2.2-4`)
	req.Header.Set("wolai-client-platform", `web`)
	req.Header.Set("x-client-timeoffset", `-480`)
	req.Header.Set("wolai-client-version", ``)
	req.Header.Set("reqId", "5Kq5MgZQf2TZhn94x2C6HZ")
	req.Header.Set("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {

		return sb, fmt.Errorf("内部错误:%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {

		return sb, fmt.Errorf("状态码:%d", resp.StatusCode)
	}

	// 检查响应的Content-Encoding头部字段
	if resp.Header.Get("Content-Encoding") == "gzip" {
		// 如果是gzip压缩的数据，使用gzip.Reader进行解压
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			// 处理错误
			return sb, err
		}
		defer reader.Close()

		// 将解压后的数据读取到一个字符串中
		_, err = io.Copy(&sb, reader)
		if err != nil {
			// 处理错误
			return sb, err
		}
	}

	return sb, nil
}
func getTeamStruct(data strings.Builder) (teamStruct, error) {
	var t teamStruct
	err := json.Unmarshal([]byte(data.String()), &t)
	if err != nil {
		return teamStruct{}, err
	}
	return t, nil
}

type teamStruct struct {
	Code    int        `json:"code"`
	Data    []TeamData `json:"data"`
	Message string     `json:"message"`
}
type TeamMembers struct {
	Type   string `json:"type"`
	Role   string `json:"role"`
	UserID string `json:"user_id"`
}
type TeamPermissions struct {
	Type        string `json:"type"`
	TeamID      string `json:"team_id,omitempty"`
	Role        string `json:"role"`
	WorkspaceID string `json:"workspace_id,omitempty"`
}
type Settings struct {
	Visibility                 string `json:"visibility"`
	JoinType                   string `json:"join_type"`
	InviteAccess               string `json:"invite_access"`
	DisableExportPage          bool   `json:"disable_export_page"`
	DisableSharePage           bool   `json:"disable_share_page"`
	DisableGuests              bool   `json:"disable_guests"`
	DisableChangePublicTopPage bool   `json:"disable_change_public_top_page"`
}
type TeamData struct {
	ID          string            `json:"id"`
	CreatedBy   string            `json:"created_by"`
	CreatedTime int64             `json:"created_time"`
	Description string            `json:"description"`
	EditedBy    string            `json:"edited_by"`
	EditedTime  int64             `json:"edited_time"`
	IsDefault   bool              `json:"is_default"`
	Members     []TeamPermissions `json:"members"`
	Name        string            `json:"name"`
	Permissions []Permissions     `json:"permissions"`
	Settings    Settings          `json:"settings"`
	Status      int               `json:"status"`
	TeamPages   []string          `json:"team_pages"`
	Version     int               `json:"version"`
	// WorkspaceIcon []string          `json:"workspace_icon"`
	WorkspaceID string `json:"workspace_id"`
}
