package init

import (
    //"fmt"
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/DoersGuild/paypal"
    "appengine"
    "strings"
    
    "appengine/urlfetch"
)

// Acommon variable with data for templates
var config = make(map[string]interface{})
var configFuncs = make(template.FuncMap)

var site_tagline="Your Web-Anywhere Experts"
var site_description = strings.Join([]string{"Doers' Guild", site_tagline}, " : ")
var site_image = "/favicon.ico"

const (
	PAYPAL_SANDBOX = true
	PAYPAL_USERNAME = "sathvik-facilitator_api1.doersguild.com"
	PAYPAL_PASSWORD = "1370333952"
	PAYPAL_SECRET = "AiPC9BjkCyDFQXbSkoZcgqH3hpacAbn-CXvfamuR2QddE1dzadDUDG7V"
	PAYPAL_CURRENCY = "USD"
)

func setupConfigDefaults() {
	// Setup common config vars to their default values
	
	// Reset config
	config = make(map[string]interface{})
	
	config["tagline"] = site_tagline
	config["title"] = site_tagline
	
	config["metaImage"] = site_image
	config["metaDescription"] = site_description
}


func init() {
	// The main router
	
	// Setup one-time common config vars
	setupConfigDefaults()
	
	// Setup one-time common template functions
	configFuncs["URLQueryEscaper"] = template.URLQueryEscaper
	
	router:=mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/init/{path:.*}", redirectOlderHandler)
	router.HandleFunc("/img/{path:.*}", imageRedirectHandler)
	router.HandleFunc("/portfolio", portfolioHandler)
	router.HandleFunc("/portfolio/{category}", portfolioCategoryHandler)
	router.HandleFunc("/portfolio/{category}/{project}", portfolioDetailsHandler)
	router.HandleFunc("/store/checkout/{previousPagePath:.*}", paypalCheckoutHandler)
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

func redirectOlderHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect images to the unlimited bandwidth hostinolder pages to the new home
	basePath := strings.Join([]string{"http://", r.Host}, "")
	http.Redirect(w, r, basePath, http.StatusMovedPermanently)
}

func executeSimpleTemplate(w http.ResponseWriter, r *http.Request, tmplFile string) {
	// Load and execute the given template
	
	// Log stats
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)
	c.Infof("Loading Template: %v", tmplFile)
	
	basePath := strings.Join([]string{"http://", r.Host}, "")
	config["basePath"] = basePath
	config["currentPath"] = r.URL.String()
	config["currentURL"] = strings.Join([]string{basePath, r.URL.String()}, "")
	
	var listTmpl = template.Must(template.New("mainTemplate").Funcs(configFuncs).ParseFiles("tmpl/base.html", "tmpl/blocks.html", tmplFile))
	
	if err := listTmpl.Execute(w, config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	// Restore default config
	setupConfigDefaults()
  
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// The 404 page
	config["title"] = "Page Not Found"
	executeSimpleTemplate(w, r, "tmpl/content/404.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// The home page
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
	config["metaDescription"] = portfolioCategories[category].Desc
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
	config["metaImage"] = portfolioItems[project].Icon
	config["metaDescription"] = portfolioItems[project].ShortDesc
	
	config["useSocialPlugins"] = true
	
	executeSimpleTemplate(w, r, "tmpl/content/portfolio_details.html")
}



func paypalCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	// Paypal express checkout page

	vars := mux.Vars(r)
	
	previousPagePath := vars["previousPagePath"]
	
	c := appengine.NewContext(r)
	httpClient := urlfetch.Client(c)
	
	client := paypal.NewClient(PAYPAL_USERNAME, PAYPAL_PASSWORD, PAYPAL_SECRET, PAYPAL_SANDBOX, httpClient)
	
	goods := make([]paypal.PayPalDigitalGood, 1)
	good := new(paypal.PayPalDigitalGood)
	good.Name, good.Amount, good.Quantity = "Test Good", 255, 1
	goods[0] = *good
	
	basePath := strings.Join([]string{"http://", r.Host}, "")
	returnURL := strings.Join([]string{basePath, "/", previousPagePath}, "")
	cancelURL := strings.Join([]string{basePath, "/", previousPagePath}, "")
	
	c.Infof("Paypal Return URL: %v \n", returnURL)
	
	response, err := client.SetExpressCheckoutDigitalGoods(255, PAYPAL_CURRENCY, returnURL, cancelURL, goods)

	c.Infof("Paypal Stuff: %v", response)
	c.Infof("Paypal Values: %v", response.Values)
	c.Infof("Paypal Error: %v", err)
	c.Infof("Paypal Checkout URL: %v", response.CheckoutUrl())
	
	http.Redirect(w, r, response.CheckoutUrl(), 301)
}