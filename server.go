package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lima1909/goheroes-appengine/db"
	"github.com/lima1909/goheroes-appengine/service"

	"google.golang.org/appengine"
)

var (
	app  *service.App
	info Info
)

// Info to the current system
type Info struct {
	HersoService       string
	EnvHeroServiceImpl string
}

func main() {
	info = Info{EnvHeroServiceImpl: os.Getenv("HERO_SERVICE_IMPL")}

	router := mux.NewRouter()

	router.Handle("/", http.RedirectHandler("/info", http.StatusFound))

	router.HandleFunc("/info", infoPage)
	router.HandleFunc("/api/heroes", heroes)
	router.HandleFunc("/api/heroes/", searchHeroes)
	router.HandleFunc("/api/heroes/{id:[0-9]+}", heroesID)

	http.Handle("/", router)

	if info.EnvHeroServiceImpl == "datastore" {
		app = service.NewApp(db.DatastoreService{})
		info.HersoService = "DatastoreService"
	} else {
		app = service.NewApp(db.NewMemService())
		info.HersoService = "MemService"
		log.Println("HeroServicem is MemService")
		log.Println("Start server on: http://localhost:8080")
	}

	appengine.Main()
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/info.html")
	if err != nil {
		fmt.Fprintf(w, "Err: %v\n", err)
		return
	}

	err = t.Execute(w, info)
	if err != nil {
		fmt.Fprintf(w, "Err: %v\n", err)
		return
	}
}

func heroes(w http.ResponseWriter, r *http.Request) {
	setHeaderOptions(w)

	switch r.Method {
	case "GET":
		loadHeroes(w, r)
	case "OPTIONS":
		fmt.Fprintf(w, string(http.StatusOK))
	case "PUT":
		updateHero(w, r)
	case "POST":
		addHero(w, r)
	default:
		http.Error(w, "invalid method: "+r.Method, http.StatusBadRequest)
	}
}

func heroesID(w http.ResponseWriter, r *http.Request) {
	setHeaderOptions(w)

	switch r.Method {
	case "GET":
		getHeroByID(w, r)
	case "OPTIONS":
		fmt.Fprintf(w, string(http.StatusOK))
	case "DELETE":
		deleteHero(w, r)
	default:
		http.Error(w, "invalid method: "+r.Method, http.StatusBadRequest)
	}
}

func loadHeroes(w http.ResponseWriter, r *http.Request) {
	//TODO is this empty string needed?
	heroes, err := app.List(appengine.NewContext(r), "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(heroes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func getHeroByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varID := vars["id"]
	i, err := strconv.Atoi(varID)

	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id: %v", varID), http.StatusBadRequest)
		return
	}

	hero, err := app.GetByID(appengine.NewContext(r), int64(i))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(hero)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func updateHero(w http.ResponseWriter, r *http.Request) {
	hero, err := getHeroFromRequest(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h, err := app.Update(appengine.NewContext(r), hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO no need to return the hero to client!!!
	fmt.Fprintf(w, string(b))
}

func addHero(w http.ResponseWriter, r *http.Request) {
	hero, err := getHeroFromRequest(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h, err := app.Add(appengine.NewContext(r), hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func deleteHero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varID := vars["id"]
	i, err := strconv.Atoi(varID)

	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id: %v", varID), http.StatusBadRequest)
		return
	}

	hero, err := app.Delete(appengine.NewContext(r), int64(i))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO no need to return the hero to client!!!
	fmt.Fprintf(w, string(b))
}

func searchHeroes(w http.ResponseWriter, r *http.Request) {
	setHeaderOptions(w)

	searchString := ""
	names, ok := r.URL.Query()["name"]

	if !ok || len(names) < 1 {
		http.Error(w, "Url Param 'key' is missing", http.StatusInternalServerError)
		return
	}
	searchString = names[0]

	findHeroes, err := app.List(appengine.NewContext(r), searchString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//convert to json
	b, err := json.Marshal(findHeroes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func setHeaderOptions(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getHeroFromRequest(w http.ResponseWriter, r *http.Request) (service.Hero, error) {
	defer r.Body.Close()

	hero := service.Hero{}
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		return service.Hero{}, err
	}

	return hero, nil
}
