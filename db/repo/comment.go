package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Comment struct {
	Id              int64            `db:"id"`
	PostId          int64            `db:"post_id"`
	ParentCommentId int64            `db:"parent_comment_id"`
	Content         *json.RawMessage `db:"content"`
}

type Comments struct {
	Items []*Comment
}

func (r *Repo) GetCommentsByPostId(id int64) (*Comments, error) {
	comments := []*Comment{}
	err := r.Db.Get(comments, "SELECT * FROM social_comment where post_id = ?", id)
	if err != nil {
		return nil, err
	}

	return &Comments{
		Items: comments,
	}, nil
}

func (r *Repo) GetCommentById(id int64) (*Comment, error) {
	comment := &Comment{}
	err := r.Db.Get(comment, "SELECT * FROM social_comment where id = ?", id)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *Repo) InsertComment(p *Comment) (int64, error) {
	if p.PostId == 0 || p.Content == nil {
		return 0, fmt.Errorf("invalid req")
	}
	var res sql.Result
	var err error
	if p.ParentCommentId != 0 {
		res, err = r.Db.NamedExec(`
		INSERT INTO social_comment (
			post_id,
			parent_comment_id,
		    content
		) VALUES (
			:post_id,
			:parent_comment_id,
		    :content
		);`, p)
	} else {
		res, err = r.Db.NamedExec(`
		INSERT INTO social_comment (
			post_id,
		    content
		) VALUES (
			:post_id,
		    :content
		);`, p)
	}
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()

}
