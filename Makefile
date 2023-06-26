build:
	go build -o bin/
build-prod:
	go build -ldflags "-s -w" -o bin/ -tags prod
run:
	go run .
tests:
	go test ./...
clean:
	rm -rf bin/

build-mac: clean
	mkdir bin	
	fyne package -os darwin -icon Icon.png --release --tags prod
	mv wsui.app bin/

build-linux: clean
	mkdir bin
	fyne package -os linux -icon Icon.png --release --tags prod
	mv wsui.tar.xz bin/

build-win: clean
	mkdir bin
	fyne package -os windows -icon Icon.png --release --tags prod --appID websockets.wsui
	mv wsui.exe bin/