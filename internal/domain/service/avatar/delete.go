package avatar

import (
	"context"
)

func (s *AvatarService) DeleteProfileAvatar(ctx context.Context, profileID int64, avatarID int64) error {
	const fn = "ChangeProfileAvatar"

	if err := s.avatarRepository.Delete(ctx, profileID, avatarID); err != nil {
		return handleError(err, fn, profileID)
	}

	return nil
}
