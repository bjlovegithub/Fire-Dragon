package models

import (
	"fmt"
)

/**
CREATE TABLE user_auth (
  user_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  jwt_sub VARCHAR(128) NOT NULL,
  email   VARCHAR(128) NOT NULL,
  jwt     VARCHAR(2048) NOT NULL,
  jwt_exp TIMESTAMP,
  auth_type VARCHAR(128),
  UNIQUE KEY(jwt_sub, auth_type)
)
*/

type UserAuth struct {
	UserId   int
	JWTSub   string
	Email    string
	JWT      string
	JWTExp   int64
	AuthType string
}

func (u *UserAuth) UpsertSQL() string {
	return fmt.Sprintf(`
    INSERT INTO user_auth(jwt_sub, email, jwt, jwt_exp, auth_type)
    VALUES('%s', '%s', '%s', FROM_UNIXTIME(%d), '%s')
    ON DUPLICATE KEY UPDATE jwt = '%s', jwt_exp = FROM_UNIXTIME(%d)`,
		u.JWTSub, u.Email, u.JWT, u.JWTExp, u.AuthType, u.JWT, u.JWTExp)
}

func (u *UserAuth) String() string {
	return fmt.Sprintf("User(%s), JWTSub(%s), Email(%s), JWT(%s), JWTExp(%d)",
		u.UserId, u.JWTSub, u.Email, u.JWT, u.JWTExp)
}
