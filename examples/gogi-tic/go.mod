module github.com/yjp20/examples/gogi-tic

go 1.13

require (
	github.com/gorilla/sessions v1.2.0
	github.com/jinzhu/gorm v1.9.12
	github.com/yjp20/gogi v0.0.0-20200124082227-56534d0b2a5d
	golang.org/x/crypto v0.0.0-20200128174031-69ecbb4d6d5d // indirect
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9 // indirect
)

replace github.com/yjp20/gogi => ../..
