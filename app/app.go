package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Nishith-Savla/Random-Wordlist-Api/domain"
	"github.com/Nishith-Savla/Random-Wordlist-Api/service"
	"github.com/gorilla/mux"
)

func Start() {
	wordlistRepo, err := domain.NewWordlistRepositoryFromFile("wordlist.json")
	if err != nil {
		log.Fatalln(err)
	}
	//wordlistRepo := &domain.WordlistRepositoryStub{Pointer: 0, Words: []string{"abreact", "abreacted", "abreacting", "abreaction", "abreactions", "abreacts", "abreast", "abri", "abridge", "abridged", "abridgement", "abridgements", "abridger", "abridgers", "abridges", "abridging", "abridgment", "abridgments", "abris", "abroach", "abroad", "abrogable", "abrogate", "abrogated", "abrogates", "abrogating", "abrogation", "abrogations", "abrogator", "abrogators", "abrosia", "abrosias", "abrupt", "abrupter", "abruptest", "abruption", "abruptions", "abruptly", "abruptness", "abruptnesses", "abs", "abscess", "abscessed", "abscesses", "abscessing", "abscise", "abscised", "abscises", "abscisin", "abscising", "abscisins", "abscissa"}}

	wordlistService := service.NewDefaultWordlistService(wordlistRepo)

	wh := WordlistHandler{service: wordlistService}

	r := mux.NewRouter()

	allowedLimits := []uint16{10, 20, 50, 100, 200}
	allowedLimitsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(allowedLimits)), "|"), "[]")
	limitQueryParam := fmt.Sprintf("{limit:(?:%v)}", allowedLimitsString)

	r.Use(limit, authorizationHandler(os.Getenv("API_KEY")))
	r.HandleFunc("/words", wh.getWords).
		Queries("limit", limitQueryParam).
		Methods("GET")

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	log.Fatalln(http.ListenAndServe(port, r))
}
