package util

var (
	countryRegionAliasMap = map[string]string{
		"Türkiye":                    "Turkey",
		"Trinidad and Tobago":        "Trinidad & Tobago",
		"Taiwan (Province of China)": "Taiwan",
		"Syrian Arab Republic":       "Syria",
		"Réunion":                    "Reunion",
		"Iran (Islamic Republic of)": "Iran",
		"Hong Kong SAR China":        "Hong Kong",
		"Hong Kong (Special Administrative Region of China)": "Hong Kong",
		"Côte d’Ivoire":          "Cote d'Ivoire",
		"Bosnia and Herzegovina": "Bosnia & Herzegovina",
	}
)

func CountryRegionConvert(countryRegion string) string {
	name, ok := countryRegionAliasMap[countryRegion]
	if ok {
		return name
	} else {
		return countryRegion
	}
}
