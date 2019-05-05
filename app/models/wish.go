package models

import (
	"errors"
	"fmt"
)

/**
CREATE TABLE wish (
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  wish VARCHAR(1024) NOT NULL,
  font_family   VARCHAR(128) NOT NULL,
  font_size INT(10),
  font_color     VARCHAR(128) NOT NULL,
  background_pic VARCHAR(128) NOT NULL,
  thumbs    INT,
  created_at INT,
  updated_at INT
)
*/

type Wish struct {
	Id               int    `json:"id"`
	UserId           int    `json:"userId"`
	UserEmail        string `json:"userEmail"`
	Wish             string `json:"wish"`
	FontFamily       string `json:"fontFamily"`
	FontSize         int    `json:"fontSize"`
	FontColor        string `json:"fontColor"`
	BackgroundPic    string `json:"backgroundPic"`
	Thumbs           int    `json:"thumbs"`
	CreatedTimestamp int    `json:"createdTimestamp"`
	UpdatedTimestamp int    `json:"updatedTimestamp"`
}

func (u *Wish) String() string {
	return fmt.Sprintf("")
}

func (u *Wish) UpsertSQL() (string, []interface{}, error) {
	if u.Wish == "" {
		return "", []interface{}{}, errors.New("No wish")
	}

	var sql = ""
	var parameters []interface{}
	if u.Id != 0 {
		sql = "UPDATE wish SET wish = ?, font_family = ?, font_size = ?, font_color = ?, background_pic = ?, updated_at = ? WHERE id = ?"
		parameters = []interface{}{u.Wish, u.FontFamily, u.FontSize, u.FontColor, u.BackgroundPic, u.UpdatedTimestamp, u.Id}
	} else {
		sql = "INSERT INTO wish(user_id, wish, thumbs, font_family, font_size, font_color, background_pic, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
		parameters = []interface{}{u.UserId, u.Wish, u.Thumbs, u.FontFamily, u.FontSize, u.FontColor, u.BackgroundPic, u.CreatedTimestamp, u.UpdatedTimestamp}
	}

	return sql, parameters, nil
}
