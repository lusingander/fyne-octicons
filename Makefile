.PHONY: bundle

bundle:
	fyne bundle -package octicons ./resource/icons/ > ./bundle.go
