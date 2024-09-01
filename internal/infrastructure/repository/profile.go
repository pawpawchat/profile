package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
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
	sql := `INSERT INTO profiles 
			DEFAULT VALUES 
			RETURNING profile_id, username, last_seen, created_at`

	return r.db.GetContext(ctx, profile, sql)
}

func (r *ProfileRepository) GetByID(ctx context.Context, id int64) (*model.Profile, error) {
	// profile with biography struct
	pwb := new(struct {
		model.Profile
		model.Biography
	})

	query, args := squirrel.Select("*").
		From("profiles").
		Join("profile_biographies USING (profile_id)").
		Where(squirrel.Eq{"profile_id": id}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	if err := r.db.GetContext(ctx, pwb, query, args...); err != nil {
		return nil, err
	}

	pwb.Profile.Biography = pwb.Biography
	return &pwb.Profile, nil
}

func (r *ProfileRepository) GetByUsername(ctx context.Context, username string) (*model.Profile, error) {
	pwb := new(struct {
		model.Profile
		model.Biography
	})

	query, args := squirrel.Select("*").
		From("profiles").
		Join("profile_biographies USING (profile_id)").
		Where(squirrel.Eq{"username": username}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	if err := r.db.GetContext(ctx, pwb, query, args...); err != nil {
		return nil, err
	}

	pwb.Profile.Biography = pwb.Biography
	return &pwb.Profile, nil
}

func (r *ProfileRepository) UpdateProfileData(ctx context.Context, data *model.UpdateProfileData) error {
	query := squirrel.Update("profiles").
		Where(squirrel.Eq{"profile_id": data.ID}).
		Suffix("RETURNING profile_id").
		PlaceholderFormat(squirrel.Dollar)

	if data.Description != "" {
		query = query.Set("description", data.Description)
	}

	if data.Username != "" {
		query = query.Set("username", data.Username)
	}

	sql, args := query.MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&data.ID)
}
