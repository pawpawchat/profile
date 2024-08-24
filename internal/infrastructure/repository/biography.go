package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type BiographyRepository struct {
	db *sqlx.DB
}

func NewBiographyRepository(db *sqlx.DB) *BiographyRepository {
	return &BiographyRepository{db}
}

func (r *BiographyRepository) Create(ctx context.Context, bio *model.Biography) error {
	sql, args := squirrel.Insert("profile_biographies").
		Columns("profile_id", "first_name", "second_name", "birthday").
		Values(bio.ProfileID, bio.FirstName, bio.SecondName, bio.Birthday).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING biography_id").
		MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&bio.ID)
}
