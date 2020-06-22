package wxregion

import "strings"

type Region struct {
	CountryOrRegion string
	Province        string
	City            string
}

func WxRegionToCN(countryOrRegion, province, city string) Region {
	region := Region{
		CountryOrRegion: countryOrRegion,
		Province:        province,
		City:            city,
	}
	if countryOrRegion == "" {
		return region
	}

	countryOrRegionLowcase := strings.ToLower(countryOrRegion)
	countryOrRegionCN := RegionMap[countryOrRegionLowcase]
	if countryOrRegionCN == "" {
		return region
	}
	region.CountryOrRegion = countryOrRegionCN

	if province == "" {
		return region
	}
	provinceLowcase := strings.ToLower(province)
	cProv := strings.Join([]string{countryOrRegionLowcase, provinceLowcase}, "_")
	provinceCN := RegionMap[cProv]
	if provinceCN == "" {
		return region
	}
	region.Province = provinceCN

	if city == "" {
		return region
	}
	cityLowcase := strings.ToLower(city)
	cProvCity := strings.Join([]string{countryOrRegionLowcase, provinceLowcase, cityLowcase}, "_")
	cityCN := RegionMap[cProvCity]
	if cityCN == "" {
		return region
	}

	region.City = cityCN

	return region
}
