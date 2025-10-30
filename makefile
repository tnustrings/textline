all: bin/textline

bin/textline: textline.ct
	ct textline.ct
	go build -o bin/textline

.PHONY deb:
	cd deb; make
