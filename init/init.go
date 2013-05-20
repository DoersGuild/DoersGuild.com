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
    router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.Handle("/", router)
}

func executeSimpleTemplate(w http.ResponseWriter, tmplFile string) {
	// Load and execute the given template
    
  var listTmpl = template.Must(template.ParseFiles("tmpl/base.html", tmplFile))
  
  if err := listTmpl.Execute(w, config); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// The 404 page
	executeSimpleTemplate(w, "tmpl/content/404.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// The home page
	executeSimpleTemplate(w, "tmpl/content/index.html")
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio page
	config["portfolioItems"] = portfolioItems
	config["portfolioCategories"] = portfolioCategories
	executeSimpleTemplate(w, "tmpl/content/portfolio.html")
}