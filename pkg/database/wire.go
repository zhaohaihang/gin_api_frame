package database

import "github.com/google/wire"

var DatabaseProviderSet = wire.NewSet(NewDatabase,NewGormLogger)
