package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/LKarataev/Film-Library/internal/auth"
	"github.com/LKarataev/Film-Library/internal/dao"
	"github.com/LKarataev/Film-Library/internal/handlers"
	"github.com/LKarataev/Film-Library/internal/parser"
	_ "github.com/lib/pq"
)

type FilmLibraryApi struct {
	actorsRepo   *dao.ActorsRepository
	filmsRepo    *dao.FilmsRepository
	accountsRepo *dao.AccountsRepository
}

func NewFilmLibraryApi(db *sql.DB) (*FilmLibraryApi, error) {
	PrepareData(db)
	ar := dao.ActorsRepository{DB: db}
	fr := dao.FilmsRepository{DB: db}
	accr := dao.AccountsRepository{DB: db}
	api := FilmLibraryApi{actorsRepo: &ar, filmsRepo: &fr, accountsRepo: &accr}
	return &api, nil
}

func PrepareData(db *sql.DB) {
	scriptBytes, err := os.ReadFile("_postgres/sample_db.sql")
	if err != nil {
		log.Println("PrepareData error: ", err)
		return
	}
	script := string(scriptBytes)
	_, err = db.Exec(script)
	if err != nil {
		log.Println("PrepareData error: ", err)
	}
}

func (api *FilmLibraryApi) ConfigureRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/films", api.films)
	mux.HandleFunc("/films/", api.filmsId)
	mux.HandleFunc("/actors", api.actors)
	mux.HandleFunc("/actors/", api.actorsId)
	mux.HandleFunc("/authenticate", api.authenticateUser)
	return mux
}

