package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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

func (c WishApp) GetMyWish() revel.Result {
	userId, err := strconv.Atoi(c.Params.Get("user_id"))
	if err != nil {
		c.Log.Panic(fmt.Sprintf("Invalid User Id: %s", c.Params.Get("user_id")))
	}
	sql := "SELECT id, user_id, wish, font_family, font_size, font_color, background_pic, thumbs, created_at, updated_at FROM wish WHERE user_id = ?"
	rows, err := app.DB.Query(sql, userId)
	if err != nil {
		c.Log.Panic(fmt.Sprintf("Get wishes for %d failed, error: %s", userId, err.Error))
	}
	defer rows.Close()

	wishes := make([]models.Wish, 0)
	for rows.Next() {
		wish := models.Wish{}
		if err := rows.Scan(&wish.Id, &wish.UserId, &wish.Wish, &wish.FontFamily, &wish.FontSize, &wish.FontColor, &wish.BackgroundPic, &wish.Thumbs, &wish.CreatedTimestamp, &wish.UpdatedTimestamp); err != nil {
			c.Log.Fatal(err.Error())
		}
		wishes = append(wishes, wish)
	}
	if err := rows.Err(); err != nil {
		c.Log.Fatal(err.Error())
	}
	return c.RenderJSON(wishes)
}

func (c WishApp) DeleteWish(wishId int64) revel.Result {
	c.Log.Info(c.Request.Header.Get("Authorization"))

	return c.RenderJSON("{}")
}

func query(sql string, parameters []interface{}, c WishApp) {
	_, err := app.DB.Query(sql, parameters[:]...)
	if err != nil {
		c.Log.Error(fmt.Sprintf("Query DB error: %s, (%s)", err.Error, sql))
		panic(err)
	}
}

func (c WishApp) PutWish() revel.Result {
	// get wish id from post data(if there is a id for the wish, for update).
	wish := models.Wish{Thumbs: 0, CreatedTimestamp: int(time.Now().Unix()), UpdatedTimestamp: int(time.Now().Unix())}
	err := json.Unmarshal(c.Params.JSON, &wish)
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Parse wish failed."})
	}
	wish.UserId, err = strconv.Atoi(c.Request.Header.Get("User-Id"))
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Invalid User Id."})
	}

	bytes, _ := json.Marshal(wish)
	print(string(bytes))

	sql, parameters, err := wish.UpsertSQL()
	query(sql, parameters, c)

	return c.RenderJSON("{}")
}

func (c WishApp) UpdateThumb() revel.Result {
	// update thumbs for the wish.
	wishId, err := strconv.Atoi(c.Params.Get("wish_id"))
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Invalid Wish Id."})
	}
	userId, err := strconv.Atoi(c.Request.Header.Get("User-Id"))
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Invalid User Id."})
	}

	sql := "UPDATE wish SET thumbs = thumbs + 1 WHERE id = ?"
	_, err = app.DB.Query(sql, wishId)
	if err != nil {
		c.Log.Error(fmt.Sprintf("Query DB error: %s, (%s)", err.Error, sql))
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(map[string]interface{}{"message": "Query DB Failed."})
	}

	c.Log.Info(string(userId))

	return c.RenderJSON("{}")
}
