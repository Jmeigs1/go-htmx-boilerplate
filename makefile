SOURCES = $(wildcard templates/*.templ)
TEMPLATES = $(SOURCES:.templ=_templ.go)

.PHONY: templates run build serve clean watch

all: templates build

run: templates serve

templates: $(TEMPLATES)

$(TEMPLATES): $(SOURCES) 
	templ generate

watch:
	templ generate --watch templates

build:
	go build

serve:
	go run main.go

clean:
	rm -f $(TEMPLATES) rename-me
