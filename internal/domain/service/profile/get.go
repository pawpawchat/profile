package profile

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
)

func (s *ProfileService) GetProfileByUsername(ctx context.Context, username string) (*model.Profile, error) {
	const fn = "GetProfileByUsername"
	profile, err := s.profileRepository.GetByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.NotFound("profile not found", "username", username)
		}

		slog.Error("unexpected", "ctx", fn, "msg", err.Error())
		return nil, status.Internal("internal server error")
	}

	return profile, nil
}

func (s *ProfileService) GetProfileByID(ctx context.Context, id int64) (*model.Profile, error) {
	const fn = "GetProfileByID"

	prf, err := s.profileRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.NotFound("profile not found", "id", id)
		}

		slog.Error("unexpected", "ctx", fn, "msg", err.Error())
		return nil, status.Internal(err.Error())
	}

	return prf, nil
}
