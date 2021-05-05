package entity

import (
	"time"

	"go-clean-arch/src/domain"
)

type Entity struct {
	repo    domain.Repo
	timeout time.Duration
}

func InitEntity(a domain.Repo, t time.Duration) domain.Entity {
	return &Entity{
		repo:    a,
		timeout: t,
	}
}
