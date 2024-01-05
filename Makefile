ifeq ($(OS),Windows_NT)
    TARGET_CLI = smecli.exe
else
    TARGET_CLI = smecli
endif

all: $(TARGET_CLI)

$(TARGET_CLI):
	cd smecli && go build -o $(TARGET_CLI) main.go devices.go values.go 

clean:
	rm -f smecli/$(TARGET_CLI)