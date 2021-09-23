.PHONY: buildForAWS buildForOSX copyJson clean deploy gomodgen

buildForWin: gomodgen
	export GO111MODULE=on
	env GOOS=windows go build -ldflags="-s -w" -o build/rsr.exe recursivesubdirremove/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

testWin: clean copyJson buildForWin
	go test -v -covermode count ./...

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh