package main

import (
	"fmt"

	db "backend-api/db"
)

type ProductListing struct {
	Id    string
	Photo string
	Name  string
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, Woqwdqqwewrld"))
	// })

	// http.ListenAndServe(":8080", nil)
	var a []ProductListing

	rows, err := db.Db().Query("SELECT * FROM product_listing")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var p_listing ProductListing

		rows.Scan(&p_listing.Id, &p_listing.Name, &p_listing.Photo)

		a = append(a, p_listing)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	// rows.Close()
	// db.Db().Close()

}
