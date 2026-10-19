package api

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"xs-coin-checker/internal/middleware"
)

func Handler(router *chi.Mux){
	//Global Middleware
	router.Use(chimiddle.StripSlashes)

	//account route and middleware
	router.Route("/account",func(route chi.Router){
		route.Use(middleware.Authorization)
		route.Get("/xcoins",GetxCoinBalance)
	})
}