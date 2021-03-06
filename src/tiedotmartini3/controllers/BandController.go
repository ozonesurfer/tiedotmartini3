// BandController
package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"tiedotmartini3"
	"tiedotmartini3/model"
)

/*func main() {
	fmt.Println("Hello World!")
}
*/

func BandAdd(r http.ResponseWriter, rq *http.Request) {
	locations := model.GetAll(tiedotmartini3.LOCATION_COL)
	t, err := template.ParseFiles("src/tiedotmartini3/views/band/add.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, struct {
			Title     string
			Locations []model.DocWithID
		}{Title: "Adding A Band", Locations: locations})
	}
}

func BandVerify(r http.ResponseWriter, rq *http.Request) {
	name := rq.FormValue("name")
	locType := rq.FormValue("loctype")
	//	var locationId uint64
	var locationId interface{}
	errorString := "no errors"
	var err error
	switch locType {
	case "existing":
		locationIdString := rq.FormValue("location_id")
		if locationIdString == "" {
			errorString = "No location was selected"
		} else {
			locationId = model.ToObjectId(locationIdString)
			fmt.Println("Using location id", locationId)
		}
		break
	case "new":
		if rq.FormValue("country") == "" {
			errorString = "Country is required"
		} else {
			city := rq.FormValue("city")
			state := rq.FormValue("state")
			country := rq.FormValue("country")
			location := map[string]interface{}{
				"city":    city,
				"state":   state,
				"country": country}
			/*	location := map[string]interface{}{"city": rq.FormValue("city"),
				"state":   rq.FormValue("state"),
				"country": rq.FormValue("country")} */

			locationId, err = model.AddDoc(location, tiedotmartini3.LOCATION_COL)
			fmt.Println("Adding location of Id", locationId)
			if err != nil {
				errorString = "error on location addition: " + err.Error()
				//	locationId = 0
			}
		}
		break
	default:
		errorString = "You need to select an option"
		//	locationId = 0
	}
	if errorString == "no errors" {
		var albums []model.Album
		fmt.Println("Adding location Id", locationId)
		band := map[string]interface{}{"name": name,
			"location_id": fmt.Sprint(locationId),
			"albums":      albums}
		_, err := model.AddDoc(band, tiedotmartini3.BAND_COL)
		if err != nil {
			errorString = fmt.Sprintf("Add band error: %s", err.Error())
		}
	}
	t, err := template.ParseFiles("src/tiedotmartini3/views/band/verify.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, struct {
			Title   string
			Message string
		}{Title: "Verifying Band",
			Message: errorString})
	}
}
