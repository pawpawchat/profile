package repository_test

import (
	"context"
	"flag"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/profile/config"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func testingProfile() *model.Profile {
	return &model.Profile{}
}

func getTestingDB(t *testing.T) *sqlx.DB {
	flag.Set("env", "testing")
	flag.Parse()

	config, err := config.LoadDefaultConfig()
	assert.NoError(t, err)

	db, err := sqlx.Connect("pgx", config.Env().DB_URL)

	assert.NoError(t, err)
	return db
}
func TestProfileRepository_CreateProfile(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	r := repository.NewProfileRepository(db)

	testCases := []struct {
		desc    string
		valid   bool
		profile func() *model.Profile
	}{
		{
			"correct data",
			true,
			func() *model.Profile {
				return testingProfile()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			//exec query
			profile := tc.profile()
			err := r.Create(context.Background(), profile)

			// check result
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotZero(t, profile.ID, profile.Username, profile.CreatedAt, profile.LastSeen)
			case false:
				assert.Error(t, err)
			}
		})
	}
}

func TestProfileRepository_GetByID(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	r := repository.NewProfileRepository(db)
	br := repository.NewBiographyRepository(db)

	profile := testingProfile()
	r.Create(context.Background(), profile)

	biography := testingBiography()
	biography.ProfileID = profile.ID
	br.Create(context.Background(), biography)

	testCases := []struct {
		desc  string
		id    int64
		valid bool
	}{
		{"profile exists", profile.ID, true},
		{"profile not found", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// query
			profile, err := r.GetByID(context.Background(), tc.id)
			// check result
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotNil(t, profile)
				assert.NotZero(t, profile.Biography)
			case false:
				assert.Error(t, err)
			}
		})
	}
}

func TestProfileRepository_GetByUsername(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	r := repository.NewProfileRepository(db)
	br := repository.NewBiographyRepository(db)

	profile := testingProfile()
	r.Create(context.Background(), profile)

	biography := testingBiography()
	biography.ProfileID = profile.ID
	br.Create(context.Background(), biography)

	testCases := []struct {
		desc     string
		username string
		valid    bool
	}{
		{"profile exists", profile.Username, true},
		{"profile not found", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// query
			profile, err := r.GetByUsername(context.Background(), tc.username)
			// check result
			switch tc.valid {
			case true:
				assert.NoError(t, err)
				assert.NotNil(t, profile)
				assert.NotEmpty(t, profile.Biography)
			case false:
				assert.Error(t, err)
			}
		})
	}
}

func TestProfileRepository_DeleteProfileById(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	r := repository.NewProfileRepository(db)

	fmt.Println(r.DeleteProfileById(context.Background(), 2))
}

func TestProfileRepository_UpdateData(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	r := repository.NewProfileRepository(db)

	fmt.Println(r.UpdateProfileData(context.Background(), &model.UpdateProfileData{ID: -1, Username: "s"}))
}
