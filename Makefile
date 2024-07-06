linux:
	go build -o encrypt main.go
	@echo "#### Moving App to /usr/bin/ ####"
	sudo mv encrypt /usr/bin/

macos:
	go build -o encrypt main.go
	@echo "#### Moving App to /usr/local/bin/ ####"
	sudo mv encrypt /usr/local/bin/

windows:
	go build -o encrypt.exe main.go
	@echo "#### Moving App to C:\Windows\System32\ ####"
	move encrypt.exe C:\Windows\System32\
