all: setup-cu setup-du

SOURCE = cmd
INSTDIR = bin

setup-cu:
	go build -o $(INSTDIR)/setup-cu $(SOURCE)/cu-config.go
setup-du:
	go build -o $(INSTDIR)/setup-du $(SOURCE)/du-config.go
clean:
	rm -rf bin
