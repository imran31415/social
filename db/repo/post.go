package repo

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	FieldNameSocialPostUserId SocialPostFieldName = "user_id"
	FieldNameSocialPostId     SocialPostFieldName = "id"
)

type SocialPostFieldName string
type Post struct {
	Id      int64            `db:"id"`
	Content *json.RawMessage `db:"content"`
	UserId  int64            `db:"user_id"`
}

type Posts struct {
	Items []*Post
}

func (r *Repo) GetPostsIds(ids []int64, fieldName SocialPostFieldName) (*Posts, error) {
	posts := []*Post{}
	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM social_post WHERE %s IN (?) ORDER BY ID DESC", fieldName), ids)
	if err != nil {
		return nil, err
	}
	query = r.Db.Rebind(query)
	err = r.Db.Select(&posts, query, args...)

	if err != nil {
		return nil, err
	}
	return &Posts{
		Items: posts,
	}, nil
}

func (r *Repo) InsertPost(p *Post) (int64, error) {
	if p.Content == nil {
		return 0, fmt.Errorf("err, content is empty")
	}
	res, err := r.Db.NamedExec(`
		INSERT INTO social_post (
			content,
			user_id
		) VALUES (
			:content,
			:user_id
		);`, p)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
