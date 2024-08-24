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

	br := repository.NewBiographyRepository(db)
	bio := testingBiography()

	assert.NoError(t, br.Create(context.Background(), bio))
	assert.NotNil(t, bio.ID)
}
