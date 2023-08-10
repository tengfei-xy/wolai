package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	log "github.com/tengfei-xy/go-log"
)

type downloadMD struct {
	workSpaceName     string
	workSpaceID       string
	workSpacePageID   []string
	workSpacePageName []string
}

func getPagesList(cookie string) ([]downloadMD, bool) {

	// 获取工作区的所有名称、ID 和 工作区页面的ID
	data, ok := getWorkspaceData(cookie)
	if !ok {
		return nil, ok
	}
	d, ok := pagesDataDeal(data)
	if !ok {
		return nil, ok
	}

	// 补充工作区页面的名称
	for _, space := range d {
		data, ok := getWorkspacePagesData(cookie, &space)

		if !ok {
			break
		}
		getWorkspacePagesDeal(data, &space)
		for i, _ := range space.workSpacePageID {
			log.Infof("发现 工作区名称:%s 页面名称:%s", space.workSpaceName, space.workSpacePageName[i])
		}
	}

	return d, true
}
func getWorkspaceData(cookie string) ([]byte, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/workspace/getWorkspaceData`, nil)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	req.Header.Set("Referer", `https://www.wolai.com/`)
	req.Header.Set("Cookie", cookie)
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
	req.Header.Set("wolai-app-version", `1.1.3-1`)
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

func pagesDataDeal(data []byte) ([]downloadMD, bool) {
	var p workspaceDataStruct
	err := json.Unmarshal(data, &p)
	if err != nil {
		log.Error(err)
		return nil, false
	}
	if p.Code != 1000 {
		log.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
		return nil, false
	}
	dmd := make([]downloadMD, len(p.Data.Workspaces))
	for i, j := range p.Data.Workspaces {
		dmd[i].workSpaceName = j.Name
		dmd[i].workSpaceID = j.ID
		dmd[i].workSpacePageID = j.Pages
		dmd[i].workSpacePageName = make([]string, len(j.Pages))
	}

	return dmd, true
}
func getWorkspacePagesData(cookie string, dmd *downloadMD) ([]byte, bool) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/workspace/getWorkspacePages`, strings.NewReader(fmt.Sprintf(`{"spaceId":"%s"}`, *&dmd.workSpaceID)))
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	req.Header.Set("Referer", `https://www.wolai.com/`)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`)
	req.Header.Set("host", `api.wolai.com`)
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

func getWorkspacePagesDeal(data []byte, dmd *downloadMD) bool {
	var p workSpacePageList
	err := json.Unmarshal(data, &p)
	if err != nil {
		log.Error(err)
		return false
	}
	if p.Code != 1000 {
		log.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
		return false
	}
	for i, _ := range (*dmd).workSpacePageName {
		(*dmd).workSpacePageName[i] = p.Data.Blocks[(*dmd).workSpacePageID[i]].Value.Attributes.Title[0][0]
	}
	return true
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

type workspaceDataStruct struct {
	Code    int           `json:"code"`
	Data    workspaceData `json:"data"`
	Message string        `json:"message"`
}
type Members struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
type PageDefaultSetting struct {
	FullWidth bool `json:"full_width"`
	SmallText bool `json:"small_text"`
	Toc       bool `json:"toc"`
}

type Plan struct {
	Type      string `json:"type"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
}
type Workspaces struct {
	ID                         string        `json:"id"`
	Active                     bool          `json:"active"`
	BilinkColor                string        `json:"bilink_color"`
	CloseSharePageAd           bool          `json:"close_share_page_ad"`
	CloseSharePageBottomLogo   bool          `json:"close_share_page_bottom_logo"`
	CreatedBy                  string        `json:"created_by"`
	CreatedTime                int64         `json:"created_time"`
	DateFormatType             string        `json:"date_format_type"`
	DisableChangePublicTopPage bool          `json:"disable_change_public_top_page"`
	DisableCopy                bool          `json:"disable_copy"`
	DisableGuests              bool          `json:"disable_guests"`
	DisableImportPage          bool          `json:"disable_import_page"`
	DisableMemberCreateTeam    bool          `json:"disable_member_create_team"`
	DisableMovePage            bool          `json:"disable_move_page"`
	DisableSharePage           bool          `json:"disable_share_page"`
	Domain                     string        `json:"domain,omitempty"`
	DomainIndexPage            string        `json:"domain_index_page"`
	EditedBy                   string        `json:"edited_by"`
	EditedTime                 int64         `json:"edited_time"`
	FaqPages                   []interface{} `json:"faq_pages"`
	Icon                       string        `json:"icon"`
	ImageAutoOcr               bool          `json:"image_auto_ocr"`
	LastActiveTime             int64         `json:"last_active_time"`
	Members                    []Members     `json:"members"`
	Name                       string        `json:"name"`
	Pages                      []string      `json:"pages"`
	PlanPricePerCapitaPerDay   string        `json:"plan_price_per_capita_per_day,omitempty"`
	PlanType                   string        `json:"plan_type"`
	ShowWatermark              bool          `json:"show_watermark"`
	StartDayOfWeek             int           `json:"start_day_of_week"`
	StorageCount               int           `json:"storage_count,omitempty"`
	StorageLimit               int           `json:"storage_limit,omitempty"`
	TeamType                   string        `json:"team_type"`
	TimeZoneType               string        `json:"time_zone_type"`
	Version                    int           `json:"version"`
	// WorkspaceIcon            []interface{}      `json:"workspace_icon"`
	TeamSpaces         []interface{}      `json:"team_spaces"`
	DbTotalRowCount    string             `json:"db_total_row_count,omitempty"`
	PageDefaultSetting PageDefaultSetting `json:"page_default_setting,omitempty"`
	Plan               Plan               `json:"plan,omitempty"`
}

type SecureConfig struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
type workspaceData struct {
	Workspaces []Workspaces `json:"workspaces"`
	// SpaceViews        `json:"spaceViews"`
	// SecureConfig []SecureConfig `json:"secureConfig"`
}
