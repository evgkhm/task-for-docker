.PHONY: \
	run \
	run swarm \

run:
	docker-compose up --build

run-swarm:
	docker swarm init
	docker stack deploy -c docker-stack.yml api