package user

import "context"

type UseCase interface {
	GetUserById(ctx context.Context, id uint32) (interface{}, error)
}

type UserRepository interface {
	GetUserById(ctx context.Context, id uint32) (interface{}, error)
}