func (api *FilmLibraryApi) authenticateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	req := handlers.AuthenticateUserRequest{
		Username: username,
		Password: password,
	}

	ctx := context.Background()
	signedToken, err := handlers.NewAuthenticateUserHandler(api.accountsRepo).Handle(ctx, req)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"X-Auth-Token": "` + signedToken + `"}`))
}

func (api *FilmLibraryApi) films(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getFilms(w, r)
	case http.MethodPost:
		api.insertFilm(w, r)
	case http.MethodPut:
		api.updateFilm(w, r)
	case http.MethodDelete:
		api.deleteFilm(w, r)
	default:
		http.Error(w, `{"error":"bad method"}`, http.StatusNotAcceptable)
	}
}

func (api *FilmLibraryApi) filmsId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getFilmById(w, r)
	default:
		http.Error(w, `{"error":"bad method"}`, http.StatusNotAcceptable)
	}
}

func (api *FilmLibraryApi) getFilmById(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 && account.Role != 1 {
		http.Error(w, `{"error": "you need to be authorized to request this method"}`, http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Path[len("/films/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.GetFilmByIdRequest{
		Id: id,
	}
	ctx := context.Background()
	film, err := handlers.NewGetFilmByIdHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)
}

func (api *FilmLibraryApi) getFilms(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 && account.Role != 1 {
		http.Error(w, `{"error": "you need to be authorized to request this method"}`, http.StatusUnauthorized)
		return
	}

	defaultLimit := 5
	defaultOffset := 0

	limit, err := strconv.Atoi(r.FormValue("limit"))

	if err != nil && r.FormValue("limit") != "" {
		errStr := fmt.Sprintf("Error: limit should be a number: `%s`", err.Error())
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if limit <= 0 {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil && r.FormValue("offset") != "" {
		errStr := fmt.Sprintf("Error: offset should be a number: `%s`", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if offset < 0 {
		offset = defaultOffset
	}

	search := r.FormValue("search")
	sort := r.FormValue("sort")
	order := r.FormValue("order")

	req := handlers.GetFilmsRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
		Sort:   sort,
		Order:  order,
	}

	ctx := context.Background()
	films, err := handlers.NewGetFilmsHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}

func (api *FilmLibraryApi) insertFilm(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 {
		http.Error(w, `{"error": "you need to be authorized as administrator to request this method"}`, http.StatusUnauthorized)
		return
	}

	values, actorsId, err := parser.ParseFilmActorsIdsOptions(r.Body)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.InsertFilmRequest{
		Values:   values,
		ActorsId: actorsId,
	}

	ctx := context.Background()
	err = handlers.NewInsertFilmHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "/films", http.StatusSeeOther)
}

func (api *FilmLibraryApi) updateFilm(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 {
		http.Error(w, `{"error": "you need to be authorized as administrator to request this method"}`, http.StatusUnauthorized)
		return
	}

	filmid, values, err := parser.ParseFilmOptions(r.Body)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.UpdateFilmRequest{
		Id:     filmid,
		Values: values,
	}

	ctx := context.Background()
	err = handlers.NewUpdateFilmHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/films", http.StatusSeeOther)
}

func (api *FilmLibraryApi) deleteFilm(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 {
		http.Error(w, `{"error": "you need to be authorized as administrator to request this method"}`, http.StatusUnauthorized)
		return
	}

	filmId, _, err := parser.ParseFilmOptions(r.Body)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.DeleteFilmRequest{
		Id: filmId,
	}

	ctx := context.Background()
	err = handlers.NewDeleteFilmHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/films", http.StatusSeeOther)
}

func (api *FilmLibraryApi) actors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getActors(w, r)
	case http.MethodPut:
		api.updateActor(w, r)
	case http.MethodPost:
		api.insertActor(w, r)
	case http.MethodDelete:
		api.deleteActor(w, r)
	default:
		http.Error(w, `{"error":"bad method"}`, http.StatusNotAcceptable)
	}
}

func (api *FilmLibraryApi) actorsId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getActorById(w, r)
	default:
		http.Error(w, `{"error":"bad method"}`, http.StatusNotAcceptable)
	}
}

func (api *FilmLibraryApi) getActorById(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 && account.Role != 1 {
		http.Error(w, `{"error": "you need to be authorized to request this method"}`, http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Path[len("/actors/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.GetActorByIdRequest{
		Id: id,
	}
	ctx := context.Background()
	actor, err := handlers.NewGetActorByIdHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actor)
}

func (api *FilmLibraryApi) getActors(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 && account.Role != 1 {
		http.Error(w, `{"error": "you need to be authorized to request this method"}`, http.StatusUnauthorized)
		return
	}

	defaultLimit := 5
	defaultOffset := 0

	limit, err := strconv.Atoi(r.FormValue("limit"))

	if err != nil && r.FormValue("limit") != "" {
		errStr := fmt.Sprintf("Error: limit should be a number: `%s`", err.Error())
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if limit <= 0 {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil && r.FormValue("offset") != "" {
		errStr := fmt.Sprintf("Error: offset should be a number: `%s`", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if offset < 0 {
		offset = defaultOffset
	}

	req := handlers.GetActorsRequest{
		Limit:  limit,
		Offset: offset,
	}

	ctx := context.Background()
	actors, err := handlers.NewGetActorsHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actors)
}

func (api *FilmLibraryApi) insertActor(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 {
		http.Error(w, `{"error": "you need to be authorized as administrator to request this method"}`, http.StatusUnauthorized)
		return
	}

	_, values, err := parser.ParseActorOptions(r.Body)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.InsertActorRequest{
		Values: values,
	}

	ctx := context.Background()
	err = handlers.NewInsertActorHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "/actors", http.StatusSeeOther)
}

func (api *FilmLibraryApi) updateActor(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 {
		http.Error(w, `{"error": "you need to be authorized as administrator to request this method"}`, http.StatusUnauthorized)
		return
	}

	actorId, values, err := parser.ParseActorOptions(r.Body)

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.UpdateActorRequest{
		Id:     actorId,
		Values: values,
	}

	ctx := context.Background()
	err = handlers.NewUpdateActorHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/actors", http.StatusSeeOther)
}

func (api *FilmLibraryApi) deleteActor(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("X-Auth-Token")
	account, err := auth.ValidateToken(authToken)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	if account.Role != 0 {
		http.Error(w, `{"error": "you need to be authorized as administrator to request this method"}`, http.StatusUnauthorized)
		return
	}

	actorId, _, err := parser.ParseActorOptions(r.Body)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	req := handlers.DeleteActorRequest{
		Id: actorId,
	}

	ctx := context.Background()
	err = handlers.NewDeleteActorHandler(api.actorsRepo, api.filmsRepo).Handle(ctx, req)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	http.Redirect(w, r, "/actors", http.StatusSeeOther)
}
