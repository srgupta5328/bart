package config

var (
	stations = map[string]string{
		"12th": "12th St. Oakland City Center",
		"16th": "16th St. Mission (SF)",
		"19th": "19th St. Oakland",
		"24th": "24th St. Mission (SF)",
		"ashb": "Ashby (Berkeley)",
		"antc": "Antioch",
		"balb": "Balboa Park (SF)",
		"bayf": "Bay Fair (San Leandro)",
		"cast": "Castro Valley",
		"civc": "Civic Center (SF)",
		"cols": "Coliseum",
		"colm": "Colma",
		"conc": "Concord",
		"daly": "Daly City",
		"dbrk": "Downtown Berkley",
		"dubl": "Dublin/Pleasenton",
		"deln": "El Cerrito del Norte",
		"plza": "El Cerrito Plaza",
		"embr": "Embarcadero",
		"frmt": "Fremont",
		"ftvl": "Fruitvale (Oakland)",
		"glen": "Glen Park (SF)",
		"hayw": "Hayward",
		"lafy": "Lafayette",
		"lake": "Lake Merrit (Oakland)",
		"mcar": "MacArthur (Oakland)",
		"mlbr": "Millbrae",
		"mont": "Montgomery St. (SF)",
		"nbrk": "North Berkley",
		"ncon": "Noth Concord/Martinez",
		"oakl": "Oakland Int'l Airport",
		"orin": "Orinda",
		"pitt": "Pittsburg/Bay Point",
		"pctr": "Pittsburg Center",
		"phil": "Pleasant Hill",
		"powl": "Powell St. (SF)",
		"rich": "Richmond",
		"rock": "Rockridge (Oakland)",
		"sbrn": "San Bruno",
		"sfia": "San Francisco Int'l Airport",
		"sanl": "San Leandro",
		"shay": "Sout Hayward",
		"ssan": "South San Francisco",
		"ucty": "Union City",
		"warm": "Warm Springs/South Fremont",
		"wcrk": "Walnut Creek",
		"wdub": "West Dublin",
		"woak": "West Oakland",
	}

	//key is the key used to access the BART API
	key = "key=MW9S-E7SL-26DU-VV8V"

	//jsonVal represents the payload in JSON format from the API
	jsonVal = "json=y"

	//cmd = "Different command values per API"
	cmd = []string{"bsa", "count", "elev", "help", "etd", "routes", "routeinfo"}

	//aspx = "determines the aspx per API"
	aspx = []string{"bsa.aspx", "etd.aspx", "route.aspx"}

	//Route Numbers
	route = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "all"}

	//baseURL API
	baseURL = "http://api.bart.gov/api/"
)
