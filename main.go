//Config API
//
//    Title: Config API
//
//    Schemes: http
//    Version: 0.0.1
//    BasePath: /
//
//    Produces:
//      - application/json
//
// swagger:meta
package main

import (
	"context"
	"log"
	"main/configstore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	store, err := configstore.New()
	if err != nil {
		log.Fatal(err)
	}

	server := configServer{
		store: store,
	}

	router.HandleFunc("/configs/", server.getAllConfigsHandler).Methods("GET")
	router.HandleFunc("/cfgroups/", server.getAllCfGroupsHandler).Methods("GET")
	router.HandleFunc("/cfgroup/", server.createConfigGroupHandler).Methods("POST")
	router.HandleFunc("/config/", server.createConfigHandler).Methods("POST")
	router.HandleFunc("/cfgroup/{id}/config/", server.expandConfigGroupHandler).Methods("PUT")
	router.HandleFunc("/config/{id}/{version}/", server.getConfigByIdAndVersionHandler).Methods("GET")
	router.HandleFunc("/config/{id}/", server.getConfigByIdHandler).Methods("GET")
	router.HandleFunc("/config/{id}/", server.deleteConfigByIdHandler).Methods("DELETE")
	router.HandleFunc("/config/{id}/{version}/", server.deleteConfigByIdAndVersionHandler).Methods("DELETE")
	router.HandleFunc("/cfgroup/{id}/", server.getCfGroupByIdHandler).Methods("GET")
	router.HandleFunc("/cfgroup/{id}/{version}/", server.getCfGroupByIdAndVersionHandler).Methods("GET")
	router.HandleFunc("/cfgroup/{id}/", server.deleteGroupByIdHandler).Methods("DELETE")
	router.HandleFunc("/cfgroup/{id}/{version}/", server.deleteGroupByIdAndVersionHandler).Methods("DELETE")
	router.HandleFunc("/cfgroup/{groupId}/{version}/config/{label}", server.getGroupConfigByLabelHandler).Methods("GET")
	router.HandleFunc("/cfgroup/{groupId}/{version}/config/{label}/{configId}/", server.getGroupConfigByIdAndLabelHandler).Methods("GET")
	router.HandleFunc("/cfgroup/{groupId}/{version}/config/{label}/{configId}/", server.deleteGroupConfigByLabelAndIdHandler).Methods("DELETE")
	router.HandleFunc("/config/{id}/", server.putConfigHandler).Methods("PUT")

	router.HandleFunc("/swagger.yaml", server.swaggerHandler).Methods("GET")

	// SwaggerUI
	optionsDevelopers := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	developerDocumentationHandler := middleware.SwaggerUI(optionsDevelopers, nil)
	router.Handle("/docs", developerDocumentationHandler)

	// ReDoc
	optionsShared := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sharedDocumentationHandler := middleware.Redoc(optionsShared, nil)
	router.Handle("/docs", sharedDocumentationHandler)

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
