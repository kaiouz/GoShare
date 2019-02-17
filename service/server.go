package service

import (
	"fmt"
	"github.com/kaiouz/GoShare/res"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Server interface {
	Start() error
}

type server struct {
	port    int
	res     []res.Resource
	resMap  map[string]res.Resource
	typeMap map[res.ResourceType][]res.Resource
	app     *iris.Application
}

func (s *server) Start() error {
	s.app = iris.Default()

	resMap := make(map[string]res.Resource, 0)
	typeMap := make(map[res.ResourceType][]res.Resource, 0)

	for _, resource := range s.res {
		resMap[resource.Id] = resource
		typeMap[resource.ResType] = append(typeMap[resource.ResType], resource)
	}
	s.resMap = resMap
	s.typeMap = typeMap

	s.app.Get("/resources", func(ctx context.Context) {
		resType := ctx.URLParam("type")
		if resType != "" {
			ctx.JSON(s.typeMap[res.ResourceType(resType)])
		} else {
			ctx.JSON(s.res)
		}
	})

	s.app.Get("resources/{id:string}", func(ctx context.Context) {
		id := ctx.Params().Get("id")
		if resource, ok := s.resMap[id]; ok {
			ctx.SendFile(resource.Path, resource.Name)
		} else {
			ctx.NotFound()
		}
	})

	return s.app.Run(iris.Addr(fmt.Sprintf(":%d", s.port)))
}

func CreateServer(port int, resources []res.Resource) Server {
	s := &server{
		res:  resources,
		port: port,
	}
	return s
}
