package v1

import "github.com/google/wire"

var ControllerProviderSet = wire.NewSet(
	NewUserContrller,
)
