package init

import (
    //"fmt"
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
    "appengine"
    "strings"
)

// Acommon variable with data for templates
var config = make(map[string]interface{})

func init() {
	// The main router
	
	// Setup common config vars
	config["tagline"] = "Your Cross-Platform App Development Partners"
	config["title"] = config["tagline"]
	
	router:=mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/img/{path:.*}", imageRedirectHandler)
	router.HandleFunc("/portfolio", portfolioHandler)
	router.HandleFunc("/portfolio/{category}", portfolioCategoryHandler)
	router.HandleFunc("/portfolio/{category}/{project}", portfolioDetailsHandler)
    router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.Handle("/", router)
}

func imageRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect images to the unlimited bandwidth hosting
	basePath := "http://store.doersguild.com/WebsiteAssets/img"
	vars := mux.Vars(r)
	imagePath:=vars["path"]
	path := strings.Join([]string{basePath, imagePath}, "/")
	http.Redirect(w, r, path, http.StatusMovedPermanently)
}

func executeSimpleTemplate(w http.ResponseWriter, r *http.Request, tmplFile string) {
	// Load and execute the given template
	
	// Log stats
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)
	c.Infof("Loading Template: %v", tmplFile)
    
	var listTmpl = template.Must(template.ParseFiles("tmpl/base.html", "tmpl/blocks.html", tmplFile))
	  
	if err := listTmpl.Execute(w, config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
  
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// The 404 page
	config["title"] = "Page Not Found"
	executeSimpleTemplate(w, r, "tmpl/content/404.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// The home page
	// config["title"] = "Welcome"
	config["homePageItems"] = homePageItems
	config["homePageFeedback"] = homePageFeedback
	config["portfolioItems"] = portfolioItems
	config["portfolioClients"] = portfolioClients
	executeSimpleTemplate(w, r, "tmpl/content/index.html")
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio page
	config["title"] = "Our Work"
	config["portfolioCategories"] = portfolioCategories
	executeSimpleTemplate(w, r, "tmpl/content/portfolio.html")
}

func portfolioCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio category listing page
	vars := mux.Vars(r)
	category:=vars["category"]
	cat_ok:=false
	config["portfolioCategory"], cat_ok = portfolioCategories[category]
	if(!cat_ok) {
		notFoundHandler(w, r)
		return
	}
	config["title"] = strings.Join([]string{"Our", portfolioCategories[category].Name}, " ")
	executeSimpleTemplate(w, r, "tmpl/content/portfolio_category.html")
}

func portfolioDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// The portfolio details page
	vars := mux.Vars(r)
	
	category:=vars["category"]
	cat_ok:=false
	config["portfolioCategory"], cat_ok = portfolioCategories[category]
	if(!cat_ok) {
		notFoundHandler(w, r)
		return
	}
	
	project:=vars["project"]
	item_ok := false
	config["portfolioItem"], item_ok = portfolioItems[project]
	if(!item_ok) {
		notFoundHandler(w, r)
		return
	}
	
	config["portfolioClients"] = portfolioClients
	
	config["title"] = portfolioItems[project].Title
	
	executeSimpleTemplate(w, r, "tmpl/content/portfolio_details.html")
}