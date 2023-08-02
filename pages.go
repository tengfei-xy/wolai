package main

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	log "github.com/tengfei-xy/go-log"
)

type downloadMD struct {
	workSpacePageID   string
	workSpacePageName string
	url               string
	filename          string
}

func getPagesList(cookie string) ([]downloadMD, bool) {

	data, ok := getWorkspacePages(cookie)
	if !ok {
		return nil, ok
	}
	return pagesDeal(data)
}
func getWorkspacePages(cookie string) ([]byte, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/workspace/getWorkspacePages`, nil)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	req.Header.Set("Referer", `https://www.wolai.com/`)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("host", `api.wolai.com`)
	// req.Header.Set("Accept-Encoding", `gzip, deflate, br`)
	req.Header.Set("Origin", `https://www.wolai.com`)
	req.Header.Set("Sec-Fetch-Dest", `empty`)
	req.Header.Set("Sec-Fetch-Site", `same-site`)
	req.Header.Set("Accept-Language", `en-US,en;q=0.9`)
	req.Header.Set("Accept", `application/json, text/plain, */*`)
	req.Header.Set("Content-Type", ` application/json`)
	req.Header.Set("Sec-Fetch-Mode", `cors`)
	req.Header.Set("wolai-os-platform", `mac`)
	req.Header.Set("x-client-timezone", `Asia/Shanghai`)
	req.Header.Set("wolai-app-version", `1.1.2-3`)
	req.Header.Set("wolai-client-platform", `web`)
	req.Header.Set("x-client-timeoffset", `-480`)
	req.Header.Set("wolai-client-version", ``)
	req.Header.Set("reqId", "7NYkp7BR1wdJm1oXFRevfJ")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Errorf("状态码:%d", resp.StatusCode)
		return nil, false
	}

	resp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("内部错误:%v", err)
		return nil, false
	}
	return resp_data, true
}
func pagesDeal(data []byte) ([]downloadMD, bool) {
	var p workSpacePageList
	err := json.Unmarshal(data, &p)
	if err != nil {
		log.Error(err)
		return nil, false
	}
	if p.Code != 1000 {
		log.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
		return nil, false
	}
	dmd := make([]downloadMD, len(p.Data.Blocks))
	v := reflect.ValueOf(p.Data.Blocks)
	for i, id := range v.MapKeys() {
		dmd[i].workSpacePageID = id.String()
	}
	for i, id := range dmd {
		dmd[i].workSpacePageName = p.Data.Blocks[id.workSpacePageID].Value.Attributes.Title[0][0]
	}
	for _, id := range dmd {
		log.Infof("发现工作区页面 名称:%s", id.workSpacePageName)
	}
	return dmd, true
}
func pageInfoToExportReqJsonAll(d downloadMD) ([]byte, error) {
	var e exportUP
	e.PageID = d.workSpacePageID
	e.PageTitle = d.workSpacePageName
	e.Options.RecoverTree = true
	e.Options.GenerateToc = "all"
	e.Options.IncludeSubPage = true
	return json.Marshal(e)
}

type workSpacePageList struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Attributes struct {
	Title [][]string `json:"title"`
}

type Permissions struct {
	Type   string `json:"type"`
	Role   string `json:"role"`
	UserID string `json:"user_id"`
}
type Setting struct {
	BlockFullWidth                bool `json:"block_full_width"`
	BlockIsShowDirectorySuspended bool `json:"block_isShowDirectorySuspended"`
	BlockSmallText                bool `json:"block_small_text"`
	// PageIcon                      []string `json:"page_icon"`
}
type VisitRecords struct {
	UserID      string `json:"user_id"`
	VisitedTime int64  `json:"visited_time"`
}

type Value struct {
	ID                 string         `json:"id"`
	Active             bool           `json:"active"`
	Attributes         Attributes     `json:"attributes"`
	BlockDiscussIds    []interface{}  `json:"block_discuss_ids"`
	CreatedBy          string         `json:"created_by"`
	CreatedTime        int64          `json:"created_time"`
	EditedBy           string         `json:"edited_by"`
	EditedTime         int64          `json:"edited_time"`
	PageID             string         `json:"page_id"`
	ParentID           string         `json:"parent_id"`
	ParentType         string         `json:"parent_type"`
	Permissions        []Permissions  `json:"permissions"`
	Setting            Setting        `json:"setting"`
	Status             int            `json:"status"`
	SubNodes           []string       `json:"sub_nodes"`
	Type               string         `json:"type"`
	Version            int            `json:"version"`
	ViewCount          int            `json:"view_count"`
	VisitRecords       []VisitRecords `json:"visit_records"`
	WorkspaceID        string         `json:"workspace_id"`
	ResolvedDiscussIds []interface{}  `json:"resolved_discuss_ids"`
	Tableviews         []interface{}  `json:"tableviews"`
	SubPages           []string       `json:"sub_pages"`
}

type DatabaseViews struct {
}
type DatabaseTables struct {
}
type Data struct {
	Blocks         map[string]Zonelist `json:"blocks"`
	DatabaseViews  DatabaseViews       `json:"database_views"`
	DatabaseTables DatabaseTables      `json:"database_tables"`
}
type Zonelist struct {
	Role   string `json:"role"`
	Active bool   `json:"active"`
	Value  Value  `json:"value"`
}
