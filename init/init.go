package init

import (
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
    "appengine"
    "strings"
    "math/rand"
)

// Acommon variable with data for templates
var config = make(map[string]interface{})
var configFuncs = make(template.FuncMap)

const (
    SITE_TAGLINE="Your Web-Anywhere Experts"
    SITE_DESCRIPTION = "Doers' Guild : Your Web-Anywhere Experts"
    SITE_IMAGE = "/favicon.ico"
)

func setupConfigDefaults() {
    // Setup common config vars to their default values

    // Reset config
    config = make(map[string]interface{})

    config["tagline"] = SITE_TAGLINE
    config["title"] = SITE_TAGLINE

    config["metaImage"] = SITE_IMAGE
    config["metaDescription"] = SITE_DESCRIPTION
}

func noHTMLEscape(text string) template.HTML {
    // Return an unsanitized version of the HTML string
    return template.HTML(text)
}

func init() {
    // The main router


    // Randomize stuff for display (Runs once each time the app starts)

    // New random home page entries
    homePageItems = []PortfolioItem {}
    homePageItemKeys := []string {}
    currentItemCount = 0
    for currentItemCount < maxItems {
        for key,value := range portfolioItems {
            i := rand.Intn(100)
            j := rand.Intn(100)
            if(i < j && currentItemCount < maxItems) {
                exists := false
                for _,item := range homePageItemKeys {
                    exists = exists || (item == key)
                }
                if(!exists) {
                    homePageItems = append(homePageItems, value)
                    homePageItemKeys = append(homePageItemKeys, key)
                    currentItemCount += 1;
                }
            }
        }
    }
    // New random home page feedback entries
    currentFeedbackCount = 0
    homePageFeedback = []Feedback {}
    homePageFeedbackKeys := []string {}
    for currentFeedbackCount < maxItems {
        for key,value := range portfolioItems {
            i := rand.Intn(100)
            j := rand.Intn(100)
            feedback := value.Feedback
            feedbackCount := len(feedback)
            if(feedbackCount>0 && i < j && currentFeedbackCount < maxItems) {
                selected := feedback[rand.Intn(feedbackCount)]
                exists := false
                for _,item := range homePageFeedbackKeys {
                    exists = exists || (item == key)
                }
                if(!exists) {
                    homePageFeedback = append(homePageFeedback, selected)
                    homePageFeedbackKeys = append(homePageFeedbackKeys, key)
                    currentFeedbackCount += 1;
                }
            }
        }
    }

    // Setup one-time common config vars
    setupConfigDefaults()

    // Setup one-time common template functions
    configFuncs["URLQueryEscaper"] = template.URLQueryEscaper
    configFuncs["noHTMLEscape"] = noHTMLEscape

    router:=mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    router.HandleFunc("/init/{path:.*}", redirectOlderHandler)
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
    query := r.URL.Query()
    if query["target_width"] != nil {
        path = strings.Join([]string{path, "target_width"}, "?")
        width := strings.Join(query["target_width"], "")
        path = strings.Join([]string{path, width}, "=")
    }
    // Log stats
    c := appengine.NewContext(r)
    c.Infof("Redirecting to image: %v", path)
    // Pass along the resolution cookie, if exists
    resolutionCookie, resErr := r.Cookie("resolution")
    if resErr == nil {
        http.SetCookie(w, resolutionCookie)
        path = strings.Join([]string{path, resolutionCookie.String()}, "?")
        c.Infof("Setting resolution Cookie: %v", resolutionCookie.String())
    }
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

    protocol := "https://"
    if(r.TLS == nil) {
        // r.TLS is nil when no HTTPS connection is detected
        protocol = "http://"
    }
    basePath := strings.Join([]string{protocol, r.Host}, "")

    urlPathString := ""
    currentMuxRoute := mux.CurrentRoute(r)
    if(currentMuxRoute!=nil) {
        urlPath, urlPathError := currentMuxRoute.URLPath()

        if(urlPathError != nil) {
            // Remove the host and protocol parts of the URL to get the relative path
            urlPathString = strings.Replace(r.URL.String(), basePath, "", 1)
        } else {
            urlPathString = urlPath.String()
            c.Infof("Mux URL path: %v", urlPath)
        }
    } else {
        urlPathString = strings.Replace(r.URL.String(), basePath, "", 1)
    }

    config["basePath"] = basePath
    config["currentPath"] = urlPathString
    config["currentURL"] = strings.Join([]string{basePath, urlPathString}, "")

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
    w.WriteHeader(http.StatusNotFound)
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
    config["portfolioPageOrderedCategories"] = portfolioPageOrderedCategories
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
