package controllers

import (
	"fmt"
	"github.com/river/qishare/app/models"
	"github.com/river/qishare/app/routes"
	"github.com/robfig/revel"
)

type Application struct {
	Qbs
}
type App struct {
	Application
}

func (this *Application) inject() revel.Result {
	return nil
}
func (this App) Index() revel.Result {
	episodes, pagination := models.GetEpisodes(this.q, 1, "", "", "created", routes.App.Index())
	fmt.Println("")
	return this.Render(episodes, pagination)
}
