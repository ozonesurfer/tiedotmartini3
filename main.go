// main
package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	//	"html/template"
	"math/rand"
	"tiedotmartini3"
	"tiedotmartini3/controllers"
	"tiedotmartini3/model"
	"time"
)

func main() {
	//	fmt.Println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano())
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "home/index", struct {
			Title string
			Bands []model.DocWithID
		}{Title: "My CD Catalog", Bands: model.GetAll(tiedotmartini3.BAND_COL)})
	})
	m.Get("/home/index", func(r render.Render) {
		r.HTML(200, "home/index", struct {
			Title string
			Bands []model.DocWithID
		}{Title: "My CD Catalog", Bands: model.GetAll(tiedotmartini3.BAND_COL)})
	})
	m.Get("/band/add", func(r render.Render) {
		r.HTML(200, "band/add", struct {
			Title     string
			Locations []model.DocWithID
		}{Title: "Adding A Band", Locations: model.GetAll(tiedotmartini3.LOCATION_COL)})
	})
	m.Post("/band/verify", controllers.BandVerify)
	m.Get("/album/index/:id", func(params martini.Params, r render.Render) {
		rawId := params["id"]
		//	id, _ := strconv.ParseUint(rawId, 10, 64)
		id := model.ToObjectId(rawId)
		band, err := model.GetDoc(id, tiedotmartini3.BAND_COL)
		if err != nil {
			panic(err)
		}
		title := "Albums by " + band.Value["name"].(string)
		r.HTML(200, "album/index", struct {
			Title string
			Band  model.DocWithID
			Id    uint64
		}{Title: title, Band: band, Id: id})
	})
	m.Get("/album/add/:id", func(params martini.Params, r render.Render) {
		rawId := params["id"]
		//	id, _ := strconv.ParseUint(rawId, 10, 64)
		id := model.ToObjectId(rawId)
		genres := model.GetAll(tiedotmartini3.GENRE_COL)
		title := "Add Album"
		r.HTML(200, "album/add", struct {
			Title  string
			Genres []model.DocWithID
			Id     uint64
		}{Title: title, Genres: genres, Id: id})
	})
	m.Post("/album/verify/:id", controllers.AlbumVerify)
	m.Get("/home/genrelist", func(r render.Render) {
		r.HTML(200, "home/genrelist", struct {
			Title  string
			Genres []model.DocWithID
		}{Title: "List of Genres", Genres: model.GetAll(tiedotmartini3.GENRE_COL)})
	})
	m.Get("/home/bygenre/:id", func(params martini.Params, r render.Render) {
		id := params["id"]
		bands := model.GetBandsByGenre(id)
		genreName := model.GetGenreName(id)
		title := fmt.Sprintf("%s Albums", genreName)
		r.HTML(200, "home/bygenre", struct {
			Title string
			Bands []model.DocWithID
		}{Title: title, Bands: bands})
	})
	m.Run()

}
