package main

import (
	"fmt"
	"net/http"


	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"xs-coin-checker/internal/api"
)

func main (){
	log.SetReportCaller(true)

	router := chi.NewRouter()

	api.Handler(router)

	
	fmt.Println(`

		 ██████╗      ██████╗          █████╗     ██████╗     ██╗
		██╔════╝     ██╔═══██╗        ██╔══██╗    ██╔══██╗    ██║
		██║  ███╗    ██║   ██║        ███████║    ██████╔╝    ██║
		██║   ██║    ██║   ██║        ██╔══██║    ██╔═══╝     ██║
		╚██████╔╝    ╚██████╔╝        ██║  ██║    ██║         ██║
		 ╚═════╝      ╚═════╝         ╚═╝  ╚═╝    ╚═╝         ╚═╝
                                                         
	`)
	fmt.Println("XS-Coin-Checker API Running...")


	err := http.ListenAndServe("localhost:8000",router)
	if err != nil{
		log.Error(err)
	}

}