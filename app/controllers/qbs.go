package controllers

import (
	"fmt"
	_ "github.com/coocood/mysql"
	"github.com/coocood/qbs"
	"github.com/river/qishare/app/models"
	"github.com/robfig/config"
	"github.com/robfig/revel"
)

type Qbs struct {
	*revel.Controller
	q *qbs.Qbs
}

//开始连接
func (c *Qbs) Begin() revel.Result {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	c.q = q
	return nil
}

//关闭连接
func (c *Qbs) End() revel.Result {
	c.q.Close()
	return nil
}

func Init() {
	//uploadPath := basePath + "/public/upload/"
	c, _ := config.ReadDefault(revel.BasePath + "/conf/qishare.conf")
	driver, _ := c.String("database", "db.driver")
	dbname, _ := c.String("database", "db.dbname")
	user, _ := c.String("database", "db.user")
	password, _ := c.String("database", "db.password")
	host, _ := c.String("database", "db.host")
	connectDb(driver, dbname, user, password, host)
}

//连接数据库
func connectDb(driver, dbname, user, password, host string) {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)
	qbs.Register(driver, params, dbname, qbs.NewMysql())
	err := initTable()
	if err != nil {
		fmt.Println(err)
	}
}

//初始化表
func initTable() error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}
	defer migration.Close()

	err = migration.CreateTableIfNotExists(new(models.User))
	err = migration.CreateTableIfNotExists(new(models.Person))
	err = migration.CreateTableIfNotExists(new(models.Episode))
	/*err = migration.CreateTableIfNotExists(new(models.Reply))
	err = migration.CreateTableIfNotExists(new(models.Permissions))*/

	return err
}
