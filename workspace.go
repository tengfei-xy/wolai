package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"

	log "github.com/tengfei-xy/go-log"
)

type workspaceInfo struct {
	name     string
	id       string
	planType string
	subspace []subspace
}
type subspace struct {
	name  string
	pages []pageList
}
type pageList struct {
	name string
	id   string
}

func getWorkSpaceStruct() (workspaceDataStruct, error) {
	var p workspaceDataStruct

	// 获取工作区信息
	h, err := getWorkspaceHtml()
	if err != nil {
		return workspaceDataStruct{}, err
	}

	err = json.Unmarshal(h, &p)
	if err != nil {
		return workspaceDataStruct{}, err
	}
	if p.Code != 1000 {
		return workspaceDataStruct{}, fmt.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
	}
	return p, nil
}
func getWorkspaceHtml() ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/workspace/getWorkspaceData`, nil)
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

func (ws *workspaceDataStruct) getWorkspaceInfo(userid string) []workspaceInfo {

	wsInfo := make([]workspaceInfo, len(ws.Data.Workspaces))
	for i, j := range ws.Data.Workspaces {
		if j.member_exist(userid) == false {
			log.Warnf("跳过工作区:%s 原因:该用户是协作访客", j.Name)
			continue
		}
		wsInfo[i].name = j.Name
		wsInfo[i].id = j.ID
		wsInfo[i].planType = j.Plan.Type
	}
	return wsInfo
}
func (wsInfo *workspaceInfo) mkdirBackupFolder() error {
	// 免费计划
	if wsInfo.is_free_plan() {
		return mkdir(filepath.Join(config.BackupPath, wsInfo.name))
	}

	for _, subspace := range wsInfo.subspace {
		if err := mkdir(filepath.Join(config.BackupPath, wsInfo.name, subspace.name)); err != nil {
			return err
		}
	}
	return nil
}
func (wsInfo *workspaceInfo) is_free_plan() bool {
	return wsInfo.planType == "free"
}
func (w *Workspaces) member_exist(userid string) bool {
	var exist bool = false
	for _, j := range w.Members {
		if j.UserID == userid {
			exist = true
		}
	}
	return exist
}

// 说明: 适用于团队模式或家庭版
func (wsInfo *workspaceInfo) getTeamSubspace() error {
	if wsInfo.is_free_plan() {
		return nil
	}
	ts, err := getTeam(wsInfo.id)
	if err != nil {
		log.Error(err)
		return err
	}

	wsInfo.subspace = make([]subspace, len(ts.Data))
	for i, j := range ts.Data {
		wsInfo.subspace[i].name = j.Name
		wsInfo.subspace[i].pages = make([]pageList, len(j.TeamPages))
		for z, line := range j.TeamPages {
			wsInfo.subspace[i].pages[z].id = line
		}
	}
	return nil
}

// 说明: 适用于团队模式或家庭版
func (wsInfo *workspaceInfo) getTermPagesMain() error {

	h, err := wsInfo.getPagesHtml()
	if err != nil {
		log.Error(err)
		return err
	}
	if wsInfo.getTermPagesDeal(h); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// 说明: 适用于团队模式或家庭版

func (wsInfo *workspaceInfo) getDefaultSubspace() error {
	wsInfo.subspace = make([]subspace, 1)
	wsInfo.subspace[0].name = ""
	h, err := wsInfo.getPagesHtml()
	if err != nil {
		log.Error(err)
		return err
	}
	if err := wsInfo.getDefaultPagesDeal(h); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (wsInfo *workspaceInfo) getPagesHtml() ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", `https://api.wolai.com/v1/workspace/getWorkspacePages`, strings.NewReader(fmt.Sprintf(`{"spaceId":"%s"}`, wsInfo.id)))
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
	req.Header.Set("Accept-Language", `en-US,en;q=0.9`)
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
		return nil, fmt.Errorf("内部错误:%v", err)
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

func (wsInfo *workspaceInfo) getTermPagesDeal(data []byte) error {
	var p workSpacePageList
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	if p.Code != 1000 {
		return fmt.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
	}

	for i, j := range wsInfo.subspace {
		for z, k := range j.pages {
			wsInfo.subspace[i].pages[z].name = p.Data.Blocks[k.id].Value.Attributes.Title[0][0]
		}
	}

	return nil
}
func (wsInfo *workspaceInfo) getDefaultPagesDeal(data []byte) error {
	var p workSpacePageList
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	if p.Code != 1000 {
		return fmt.Errorf("请求异常 状态码:%d 消息:%s", p.Code, p.Message)
	}

	v := reflect.ValueOf(p.Data.Blocks)
	wsInfo.subspace[0].pages = make([]pageList, len(v.MapKeys()))
	for i, id := range v.MapKeys() {
		// 忽略星标置顶
		if p.Data.Blocks[id.String()].Value.ParentType != "workspace" {
			continue
		}
		wsInfo.subspace[0].pages[i].id = id.String()
		wsInfo.subspace[0].pages[i].name = p.Data.Blocks[id.String()].Value.Attributes.Title[0][0]
	}
	return nil
}
func (wsInfo *workspaceInfo) output() {
	var msg string
	for _, j := range wsInfo.subspace {
		for _, k := range j.pages {
			if wsInfo.is_free_plan() {
				msg = fmt.Sprintf("发现工作区:%s 页面:%s", wsInfo.name, k.name)
			} else {
				msg = fmt.Sprintf("发现工作区:%s 子空间:%s 页面:%s", wsInfo.name, j.name, k.name)
			}
			log.Infof(msg)
		}
	}
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
	ShowWateamark              bool          `json:"show_wateamark"`
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
