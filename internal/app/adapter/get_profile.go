package adapter

import (
	"context"
	"sync"

	"github.com/pawpawchat/profile/api/pb"
	"github.com/pawpawchat/profile/internal/domain/model"
	"github.com/pawpawchat/profile/pkg/status"
	"github.com/pawpawchat/profile/pkg/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileGetter interface {
	GetProfileByID(context.Context, int64) (*model.Profile, error)
	GetProfileByUsername(context.Context, string) (*model.Profile, error)
}

type AvatarGetter interface {
	GetAllProfileAvatars(context.Context, int64) ([]*model.Avatar, error)
}

func GetProfileAdapter(ctx context.Context, req *pb.GetProfileRequest, pg ProfileGetter, ag AvatarGetter) (*pb.GetProfileResponse, error) {
	if emptyFields := validation.GetZeroFields(req); len(emptyFields) != 0 {
		return nil, status.MissingFields(emptyFields)
	}

	var err1, err2 error
	var profile *model.Profile
	var avatars []*model.Avatar

	switch by := req.SearchBy.(type) {
	case *pb.GetProfileRequest_Id:
		var wg sync.WaitGroup
		cancelableCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		// fetch profile service
		wg.Add(1)
		go func() {
			defer wg.Done()
			profile, err1 = pg.GetProfileByID(cancelableCtx, by.Id)
			if err1 != nil {
				cancel()
			}
		}()

		// fetch avatar service
		wg.Add(1)
		go func() {
			defer wg.Done()
			avatars, err2 = ag.GetAllProfileAvatars(cancelableCtx, by.Id)
			if err2 != nil {
				cancel()
			}
		}()
		wg.Wait()
		// handle errors
		if err1 != nil {
			return nil, err1
		}
		if err2 != nil {
			return nil, err2
		}

	case *pb.GetProfileRequest_Username:
		profile, err1 = pg.GetProfileByUsername(ctx, by.Username)
		if err1 != nil {
			return nil, err1
		}
		avatars, err2 = ag.GetAllProfileAvatars(ctx, profile.ID)
		if err2 != nil {
			return nil, err2
		}
	}

	return modelToGetProfileResponse(profile, avatars)
}

func modelToGetProfileResponse(prf *model.Profile, avs []*model.Avatar) (*pb.GetProfileResponse, error) {
	avatars, selected := parseAvatarsModel(avs)
	profile := modelToProfile(prf)

	return &pb.GetProfileResponse{
		Profile:        profile,
		Avatars:        avatars,
		SelectedAvatar: selected,
	}, nil
}

func parseAvatarsModel(avs []*model.Avatar) ([]*pb.Avatar, *pb.Avatar) {
	var avatars []*pb.Avatar
	var selected *pb.Avatar

	for idx := range avs {
		avatars = append(avatars, modelToAvatar(avs[idx]))
		if avs[idx].IsSelected {
			selected = avatars[idx]
		}
	}

	return avatars, selected
}

func modelToAvatar(avatar *model.Avatar) *pb.Avatar {
	return &pb.Avatar{
		AvatarId: avatar.ID,
		OrigUrl:  avatar.OrigURL,
		AddedAt:  timestamppb.New(avatar.AddedAt),
	}
}

func modelToProfile(prf *model.Profile) *pb.Profile {
	var birthday *timestamppb.Timestamp

	if prf.Biography.Birthday != nil {
		birthday = timestamppb.New(*prf.Biography.Birthday)
	}

	return &pb.Profile{
		Id:          prf.ID,
		Username:    prf.Description,
		Description: prf.Description,
		NumFriends:  prf.NumFriends,
		Online:      prf.Online,
		LastSeen:    timestamppb.New(prf.LastSeen),
		CreatedAt:   timestamppb.New(prf.CreatedAt),
		Biography: &pb.Biography{
			FirstName:  prf.Biography.FirstName,
			SecondName: prf.Biography.SecondName,
			Birthday:   birthday,
		},
	}
}
