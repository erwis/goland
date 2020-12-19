package main

import (
	"fmt"
	"goland/database"

	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// )

func main() {

	db := database.InitDB()
	defer db.Close()
	fmt.Println(db)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
