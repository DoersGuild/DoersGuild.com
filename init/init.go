package init

import (
    "fmt"
    "html/template"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/temp", temp)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!!!!")
}

func temp(w http.ResponseWriter, r *http.Request) {

  var listTmpl = template.Must(template.ParseFiles("tmpl/welcome.html"))

  tc := make(map[string]interface{})
  tc["Store"] = "HEY"
  
  if err := listTmpl.Execute(w, tc); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  
}