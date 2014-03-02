package models

import (
	"crypto/md5"
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"io"
)

//用户
type User struct {
	Id             int64
	UserName       string `qbs:"size:32,notnull"`
	Password       string `qbs:"size:32,notnull"`
	Email          string `qbs:"size:80,unique,notnull"`
	AffirmPassword string `qbs:"-"`
	ValidateCode   string `qbs:"size:80,unique"`
	IsActive       bool
	Role           string `qbs:"size:2"`
	Status         int32  `qbs:"size:1"`
}

//保存
func (u *User) Save(q *qbs.Qbs) bool {
	if u.Password != "" {
		//u.HashedPassword = EncryptPassword(u.Password)
	}

	_, err := q.Save(u)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//校验
func (user *User) Validation(q *qbs.Qbs, v *revel.Validation) {

	v.Required(user.UserName).Message("请输入用户名")
	v.Required(user.Email).Message("请输入邮箱")
	v.Required(user.Password).Message("请输入密码")
	valid := v.Email(user.Email).Message("邮箱格式错误")

	if valid.Ok {
		if user.HasEmail(q) {
			err := &revel.ValidationError{
				Message: "该邮件已经注册过",
				Key:     "user.Email",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(user.Password == user.AffirmPassword).Message("两次密码输入不一致")

}

func (u *User) HasEmail(q *qbs.Qbs) bool {
	user := new(User)
	q.WhereEqual("email", u.Email).Find(user)

	return user.Id > 0
}

// 加密密码,转成md5
func EncryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}
