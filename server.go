package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
<<<<<<< HEAD
	"strings"
=======
>>>>>>> mem-impl from HeroService and add methods Add and GetByID

	"github.com/gorilla/mux"
	"github.com/lima1909/goheroes-appengine/db"
	"github.com/lima1909/goheroes-appengine/service"

	"google.golang.org/appengine"
)

var (
	app *service.App
)

func main() {

	app = service.NewApp(db.MemService{})

	router := mux.NewRouter()

<<<<<<< HEAD
	router.HandleFunc("/api/heroes", heroes)
	router.HandleFunc("/api/heroes/", searchHeroes)
	router.HandleFunc("/api/heroes/{id:[0-9]+}", heroesID)
=======
	router.Handle("/", http.RedirectHandler("/api/heroes", http.StatusFound))
>>>>>>> mem-impl from HeroService and add methods Add and GetByID

	router.HandleFunc("/api/heroes", heroes)
	router.Methods("GET").Path("/api/heroes/{id:[0-9]+}").HandlerFunc(heroID)
	http.Handle("/", router)

	log.Println("Start Server: http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
	// appengine.Main()
}

func heroes(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	setHeaderOptions(w)

	switch r.Method {
	case "GET":
		// loadHeroes(w)
		h, err := db.ListHeroes(appengine.NewContext(r))
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		fmt.Fprintln(w, toJSON(h))
	case "OPTIONS":
		fmt.Fprintf(w, string(http.StatusOK))
	case "PUT":
		updateHero(w, r)
	case "POST":
		addHero(w, r)
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
	}
}

func loadHeroes(w http.ResponseWriter) {
	b, err := json.Marshal(Heroes)

	if err != nil {
		fireError(w, "Err by marshal heroes: %v", err)
		return
=======
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		heroList(w, r)
	case "POST":
		addHero(w, r)
	default:
		http.Error(w, "invalid method: "+r.Method, http.StatusBadRequest)
>>>>>>> mem-impl from HeroService and add methods Add and GetByID
	}
}

<<<<<<< HEAD
	fmt.Fprintf(w, string(b))
}

