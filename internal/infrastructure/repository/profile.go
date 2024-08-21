package repository

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type ProfileRepository struct {
	db *sqlx.DB
}

func NewProfile(db *sqlx.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return rollback(err, tx)
	}

	query := `INSERT INTO profiles DEFAULT VALUES RETURNING username, profile_id, last_seen, created_at`

	if err := r.db.Get(profile, query); err != nil {
		slog.Error("error inserting profile", "err", err)
		return rollback(err, tx)
	}

	query, args, err := sq.Insert("profile_biographies").
		Columns("profile_id", "first_name", "second_name").
		Values(profile.Id, profile.Biography.FirstName, profile.Biography.SecondName).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		slog.Error("error inserting bigraphy", "err", err)
		return rollback(err, tx)
	}

	if err := tx.Commit(); err != nil {
		slog.Error("error commiting create profile tx", "err", err)
		return rollback(err, tx)
	}

	return nil
}

func (r *ProfileRepository) GetById(ctx context.Context, id uint64) (*model.Profile, error) {
	// тип, для заполнения данных с join запроса
	pwb := new(struct {
		model.Profile
		model.Biography
	})

	query, args, err := sq.Select("*").
		From("profiles").
		Join("profile_biographies USING (profile_id)").
		Where(sq.Eq{"profile_id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := r.db.Get(pwb, query, args...); err != nil {
		return nil, err
	}

	pwb.Profile.Biography = pwb.Biography

	return &pwb.Profile, nil
}

func rollback(e error, tx *sqlx.Tx) error {
	if tx != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}
	return e
}
