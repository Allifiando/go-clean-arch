package entity

import (
	"context"
	"go-clean-arch/src/domain"
)

func (a *Entity) Login(c context.Context, b domain.Login) (res domain.UserLogin, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.Login(ctx, b)
	if err != nil {
		return res, err
	}
	return
}

func (a *Entity) GetUserByUuid(c context.Context, uid string) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, a.timeout)
	defer cancel()

	res, err = a.repo.GetUserByUuid(ctx, uid)
	if err != nil {
		return res, err
	}
	return
}
