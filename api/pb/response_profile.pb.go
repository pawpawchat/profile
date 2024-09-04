// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: response_profile.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CreateProfileResponse) Reset() {
	*x = CreateProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileResponse) ProtoMessage() {}

func (x *CreateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileResponse.ProtoReflect.Descriptor instead.
func (*CreateProfileResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProfileResponse) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile  `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Avatars []*Avatar `protobuf:"bytes,2,rep,name=avatars,proto3" json:"avatars,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *GetProfileResponse) GetAvatars() []*Avatar {
	if x != nil {
		return x.Avatars
	}
	return nil
}

type UpdateProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProfileResponse) Reset() {
	*x = UpdateProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileResponse) ProtoMessage() {}

func (x *UpdateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{2}
}

type DeleteProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProfileResponse) Reset() {
	*x = DeleteProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileResponse) ProtoMessage() {}

func (x *DeleteProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileResponse.ProtoReflect.Descriptor instead.
func (*DeleteProfileResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{3}
}

type AddProfileAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar *Avatar `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *AddProfileAvatarResponse) Reset() {
	*x = AddProfileAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProfileAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProfileAvatarResponse) ProtoMessage() {}

func (x *AddProfileAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProfileAvatarResponse.ProtoReflect.Descriptor instead.
func (*AddProfileAvatarResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{4}
}

func (x *AddProfileAvatarResponse) GetAvatar() *Avatar {
	if x != nil {
		return x.Avatar
	}
	return nil
}

type DeleteProfileAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProfileAvatarResponse) Reset() {
	*x = DeleteProfileAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileAvatarResponse) ProtoMessage() {}

func (x *DeleteProfileAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileAvatarResponse.ProtoReflect.Descriptor instead.
func (*DeleteProfileAvatarResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{5}
}

type ChangeProfileAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar *Avatar `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *ChangeProfileAvatarResponse) Reset() {
	*x = ChangeProfileAvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeProfileAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeProfileAvatarResponse) ProtoMessage() {}

func (x *ChangeProfileAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_response_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeProfileAvatarResponse.ProtoReflect.Descriptor instead.
func (*ChangeProfileAvatarResponse) Descriptor() ([]byte, []int) {
	return file_response_profile_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeProfileAvatarResponse) GetAvatar() *Avatar {
	if x != nil {
		return x.Avatar
	}
	return nil
}

var File_response_profile_proto protoreflect.FileDescriptor

var file_response_profile_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x62, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x6f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x07, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x1d, 0x0a, 0x1b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x0a, 0x1b, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x08, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_response_profile_proto_rawDescOnce sync.Once
	file_response_profile_proto_rawDescData = file_response_profile_proto_rawDesc
)

func file_response_profile_proto_rawDescGZIP() []byte {
	file_response_profile_proto_rawDescOnce.Do(func() {
		file_response_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_profile_proto_rawDescData)
	})
	return file_response_profile_proto_rawDescData
}

var file_response_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_response_profile_proto_goTypes = []interface{}{
	(*CreateProfileResponse)(nil),       // 0: profilepb.CreateProfileResponse
	(*GetProfileResponse)(nil),          // 1: profilepb.GetProfileResponse
	(*UpdateProfileResponse)(nil),       // 2: profilepb.UpdateProfileResponse
	(*DeleteProfileResponse)(nil),       // 3: profilepb.DeleteProfileResponse
	(*AddProfileAvatarResponse)(nil),    // 4: profilepb.AddProfileAvatarResponse
	(*DeleteProfileAvatarResponse)(nil), // 5: profilepb.DeleteProfileAvatarResponse
	(*ChangeProfileAvatarResponse)(nil), // 6: profilepb.ChangeProfileAvatarResponse
	(*Profile)(nil),                     // 7: profilepb.Profile
	(*Avatar)(nil),                      // 8: profilepb.Avatar
}
var file_response_profile_proto_depIdxs = []int32{
	7, // 0: profilepb.CreateProfileResponse.profile:type_name -> profilepb.Profile
	7, // 1: profilepb.GetProfileResponse.profile:type_name -> profilepb.Profile
	8, // 2: profilepb.GetProfileResponse.avatars:type_name -> profilepb.Avatar
	8, // 3: profilepb.AddProfileAvatarResponse.avatar:type_name -> profilepb.Avatar
	8, // 4: profilepb.ChangeProfileAvatarResponse.avatar:type_name -> profilepb.Avatar
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_response_profile_proto_init() }
func file_response_profile_proto_init() {
	if File_response_profile_proto != nil {
		return
	}
	file_messages_profile_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_response_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_response_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_response_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_response_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_response_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProfileAvatarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_response_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileAvatarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_response_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeProfileAvatarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_profile_proto_goTypes,
		DependencyIndexes: file_response_profile_proto_depIdxs,
		MessageInfos:      file_response_profile_proto_msgTypes,
	}.Build()
	File_response_profile_proto = out.File
	file_response_profile_proto_rawDesc = nil
	file_response_profile_proto_goTypes = nil
	file_response_profile_proto_depIdxs = nil
}
