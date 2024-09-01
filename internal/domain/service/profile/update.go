package profile

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
)

func (s *ProfileService) UpdateProfile(ctx context.Context, data *model.UpdateProfileData) error {

	if data.ID == 0 {
		return status.MissingFields([]string{"profile_id"})
	}

	if data.Biography != nil {

		if err := s.biographyRepository.UpdateProfileBiography(ctx, data); err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				return status.NotFound("profile not found", "id", data.ID)
			}

			slog.Error("specific database error", "msg", err.Error())
			return status.Unexpected(err.Error())
		}

	}

	if data.Username != "" || data.Description != "" {

		if err := s.profileRepository.UpdateProfileData(ctx, data); err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				return status.NotFound("profile not found", "id", data.ID)
			}

			slog.Error("specific database error", "msg", err.Error())
			return status.Unexpected(err.Error())
		}
	}

	return nil
}
