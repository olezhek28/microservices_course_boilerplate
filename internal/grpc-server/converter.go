package grpc_server

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	usecases "github.com/neracastle/auth/internal/usecases/models"
	"github.com/neracastle/auth/pkg/user_v1"
)

// FromGrpcToCreateUsecase преобразует grpc-запрос в дто сервисного слоя
func FromGrpcToCreateUsecase(req *user_v1.CreateRequest) usecases.CreateDTO {
	dto := usecases.CreateDTO{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Name:            req.Name,
		IsAdmin:         false,
	}

	if req.Role == user_v1.Role_ADMIN {
		dto.IsAdmin = true
	}

	return dto
}

// FromGrpcToUpdateUsecase преобразует grpc-запрос в дто сервисного слоя
func FromGrpcToUpdateUsecase(req *user_v1.UpdateRequest) usecases.UpdateDTO {
	dto := usecases.UpdateDTO{
		ID:    req.Id,
		Email: req.GetEmail(),
		Name:  req.GetName(),
		Role:  0,
	}

	switch req.GetRole() {
	case user_v1.Role_ADMIN:
		dto.Role = 1
	case user_v1.Role_USER:
		dto.Role = 2
	}

	return dto
}

// FromUsecaseToGetResponse преобразует дто сервисного слоя в grpc-ответ
func FromUsecaseToGetResponse(dto usecases.UserDTO) *user_v1.GetResponse {
	rsp := &user_v1.GetResponse{
		Id:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
		Role:      user_v1.Role_USER,
		CreatedAt: timestamppb.New(dto.CreatedAt),
	}

	if dto.IsAdmin {
		rsp.Role = user_v1.Role_ADMIN
	}

	return rsp
}
