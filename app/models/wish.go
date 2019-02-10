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

func (u *Wish) UpsertSQL() (string, error) {
	if u.Wish == "" {
		return "", errors.New("No wish")
	}

	var sql = ""
	if u.Id != 0 {
		sql = fmt.Sprintf("UPDATE wish SET wish = \"%s\", font_family = \"%s\", font_size = %d, font_color = \"%s\", background_pic = \"%s\", updated_at = now() WHERE id = %d", u.Wish, u.FontFamily, u.FontSize, u.FontColor, u.BackgroundPic, u.Id)
	} else {
		sql = fmt.Sprintf("INSERT INTO wish(user_id, wish, font_family, font_size, font_color, background_pic, created_at) VALUES(%d, \"%s\", \"%s\", %d, \"%s\", \"%s\", now())", u.UserId, u.Wish, u.FontFamily, u.FontSize, u.FontColor, u.BackgroundPic)
	}

	return sql, nil
}
