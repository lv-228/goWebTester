package core_html_domobjs

type Input struct{
	Main_data Html_object_data
}

func NewInput(name string, value string) Input{
	new_input := Input{}
	new_input.Main_data.Name = name
	new_input.Main_data.Value = value
	return new_input
}