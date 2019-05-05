package controllers

import (
	"Fire-Dragon/app"
	"Fire-Dragon/app/models"
	"encoding/json"
	"github.com/revel/revel"
	"net/http"
	"strconv"

	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetBoardWish() revel.Result {
	sql := "SELECT id, user_id, wish, font_family, font_size, font_color, background_pic, thumbs, created_at, updated_at FROM wish ORDER BY RAND() LIMIT 3"

	rows, err := app.DB.Query(sql)
	if err != nil {
		c.Log.Panic(fmt.Sprintf("Get board wishes failed, error: %s", err.Error()))
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

func (c App) Feedback() revel.Result {
	feedback := models.Feedback{}
	err := json.Unmarshal(c.Params.JSON, &feedback)
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Parse feedback failed."})
	}

	feedback.UserId, err = strconv.Atoi(c.Request.Header.Get("User-Id"))
	if err != nil {
		c.Log.Error(err.Error())
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(map[string]interface{}{"message": "Invalid User Id."})
	}

	bytes, _ := json.Marshal(feedback)
	print(string(bytes))

	sql, parameters := feedback.UpsertSQL()

	_, err = app.DB.Query(sql, parameters[:]...)
	if err != nil {
		c.Log.Error(fmt.Sprintf("Query DB error: %s, (%s)", err.Error, sql))
		panic(err)
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(map[string]string{"status": "ok"})
}
