package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func testingAvatar() *model.Avatar {
	return &model.Avatar{
		OrigURL: "https://yandex.ru/dickpic",
		AddedAt: time.Now(),
	}
}

func TestAvatarRepository_Create(t *testing.T) {
	db := getTestingDB(t)
	defer db.Close()

	ar := repository.NewAvatarsRepository(db)
	avatar := testingAvatar()

	assert.NoError(t, ar.Create(context.Background(), avatar))
	assert.NotNil(t, avatar.ID)

}
