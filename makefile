SOURCES = $(wildcard templates/*.templ)
TEMPLATES = $(SOURCES:.templ=_templ.go)

.PHONY: templates run build serve clean

all: templates build

run: templates serve

templates: $(TEMPLATES)

$(TEMPLATES): $(SOURCES) 
	templ generate

build:
	go build

serve:
	go run main.go

clean:
	rm -f $(wildcard templates/*.go) rename-me
