all: bin/lingobin

bin/lingobin: lingobin.ct
	ct lingobin.ct
	go build -o bin/lingobin

.PHONY deb:
	cd deb; make
