package avatar

import (
	"context"
	"log/slog"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
)

func (s *AvatarService) AddProfileAvatar(ctx context.Context, avatar *model.Avatar) error {
	const fn = "AddProfileAvatar"
	if err := s.avatarRepository.Create(ctx, avatar); err != nil {
		return handleError(err, fn, avatar.ProfileID)
	}

	if err := s.avatarRepository.SetAvatar(ctx, avatar.ProfileID, avatar.ID); err != nil {
		return handleError(err, fn, avatar.ProfileID)
	}

	return nil
}

func handleError(err error, fn string, payload interface{}) error {
	if pgerr, ok := err.(*pgconn.PgError); ok {

		if pgerr.Code == pgerrcode.ForeignKeyViolation {
			return status.NotFound("profile not found", "id", payload)
		}

		if pgerr.Code == pgerrcode.UniqueViolation {
			return status.Exists("profile already exists")
		}

		slog.Error("unexpected database error", "ctx", fn, "msg", pgerr.Message)
		return status.Internal(pgerr.Message)
	}

	slog.Error("unexpected", "ctx", fn, "msg", err.Error())
	return status.Internal(err.Error())
}
