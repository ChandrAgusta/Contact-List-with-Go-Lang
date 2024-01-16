package contactcontroller

import (
	"go-web-native/entities"
	"go-web-native/model/contactmodel"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
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

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var contact entities.Contact

		contact.Nama = r.FormValue("nama")
		contact.Nomor = r.FormValue("nomor")

		if ok := contactmodel.Create(contact); !ok {
			temp, _ := template.ParseFiles("views/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		contacts := contactmodel.Detail(id)
		data := map[string]any{
			"contacts": contacts,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var contacts entities.Contact
		idString := r.FormValue("id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		contacts.Nama = r.FormValue("nama")
		contacts.Nomor = r.FormValue("nomor")

		if ok := contactmodel.Update(id, contacts); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Hapus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil{
			panic(err)
		}

		if ok := contactmodel.Delete(id); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
