package model

import (
	"github.com/pawpawchat/profile/api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (p *Profile) ToPb() *pb.Profile {
	var birthday *timestamppb.Timestamp

	if p.Biography.Birthday != nil {
		birthday = timestamppb.New(*p.Biography.Birthday)
	}

	return &pb.Profile{
		Id:          p.ID,
		Username:    p.Username,
		Description: p.Description,
		NumFriends:  p.NumFriends,
		Online:      p.Online,
		LastSeen:    timestamppb.New(p.LastSeen),
		CreatedAt:   timestamppb.New(p.CreatedAt),
		Biography: &pb.Biography{
			FirstName:  p.Biography.FirstName,
			SecondName: p.Biography.SecondName,
			Birthday:   birthday,
		},
	}
}

func (a Avatars) ToPb() []*pb.Avatar {
	if a != nil && len(a) == 0 || a == nil {
		return nil
	}
	var avatars []*pb.Avatar
	for _, avatar := range a {
		avatars = append(avatars, avatar.ToPb())
	}
	return avatars
}

func (a *Avatar) ToPb() *pb.Avatar {
	return &pb.Avatar{
		AvatarId: a.ID,
		OrigUrl:  a.OrigURL,
		AddedAt:  timestamppb.New(a.AddedAt),
	}
}
