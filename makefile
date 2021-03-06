
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
system:
	# ensure make proxy network
	docker stack deploy -c system.yml system
hint:
	@echo export DOCKER_HOST=home.bluemir.me:12376

.PHONY: system
