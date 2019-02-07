package controllers

import (
	// "encoding/json"	
	
	"github.com/revel/revel"

	"Fire-Dragon/app"
)

type WishApp struct {
	*revel.Controller
}

func (c WishApp) Index() revel.Result {
	return c.Render()
}

func (c WishApp) GetMyWish(userId int64) revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))

	wish := Wish{
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
	
	return c.RenderJSON([3]Wish{wish, wish, wish})
}

func (c WishApp) GetWish() revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))
	
	wish := Wish{
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

	return c.RenderJSON(wish)
}

func (c WishApp) DeleteWish(wishId int64) revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))
	
	return c.RenderJSON("{}")
}

func query() {
	app.DB.Query("select 1")
}

func (c WishApp) PutWish() revel.Result {
	// get wish id from post data(if there is a id for the wish, for update).
	jsonData := make(map[string]interface{})
	c.Params.BindJSON(&jsonData)

	/*
	err := json.Unmarshal(c.Params.JSON, &jsonData)
	if err != nil {
		panic(err)
	}
*/

	query()
	print(jsonData["wish"].(string))
	
	return c.RenderJSON("{}")
}

func (c WishApp) UpdateThumb(wishId int64) revel.Result {
	// update thumbs for the wish.
	c.Log.Info(c.Request.Header.Get("Authorization"))
	
	return c.RenderJSON("{}")
}
