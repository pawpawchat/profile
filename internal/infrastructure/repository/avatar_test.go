package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func testingAvatar() *model.Avatar {
	return &model.Avatar{
		OrigURL:   "https://yandex.ru/example.png",
		ProfileID: 1,
		AddedAt:   time.Now(),
	}
}

func addTestingProfileWithAvatar(t *testing.T) (*sqlx.DB, *model.Profile, *model.Avatar) {
	db := getTestingDB(t)

	p := testingProfile()
	a := testingAvatar()

	pr := repository.NewProfileRepository(db)
	assert.NoError(t, pr.Create(context.Background(), p))

	a.ProfileID = p.ID
	ar := repository.NewAvatarsRepository(db)
	assert.NoError(t, ar.Create(context.Background(), a))
	assert.NoError(t, ar.SetAvatar(context.Background(), a.ProfileID, a.ID))

	return db, p, a
}

func TestAvatarRepository_Create(t *testing.T) {
	db, p, _ := addTestingProfileWithAvatar(t)
	defer db.Close()

	testCases := []struct {
		desc   string
		avatar func() *model.Avatar
		valid  bool
	}{
		{
			"avatar and user exist",
			func() *model.Avatar {
				av := testingAvatar()
				av.ProfileID = p.ID
				return av
			},
			true,
		},
		{
			"profile not found",
			func() *model.Avatar {
				av := testingAvatar()
				av.ProfileID = 0
				return av
			},
			false,
		},
	}

	ar := repository.NewAvatarsRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			avatar := tc.avatar()
			err := ar.Create(context.Background(), avatar)

			if tc.valid {
				assert.NoError(t, err)
				assert.NotZero(t, avatar.ID)

			} else {
				assert.Error(t, err)

			}
		})
	}
}

func TestAvatarRepository_SetAvatar(t *testing.T) {
	db, p, a := addTestingProfileWithAvatar(t)
	defer db.Close()

	testCases := []struct {
		desc      string
		profileID int64
		avatarID  int64
		valid     bool
	}{
		{"avatar and user exist", p.ID, a.ID, true},
		{"profile not found", 0, a.ID, false},
	}

	ar := repository.NewAvatarsRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := ar.SetAvatar(context.Background(), tc.profileID, tc.avatarID)

			if tc.valid {
				assert.NoError(t, err)

			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestAvatarRepository_GetProfileAvatars(t *testing.T) {
	db, p, _ := addTestingProfileWithAvatar(t)
	defer db.Close()

	testCases := []struct {
		desc      string
		profileID int64
		valid     bool
	}{
		{"profile exists", p.ID, true},
		{"profile doesn't exist", 0, false},
	}

	ar := repository.NewAvatarsRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			avatars, err := ar.GetProfileAvatars(context.Background(), tc.profileID)

			if tc.valid {
				assert.NoError(t, err)
				assert.NotEmpty(t, avatars)

			} else {
				assert.NoError(t, err)
				assert.Empty(t, avatars)
			}
		})
	}
}

func TestAvatarRepository_DeleteProfileAvatar(t *testing.T) {
	db, p, a := addTestingProfileWithAvatar(t)
	defer db.Close()

	testCases := []struct {
		desc      string
		profileID int64
		avatarID  int64
		valid     bool
	}{
		{"delete one profile avatars", p.ID, a.ID, true},
		{"profile doesn't exist", 0, 0, false},
	}

	ar := repository.NewAvatarsRepository(db)
	// ar.Create(context.Background(), a)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			err := ar.DeleteProfileAvatar(context.Background(), tc.profileID, tc.avatarID)

			if tc.valid {
				assert.NoError(t, err)

			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestAvatarRepository_DeleteProfileAvatars(t *testing.T) {
	db, p, _ := addTestingProfileWithAvatar(t)
	defer db.Close()

	testCases := []struct {
		desc      string
		profileID int64
		valid     bool
	}{
		{"profile exists", p.ID, true},
		{"profile doesn't exist", 0, false},
	}

	ar := repository.NewAvatarsRepository(db)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			err := ar.DeleteProfileAvatars(context.Background(), tc.profileID)

			if tc.valid {
				assert.NoError(t, err)

			} else {
				assert.Error(t, err)
			}
		})
	}
}
