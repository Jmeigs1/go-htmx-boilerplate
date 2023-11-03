SOURCES = $(shell find . -type f -name '*.templ')
TEMPLATES = $(SOURCES:.templ=_templ.go)

.PHONY: templates run build serve clean watch

all: templates build

run: templates serve
	
templ:
	go install github.com/a-h/templ/cmd/templ@latest && cp $(shell go env GOPATH)/bin/templ .

templates: templ $(TEMPLATES)

$(TEMPLATES): $(SOURCES) 
	./templ generate

watch: templ
	./templ generate --watch .

build:
	go build

serve:
	go run main.go

clean:
	rm -f $(shell find . -type f -name '*_templ.go') templ rename-me
