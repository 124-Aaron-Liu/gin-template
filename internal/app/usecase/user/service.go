package user

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.SugaredLogger
	//userRepo UserRepository
}

func NewService(
	logger *zap.SugaredLogger,
// userRepo UserRepository,
) Service {
	return Service{
		logger,
		//userRepo,
	}
}

func (s Service) GetUserById(ctx context.Context, id uint32) (interface{}, error) {
	//user, err := s.userRepo.GetUserById(ctx, pfID)
	//if err != nil {
	//	return nil, err
	//}
	return []string{"user id is", fmt.Sprint(id)}, nil
}

var ProviderSet = wire.NewSet(NewService)
