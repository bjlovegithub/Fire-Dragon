package controllers

import (
	// "encoding/json"

	"Fire-Dragon/app"
	"Fire-Dragon/app/models"
	"github.com/revel/revel"
	"net/http"
)

type WishApp struct {
	*revel.Controller
}

func (c WishApp) Index() revel.Result {
	return c.Render()
}

func (c WishApp) GetMyWish(userId int64) revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))

	wish := models.Wish{
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

	return c.RenderJSON([3]models.Wish{wish, wish, wish})
}

func (c WishApp) GetWish() revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))

	wish := models.Wish{
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

	return c.RenderJSON(wish)
}

func (c WishApp) DeleteWish(wishId int64) revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))

	return c.RenderJSON("{}")
}

func query(sql string) {
	app.DB.Query(sql)
}

func (c WishApp) PutWish() revel.Result {
	// get wish id from post data(if there is a id for the wish, for update).
	jsonData := make(map[string]interface{})
	err := c.Params.BindJSON(&jsonData)
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Parse wish failed."})
	}
	jsonData["user_id"] = c.Request.Header.Get("User-Id")

	print(string(c.Params.JSON))

	wish := models.Wish{}
	sql, err := wish.UpsertSQL(jsonData)
	print(sql)
	query(sql)
	print(jsonData["wish"].(string))

	return c.RenderJSON("{}")
}

func (c WishApp) UpdateThumb(wishId int64) revel.Result {
	// update thumbs for the wish.
	c.Log.Info(c.Request.Header.Get("Authorization"))

	return c.RenderJSON("{}")
}
