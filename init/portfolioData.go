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
	"Unitu": {"Unitu","Anish","Unitu","https://www.elance.com/e/baggaanish/", ""},
	"HihnSight": {"HihnSight","HihnSight","","https://www.elance.com/e/HihnSight/", ""},
	"MigRyes": {"MigRyes","Mig Ryes","37Signals","http://migreyes.com/", "http://migreyes.com/assets/images/mig.jpg"},
	"Rebel316": {"Rebel316","BigPlans","Rebel316","https://www.elance.com/e/bigplans/", ""},
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
	"Unitu": {
		"Unitu",
		"Unitu Cross-Browser Compatiability",
		"We helped fix a lot of cross-browser compatiability issues on Unitu's newest responsive template and ensured that it supported almost every modern platform including IE8.",
		"February, 2013",
		"Bugfixes",
		map[string]string{"Elance" : "https://www.elance.com/j/front-end-developer-needed-make-fixes-current-website/36709339/"},
		"",
		[]string{},
		[]string{"Responsive", "Cross-Browser", "IE8"},
		portfolioClients["Unitu"],
		[]Feedback{
			{"Unitu","Unitu", "Contractor was fantastic! Was an expert in his field and very responsive, kind and patient. Will definitely be working with him again!"},
		},
	},
	"PassHasher": {
		"PassHasher",
		"Password Hasher",
		"Password Hasher is an open-source browser extension for Chrome, Firefox, Safari and IE that makes it easy for you to create a unique, highly secure password for every domain, by hashing your password, combined with the current domain using the PBKDF2 algorithm. It is also available as a web-app.",
		"November, 2012",
		"BrowserExtensions",
		map[string]string{"Password Hasher Web-App":"http://passhasher.com/", "Blog Post":"http://blog.doersguild.com/post/42079891079/beat-the-domino-effect-secure-your-passwords","Elance" : "https://www.elance.com/j/simple-browser-plugin-but-needed-ff-chrome-ie/33109099/"},
		"/img/projects/PassHasher/logo.png",
		[]string{"/img/projects/PassHasher/desktop1.png", "/img/projects/PassHasher/desktop2.png", "/img/projects/PassHasher/tablet1.png", "/img/projects/PassHasher/mobile1.png"},
		[]string{"CrossRider", "Twitter Bootstrap", "Offline Web-App", "Security"},
		portfolioClients["HihnSight"],
		[]Feedback{
			{"PassHasher","HihnSight", "Absolute joy to work with. No communications issues. Delivered better than promised."},
		},
	},
	"ToggleText": {
		"ToggleText",
		"ToggleText",
		"Toggle-Text is a cross-browser extension that lets you hide the text on a page, so that you can better appreciate it’s design.",
		"May, 2012",
		"BrowserExtensions",
		map[string]string{"Chrome Webstore":"https://chrome.google.com/webstore/detail/toggle-text/hfccicooedcnpfdpkdbokgehcnhchfaj", "Blog Post":"http://blog.doersguild.com/post/48191208744/toggle-text-look-at-websites-without-their-text","CrossRider Installation Page" : "http://crossrider.com/install/31274-toggle-text"},
		"/img/projects/ToggleText/logo.png",
		[]string{"/img/projects/ToggleText/desktop1.png"},
		[]string{"CrossRider", "Chrome Extension", "Design"},
		portfolioClients["DoersGuild"],
		[]Feedback{
			{"ToggleText","MigRyes", "Wow, thanks so much. This is super cool. Sharing with the team!"},
		},
	},
	"Rebel316Blog": {
		"Rebel316Blog",
		"Rebel316 Blogspot",
		"We designed a Blogspot theme for Rebel316, keeping in line with their website.",
		"July, 2012",
		"Websites",
		map[string]string{"Live Site":"http://blog.rebel316.com/", "Elance" : "https://www.elance.com/j/create-sites-blog/32049124/"},
		"/img/projects/Rebel316/desktop1.png",
		[]string{"/img/projects/Rebel316/desktop1.png"},
		[]string{"Blogspot"},
		portfolioClients["Rebel316"],
		[]Feedback{
			{"Rebel316Blog", "Rebel316", "Ravi has continually performed incredibly well and exceeded standards for a job perfectly executed! Well done!"},
		},
	},
	"DoersWebsite": {
		"DoersWebsite",
		"This website",
		"We designed and handcoded this beautiful, fully responsive website, with Twitter Bootstrap for the frontent and GoLang on the backend.",
		"May, 2012",
		"Websites",
		map[string]string{"Live Site":"http://www.doersguild.com/"},
		"/img/projects/DoersWebsite/desktop1.png",
		[]string{
			"/img/projects/DoersWebsite/desktop1.png", "/img/projects/DoersWebsite/tablet1.png", "/img/projects/DoersWebsite/mobile1.png",
			"/img/projects/DoersWebsite/desktop2.png", "/img/projects/DoersWebsite/tablet2.png", "/img/projects/DoersWebsite/mobile2.png",
			"/img/projects/DoersWebsite/desktop3.png", "/img/projects/DoersWebsite/tablet3.png", "/img/projects/DoersWebsite/mobile3.png",
		},
		[]string{"Twitter Bootstrap", "Responsive", "GoLang", "FLatUI"},
		portfolioClients["Rebel316"],
		[]Feedback{ },
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
			portfolioItems["PassHasher"],
		},
	},
	"Websites" : Category {
		"Websites",
		"Websites and Templates",
		[]PortfolioItem {
			portfolioItems["DoersWebsite"],
			portfolioItems["IXXOCart"],
			portfolioItems["CheriyalAndPembarthy"],
			portfolioItems["Rebel316Blog"],
		},
	},
	"Bugfixes" : Category {
		"Bugfixes",
		"Bugfixes and Upgrades",
		[]PortfolioItem {
			portfolioItems["Unitu"],
		},
	},
	"BrowserExtensions" : Category {
		"BrowserExtensions",
		"Browser Extensions",
		[]PortfolioItem {
			portfolioItems["PassHasher"],
			portfolioItems["ToggleText"],
		},
	},
}

/* Items to be displayed on the home page */
var homePageItems = []PortfolioItem {
	portfolioItems["IXXOCart"],
	portfolioItems["GeneralKnowitallKnowledge"],
	portfolioItems["PassHasher"],
	portfolioItems["CheriyalAndPembarthy"],
}

/* Client feedback to be highlighted on the home page */
var homePageFeedback = []Feedback {
	portfolioItems["Unitu"].Feedback[0],
	portfolioItems["PassHasher"].Feedback[0],
	portfolioItems["Rebel316Blog"].Feedback[0],
	portfolioItems["ToggleText"].Feedback[0],
}