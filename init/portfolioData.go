package init

type Person struct {
    ID, Name, Company, URL, IconURL string
}

var portfolioClients = map[string]Person {
    "DoersGuild" : {"DoersGuild", "Doers' Guild", "Doers' Guild", "http://doersguild.com", ""},
    "WissamSaleh" : {"WissamSaleh", "Wissam Saleh", "", "http://ientertain.mobi", ""},
    "IanMcLean" : {"IanMcLean", "Ian McLean", "", "https://www.linkedin.com/profile/view?id=13404532",""},
    "IXXOCart" : {"IXXOCart", "IXXO Cart", "IXXO Cart", "http://ixxocart.com", "/img/projects/IXXOCart/logo.png"},
    "Tharuni" : {"Tharuni","Tharuni","Tharuni","http://tharuni.org","/img/projects/Tharuni/logo.png"},
    "GKK_Rahul" : {"GKK_Rahul", "Rahul", "", "https://play.google.com/store/apps/details?id=com.doersguild.generalknowledge&reviewId=bGc6QU9xcFRPSDZUNEdCRlVLWjBGaF9WbGtaWDNVdFZFcjN2ZkZacXk5OVFpWEFwbkpwVk1HNmRKSlJlTTE5bEEwcS1rVG1sLThRRUlOYmRURU1Ca2F6NGFz", ""},
    "Unitu": {"Unitu","Anish","Unitu","https://www.elance.com/e/baggaanish/", ""},
    "HihnSight": {"HihnSight","HihnSight","","https://www.elance.com/e/HihnSight/", ""},
    "MigRyes": {"MigRyes","Mig Ryes","37Signals","http://migreyes.com/", ""},
    "Rebel316": {"Rebel316","BigPlans","Rebel316","https://www.elance.com/e/bigplans/", ""},
    "TouchPoint": {"TouchPoint", "MeDine", "TouchPoint Data Sciences Pvt. Ltd", "http://medine.in/", ""},
    "Tasheelat": {"Tasheelat", "Tasheelat", "Tasheelat", "http://TasheelatIraq.com/", ""},
    "FrenchBob": {"FrenchBob", "FrenchBob", "FrenchBob", "http://FrenchBob.com/", "http://frenchbob.com/modules/frenchbob/css/front/productSuggestionEngine/images/frenchbob.png"},
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
        map[string]string{"Play Store" : "https://play.google.com/store/apps/details?id=com.doersguild.generalknowledge", "Ad-Free Edition":"https://play.google.com/store/apps/details?id=com.doersguild.gkkplus", "Firefox OS": "https://marketplace.firefox.com/app/general-knowitall-knowledge?src=search"},
        "/img/projects/GeneralKnowitallKnowledge/logo.png",
        []string{"/img/projects/GeneralKnowitallKnowledge/mobile1.png", "/img/projects/GeneralKnowitallKnowledge/tablet1.png"},
        []string{"PhoneGap", "Android", "Twitter Bootstrap", "Firefox OS"},
        portfolioClients["DoersGuild"],
        []Feedback{
            {"GeneralKnowitallKnowledge", "GKK_Rahul", "Great app, brilliant questions nd i am sure u gonna love it..try it out..already in luv with this app <3 <3",},
        },
    },
    "iEntertain" : {
        "iEntertain",
        "i-Entertain",
        "A Mobile Gallery Application that shows its users, incredibly funny images from around the world, leaving them fully entertained. It is designed to run as both a website and a native android application and was built with the standard web technologies and packaged with PhoneGap.",
        "September, 2012",
        "MobileApps",
        map[string]string{"Play Store" : "https://play.google.com/store/apps/details?id=com.WissamSaleh.iEntertain"},
        "/img/projects/iEntertain/logo.jpg",
        []string{"/img/projects/iEntertain/I-Entertain_mobile11.png", "/img/projects/iEntertain/I-Entertain_mobile22.png",
        "/img/projects/iEntertain/I-Entertain_mobile33.png", "/img/projects/iEntertain/I-Entertain_mobile44.png"},
        []string{"PhoneGap","Android", "jQuery Mobile", "iOS"},
        portfolioClients["WissamSaleh"],
        []Feedback{},
    },
    "ProGrade" : {
        "ProGrade",
        "ProGrade",
        "Programmer Competency Calculator, based on the Programmer Competency Matrix from sijinjoseph.com We built it because we were bored and wanted to play around with the latest PhoneGap & CoffeeScript",
        "September, 2013",
        "MobileApps",
        map[string]string{"Play Store" : "https://play.google.com/store/apps/details?id=com.doersguild.procomp"},
        "/img/projects/ProGrade/logo.jpg",
        []string{"/img/projects/ProGrade/mobile1.jpg", "/img/projects/ProGrade/mobile2.jpg"},
        []string{"PhoneGap","Android", "Bootstrap", "CoffeeScript"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
     "MindOverMedia" : {
        "MindOverMedia",
        "Mind Over Media",
        "An application developed for the United States Holocaust Memorial Museum's Online Exhibit on Nazi Propaganda. This application allows museum visitors to view examples of propaganda and rate them in terms of how dangerous it is. It also allows them to submit pieces of propaganda that they see around them in their day to day lives.",
        "June, 2012",
        "MobileApps",
        map[string]string{"Elance" : "https://www.elance.com/j/jquery-mobile-application/30994515/"},
        "/img/projects/MindOverMedia/logo.jpg",
        []string{"/img/projects/MindOverMedia/mom_loading.png"},
        []string{"PhoneGap", "Android", "jQuery Mobile", "iOS"},
        portfolioClients["IanMcLean"],
        []Feedback{
            {"MindOverMedia", "IanMcLean", "Ravi and his team have done a great job building my mobile application. I'll will definitely continue to use them in the future.",},
        },
    },
    "MeDine" : {
        "MeDine",
        "MeDine",
        "MeDine is a Android application that makes finding and ordering your ideal meal all the more easier with a searchable interactive menu, exclusive offers and recommendations based on your tastes and reviews by those who have already tried a dish. We handled the app’s production for this complete redesign, which brings you a faster and more streamlined new UX with a bunch of new features that are sure to make your restaurant experience even more enjoyable.",
        "September, 2013",
        "MobileApps",
        map[string]string{"Play Store" : "http://play.google.com/store/apps/details?id=com.iglulabs.medine"},
        "/img/projects/MeDine/mobile1.jpg",
        []string{"/img/projects/MeDine/mobile1.jpg","/img/projects/MeDine/mobile2.jpg"},
        []string{"PhoneGap","Android", "jQuery Mobile", "Facebook"},
        portfolioClients["TouchPoint"],
        []Feedback{},
    },
    "IXXOCart" : {
        "IXXOCart",
        "IXXO Cart",
        "We’ve been working with IXXOCart to modernize their cart and make it the most powerful & comprehensive e-commerce software out there. We helped them upgrade major parts of the cart, created several new templates, & modernized and made all their interfaces responsive, across several of their offerings including their flagship multi-vendor software.",
        "May, 2013",
        "Websites",
        map[string]string{"IXXO Cart" : "http://www.ixxocart.com/"},
        "/img/projects/IXXOCart/logo.png",
        []string{"/img/projects/IXXOCart/desktop1.png","/img/projects/IXXOCart/tablet1.png","/img/projects/IXXOCart/desktop2.png"},
        []string{"PHP", "Smarty", "Twitter Bootstrap", "IXXOCart", "NVD3", "jQuery"},
        portfolioClients["IXXOCart"],
        []Feedback{},
    },
    "CheriyalAndPembarthy": {
        "CheriyalAndPembarthy",
        "Cheriyal & Pembarthy Websites",
        "We built responsive e-commerce websites for the artisans of Cheriyal and Pembarthy villages to better educate the world about their artwork and sell it online.",
        "January, 2013",
        "Websites",
        map[string]string{},
        "/img/projects/CheriyalAndPembarthy/desktop2.png",
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
    "CodePad": {
        "CodePad",
        "Code Pad",
        "CodePad is a web-based Real-time Collaborative editor. It allows one to code/write together with other people, from almost any computer connected to the Internet. Whatever one of the participants writes is instantly shared with all of their peers, making it very easy to write things collaboratively. It has support for highlighting the syntax of multiple programming languages.",
        "May, 2012",
        "WebApps",
        map[string]string{"CodePad" : "http://notepyd.appspot.com"},
        "/img/projects/CodePad/laptop1.jpg",
        []string{"/img/projects/CodePad/Codepadinfo.jpg"},
        []string{"Python", "GAE", "MobWrite", "CodeMirror"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "FreeWorkExchange": {
        "FreeWorkExchange",
        "Free Work Exchange",
        "A StackOverflow-like platform for FreeWork, where people can share work to be done and show-off their solutions to the world. Free-Work is about working freely without an employee/employer relationship. It allows you to work on real-world problems with real people, at your own time and in your own pace.",
        "May, 2012",
        "WebApps",
        map[string]string{"Free Work Exchange" : "http://freeworkexchange.appspot.com"},
        "/img/projects/FreeWorkExchange/freeworkxchange_screenshot.png",
        []string{"/img/projects/FreeWorkExchange/freeworkxchange_screenshot.png"},
        []string{"Python", "Web2Py", "GAE", "JavaScript", "CSS3"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "Xervmon": {
        "Xervmon",
        "Xervmon",
        "Xervmon is a cloud asset management solution that helps enterprises manage, monitor, plan and deploy cloud servers, databases, etc. We have been handling most of their UI, building complex interfaces and reusable components that form the fore-front of their offering. We have also been helping them integrate better with various cloud vendors and their monitoring APIs.",
        "September, 2012",
        "WebApps",
        map[string]string{"Xervmon" : "http://xervmon.com"},
        "/img/projects/Xervmon/xervmon_logo.png",
        []string{"/img/projects/Xervmon/Xervmon_main.jpg"},
        []string{"PHP", "CSS3", "jQuery", "NVD3", "Twitter Bootstrap"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "Tharunopayam": {
        "Tharunopayam",
        "Tharunopayam",
        "Tharunopayam is an SMS based helpline, currently targeted at helping adolescent girls, women and the aged in Andhra Pradesh. It uses an android device as an SMS server and a laravel based responsive front-end which allows our experts in various fields like nutrition, psychology, law, etc to answer peoples' queries from wherever and whenever they find it most comfortable.",
        "September, 2013",
        "WebApps",
        map[string]string{"Tharunopayam" : "http://upayam.tharuni.org/"},
        "/img/projects/Tharunopayam/Tharunopayam_logo.jpg",
        []string{"/img/projects/Tharunopayam/Tharunopayam.jpg"},
        []string{"Android", "Laravel", "Google Language tools", "Mandrill", "Twitter Bootstrap"},
        portfolioClients["Tharuni"],
        []Feedback{},
    },
    "Unitu": {
        "Unitu",
        "Unitu Cross-Browser Compatiability",
        "We helped fix a lot of cross-browser compatibility issues on Unitu's newest responsive template and ensured that it supported almost every modern platform including IE8.",
        "February, 2013",
        "Bugfixes",
        map[string]string{"Elance" : "https://www.elance.com/j/front-end-developer-needed-make-fixes-current-website/36709339/"},
        "https://lh4.googleusercontent.com/-a9t8aG8hMdc/AAAAAAAAAAI/AAAAAAAAABQ/UsB6EGRRDoE/s120-c/photo.jpg",
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
        "Toggle-Text is a cross-browser extension that lets you hide the text on a page, so that you can better appreciate its design.",
        "May, 2012",
        "BrowserExtensions",
        map[string]string{"Chrome Webstore":"https://chrome.google.com/webstore/detail/toggle-text/hfccicooedcnpfdpkdbokgehcnhchfaj", "Blog Post":"http://blog.doersguild.com/post/48191208744/toggle-text-look-at-websites-without-their-text","CrossRider Installation Page" : "http://crossrider.com/install/31274-toggle-text"},
        "/img/projects/ToggleText/logo.png",
        []string{"/img/projects/ToggleText/desktop1.png"},
        []string{"CrossRider", "Chrome Extension", "Design"},
        portfolioClients["DoersGuild"],
        []Feedback{
            {"ToggleText", "MigRyes", "Wow, thanks so much. This is super cool. Sharing with the team!"},
        },
    },
    "Rebel316Blog": {
        "Rebel316Blog",
        "Rebel316 Blogspot",
        "We designed a Blogspot theme for Rebel316, keeping in line with their website.",
        "July, 2012",
        "Websites",
        map[string]string{"Elance" : "https://www.elance.com/j/create-sites-blog/32049124/"},
        "/img/projects/Rebel316/desktop1.png",
        []string{"/img/projects/Rebel316/desktop1.png"},
        []string{"Blogspot"},
        portfolioClients["Rebel316"],
        []Feedback{
            {"Rebel316Blog", "Rebel316", "Ravi has continually performed incredibly well and exceeded standards for a job perfectly executed! Well done!"},
        },
    },
    "Tharuni": {
        "Tharuni",
        "Tharuni.org",
        "Tharuni is a non-profit working towards empowerment of adolescent girls, women and the aged. We designed their web-presence, with a fully responsive website that runs virtually everywhere.",
        "September, 2012",
        "Websites",
        map[string]string{"Live Site" : "http://tharuni.org/"},
        "/img/projects/Tharuni/logo.jpg",
        []string{"/img/projects/Tharuni/Tharuni.jpg"},
        []string{"jQuery", "Twitter Bootstrap"},
        portfolioClients["Tharuni"],
        []Feedback{},
    },
    "tasheelatiraq": {
        "tasheelatiraq",
        "tasheelatiraq.com",
        "We helped develop TasheelatIraq.com in English and Kurdish languages for an Iraqi car-dealership",
        "May, 2013",
        "Websites",
        map[string]string{"Live Site": "http://tasheelatiraq.com/"},
        "/img/projects/tasheelatiraq/laptop1.jpg",
        []string{"/img/projects/tasheelatiraq/Thasheelath.jpg"},
        []string{"jQuery", "PHP", "Twitter Bootstrap"},
        portfolioClients["Tasheelat"],
        []Feedback{},
    },
    "DoersWebsite": {
        "DoersWebsite",
        "This website",
        "We designed and handcoded this beautiful, fully responsive website, with Twitter Bootstrap for the frontend and GoLang on the backend.",
        "May, 2012",
        "Websites",
        map[string]string{"Live Site":"http://www.doersguild.com/"},
        "/img/projects/DoersWebsite/desktop1.png",
        []string{
            "/img/projects/DoersWebsite/desktop1.png", "/img/projects/DoersWebsite/tablet1.png", "/img/projects/DoersWebsite/mobile1.png",
            "/img/projects/DoersWebsite/desktop2.png", "/img/projects/DoersWebsite/tablet2.png", "/img/projects/DoersWebsite/mobile2.png",
            "/img/projects/DoersWebsite/desktop3.png", "/img/projects/DoersWebsite/tablet3.png", "/img/projects/DoersWebsite/mobile3.png",
        },
        []string{"Twitter Bootstrap", "Responsive", "GoLang", "Mostly-FlatUI"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "AutoPostingPlugin": {
        "AutoPostingPlugin",
        "WordPress Auto Posting Plugin",
        "A WordPress Plugin for automatically posting one article per day and a theme to go with it. This plugin is designed to create search-engine-optimized URLs for the posts. We built this so that Rebel316.com can have a fresh item-of-the-day, posted automatically everyday.",
        "June, 2012",
        "Plugins",
        map[string]string{"Elance": "https://www.elance.com/j/add-scripted-blog-webpage-that-will-show-one-main-large-picture/30948366/"},
        "/img/projects/AutoPostingPlugin/logo.png",
        []string{},
        []string{"WordPress", "PHP"},
        portfolioClients["Rebel316"],
        []Feedback{
            {"AutoPostingPlugin", "Rebel316", "Excellent job, saw the work through to the finish"},
        },
    },	
    "jQueryPrint": {
        "jQueryPrint",
        "jQuery.print",
        "Easy to use, Element Printing Plugin for jQuery",
        "July, 2012",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/jQuery.print"},
        "/img/projects/jQuery_print/jquery_print.png",
        []string{},
        []string{"jQuery"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "SquirtCrossrider": {
        "SquirtCrossrider",
        "Squirt",
        "Spritz/Squirt.io's speed-reading bookmarklet as a browser button for Chrome, Firefox & IE",
        "March, 2014",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/squirt-crossrider"},
        "/img/projects/SquirtCrossrider/crossrider.png",
        []string{},
        []string{"Cross-Rider", "Squirt.io"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "BracketsPHPLint": {
        "BracketsPHPLint",
        "PHPLint for Brackets",
        "Command line php -l powered PHP-linting extension for brackets",
        "February, 2014",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/brackets-phplint"},
        "/img/projects/BracketsPHPLint/logo.png",
        []string{},
        []string{"Brackets", "NodeJS", "PHP"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "GremlinsCrossrider": {
        "GremlinsCrossrider",
        "Gremlins.js Crossrider Plugin",
        "Monkey testing library marmelab/gremlins.js, as an extension for Chrome, Firefox & IE",
        "March, 2014",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/gremlins.js-crossrider"},
        "/img/projects/GremlinsCrossrider/crossrider.png",
        []string{},
        []string{"Cross-Rider", "Gremlins.js"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "SimpleImageSlider": {
        "SimpleImageSlider",
        "Simple Image Slider",
        "An extremely simple image slider powered by jQuery, built a part of a teaching exercise.",
        "July, 2013",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/jQuery.simpleImageSlider"},
        "/img/projects/SimpleImageSlider/jquery.png",
        []string{},
        []string{"jQuery", "CSS"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "HoverHeight": {
        "HoverHeight",
        "Hover Height",
        "A simple jQuery plugin that shortens elements to a specified height and expands them on hover",
        "May, 2013",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/jQuery.hoverHeight"},
        "/img/projects/HoverHeight/jquery.png",
        []string{},
        []string{"jQuery"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },
    "CordovaAndroidAnalytics": {
        "CordovaAndroidAnalytics",
        "Cordova-Android-Analytics",
        "Google Analytics v2 plugin for Cordova Android 1.9+",
        "November, 2012",
        "Plugins",
        map[string]string{"Github":"https://github.com/DoersGuild/Cordova-Android-Analytics"},
        "/img/projects/CordovaAndroidAnalytics/Cordova-Android-Analytics.png",
        []string{},
        []string{"Android", "Java", "Cordova"},
        portfolioClients["DoersGuild"],
        []Feedback{},
    },	
    "FrenchBob": {
        "FrenchBob",
        "FrenchBob",
        "FrenchBob is a novel e-commerce gifting platform that helps you select and buy the right gift. We handled the tech aspects of the store.",
        "October, 2014",
        "WebApps",
        map[string]string{"FrenchBob" : "http://FrenchBob.com"},
        "http://frenchbob.com/modules/frenchbob/css/front/productSuggestionEngine/images/frenchbob.png",
        []string{},
        []string{"PHP", "PrestaShop", "Twitter Bootstrap"},
        portfolioClients["FrenchBob"],
        []Feedback{},
    },
}


type Category struct {
    ID, Name string
    Items []PortfolioItem
    Desc string
}

/* The categories & the projects under them */
var portfolioCategories = map[string]Category {
    "MobileApps" : Category {
        "MobileApps",
        "Mobile Applications",
        []PortfolioItem {
            portfolioItems["GeneralKnowitallKnowledge"],
            portfolioItems["iEntertain"],
            portfolioItems["MindOverMedia"],
            portfolioItems["ProGrade"],
            portfolioItems["MeDine"],
        },
        "Want to build an app that you can release on Android, iPhone and Blackberry, all at once? Want an app that'll look great and work well no matter what kind of a screen the user has? We can build you a hybrid app that'll combine the speed of developing web-apps with the experience of native-smartphone-apps to give your users an equally great experience, no matter what device they use.",
    },
    "WebApps" : Category {
        "WebApps",
        "Web Applications",
        []PortfolioItem {
            portfolioItems["GoSketch"],
            portfolioItems["CodePad"],
            portfolioItems["FreeWorkExchange"],
            portfolioItems["Xervmon"],
            portfolioItems["Tharunopayam"],
            portfolioItems["FrenchBob"],
        },
        "Got a neat idea in mind? Want to build a scalable service? Looking to engage users, no matter what device they are on?  We'll build you a responsive, scalable web-based application that is designed to run beautifully on every modern device out there, interacting asynchronously with the backend through a simple RESTful API and with the power to save data locally when the user is offline, to give your users as smooth an experience as possible.",
    },
    "Websites" : Category {
        "Websites",
        "Websites and Templates",
        []PortfolioItem {
            portfolioItems["DoersWebsite"],
            portfolioItems["IXXOCart"],
            portfolioItems["CheriyalAndPembarthy"],
            portfolioItems["Rebel316Blog"],
            portfolioItems["Tharuni"],
            portfolioItems["tasheelatiraq"],
        },
        "Looking to interact better with your users on the various devices out there & improve conversions? Looking to upgrade your clunky old website to something modern and truly cross-platform? Have us build you a website that'll run really well on every modern-device, powered by the best practices in web-development & design.",
    },
    "Bugfixes" : Category {
        "Bugfixes",
        "Bugfixes and Upgrades",
        []PortfolioItem {
            portfolioItems["Unitu"],
        },
        "Does your app/site need an upgrade? Are you having trouble getting it to work on all the various devices out there? Is there a bug that's been troubling your users? Call us! And we'll take care of it in no time",
    },
    "BrowserExtensions" : Category {
        "BrowserExtensions",
        "Browser Extensions",
        []PortfolioItem {
            portfolioItems["PassHasher"],
            portfolioItems["ToggleText"],
            portfolioItems["SquirtCrossrider"],
            portfolioItems["GremlinsCrossrider"],
        },
        "Want to empower your users with a browser extension? Want to upgrade their experience on a particular website? Or help them find your business & your affiliates faster on Google? We can build you an easy to use browser extension for most of the popular desktop browsers out there!",
    },
    "Plugins" : Category {
        "Plugins",
        "Plugins",
        []PortfolioItem {
            portfolioItems["AutoPostingPlugin"],
            portfolioItems["jQueryPrint"],
            portfolioItems["BracketsPHPLint"],
            portfolioItems["SimpleImageSlider"],
            portfolioItems["HoverHeight"],
            portfolioItems["CordovaAndroidAnalytics"],			
        },
        "Want to empower your users with a cool plugin? Need jQuery or WordPress Plugins to enhance your site's performance? We can build them for you!",
    },
}

var itemCount = len(portfolioItems)
var maxItems = 4
var currentItemCount = 0
var currentFeedbackCount = 0

/* Items to be displayed on the home page */
var homePageItems = []PortfolioItem {
    portfolioItems["IXXOCart"],
    portfolioItems["GeneralKnowitallKnowledge"],
    portfolioItems["PassHasher"],
    portfolioItems["CheriyalAndPembarthy"],
}


/* Client feedback to be highlighted on the home page */
var homePageFeedback = []Feedback {
    portfolioItems["Rebel316Blog"].Feedback[0],
    portfolioItems["PassHasher"].Feedback[0],
    portfolioItems["Unitu"].Feedback[0],
    portfolioItems["ToggleText"].Feedback[0],
}

/* Ordered list of categories on the portfolio page */
var portfolioPageOrderedCategories = []Category {
    portfolioCategories["MobileApps"],
    portfolioCategories["WebApps"],
    portfolioCategories["Websites"],
    portfolioCategories["BrowserExtensions"],
    portfolioCategories["Bugfixes"],
    portfolioCategories["Plugins"],
}
