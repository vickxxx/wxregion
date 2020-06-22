package wxregion

import "strings"

type Region struct {
	Country  string
	Province string
	City     string
}

func WxRegionToCN(country, province, city string) Region {
	region := Region{
		Country:  country,
		Province: province,
		City:     city,
	}
	if country == "" {
		return region
	}

	countryLowcase := strings.ToLower(country)
	countryCN := RegionMap[countryLowcase]
	if countryCN == "" {
		return region
	}
	region.Country = countryCN

	if province == "" {
		return region
	}
	provinceLowcase := strings.ToLower(province)
	cProv := strings.Join([]string{countryLowcase, provinceLowcase}, "_")
	provinceCN := RegionMap[cProv]
	if provinceCN == "" {
		return region
	}
	region.Province = provinceCN

	if city == "" {
		return region
	}
	cityLowcase := strings.ToLower(city)
	cProvCity := strings.Join([]string{countryLowcase, provinceLowcase, cityLowcase}, "_")
	cityCN := RegionMap[cProvCity]
	if cityCN == "" {
		return region
	}

	region.City = cityCN

	return region
}
