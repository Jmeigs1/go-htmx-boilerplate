TEMPL_SOURCES = $(shell find . -type f -name '*.templ')
TEMPLATES = $(TEMPL_SOURCES:.templ=_templ.go)

GO_SOURCES = $(shell find . -type f -name '*.go')

.PHONY: templates run build serve clean watch

all: templates build

run: templates serve
	
templ:
	go install github.com/a-h/templ/cmd/templ@latest && cp $(shell go env GOPATH)/bin/templ .

templates: templ $(TEMPLATES)

$(TEMPLATES): $(TEMPL_SOURCES) 
	./templ generate

watch: templ
	./templ generate --watch .

build: rename-me

rename-me: $(GO_SOURCES) $(TEMPL_SOURCES)
	go build

serve:
	go run main.go

clean:
	rm -f $(shell find . -type f -name '*_templ.go') templ rename-me
