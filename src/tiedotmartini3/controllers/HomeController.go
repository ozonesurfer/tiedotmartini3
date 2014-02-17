// HomeController
package controllers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"html/template"
	"net/http"
	//	"strconv"
	"tiedotmartini3"
	"tiedotmartini3/model"
)

/* func main() {
	fmt.Println("Hello World!")
} */

func HomeIndex(r http.ResponseWriter, rw *http.Request) {
	bands := model.GetAll(tiedotmartini3.BAND_COL)
	t, err := template.ParseFiles("src/tiedotmartini3/views/home/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(r, struct {
		Title string
		Bands []model.DocWithID
	}{Title: "My CD Catalog", Bands: bands})
}

func HomeGenreList(r http.ResponseWriter, rw *http.Request) {
	genres := model.GetAll(tiedotmartini3.GENRE_COL)
	t, err := template.ParseFiles("src/tiedotmartini3/views/home/genrelist.html")
	if err != nil {
		panic(err)
	}
	t.Execute(r, struct {
		Title  string
		Genres []model.DocWithID
	}{Title: "List of Genres", Genres: genres})
}

func HomeByGenre(params martini.Params, r http.ResponseWriter, rw *http.Request) {
	id := params["id"]
	//	id, _ := strconv.ParseUint(rawId, 10, 64)
	//	id := model.ToObjectId(rawId)
	bands := model.GetBandsByGenre(id)
	genreName := model.GetGenreName(id)
	//	title := genreName + " Albums"
	title := fmt.Sprintf("%s Albums", genreName)
	t, err := template.ParseFiles("src/tiedotmartini3/views/home/bygenre.html")
	if err != nil {
		panic(err)
	}
	t.Execute(r, struct {
		Title string
		Bands []model.DocWithID
	}{Title: title, Bands: bands})
}
