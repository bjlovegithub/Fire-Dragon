package controllers

import (
	"Fire-Dragon/app"
	"Fire-Dragon/app/models"
	"github.com/revel/revel"
	"net/http"

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
		fmt.Printf("------------\nid %d name is %s\n", wish.Id, wish.Wish)
		wishes = append(wishes, wish)
	}
	if err := rows.Err(); err != nil {
		c.Log.Fatal(err.Error())
	}

	return c.RenderJSON(wishes)
}

func (c App) Feedback() revel.Result {
	c.Response.Status = http.StatusOK
	return c.RenderJSON(map[string]string{"status": "ok"})
}
