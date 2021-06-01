package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FeedItem struct {
	Id      int64 `db:"id"`
	OwnerId int64 `db:"owner_id"`
	PostId  int64 `db:"post_id"`
}

type Feed struct {
	Items []*FeedItem
}

func (r *Repo) InsertFeedItem(p *FeedItem) (int64, error) {
	if p.OwnerId == 0 || p.PostId == 0 {
		return 0, fmt.Errorf("invalid req")
	}
	res, err := r.Db.NamedExec(`
		INSERT INTO social_feed (
			owner_id,
			post_id
		) VALUES (
			:owner_id,
			:post_id
		);`, p)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *Repo) GetFeedItemById(id int64) (*FeedItem, error) {
	c := &FeedItem{}
	err := r.Db.Get(c, "SELECT * FROM social_feed where id = ?", id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repo) GetFeedByOwnerId(id int64) (*Feed, error) {
	items := []*FeedItem{}
	query, args, err := sqlx.In("SELECT * FROM social_feed WHERE owner_id = ? ORDER BY ID DESC LIMIT 1000", id)
	if err != nil {
		return nil, err
	}
	query = r.Db.Rebind(query)
	err = r.Db.Select(&items, query, args...)

	if err != nil {
		return nil, err
	}
	return &Feed{
		Items: items,
	}, nil
}
