package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type AvatarsRepository struct {
	db *sqlx.DB
}

func NewAvatarsRepository(db *sqlx.DB) *AvatarsRepository {
	return &AvatarsRepository{db}
}

func (r *AvatarsRepository) Create(ctx context.Context, avatar *model.Avatar) error {
	sql, args := squirrel.Insert("profile_avatars").
		Columns("orig_url").
		Values(avatar.OrigURL).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING avatar_id").
		MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&avatar.ID)
}
