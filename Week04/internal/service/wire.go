// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package service

import (
	"dialogue/internal/biz"
	"dialogue/internal/data"

	"github.com/google/wire"
)

//go:generate
func InitDialogueService() (*DialogueService, func(), error) {
	panic(wire.Build(data.Provider, biz.Provider, NewDialogueService))
}
