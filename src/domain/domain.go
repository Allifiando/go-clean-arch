package domain

import "context"

type Entity interface {
	Login(ctx context.Context, b Login) (res UserLogin, err error)
	GetUserByUuid(ctx context.Context, uid string) (res User, err error)
}
type Repo interface {
	Login(ctx context.Context, b Login) (res UserLogin, err error)
	GetUserByUuid(ctx context.Context, uid string) (res User, err error)
}
