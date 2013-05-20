package init

import (
    //"fmt"
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
    "appengine"
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
	router.HandleFunc("/portfolio/{category}", portfolioCategoryHandler)
	router.HandleFunc("/portfolio/{category}/{project}", portfolioDetailsHandler)
    router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.Handle("/", router)
}

func executeSimpleTemplate(w http.ResponseWriter, r *http.Request, tmplFile string) {
	// Load and execute the given template
	
	// Log stats
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)
	c.Infof("Loading Template: %v", tmplFile)
    
	var listTmpl = template.Must(template.ParseFiles("tmpl/base.html", tmplFile))
	  
	if err := listTmpl.Execute(w, config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
  
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// The 404 page
	executeSimpleTemplate(w, r, "tmpl/content/404.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// The home page
	executeSimpleTemplate(w, r, "tmpl/content/index.html")
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio page
	config["portfolioItems"] = portfolioItems
	config["portfolioCategories"] = portfolioCategories
	executeSimpleTemplate(w, r, "tmpl/content/portfolio.html")
}

func portfolioCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio category listing page
	vars := mux.Vars(r)
	category, cat_ok:=vars["category"]
	if(!cat_ok) {
		portfolioHandler(w, r)
		return
	}
	config["portfolioItems"] = portfolioItems[category]
	config["portfolioCategory"] = portfolioCategories[category]
	executeSimpleTemplate(w, r, "tmpl/content/portfolio_category.html")
}

func portfolioDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio details page
	vars := mux.Vars(r)
	category:=vars["category"]
	project:=vars["project"]
	config["portfolioItem"] = portfolioItems[category].(map[string]PortfolioItem)[project]
	config["portfolioCategory"] = portfolioCategories[category]
	executeSimpleTemplate(w, r, "tmpl/content/portfolio_details.html")
}