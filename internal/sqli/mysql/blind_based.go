package sqli_mysql

type Time_based struct{
	Funcs map[string][]string
}

type Boolean_based struct{
	Funcs map[string][]string
}

func NewBooleanBased() Boolean_based{
	new := Boolean_based{}

	funcs := make(map[string][]string, 3)

	funcs["ascii"] = []string{"ASCII",}
	funcs["substring"] = []string{"SUBSTRING",}
	funcs["lenght"] = []string{"LENGHT",}

	return new
}