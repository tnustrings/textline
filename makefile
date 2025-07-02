all: bin/lily

bin/lily: lily.ct
	ct lily.ct
	go build -o bin/lily

.PHONY deb:
	cd deb; make
