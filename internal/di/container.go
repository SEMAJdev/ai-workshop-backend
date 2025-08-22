package di

import (
	"time"

	jwtauth "workshop-cursor/backend/internal/adapter/auth/jwt"
	"workshop-cursor/backend/internal/adapter/http/handler"
	memrepo "workshop-cursor/backend/internal/adapter/repository/memory"
	sqliterepo "workshop-cursor/backend/internal/adapter/repository/sqlite"
	"workshop-cursor/backend/internal/config"
	usercore "workshop-cursor/backend/internal/core/user"
	dbprovider "workshop-cursor/backend/internal/infra/sqlite"
	authuc "workshop-cursor/backend/internal/usecase/auth"
	hellouc "workshop-cursor/backend/internal/usecase/hello"

	"golang.org/x/crypto/bcrypt"
)

type Container struct {
	HelloHandler *handler.HelloHandler
	AuthHandler  *handler.AuthHandler
	AuthMW       func(token string) (int64, error)
}

func NewContainer(cfg *config.Config) *Container {
	// repositories
	greeterRepo := memrepo.NewInMemoryGreeterRepository()

	// database
	db, err := dbprovider.Open(cfg.SQLitePath)
	if err != nil {
		panic(err)
	}
	userRepo := sqliterepo.NewSQLiteUserRepository(db)
	if err := userRepo.InitSchema(); err != nil {
		panic(err)
	}

	// seed user if empty
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	seed := &usercore.User{
		Email:           "somchai@example.com",
		PasswordHash:    string(hash),
		FirstName:       "สมชาย",
		LastName:        "ใจดี",
		Phone:           "081-234-5678",
		MemberCode:      "LBK001234",
		MembershipLevel: "Gold",
		Points:          15420,
		JoinedAt:        time.Now().AddDate(-1, 0, 0),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := userRepo.SeedInitialUserIfEmpty(seed); err != nil {
		panic(err)
	}

	// use cases
	helloUC := hellouc.NewUseCase(greeterRepo)
	signer := jwtauth.NewHS256Signer(cfg.JWTSecret)
	authUC := authuc.NewUseCase(userRepo, signer)

	// handlers
	helloHandler := handler.NewHelloHandler(helloUC)
	authHandler := handler.NewAuthHandler(authUC)

	return &Container{
		HelloHandler: helloHandler,
		AuthHandler:  authHandler,
		AuthMW:       signer.Verify,
	}
}
