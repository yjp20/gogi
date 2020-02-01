module github.com/yjp20/examples/gogi-tic

go 1.13

require (
	github.com/gorilla/sessions v1.2.0
	github.com/jinzhu/gorm v1.9.12
	github.com/yjp20/gogi v0.0.0-20200124082227-56534d0b2a5d
)

replace github.com/yjp20/gogi => ../..
