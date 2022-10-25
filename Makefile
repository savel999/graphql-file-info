GQLGEN ?= github.com/99designs/gqlgen

################################################################################################################

.PHONY: gen
gen:
	go get -d github.com/99designs/gqlgen
	go run $(GQLGEN) generate --config ./gqlgen.yml


################################################################################################################

.PHONY: gen
run:
	go run server.go