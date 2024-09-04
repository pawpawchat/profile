package avatar

import (
	"context"
)

func (s *AvatarService) ChangeProfileAvatar(ctx context.Context, profileID int64, avatarID int64) error {
	const fn = "ChangeProfileAvatar"

	if err := s.avatarRepository.Unselect(ctx, profileID); err != nil {
		return handleError(err, fn, profileID)
	}

	if err := s.avatarRepository.Select(ctx, profileID, avatarID); err != nil {
		return handleError(err, fn, profileID)
	}

	return nil
}
