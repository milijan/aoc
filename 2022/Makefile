SOURCES = $(wildcard */*.go)
BINARIES = $(basename $(SOURCES))
SUBDIRS = $(dir $(BINARIES))

%: %.go
	go build -o $@ $< 

.PHONY: %: $(SUBDIRS)
$(SUBDIRS): 
	go build -o $@/main $@/main.go

.PHONY: all
all: $(BINARIES)

.PHONY: clean
clean:
	@rm -f $(BINARIES)

