package data

import "github.com/google/wire"

var Provider = wire.NewSet(NewDB, NewCache, NewES)
