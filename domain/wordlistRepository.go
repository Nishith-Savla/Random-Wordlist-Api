package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type WordlistRepository interface {
	GetWords(limit int) []string
	Shuffle()
}

type WordlistRepositoryFile struct {
	Pointer int
	Words   []string
}

func (r *WordlistRepositoryFile) GetWords(limit int) []string {
	var words []string
	if r.Pointer+limit > len(r.Words) {
		words = r.Words[r.Pointer:len(r.Words)]
		words = append(words, r.Words[0:limit-len(words)]...)
		r.Pointer = 0
	} else {
		words = r.Words[r.Pointer : r.Pointer+limit]
	}
	r.Pointer = r.Pointer + limit
	fmt.Println(r.Pointer, limit)
	if r.Pointer >= len(r.Words)-limit {
		go r.Shuffle()
	}
	return words
}

func (r *WordlistRepositoryFile) Shuffle() {
	start := time.Now()
	a := make([]string, len(r.Words))
	copy(a, r.Words)
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	r.Words = a
	log.Printf("Took %s seconds\n", time.Since(start))
}

func NewWordlistRepositoryFile(filename string) (*WordlistRepositoryFile, error) {
	var wordlist map[string][]string
	var bytes []byte
	var err error

	if bytes, err = os.ReadFile(filename); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &wordlist); err != nil {
		return nil, err
	}

	r := WordlistRepositoryFile{
		Pointer: 0,
		Words:   wordlist["words"],
	}
	//r.Shuffle()
	return &r, nil

}
