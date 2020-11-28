.PHONY: build

build:
	sam build

invoke: build
	sam local invoke --event ./event.json
