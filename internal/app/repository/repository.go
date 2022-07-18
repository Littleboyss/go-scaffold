package repository

import (
	"github.com/google/wire"
	"go-scaffold/internal/app/repository/role"
	"go-scaffold/internal/app/repository/user"
)

var ProviderSet = wire.NewSet(
	wire.NewSet(wire.Bind(new(role.RepositoryInterface), new(*role.Repository)), role.NewRepository),
	wire.NewSet(wire.Bind(new(user.RepositoryInterface), new(*user.Repository)), user.NewRepository),
)
