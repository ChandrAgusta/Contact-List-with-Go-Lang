package contactmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Contact {
	rows, err := config.DB.Query("SELECT id, nama, nomor FROM contact_list")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var contacts []entities.Contact

	for rows.Next() {
		var contact entities.Contact
		err := rows.Scan(&contact.Id, &contact.Nama, &contact.Nomor)

		if err != nil {
			panic(err)
		}

		contacts = append(contacts, contact)
	}

	return contacts
}

func Create(contact entities.Contact) bool{
	result, err := config.DB.Exec(
		"INSERT INTO contact_list (nama, nomor) VALUE(?,?)",
		contact.Nama, contact.Nomor,
	)

	if err != nil{
		panic(err)
	}

	lastInsertId, err:= result.LastInsertId()
	if err!= nil{
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int)entities.Contact{
	row := config.DB.QueryRow("SELECT id, nama, nomor FROM contact_list WHERE id =?", id)

	var contact entities.Contact

	if err := row.Scan(&contact.Id, &contact.Nama, &contact.Nomor); err != nil{
		panic(err.Error())
	}
	return contact
}

func Update(id int, contacts entities.Contact)bool {
	query, err := config.DB.Exec("UPDATE contact_list SET nama = ?, nomor =? WHERE id = ?", contacts.Nama, contacts.Nomor, id)

	if err!= nil{
		panic(err)
	}

	result,err := query.RowsAffected()
	if err!= nil{
		panic(err)
	}

	return result > 0
}
