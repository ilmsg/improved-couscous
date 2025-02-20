package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/improved-couscous/note-zero/database"
	"github.com/ilmsg/improved-couscous/note-zero/handler"
	"github.com/ilmsg/improved-couscous/note-zero/model"
	"github.com/ilmsg/improved-couscous/note-zero/repo"
	"github.com/ilmsg/improved-couscous/note-zero/service"
	"gorm.io/driver/sqlite"
)

func main() {
	db := database.NewDatabase(sqlite.Open("db.sqlite"))
	db.AutoMigrate(model.Note{})

	repoNote := repo.NewRepoNote(db)
	srvNote := service.NewServiceNote(repoNote)
	hNote := handler.NewHandlerNote(srvNote)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	s := api.PathPrefix("/v1").Subrouter()

	s.HandleFunc("/notes", hNote.CreateNote).Methods("POST")
	s.HandleFunc("/notes", hNote.ListNote).Methods("GET")
	s.HandleFunc("/notes/{id}", hNote.GetNote).Methods("GET")
	s.HandleFunc("/notes/{id}", hNote.UpdateNote).Methods("PATCH")
	s.HandleFunc("/notes/{id}", hNote.DeleteNote).Methods("DELETE")

	log.Printf("server listen at *:3060")
	if err := http.ListenAndServe(":3060", r); err != nil {
		log.Fatal(err)
	}
}
