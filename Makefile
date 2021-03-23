all: clean build

build:
	cd src ; \
	go build -o timestamp ; \
	zip ../Timestamp-Alfred.alfredworkflow . -r --exclude=*.DS_Store* --exclude=.git/* --exclude=*.go --exclude=go.* --exclude="LICENSE" --exclude=".*"

clean:
	rm -f *.alfredworkflow