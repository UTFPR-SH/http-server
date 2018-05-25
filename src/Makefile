COMPILER=go
BUILD=build
INSTALL=install
SRC=main.go page.go server.go rest.go
EXEC=direc-server

all:
	$(COMPILER) $(BUILD) -o $(EXEC) $(SRC)

install:
	$(COMPILER) $(INSTALL) -o $(EXEC) $(SRC)

clear:
	rm $(EXEC)
	rm $GOPATH/bin/$(EXEC)
