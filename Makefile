CC = go
TARGET = awesome-trivia

.PHONY: build clean

$(TARGET):
	$(CC) build

build: $(TARGET)

clean:
	rm -rf $(TARGET)

