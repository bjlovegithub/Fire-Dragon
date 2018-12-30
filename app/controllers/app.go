package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Wish struct {
	Id int `json:"id"`
	UserEmail string `json:"userEmail"`
	Wish string `json:"wish"`
	FontFamily string `json:"fontFamily"`
	FontSize int `json:"fontSize"`
	FontColor string `json:"fontColor"`
	BackgroundPic string `json:"backgroundPic"`
	Thumbs int `json:"thumbs"`
	CreatedTimestamp int `json::"createdTimestamp"`
	UpdatedTimestamp int `json:"updatedTimestamp"`
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) BoardWish() revel.Result {
	wish1 := Wish{
		Id: 1,
		UserEmail: "abc@abc.com",
		Wish: "this is my wish 1\nhappy",
		FontFamily: "Helvetica",
		FontSize: 16,
		FontColor: "blue",
		BackgroundPic: "https://images.pexels.com/photos/17679/pexels-photo.jpg?w=940&h=650&dpr=2&auto=compress&cs=tinysrgb",
		Thumbs: 1,
		CreatedTimestamp: 1525506395,
		UpdatedTimestamp: 0,
	}

	wish2 := wish1
	wish2.Id = 2

	wish3 := wish1
	wish3.Id = 3

	arr := [3]Wish{ wish1, wish2, wish3 }

	return c.RenderJSON(arr)
}
