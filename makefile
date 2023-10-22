SOURCES = $(wildcard templates/*.templ)
TEMPLATES = $(wildcard templates/*_templ.go)

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
	rm $(wildcard templates/*.go) caec-dashboard
