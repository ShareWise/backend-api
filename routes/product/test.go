package routes

import (
	"backend-api/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	var a []ProductListing

	rows, err := db.Db().Query("SELECT * FROM product_listing where name = 'car'")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var p_listing ProductListing

		rows.Scan(&p_listing.Id, &p_listing.Name, &p_listing.Photo)

		a = append(a, p_listing)
	}

	jsonData, err := json.Marshal(a)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// rows.Close()
	// db.Db().Close()
	w.Write(jsonData)
}
