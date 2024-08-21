package repository_test

import (
	"context"
	"flag"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/config"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func testProfile() *model.Profile {
	return &model.Profile{
		Username: "some-username",
		Biography: model.Biography{
			FirstName:  "first-name",
			SecondName: "second-name",
		},
	}
}

func TestCreateProfile(t *testing.T) {
	db, r := getTestDbContext(t)
	defer db.Close()

	testCases := []struct {
		desc    string
		valid   bool
		profile func() *model.Profile
	}{
		{
			"correct data",
			true,
			func() *model.Profile {
				profile := testProfile()
				return profile
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			//exec query
			profile := tc.profile()
			err := r.CreateProfile(context.Background(), profile)

			// check result
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotZero(t, profile.Id, profile.Username, profile.Biography.FirstName, profile.Biography.SecondName, profile.CreatedAt)
			case false:
				assert.Error(t, err)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	db, r := getTestDbContext(t)
	defer db.Close()

	testCases := []struct {
		desc  string
		id    uint64
		valid bool
	}{
		{"profile exists", 1, true},
		// {"profile not found", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// query
			_, err := r.GetById(context.Background(), tc.id)
			// check result
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				// assert.NotNil(t, res)
				// assert.NotZero(t, res.Biography)
			case false:
				assert.Error(t, err)
			}
		})
	}
}

func getTestDbContext(t *testing.T) (*sqlx.DB, *repository.ProfileRepository) {
	flag.Set("env", "testing")
	flag.Parse()

	config, err := config.LoadConfig("../../../config.yaml")
	assert.NoError(t, err)

	db, err := sqlx.Connect("pgx", config.Env().DB_URL)

	assert.NoError(t, err)
	r := repository.NewProfile(db)

	return db, r
}
