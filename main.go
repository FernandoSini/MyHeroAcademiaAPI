package main

import (
	"MyHeroAcademiaApi/src/config"
	"MyHeroAcademiaApi/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	fmt.Println("running server")
	r := router.GenerateRoutes()

	fmt.Printf("server running on port %d", config.Porta)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
