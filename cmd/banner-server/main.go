package main

import (
	"backend-trainee-assignment-2024/internal/config"
	"backend-trainee-assignment-2024/internal/handler"
	"backend-trainee-assignment-2024/internal/middleware"
	"backend-trainee-assignment-2024/internal/service"
	"backend-trainee-assignment-2024/internal/storage/postgres"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func registerRouts(cfg *config.Config, router *chi.Mux, bannerHandler *handler.BannerHandler) {
	router.Use(middleware.UserAndAdminAuth(cfg.Server.UserToken, cfg.Server.AdminToken))
	router.Get("/user_banner", bannerHandler.GetBanner)

	// Admin allowed methods.
	router.Route("/", func(r chi.Router) {
		r.Use(middleware.AdminAuth(cfg.Server.AdminToken))

		router.Post("/banner", bannerHandler.CreateBanner)
		router.Get("/banner", bannerHandler.GetFilteredBannerList)
		router.Patch("/banner/{id:[0-9]+}", bannerHandler.UpdateBanner)
		router.Delete("/banner/{id:[0-9]+}", bannerHandler.DeleteBanner)
	})

}

func main() {
	cfg := config.MustLoad()

	database, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("can't init database: %s", err)
	}

	bannerRepo := postgres.NewRepository(database)
	//TODO: init cache: redis

	banesService := service.NewBannerService(bannerRepo, nil)
	bannerHandler := handler.NewBannerHandler(banesService)

	router := chi.NewRouter()

	registerRouts(cfg, router, bannerHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal("failed to start server")
	}
}
