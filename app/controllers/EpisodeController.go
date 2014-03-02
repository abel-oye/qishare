package controllers

import (
	"fmt"
	"github.com/coocood/qbs"
	"github.com/river/qishare/app/models"
	"github.com/river/qishare/app/routes"
	"github.com/robfig/revel"
	"path"
	"strconv"
)

/**
* 段子
*
**/
type EpisodeController struct {
	Application
}

//页面跳转
func (this *EpisodeController) Jump() revel.Result {
	return this.RenderTemplate("Article/episode.html")
}

//新增
func (this *EpisodeController) New() revel.Result {
	return this.RenderTemplate("Article/newEpisode.html")
}

//分页查询
func (this *EpisodeController) QueryList(page int) revel.Result {
	episodes, pagination :=
		models.GetEpisodes(this.q, page, "", "", "created",
			routes.EpisodeController.QueryList(page))
	this.RenderArgs["episodes"] = episodes
	this.RenderArgs["pagination"] = pagination
	return this.RenderTemplate("article/episode.html")
}

//查询明细信息
func (this *EpisodeController) QueryDetail(id int64) revel.Result {
	episode := new(models.Episode)
	condition := qbs.NewCondition("id = ?", id)
	this.q.Condition(condition).Find(episode)
	user := FindUserById(this.q, episode.Author)
	episode.User = user
	this.RenderArgs["episode"] = episode
	return this.RenderTemplate("article/detailEpisode.html")
}

//保存
func (this *EpisodeController) Save(episode models.Episode) revel.Result {
	episode.Author, _ = strconv.ParseInt(this.Session["userId"], 10, 0)
	//上传文件
	saveFile(this.Request, "attach", path.Join(revel.BasePath, "attach"))
	fmt.Println(path.Join(revel.BasePath, "attach"))
	if !episode.Save(this.q) {
		this.Flash.Error("保存错误")
	}
	return this.Redirect(routes.EpisodeController.QueryDetail(episode.Id))
}

//删除
func (this *EpisodeController) Delete(id int64) revel.Result {

	return this.Redirect(routes.EpisodeController.QueryList(1))
}

func (this *EpisodeController) Good(id int64) revel.Result {
	return nil
}
