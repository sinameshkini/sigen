build:
	go build .

install:
	mkdir ${HOME}/.sigen
	cp config.yml ${HOME}/.sigen/