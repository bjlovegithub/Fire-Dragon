package models

import (
	"fmt"
)

/**
CREATE TABLE feedback (
  id INT  UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  email   VARCHAR(128) NOT NULL,
  message TEXT NOT NULL,
  date    Timestamp
)
*/

type Feedback struct {
	UserId  int
	Email   string
	Message string
}

func (u *Feedback) UpsertSQL() string {
	// TODO - Clean message(SQL Injection, etc)
	return fmt.Sprintf(`
    INSERT INTO feedback(user_id, email, message, date)
    VALUES(%d, '%s', '%s', now())`, u.UserId, u.Email, u.Message)
}
