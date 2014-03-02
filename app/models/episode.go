package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"strings"
	"time"
)

//段子
type Episode struct {
	Id      int64
	Title   string `qbs:"size:32,notnull"`
	Content string `qbs:"size:1000,notnull"`
	Author  int64
	User    *User
	Tag     string
	DownNum int `qbs:"default:0"`
	UpNum   int `qbs:"default:0"`
	Created time.Time
}

//保存
func (this *Episode) Save(q *qbs.Qbs) bool {
	_, err := q.Save(this)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//分页查询
func GetEpisodes(q *qbs.Qbs, page int, column string, value interface{}, order string, url string) ([]*Episode, *Pagination) {
	page -= 1
	if page < 0 {
		page = 0
	}

	var episode []*Episode
	var rows int64
	if column == "" {
		fmt.Println(ItemsPerPage)
		fmt.Println(page)
		rows = q.Count("episode")
		err := q.OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&episode)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		rows = q.WhereEqual(column, value).Count("episode")
		err := q.WhereEqual(column, value).OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&episode)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(episode)
	url = url[:strings.Index(url, "=")+1]
	pagination := NewPagination(page, int(rows), url)

	return episode, pagination
}
