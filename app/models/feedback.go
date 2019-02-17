package models

import (
	"fmt"
)

/**
CREATE TABLE feedback (
  user_id INT NOT NULL,
  email   VARCHAR(128) NOT NULL,
  message VARCHAR(20480) NOT NULL,
)
*/

type Feedback struct {
	UserId  int
	Email   string
	Message string
}

func (u *Feedback) UpsertSQL() string {
	return fmt.Sprintf(`
    INSERT INTO user_auth(jwt_sub, email, jwt, jwt_exp, auth_type)
    VALUES('%s', '%s', '%s', FROM_UNIXTIME(%d), '%s')
    ON DUPLICATE KEY UPDATE jwt = '%s', jwt_exp = FROM_UNIXTIME(%d)`,
		u.JWTSub, u.Email, u.JWT, u.JWTExp, u.AuthType, u.JWT, u.JWTExp)
}

func (u *Feedback) String() string {
	return fmt.Sprintf("User(%s), JWTSub(%s), Email(%s), JWT(%s), JWTExp(%d)",
		u.UserId, u.JWTSub, u.Email, u.JWT, u.JWTExp)
}
