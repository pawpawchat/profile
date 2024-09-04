package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/internal/domain/model"
)

type AvatarRepository struct {
	db *sqlx.DB
}

func NewAvatarsRepository(db *sqlx.DB) *AvatarRepository {
	return &AvatarRepository{db}
}

func (r *AvatarRepository) Create(ctx context.Context, avatar *model.Avatar) error {
	sql, args := squirrel.Insert("profile_avatars").
		Columns("orig_url", "profile_id", "added_at").
		Values(avatar.OrigURL, avatar.ProfileID, avatar.AddedAt).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING avatar_id").
		MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&avatar.ID)
}

func (r *AvatarRepository) Select(ctx context.Context, profileID int64, avatarID int64) error {
	sql, args := squirrel.Update("profile_avatars").
		Set("is_selected", "true").
		Where(squirrel.Eq{"profile_id": profileID, "avatar_id": avatarID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING avatar_id").
		MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&avatarID)
}

func (r *AvatarRepository) Unselect(ctx context.Context, profileID int64) error {
	sql, args := squirrel.Update("profile_avatars").
		Set("is_selected", "false").
		Where(squirrel.Eq{"profile_id": profileID, "is_selected": true}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING profile_id").
		MustSql()

	return r.db.QueryRowContext(ctx, sql, args...).Scan(&profileID)
}

func (r *AvatarRepository) GetAll(ctx context.Context, profileID int64) ([]*model.Avatar, error) {
	sql, args := squirrel.Select("*").
		From("profile_avatars").
		Where(squirrel.Eq{"profile_id": profileID}).
		OrderBy("is_selected").Suffix("DESC").
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	avatars := make([]*model.Avatar, 0)
	return avatars, r.db.Select(&avatars, sql, args...)
}

func (r *AvatarRepository) Delete(ctx context.Context, profileID int64, avatarID int64) error {
	sql, args := squirrel.Delete("profile_avatars").
		Where(squirrel.Eq{"profile_id": profileID, "avatar_id": avatarID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING avatar_id").
		MustSql()

	return r.db.QueryRow(sql, args...).Scan(&avatarID)
}

func (r *AvatarRepository) DeleteAll(ctx context.Context, profileID int64) error {
	sql, args := squirrel.Delete("profile_avatars").
		Where(squirrel.Eq{"profile_id": profileID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING profile_id").
		MustSql()

	return r.db.QueryRow(sql, args...).Scan(&profileID)
}
