DEP=$(GOPATH)/bin/dep
GO=go
FINAL=/usr/local/bin/
VERSION=0.0.1-beta
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"
PROG=pokeapi

all:
	# Dep is golang's official dependancy management tool, the /vendor direcotry(created by dep ensure) can act as another GOPATH source directory
	$(DEP) ensure
	cd cmd/$(PROG); $(GO) build $(LDFLAGS)
	mv cmd/$(PROG)/$(PROG) .
	
	

install:
	mkdir -p ~/.pokeapi
	cp --preserve $(PROG) $(FINAL)

clean:
	rm -f $(PROG)
	rm -rf vendor/

uninstall:
	rm $(FINAL)/$(PROG)
