package dao

import "github.com/google/wire"

var DaoProviderSet = wire.NewSet(
	NewUserDao,
)
