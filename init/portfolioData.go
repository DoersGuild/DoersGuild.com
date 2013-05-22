package init

type Person struct {
	ID, Name, Company, URL, IconURL string
}

var portfolioClients = map[string]Person {
	"DoersGuild" : {"DoersGuild", "Doers' Guild", "Doers' Guild", "http://doersguild.com", ""},
	"WissamSaleh" : {"WissamSaleh", "Wissam Saleh", "", "http://ientertain.mobi", ""},
	"IXXOCart" : {"IXXOCart", "IXXO Cart", "IXXO Cart", "http://ixxocart.com", "/img/projects/IXXOCart/logo.png"},
	"GKK_Rahul" : {"GKK_Rahul", "Rahul", "", "https://play.google.com/store/apps/details?id=com.doersguild.generalknowledge&reviewId=bGc6QU9xcFRPSDZUNEdCRlVLWjBGaF9WbGtaWDNVdFZFcjN2ZkZacXk5OVFpWEFwbkpwVk1HNmRKSlJlTTE5bEEwcS1rVG1sLThRRUlOYmRURU1Ca2F6NGFz", ""},
}

type Feedback struct {
	ProjectID, PersonID, Message string
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
    Client Person
    Feedback []Feedback
}

/* All our projects as an array */
var portfolioItems = map[string]PortfolioItem { 
	"GeneralKnowitallKnowledge" : {
		"GeneralKnowitallKnowledge",
		"General Knowitall Knowledge",
		"General Knowitall Knowledge is a fun android app for all the quizzing enthusiasts out there. It comes loaded with fun and interesting facts about various things from around the world.",
		"November, 2012",
		"MobileApps",
		map[string]string{"PlayStore" : "http://play.google.com", "Blog":"http://blog.doersguild.com"},
		"/img/projects/GeneralKnowitallKnowledge/logo.png",
		[]string{},
		[]string{"PhoneGap","Android", "Twitter Bootstrap"},
		portfolioClients["DoersGuild"],
		[]Feedback{
			{"GeneralKnowitallKnowledge", "GKK_Rahul", "Great app, brilliant questions nd i am sure u gonna love it..try it out..already in luv with this app <3 <3",},
		},
	},
	"iEntertain" : {
		"iEntertain",
		"i-Entertain",
		"A Mobile Gallery Application that shows it’s users, incredibly funny images from around the world, leaving them fully entertained. It is designed to run as both a website and a native android application and was built with the standard web technologies and packaged with PhoneGap. This application’s UI is fully responsive, so it works extremely well on all modern devices and their various screen sizes.",
		"September, 2012",
		"MobileApps",
		map[string]string{"PlayStore" : "http://play.google.com"},
		"/img/projects/iEntertain/logo.jpg",
		[]string{},
		[]string{"PhoneGap","Android", "jQuery Mobile"},
		portfolioClients["WissamSaleh"],
		[]Feedback{
			{"iEntertain", "WissamSaleh", "Very Responsive, really enjoyed working with you guys",},
		},
	},
	"IXXOCart" : {
		"IXXOCart",
		"IXXO Templating framework",
		"We helped IXXO-cart upgrade their templating framework with site-modules, twitter-bootstrap, JSON-web-services, etc and created several new templates for their various offerings including Multi-Vendor cart, Mobile templates, etc.",
		"May, 2013",
		"WebApps",
		map[string]string{"MVDemo" : "http://ixxocart.com", "Blog":"http://blog.doersguild.com"},
		"/img/projects/IXXOCart/logo.png",
		[]string{},
		[]string{"PHP", "Smarty", "Twitter Bootstrap", "IXXOCart"},
		portfolioClients["IXXOCart"],
		[]Feedback{
			{"IXXOCart", "IXXOCart", "The best development agency I've worked with in the last 20 years.",},
		},
	},
}


type Category struct {
	ID, Name string
	Items []PortfolioItem
}

/* The categories & the projects under them */
var portfolioCategories = map[string]Category {
	"MobileApps" : Category {
		"MobileApps",
		"Mobile Applications",
		[]PortfolioItem {
			portfolioItems["GeneralKnowitallKnowledge"],
			portfolioItems["iEntertain"],
		},
	},
	"WebApps" : Category {
		"WebApps",
		"Web Applications",
		[]PortfolioItem {
			portfolioItems["IXXOCart"],
		},
	},
}

/* Items to be displayed on the home page */
var homePageItems = []PortfolioItem {
	portfolioItems["IXXOCart"],
	portfolioItems["GeneralKnowitallKnowledge"],
	portfolioItems["iEntertain"],
}

/* Client feedback to be highlighted on the home page */
var homePageFeedback = []Feedback {
	portfolioItems["IXXOCart"].Feedback[0],
	portfolioItems["GeneralKnowitallKnowledge"].Feedback[0],
	portfolioItems["iEntertain"].Feedback[0],
}