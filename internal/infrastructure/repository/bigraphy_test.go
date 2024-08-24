package repository_test

import (
	"context"
	"testing"

	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func testingBiography() *model.Biography {
	return &model.Biography{
		ProfileID:  1,
		FirstName:  "first name",
		SecondName: "second name",
	}
}

func TestBiographyRepository_Create(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	bio := testingBiography()
	prf := testingProfile()

	br := repository.NewBiographyRepository(db)
	repository.NewProfileRepository(db).Create(context.Background(), prf)

	bio.ProfileID = prf.ID

	assert.NoError(t, br.Create(context.Background(), bio))
	assert.NotNil(t, bio.ID)
}
