package service

import (
	"restapi/model"
	"restapi/repository"

	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/shared/schema"
)

type UserService interface {
	CreateUser(ctx *appcontext.AppContext, userRequest model.UserRequest) (model.UserResponse, error)
	GetUser(ctx *appcontext.AppContext, id string) (model.UserResponse, error)
	GetUserByEmail(ctx *appcontext.AppContext, email string) (model.UserResponse, error)
	UpdateUser(ctx *appcontext.AppContext, id string, userRequest model.UserRequest) (model.UserResponse, error)
	DeleteUser(ctx *appcontext.AppContext, id string) error
	ListUser(ctx *appcontext.AppContext) ([]model.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx *appcontext.AppContext, userRequest model.UserRequest) (model.UserResponse, error) {
	user := schema.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	if err := s.userRepo.Create(ctx, &user); err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		ID:    user.ID.Hex(),
		Email: user.Email,
	}, nil
}

func (s *userService) GetUser(ctx *appcontext.AppContext, id string) (model.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		ID:    user.ID.Hex(),
		Email: user.Email,
	}, nil
}

func (s *userService) GetUserByEmail(ctx *appcontext.AppContext, email string) (model.UserResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		ID:    user.ID.Hex(),
		Email: user.Email,
	}, nil
}

func (s *userService) UpdateUser(ctx *appcontext.AppContext, id string, userRequest model.UserRequest) (model.UserResponse, error) {
	user := schema.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	if err := s.userRepo.Update(ctx, &user); err != nil {
		return model.UserResponse{}, err
	}
	return model.UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
}

func (s *userService) DeleteUser(ctx *appcontext.AppContext, id string) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *userService) ListUser(ctx *appcontext.AppContext) ([]model.UserResponse, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return []model.UserResponse{}, err
	}
	var userResponses []model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, model.UserResponse{
			ID:    user.ID.Hex(),
			Email: user.Email,
		})
	}
	return userResponses, nil
}
