package main

import (
	"context"
)

type UserRepo interface {
	NewUser(ctx context.Context, name string, age int, gender string) (User, error)
}

type UserRepost struct {
	userrepo UserRepo
}

func NewUserRepo(repo UserRepo) *UserRepost {
	return &UserRepost{repo}
}

func (u *UserRepost) Prin(ctx context.Context, name string, age int, gender string) (User, error) {
	return u.userrepo.NewUser(ctx, name+"busness", age, gender)
}
