package sqli_mysql

import(
	//"core/data/replace"
)

type Error_based struct{
	Funcs map[string]string
}

func NewErrorBased() Error_based{
	new := Error_based{}

	funcs := make(map[string][]string, 3)

	funcs["Geo_functions1"] = []string{"SELECT ST_LatFromGeoHash(", ")",}
	funcs["Geo_functions2"] = []string{"SELECT ST_LongFromGeoHash(", ")",}
	funcs["Geo_functions3"] = []string{"SELECT ST_PointFromGeoHash(", ")",}

	return new
}