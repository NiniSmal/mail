docker-build:
	docker build -t mail:latest .

docker-run:
	docker run -d --name mail -p 8095:8090 --network my-network mail:latest