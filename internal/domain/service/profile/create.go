package profile

import (
	"context"
	"log/slog"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
)

func (s *ProfileService) CreateProfile(ctx context.Context, profile *model.Profile) error {
	if err := s.profileRepository.Create(ctx, profile); err != nil {
		slog.Error("database", "error", err)
		return status.Internal(err.Error())
	}

	profile.Biography.ProfileID = profile.ID

	if err := s.biographyRepository.Create(ctx, &profile.Biography); err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {

			if pgerr.Code == pgerrcode.UniqueViolation {
				return status.Exists("profile biography already exists")
			}

			if pgerr.Code == pgerrcode.ConnectionException {
				return status.Internal(err.Error())
			}

		}
		slog.Error("specific database error", "msg", err.Error())
		return status.Unexpected(err.Error())
	}

	return nil
}
