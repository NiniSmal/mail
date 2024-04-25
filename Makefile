docker-build:
	docker build --no-cache -t mail:latest .

docker-run:
	docker rm -f mail && docker run -d --name mail -p 8095:8090 --network my-network mail:latest

kafka:
	docker rm -f tmanager-kafka && docker run -d --name tmanager-kafka --hostname tmanager-kafka \
                                       --network my-network \
                                       -e KAFKA_CFG_NODE_ID=0 \
                                       -e KAFKA_CFG_PROCESS_ROLES=controller,broker \
                                       -e KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093 \
                                       -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT \
                                       -e KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@tmanager-kafka:9093 \
                                       -e KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER \
                                       bitnami/kafka:latest

docker-logs:
	docker logs -f mail
