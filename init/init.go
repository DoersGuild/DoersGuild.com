package init

import (
    //"fmt"
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
)

// Acommon variable with data for templates
var config = make(map[string]interface{})

func init() {
	// The main router
	
	// Setup common config vars
	config["title"] = "Doers' Guild"
	
	router:=mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/portfolio", portfolioHandler)
	router.HandleFunc("/portfolio/{category}", portfolioHandler)
	router.HandleFunc("/portfolio/{category}/{project}", portfolioHandler)
	http.Handle("/", router)
}

func executeSimpleTemplate(w http.ResponseWriter, tmplFile string) {
	// Load and execute the given template
    
  var listTmpl = template.Must(template.ParseFiles("tmpl/base.html", tmplFile))
  
  if err := listTmpl.Execute(w, config); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// The home page
	executeSimpleTemplate(w, "tmpl/content/index.html")
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio page
	config["portfolioItems"] = portfolioItems
	executeSimpleTemplate(w, "tmpl/content/portfolio.html")
}