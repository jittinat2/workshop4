package app

import (
	dbpkg "workshop4/internal/db"
	"workshop4/internal/entity"
	repo "workshop4/internal/repository"
	transport "workshop4/internal/transport/http"
	use "workshop4/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App  *fiber.App
	Addr string
}

func New() (*Server, error) {
	if err := dbpkg.Init("data.db"); err != nil {
		return nil, err
	}
	// automigrate
	if err := dbpkg.DB.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}

	r := repo.NewSQLiteUserRepo()
	u := use.NewUserUsecase(r)
	h := transport.New(u)

	app := fiber.New()
	// register under /api/v1/users (existing) and also expose flat /users endpoints
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")
	h.Register(users)

	// also expose endpoints at top-level /users as requested
	topUsers := app.Group("/users")
	h.Register(topUsers)

	return &Server{App: app, Addr: ":3000"}, nil
}

func (s *Server) Listen() error {
	return s.App.Listen(s.Addr)
}
