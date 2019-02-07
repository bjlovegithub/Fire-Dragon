package models

import (
	"fmt"
)

/**
CREATE TABLE user_auth (
  user_id INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  jwt_sub VARCHAR(128) NOT NULL,
  email   VARCHAR(128) NOT NULL,
  jwt     VARCHAR(2048) NOT NULL,
  jwt_exp TIMESTAMP,
  auth_type VARCHAR(128),
  UNIQUE KEY(jwt_sub, auth_type)
)
*/

type Wish struct {
	Id int
	UserEmail string
	UserId int
	Wish string
	FontFamily string
	FontSize int
	FontColor string
	BackgroundPic string
	Thumbs int
	CreatedTimestamp int
	UpdatedTimestamp int
}

func (u *Wish) String() string {
	return fmt.Sprintf("");
}
