package init

type Category struct {
	ID, Name string
}

var categories = []Category{
	{"mobileApps", "Mobile Applications"},
	{"webApps", "Web Applications"},
	{"browserExtensions", "Browser extensions"},
	{"plugins", "Plugins"},
	{"websites", "Websites & Templates"},
}

type PortfolioItem struct {
    Title string
    ShortDesc string
    Date string
    Screenshots []string
    LongDesc string
    Category string
}

var portfolioItems = []PortfolioItem {
	{
		"Hello",
		"World is wonderful",
		"Yesterday",
		[]string{"img","img2"},
		"World is wonderful, yes indeed",
		"Stuff",
	},
	{
		"Hello2",
		"World is wonderful",
		"Yesterday",
		[]string{"img","img2"},
		"World is wonderful, yes indeed!!!",
		"stuff",
	},
}