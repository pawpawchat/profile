package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileRepository struct {
	db *sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (r *ProfileRepository) Create(ctx context.Context, profile *model.Profile) error {
	sql := "INSERT INTO profiles DEFAULT VALUES RETURNING profile_id, username, last_seen, created_at"
	return r.db.GetContext(ctx, profile, sql)
}

func (r *ProfileRepository) GetById(ctx context.Context, id int64) (*model.Profile, error) {
	pwb := new(struct {
		model.Profile
		model.Biography
	})

	query, args := sq.Select("*").
		From("profiles").
		Join("profile_biographies USING (profile_id)").
		Where(sq.Eq{"profile_id": id}).
		PlaceholderFormat(sq.Dollar).
		MustSql()

	if err := r.db.GetContext(ctx, pwb, query, args...); err != nil {
		return nil, err
	}

	pwb.Profile.Biography = pwb.Biography
	return &pwb.Profile, nil
}

func (r *ProfileRepository) SetAvatar(ctx context.Context, profileID int64, avatarID int64) error {
	sql, args := squirrel.Update("profiles").
		Set("avatar_id", avatarID).
		Where(squirrel.Eq{"profile_id": profileID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING profile_id").
		MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&profileID)
}
