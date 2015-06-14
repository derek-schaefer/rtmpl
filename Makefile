IMAGE=derekschaefer/rtmpl
TAG=$(shell git describe --tags --abbrev=0)

all: release

build:
	gox -os="darwin linux" -arch="amd64" -ldflags="-X main.VERSION $(TAG)" \
	  -output="bin/{{.Dir}}-$(TAG)-{{.OS}}-{{.Arch}}/{{.Dir}}"

release: clean build
	cd bin && ./release $(TAG)

image: release
	docker build -t $(IMAGE) .
	docker tag -f $(IMAGE) $(IMAGE):$(TAG)

push:
	docker push $(IMAGE)

clean:
	rm -rf bin/rtmpl-*
