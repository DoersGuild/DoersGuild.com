package init

type Person struct {
	ID, Name, Company, URL, IconURL string
}

var portfolioClients = map[string]Person {
	"DoersGuild" : {"DoersGuild", "Doers' Guild", "Doers' Guild", "http://doersguild.com", ""},
	"WissamSaleh" : {"WissamSaleh", "Wissam Saleh", "", "http://ientertain.mobi", ""},
	"IXXOCart" : {"IXXOCart", "IXXO Cart", "IXXO Cart", "http://ixxocart.com", "/img/projects/IXXOCart/logo.png"},
	"Tharuni" : {"Tharuni","Tharuni","Tharuni","http://tharuni.org","/img/projects/Tharuni/logo.png"},
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
		map[string]string{"Play Store" : "https://play.google.com/store/apps/details?id=com.doersguild.generalknowledge", "Ad-Free Edition":"https://play.google.com/store/apps/details?id=com.doersguild.gkkplus"},
		"/img/projects/GeneralKnowitallKnowledge/logo.png",
		[]string{"/img/projects/GeneralKnowitallKnowledge/mobile1.png", "/img/projects/GeneralKnowitallKnowledge/tablet1.png"},
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
		map[string]string{"Play Store" : "https://play.google.com/store/apps/details?id=com.WissamSaleh.iEntertain"},
		"/img/projects/iEntertain/logo.jpg",
		[]string{},
		[]string{"PhoneGap","Android", "jQuery Mobile"},
		portfolioClients["WissamSaleh"],
		[]Feedback{},
	},
	"IXXOCart" : {
		"IXXOCart",
		"IXXO Templates",
		"We helped IXXO-cart upgrade their templating framework with site-modules, twitter-bootstrap, JSON-web-services, etc and created several new templates for their various offerings including Multi-Vendor cart, Mobile templates, etc.",
		"May, 2013",
		"Websites",
		map[string]string{"Multi-Vendor Demo" : "http://demo-multi-vendor-sa.ixxocart.com/index.php"},
		"/img/projects/IXXOCart/logo.png",
		[]string{"/img/projects/IXXOCart/desktop1.png","/img/projects/IXXOCart/tablet1.png","/img/projects/IXXOCart/desktop2.png"},
		[]string{"PHP", "Smarty", "Twitter Bootstrap", "IXXOCart"},
		portfolioClients["IXXOCart"],
		[]Feedback{},
	},
	"CheriyalAndPembarthy": {
		"CheriyalAndPembarthy",
		"Cheriyal & Pembarthy Websites",
		"We built responsive e-commerce websites for the artisans of Cheriyal and Pembarthy villages to better educate the world about their artwork and sell it online.",
		"January, 2013",
		"Websites",
		map[string]string{"Cheriyal Paintings" : "http://cheriyalpainitngs.com", "Pembarthy Brassware":"http://pembarthybrassware.com"},
		"http://cheriyalpaintings.com/assets/DSC01542.jpg",
		[]string{"/img/projects/CheriyalAndPembarthy/desktop2.png","/img/projects/CheriyalAndPembarthy/desktop1.png","/img/projects/CheriyalAndPembarthy/tablet1.png","/img/projects/CheriyalAndPembarthy/tablet2.png","/img/projects/CheriyalAndPembarthy/mobile1.png","/img/projects/CheriyalAndPembarthy/mobile2.png"},
		[]string{"WordPress", "RootsTheme", "Twitter Bootstrap", "EShop"},
		portfolioClients["Tharuni"],
		[]Feedback{},
	},
	"GoSketch": {
		"GoSketch",
		"GoSketch",
		"GoSketch is a collaborative sketching tool that lets you doodle with your friends. We built this to explore an interesting new Node JS framework called Meteor JS.",
		"April, 2013",
		"WebApps",
		map[string]string{"GoSketch" : "http://gosketch.meteor.com", "Blog Post":"http://blog.doersguild.com/post/48109069431/gosketch-a-meteor-app-for-collaborative-doodling"},
		"/img/projects/goSketch/tablet1.png",
		[]string{"/img/projects/goSketch/desktop1.png","/img/projects/goSketch/tablet1.png"},
		[]string{"MeteorJS", "D3.js", "Twitter Bootstrap"},
		portfolioClients["DoersGuild"],
		[]Feedback{},
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
			portfolioItems["GoSketch"],
		},
	},
	"Websites" : Category {
		"Websites",
		"Websites and Templates",
		[]PortfolioItem {
			portfolioItems["IXXOCart"],
			portfolioItems["CheriyalAndPembarthy"],
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
	portfolioItems["GeneralKnowitallKnowledge"].Feedback[0],
}