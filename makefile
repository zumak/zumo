
build:
	docker build -t bluemir/zumo-ip-report    components/ip-report
	docker build -t bluemir/zumo-dashboard    components/dashboard
	docker build -t bluemir/zumo-file-manager components/file-manager
	# special case
	docker build -t bluemir/zumo-grpc-gateway -f components/grpc-gateway/dockerfile .
run:
	docker stack deploy -c services.yml zumo
clean:
	docker stack rm zumo
go-env:
	docker run --rm -it -v `pwd`:/src --workdir /src golang bash
system: system/.htpasswd
	# ensure make proxy network
	docker stack deploy -c system.yml system
system/.htpasswd:
	htpasswd -c system/.htpasswd $(shell whoami)
hint:
	@echo export DOCKER_HOST=home.bluemir.me:12376

.PHONY: system
