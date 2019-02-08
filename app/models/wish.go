package models

import (
	"errors"
	"fmt"
)

/**
CREATE TABLE wish (
  id INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id INT(10) NOT NULL,
  wish VARCHAR(1024) NOT NULL,
  font_family   VARCHAR(128) NOT NULL,
  font_size INT(10),
  font_color     VARCHAR(128) NOT NULL,
  background_pic VARCHAR(128) NOT NULL,
  thumbs    INT(10),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
)
*/

type Wish struct {
	Id               int    `json:"id"`
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

func (u *Wish) UpsertSQL(m map[string]interface{}) (string, error) {
	userId, exists := m["user_id"]
	if !exists {
		return "", errors.New("No user id")
	}

	wish, exists := m["wish"]
	if !exists {
		return "", errors.New("No wish")
	}

	font, exists := m["fontFamily"]
	// TODO - assign default

	fontSize, exists := m["fontSize"]

	fontColor, exists := m["fontColor"]

	pic, exists := m["backgroudPic"]

	wishId, exists := m["id"]
	println("-----------------")
	println(wishId)
	var sql = ""
	if exists {
		sql = fmt.Sprintf("UPDATE wish SET wish = \"%s\", font_family = \"%s\", font_size = %d, font_color = \"%s\", background_pic = \"%s\", updated_at = now() WHERE id = %d", wish, font, fontSize.(int), fontColor, pic, wishId)
	} else {
		sql = fmt.Sprintf("INSERT INTO wish(user_id, wish, font_family, font_sie, font_color, background_pic, created_at) VALUES(%d, \"%s\", \"%s\", %d, \"%s\", \"%s\", now())", userId, wish, font, fontSize, fontColor, pic)
	}

	return sql, nil
}
