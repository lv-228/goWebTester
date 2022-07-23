module web_tester

go 1.18

require web_tester/target v0.0.0-00010101000000-000000000000

require (
	core/data v0.0.0-00010101000000-000000000000 // indirect
	core/data/json v0.0.0-00010101000000-000000000000 // indirect
	core/data/replace v0.0.0-00010101000000-000000000000 // indirect
	core/html v0.0.0-00010101000000-000000000000 // indirect
	core/html/domobjs v0.0.0-00010101000000-000000000000 // indirect
	core/http v0.0.0-00010101000000-000000000000 // indirect
	core/os v0.0.0-00010101000000-000000000000 // indirect
	core/sql v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	internal/sqli/modules v0.0.0-00010101000000-000000000000 // indirect
	internal/sqli/mysql v0.0.0-00010101000000-000000000000 // indirect
	web_tester/internal/sqli v0.0.0-00010101000000-000000000000 // indirect
	web_tester/web v0.0.0-00010101000000-000000000000 // indirect
)

replace web_tester/target => ./target

replace web_tester/internal/sqli => ./internal/sqli

replace core/sql => ./core/sql

replace core/data => ./core/data

replace core/html => ./core/html

replace core/http => ./core/http

replace core/os => ./core/os

replace web_tester/web => ./web

replace core/data/replace => ./core/data/replace_module

replace core/data/json => ./core/data/json

replace core/html/domobjs => ./core/html/DOMobjs

replace internal/sqli/modules => ./internal/sqli/modules

replace internal/sqli/mysql => ./internal/sqli/mysql
