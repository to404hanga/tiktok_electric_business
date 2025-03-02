package grpc

import (
	"context"
	userv1 "tiktok_electric_business/api/proto/gen/user/v1"
	"tiktok_electric_business/user/domain"
	"tiktok_electric_business/user/service"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceServer struct {
	userv1.UnimplementedUserServiceServer
	svc service.UserService
}

func NewUserServiceServer(svc service.UserService) *UserServiceServer {
	return &UserServiceServer{svc: svc}
}

func (u *UserServiceServer) Register(srv grpc.ServiceRegistrar) {
	userv1.RegisterUserServiceServer(srv, u)
}

func (u *UserServiceServer) FindOrCreateByPhone(ctx context.Context, req *userv1.FindOrCreateByPhoneRequest) (*userv1.FindOrCreateByPhoneResponse, error) {
	user, err := u.svc.FindOrCreateByPhone(ctx, req.GetPhone())
	return &userv1.FindOrCreateByPhoneResponse{
		User: u.convertToView(user),
	}, err
}

func (u *UserServiceServer) FindOrCreateByWechat(ctx context.Context, req *userv1.FindOrCreateByWechatRequest) (*userv1.FindOrCreateByWechatResponse, error) {
	user, err := u.svc.FindOrCreateByWechat(ctx, domain.WechatInfo{
		OpenId:  req.GetInfo().GetOpenId(),
		UnionId: req.GetInfo().GetUnionId(),
	})
	return &userv1.FindOrCreateByWechatResponse{
		User: u.convertToView(user),
	}, err
}

func (u *UserServiceServer) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	user, err := u.svc.Login(ctx, req.GetEmail(), req.GetPassword())
	return &userv1.LoginResponse{
		User: u.convertToView(user),
	}, err
}

func (u *UserServiceServer) Profile(ctx context.Context, req *userv1.ProfileRequest) (*userv1.ProfileResponse, error) {
	user, err := u.svc.Profile(ctx, req.GetId())
	return &userv1.ProfileResponse{
		User: u.convertToView(user),
	}, err
}

func (u *UserServiceServer) SignUp(ctx context.Context, req *userv1.SignUpRequest) (*userv1.SignUpResponse, error) {
	err := u.svc.SignUp(ctx, u.convertToDomain(req.GetUser()))
	return &userv1.SignUpResponse{}, err
}

func (u *UserServiceServer) UpdateNonSensitiveInfo(ctx context.Context, req *userv1.UpdateNonSensitiveInfoRequest) (*userv1.UpdateNonSensitiveInfoResponse, error) {
	err := u.svc.UpdateNonSensitiveInfo(ctx, u.convertToDomain(req.GetUser()))
	return &userv1.UpdateNonSensitiveInfoResponse{}, err
}

func (u *UserServiceServer) convertToView(user domain.User) *userv1.User {
	return &userv1.User{
		Id:        user.Id,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Phone:     user.Phone,
		UserType:  uint32(user.UserType),
		CreatedAt: timestamppb.New(user.CreatedAt),
		WechatInfo: &userv1.WechatInfo{
			OpenId:  user.WechatInfo.OpenId,
			UnionId: user.WechatInfo.UnionId,
		},
	}
}

func (u *UserServiceServer) convertToDomain(user *userv1.User) domain.User {
	res := domain.User{}
	if user != nil {
		res.Id = user.GetId()
		res.Email = user.GetEmail()
		res.Nickname = user.GetNickname()
		res.Password = user.GetPassword()
		res.Phone = user.GetPhone()
		res.UserType = domain.UserType(user.GetUserType())
		res.CreatedAt = user.GetCreatedAt().AsTime()
		if user.WechatInfo != nil {
			res.WechatInfo = domain.WechatInfo{
				OpenId:  user.GetWechatInfo().GetOpenId(),
				UnionId: user.GetWechatInfo().GetUnionId(),
			}
		}
	}
	return res
}
