package main

import (
	log "github.com/tengfei-xy/go-log"
)

func (wsInfo *workspaceInfo) outputIgnore(ws, sb, page string) {
	if wsInfo.is_free_plan() {
		log.Warnf("忽略导出 工作区:%s 页面:%s", ws, page)
	} else {
		log.Warnf("忽略导出 工作区:%s 子空间:%s 页面:%s", ws, sb, page)
	}
}
func (wsInfo *workspaceInfo) outputExport(ws, sb, page string) {
	if wsInfo.is_free_plan() {
		log.Infof("开始导出 工作区:%s 页面:%s", ws, page)
	} else {
		log.Infof("开始导出 工作区:%s 子空间:%s 页面:%s", ws, sb, page)
	}
}

type exportStruct struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Data      string `json:"data"`
	Title     string `json:"title"`
	RequestID string `json:"requestId"`
}
