
build:
	docker build -t bluemir/zumo-ip-report -f components/ip-report/dockerfile .
run:
	docker stack deploy -c services.yml zumo
clean:
	docker stack rm zumo
go-env:
	docker run --rm -it -v `pwd`:/src --workdir /src golang bash
