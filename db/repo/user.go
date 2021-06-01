package repo

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id       int64            `db:"id"`
	Password string           `db:"password"`
	Username string           `db:"username"`
	Profile  *json.RawMessage `db:"profile"`
}

type Users struct {
	Items []*User
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

func (r *Repo) GetUsersOtherThanId(id int64) (*Users, error) {
	users := []*User{}
	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM social_user WHERE id != ?"), id)
	if err != nil {
		return nil, err
	}
	query = r.Db.Rebind(query)
	err = r.Db.Select(&users, query, args...)

	if err != nil {
		return nil, err
	}
	return &Users{
		Items: users,
	}, nil
}

func (r *Repo) InsertUser(p *User) (int64, error) {
	if p.Username == "" || p.Password == "" || p.Username == p.Password {
		return 0, fmt.Errorf("invalid user req")
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
