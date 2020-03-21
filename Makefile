.PHONY: bundle generator generate

generate: bundle generator
	./generator

bundle:
	fyne bundle -package octicons ./resource/icons/ > ./bundle.go

generator:
	go build -o generator ./internal/cmd/generator

clean:
	rm generator