func getHeroByID(w http.ResponseWriter, r *http.Request) {
	setHeaderOptions(w)

	vars := mux.Vars(r)
	varID := vars["id"]
	i, err := strconv.Atoi(varID)

	if err != nil {
		fireError(w, "Err during string convert: %v", err)
		return
	}

	b, err := json.Marshal(Heroes[i-1])

	if err != nil {
		fireError(w, "Err by marshal hero: %v", err)
		return
=======
func heroList(w http.ResponseWriter, r *http.Request) {
	name := ""
	names, ok := r.URL.Query()["name"]
	if ok || len(names) == 1 {
		name = names[0]
>>>>>>> mem-impl from HeroService and add methods Add and GetByID
	}

<<<<<<< HEAD
	fmt.Fprintf(w, string(b))
}

func updateHero(w http.ResponseWriter, r *http.Request) {
	hero, err := getHeroFromRequest(r, w)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	heroes, err := svc.List(appengine.NewContext(r), name)
=======
	heroes, err := app.List(appengine.NewContext(r), name)
>>>>>>> mem-impl from HeroService and add methods Add and GetByID
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(heroes)
	if err != nil {
		fireError(w, "Err during string convert: %v", err)
		return
	}

	//update Hero in List
	Heroes[i-1] = hero

	fmt.Fprintf(w, "")
}

<<<<<<< HEAD
func addHero2(w http.ResponseWriter, r *http.Request) {
	hero, err := getHeroFromRequest(r, w)

	if err != nil {
		fireError(w, "Err by getHeroFromRequest %v", err)
		return
	}

	hero.ID = strconv.Itoa(len(Heroes) + 1)

	// func heros(w http.ResponseWriter, r *http.Request) {

	// 	switch r.Method {
	// 	case "GET":
	// 		// loadHeroes(w)
	// 		h, err := db.ListHeroes(appengine.NewContext(r))
	// 		if err != nil {
	// 			fmt.Fprintf(w, "%v", err)
	// 		}
	// 		fmt.Fprintln(w, toJSON(h))
	// 	case "OPTIONS":
	// 		writeToClient(w, string(http.StatusOK))
	// 	case "PUT":
	// 		updateHero(w, r)
	// 	case "POST":
	// 		addHero(w, r)
	// 	}
	// }

	// func toJSON(v interface{}) string {
	// 	b, err := json.Marshal(v)

	// 	if err != nil {
	// 		return fmt.Sprintf("Err by marshal heros: %v", err)
	// 	}

	// 	return string(b)
	// }

	// func loadHeroes(w http.ResponseWriter) {
	// 	b, err := json.Marshal(Heroes)

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err by marshal heros: %v", err)
	// 		return
	// 	}

	// 	writeToClient(w, string(b))
	// }

	// func getHeroByID(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	varID := vars["id"]
	// 	i, err := strconv.Atoi(varID)

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err during string convert: %v", err)
	// 		return
	// 	}

	// 	b, err := json.Marshal(Heroes[i-1])

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err by marshal hero: %v", err)
	// 		return
	// 	}

	// 	writeToClient(w, string(b))
	// }

	// func updateHero(w http.ResponseWriter, r *http.Request) {
	// 	hero, err := getHeroFromRequest(r, w)

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err by getHeroFromRequest %v", err)
	// 		return
	// 	}

	// 	i, err := strconv.Atoi(hero.ID)

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err during string convert: %v", err)
	// 		return
	// 	}

	// 	//update Hero in List
	// 	Heroes[i-1] = hero

	// 	writeToClient(w, "")
	// }

	// func addHero2(w http.ResponseWriter, r *http.Request) {
	// 	hero, err := getHeroFromRequest(r, w)

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err by getHeroFromRequest %v", err)
	// 		return
	// 	}

	// 	hero.ID = strconv.Itoa(len(Heroes) + 1)

	// 	Heroes = append(Heroes, hero)

	// 	b, err := json.Marshal(hero)

	// 	if err != nil {
	// 		fmt.Fprintf(w, "Err by marshal hero: %v", err)
	// 		return
	// 	}

	if err != nil {
		fireError(w, "Err by marshal hero: %v", err)
		return
	}

	fmt.Fprintf(w, string(b))
}

func deleteHero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varID := vars["id"]
	i, err := strconv.Atoi(varID)

	if err != nil {
		fireError(w, "Err during string convert: %v", err)
		return
	}

	//remove hero from list
	Heroes = append(Heroes[:i-1], Heroes[i:]...)

	//adjust Hero.ID
	for j := (i - 1); j < len(Heroes); j++ {
		Heroes[j].ID = strconv.Itoa(j + 1)
	}

	fmt.Fprintf(w, "")
}

func searchHeroes(w http.ResponseWriter, r *http.Request) {
	searchString, ok := r.URL.Query()["name"]

	if !ok || len(searchString) < 1 {
		fireError(w, "Url Param 'key' is missing", nil)
		return
	}

	setHeaderOptions(w)

	findHeroes := []Hero{}

	//compare Hero.Name with searchString
	for _, hero := range Heroes {
		if strings.Contains(hero.Name, searchString[0]) {
			findHeroes = append(findHeroes, hero)
		}
	}

	//convert to json
	b, err := json.Marshal(findHeroes)

	if err != nil {
		fireError(w, "Err by marshal hero: %v", err)
		return
	}

	fmt.Fprintf(w, string(b))
=======
func addHero(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	hero := service.Hero{}
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hero, err = app.Add(appengine.NewContext(r), hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
>>>>>>> mem-impl from HeroService and add methods Add and GetByID
}

func heroID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	varID := vars["id"]
	id, err := strconv.Atoi(varID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id: %v", varID), http.StatusBadRequest)
		return
	}

	hero, err := app.GetByID(appengine.NewContext(r), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

<<<<<<< HEAD
func setHeaderOptions(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func fireError(w http.ResponseWriter, message string, err error) {
	formattedError := fmt.Sprintf(message, err)
	http.Error(w, fmt.Errorf(formattedError).Error(), http.StatusInternalServerError)
=======
	fmt.Fprintf(w, "%s", string(b))
>>>>>>> mem-impl from HeroService and add methods Add and GetByID
}
