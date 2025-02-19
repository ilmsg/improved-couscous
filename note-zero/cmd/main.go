package main

import (
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
	v1 := r.PathPrefix("/v1").Subrouter()
	s := v1.PathPrefix("/notes").Subrouter()

	s.HandleFunc("/", hNote.CreateNote).Methods("POST")
	s.HandleFunc("/", hNote.ListNote).Methods("GET")
	s.HandleFunc("/{id}", hNote.GetNote).Methods("GET")
	s.HandleFunc("/{id}", hNote.UpdateNote).Methods("PATCH")
	s.HandleFunc("/{id}", hNote.DeleteNote).Methods("DELETE")

	http.ListenAndServe(":3060", r)
}
