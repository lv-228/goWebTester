module web_tester

go 1.18

replace web_tester/http_funcs => ./http_funcs

require web_tester/http_funcs v0.0.0-00010101000000-000000000000

require (
	web_tester/server_ui v0.0.0-00010101000000-000000000000 // indirect
	web_tester/target v0.0.0-00010101000000-000000000000 // indirect
)

replace web_tester/server_ui => ./server_ui

replace web_tester/target => ./target
