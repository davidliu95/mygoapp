// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"mygoapp/kratos-demo/internal/dao"
	"mygoapp/kratos-demo/internal/server/grpc"
	"mygoapp/kratos-demo/internal/server/http"
	"mygoapp/kratos-demo/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
