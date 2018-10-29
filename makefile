
build:
	docker build -f components/ip-report/dockerfile .
run:
	docker stack deploy -c service.yml zumo
clean:
	docker stack rm zumo
go-env:
	docker run --rm -it -v `pwd`:/src --workdir /src golang bash
