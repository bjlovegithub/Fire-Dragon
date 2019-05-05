package models

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

func (u *Feedback) UpsertSQL() (string, [3]interface{}) {
	sql := `
    INSERT INTO feedback(user_id, email, message, date)
    VALUES(?, ?, ?, now())`
	parameters := [3]interface{}{u.UserId, u.Email, u.Message}

	return sql, parameters
}
