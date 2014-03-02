package controllers

import (
	"github.com/robfig/revel"
	"time"
)

func init() {
	revel.OnAppStart(Init)

	revel.InterceptMethod((*Qbs).Begin, revel.BEFORE)
	revel.InterceptMethod((*Application).inject, revel.BEFORE)
	revel.InterceptMethod((*Qbs).End, revel.AFTER)

	//注册模板函数，
	//不等于
	revel.TemplateFuncs["notEq"] = func(a, b interface{}) bool { return a != b }

	revel.TemplateFuncs["dateFormat"] = func(t time.Time) string {
		return t.Format("2006-01-02")
	}
}
