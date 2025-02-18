package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type CounterHandler struct {
	mu    sync.Mutex
	count int
}

func (h *CounterHandler) loadCounter() {
	data, err := os.ReadFile("counter.txt")
	if err == nil {
		val, err := strconv.Atoi(string(data))
		if err == nil {
			h.count = val
		}
	}
}

func (h *CounterHandler) saveCounter() {
	err := os.WriteFile("counter.txt", []byte(strconv.Itoa(h.count)), 0644)
	if err != nil {
		log.Println("Error saving counter: ", err)
	}
}

func (h *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.count++
	h.saveCounter()
	h.mu.Unlock()

	data := struct {
		Count int
	}{h.count}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "could not load template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	tmpl.Execute(w, data)
}

func main() {

	counter := &CounterHandler{}
	counter.loadCounter()

	http.Handle("/counter", counter)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not set
	}

	fmt.Printf("server running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
