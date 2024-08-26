package profile

import (
	"context"

	"github.com/pawpawchat/profile/pkg/status"
)

func (s *ProfileService) DeleteProfile(ctx context.Context, profileID int64) error {
	return status.Internal("delete method doesn't imlpemented")
}
