// datainit
package main

import (
	"fmt"
	"tiedotmartini3"
	"tiedotmartini3/model"
)

func main() {
	fmt.Println("Initializing the database...")
	database := model.GetDB()
	database.Drop(tiedotmartini3.BAND_COL)
	database.Drop(tiedotmartini3.LOCATION_COL)
	database.Drop(tiedotmartini3.GENRE_COL)
	database.Create(tiedotmartini3.BAND_COL, 1)
	database.Create(tiedotmartini3.LOCATION_COL, 1)
	database.Create(tiedotmartini3.GENRE_COL, 1)
	col := database.Use(tiedotmartini3.BAND_COL)
	col.Index([]string{"albums", "genre_id"})
	database.Close()
}
