
X_APP_NAME := sparrow

.PHONY: default
default:
	echo 'Choose a command'

.PHONY: build-run
build-run:
	make build
	./target/linux/$(X_APP_NAME)

.PHONY: build
build:
	go build \
	-trimpath \
	-o target/linux/$(X_APP_NAME) main.go
	cp config/config.yml target/linux/

.PHONY: clean
clean:
	rm -rf target
