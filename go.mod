module web_tester

go 1.18

replace web_tester/http_funcs => ./http_funcs

require web_tester/http_funcs v0.0.0-00010101000000-000000000000

require (
	core/sql v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	web_tester/internal/sqli v0.0.0-00010101000000-000000000000 // indirect
	web_tester/server_ui v0.0.0-00010101000000-000000000000 // indirect
	web_tester/target v0.0.0-00010101000000-000000000000 // indirect
)

replace web_tester/server_ui => ./server_ui

replace web_tester/target => ./target

replace web_tester/internal/sqli => ./internal/sqli

replace core/sql => ./core/sql
