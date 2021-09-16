package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	iu "github.com/ArtemZar/URL_shortener/app/isurl"
	lg "github.com/ArtemZar/URL_shortener/app/linksgen"
	db "github.com/ArtemZar/URL_shortener/db"
)

func HandlesFunc() {
	rtr := mux.NewRouter()
	// Подключение статических фалов (папка с доп css)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../../frontend/css/"))))

	rtr.HandleFunc("/", index).Methods("GET", "POST")
	rtr.HandleFunc("/generate_link", generate_link).Methods("GET", "POST") // был гет стал гет пост
	http.Handle("/", rtr)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../frontend/index.html", "../../frontend/header.html", "../../frontend/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func generate_link(w http.ResponseWriter, r *http.Request) {
	var link db.MyDB
	link.LongLink = r.FormValue("longlink")
	if !iu.IsUrl(link.LongLink) {
		fmt.Fprintf(w, "link not entered")
		return
	}

	ll := db.LookForLongLink(link.LongLink)
	if ll != "" {
		link.ShortLink = ll
	} else {
		link.ShortLink = r.Host + "/" + lg.LinksGen()
		db.InsertToDB(link)
	}
	link.ClickCounter = 0

	t, err := template.ParseFiles("../../frontend/generate_link.html", "../../frontend/header.html", "../../frontend/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "generate_link", link)
	//t.Execute(w, nil)
}
