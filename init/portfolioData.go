package init

type Category struct {
	ID, Name string
}

type PortfolioItem struct {
	ID string
    Title string
    ShortDesc string
    Date string
    CategoryID string
    Links map[string]string
    Icon string
    Screenshots []string
    Tags []string
}

var portfolioCategories = map[string]Category {
	"MobileApps" : Category {"MobileApps", "Mobile Applications"},
	"WebApps" : Category {"WebApps", "Web Applications"},
}

var portfolioItems = map[string]interface{} {
	"MobileApps" : map[string]PortfolioItem {
		"GeneralKnowitallKnowledge" : {
			"GeneralKnowitallKnowledge",
			"General Knowitall Knowledge",
			"General Knowitall Knowledge is a fun android app for all the quizzing enthusiasts out there. It comes loaded with fun and interesting facts about various things from around the world.",
			"November, 2012",
			"MobileApps",
			map[string]string{"PlayStore" : "http:play.google.com", "Blog":"http://blog.doersguild.com"},
			"/img/logos/mobile/gkk.png",
			[]string{},
			[]string{"PhoneGap","Android", "Twitter Bootstrap"},
		},
		"iEntertain" : {
			"iEntertain",
			"i-Entertain",
			"A Mobile Gallery Application that shows it’s users, incredibly funny images from around the world, leaving them fully entertained. It is designed to run as both a website and a native android application and was built with the standard web technologies and packaged with PhoneGap. This application’s UI is fully responsive, so it works extremely well on all modern devices and their various screen sizes.",
			"September, 2012",
			"MobileApps",
			map[string]string{"PlayStore" : "http:play.google.com"},
			"/img/logos/mobile/i-entertain.jpg",
			[]string{},
			[]string{"PhoneGap","Android", "jQuery Mobile"},
		},
	},
	"WebApps" : map[string]PortfolioItem {
		"IXXOCart" : {
			"IXXOCart",
			"IXXO Templating framework",
			"We helped IXXO-cart upgrade their templating framework with site-modules, twitter-bootstrap, JSON-web-services, etc and created several new templates for their various offerings including Multi-Vendor cart, Mobile templates, etc.",
			"May, 2013",
			"WebApps",
			map[string]string{"MVDemo" : "http://ixxocart.com", "Blog":"http://blog.doersguild.com"},
			"/img/logos/mobile/ixxo.png",
			[]string{},
			[]string{"PHP", "Smarty", "Twitter Bootstrap", "IXXOCart"},
		},
	},
}

var homePageItems = []PortfolioItem {
	portfolioItems["WebApps"].(map[string]PortfolioItem)["IXXOCart"],
	portfolioItems["MobileApps"].(map[string]PortfolioItem)["GeneralKnowitallKnowledge"],
	portfolioItems["MobileApps"].(map[string]PortfolioItem)["iEntertain"],
}