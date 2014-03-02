package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"time"
)

//人员信息
type Person struct {
	Id           int64
	UserId       int64  `qbs:"notnull"`
	Name         string `qbs:"size:50,notnull"`
	IdType       string `qbs:"size:2"`
	IdNo         string `qbs:"size:50"`
	Sex          int32  `qbs:"size:1"`
	Birthday     time.Time
	Photo        int64
	Phone        string `qbs:"size:80"`
	Address      string `qbs:"size:120"`
	Company      string `qbs:"size:120"`
	Position     string `qbs:"size:50"`
	Industry     string `qbs:"size:50"`
	Introduction string `qbs:"size:500"`
}

//保存
func (p *Person) Save(q *qbs.Qbs) bool {
	_, err := q.Save(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//更新
func (p *Person) update(person Person, q *qbs.Qbs) bool {

	_, err := q.Update(person)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*
* 通过用户编号获得人员信息
* @param  userId 用户编号
* @return person
 */
func (p *Person) findPersonByUserId(userId string, q *qbs.Qbs) *Person {
	person := new(Person)
	q.WhereEqual("user_id", userId).Find(person)
	return person
}
