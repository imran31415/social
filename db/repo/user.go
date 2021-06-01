package repo

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       int64            `db:"id"`
	Password string           `db:"password"`
	Username string           `db:"username"`
	Profile  *json.RawMessage `db:"profile"`
}

func (r *Repo) GetUserByUserName(username string) (*User, error) {
	c := &User{}
	err := r.Db.Get(c, "SELECT * FROM social_user where username = ?", username)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repo) GetUserById(id int64) (*User, error) {
	c := &User{}
	err := r.Db.Get(c, "SELECT * FROM social_user where id = ?", id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repo) InsertUser(p *User) (int64, error) {
	if p.Username == "" || p.Password == "" || p.Username == p.Password {
		return 0, fmt.Errorf("invalid req")
	}
	res, err := r.Db.NamedExec(`
		INSERT INTO social_user (
			password,
			username
		) VALUES (
			:password,
			:username
		);`, p)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
