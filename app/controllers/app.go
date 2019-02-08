package controllers

import (
	"Fire-Dragon/app/models"
	"github.com/revel/revel"
	"net/http"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetBoardWish() revel.Result {
	wish1 := models.Wish{
		Id:               1,
		UserEmail:        "abc@abc.com",
		Wish:             "this is my wish 1\nhappy",
		FontFamily:       "Helvetica",
		FontSize:         16,
		FontColor:        "blue",
		BackgroundPic:    "https://images.pexels.com/photos/17679/pexels-photo.jpg?w=940&h=650&dpr=2&auto=compress&cs=tinysrgb",
		Thumbs:           1,
		CreatedTimestamp: 1525506395,
		UpdatedTimestamp: 0,
	}

	wish2 := wish1
	wish2.Id = 2

	wish3 := wish1
	wish3.Id = 3

	arr := [3]models.Wish{wish1, wish2, wish3}

	return c.RenderJSON(arr)
}

func (c App) Feedback() revel.Result {
	c.Response.Status = http.StatusOK
	return c.RenderJSON(map[string]string{"status": "ok"})
}
