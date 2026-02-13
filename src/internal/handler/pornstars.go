package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Pornstar struct {
	Name   string
	Rating float64
}

// type Pornstars struct {
// 	[]struct
// 	// name string
// 	// rating float64
// }

func GetPornstars(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var pornstar Pornstar
		pornstars, err := db.Query("select name, rating from public.pornstars")
		pornstars.Next()
		pornstars.Next()
		pornstars.Scan(&pornstar.Name, &pornstar.Rating)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(pornstar)

		w.Header().Set("Content-Type", "application/json")
		JWTSet("Nika", "token1")
		fmt.Println("token1 exists", JWTIsValid("Nika", "token1"))
		fmt.Println("token2 exists", JWTIsValid("Nika", "token2"))

		json.NewEncoder(w).Encode(pornstar)
	}
}
