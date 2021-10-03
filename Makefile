.PHONY: run bundle generator generate clean

generate: bundle generator
	./generator

bundle:
	fyne bundle -package octicons ./resource/icons/ > ./bundle.go

generator:
	go build -o generator ./internal/cmd/generator

run: generate
	go run internal/cmd/fyne-octicons/*.go

clean:
	rm generator