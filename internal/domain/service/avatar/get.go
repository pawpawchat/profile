package avatar

import (
	"context"
	"log/slog"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
)

func (s *AvatarService) GetAllProfileAvatars(ctx context.Context, profileID int64) (model.Avatars, error) {
	const fn = "GetAllProfileAvatars"

	avatars, err := s.avatarRepository.GetProfileAvatars(ctx, profileID)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {

			if pgerr.Code == pgerrcode.ForeignKeyViolation {
				return nil, status.NotFound("profile avatars not found", "id", profileID)
			}

			slog.Error("unexpected database error", "ctx", fn, "msg", pgerr.Message)
			return nil, status.Internal(pgerr.Message)
		}

		slog.Error("unexpected", "ctx", fn, "msg", err.Error())
		return nil, status.Internal(err.Error())
	}

	return avatars, nil
}
