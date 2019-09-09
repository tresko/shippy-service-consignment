build:
	protoc --proto_path=./proto/consignment \
	  --micro_out=proto/consignment \
	  --go_out=plugins=micro:proto/consignment \
	  consignment.proto
	docker build -t shippy-service-consignment .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		shippy-service-consignment
