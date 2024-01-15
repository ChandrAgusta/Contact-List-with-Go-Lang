package homecontroller

import (
	"go-web-native/model/contactmodel"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// temp, err := template.ParseFiles("views/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	// temp.Execute(w, nil)

	contacts := contactmodel.GetAll()

	data := map[string]any{
		"contacts": contacts,
	}

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
