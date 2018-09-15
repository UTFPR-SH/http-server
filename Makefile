COMPILER=go
BUILD=build
INSTALL=install
SRC=main.go page.go server.go rest.go
EXEC=server

ifeq ($(OS),Windows_NT)
	EXEC=server.exe
endif

# The install might not work as it is only part of the software.
go:
	$(COMPILER) $(BUILD) -o $(EXEC) $(SRC)

clear:
	rm $(EXEC)
	rm $GOPATH/bin/$(EXEC)
