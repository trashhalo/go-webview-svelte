ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

deps:
	go mod tidy
	cd frontend && npm i

build:
	cd frontend && npm run build  
	$(GOPATH)/bin/pkger -include /frontend/public
	go build .