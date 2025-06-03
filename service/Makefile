core:
	go run cmd/core-service/main.go
upload-mp4:
	go run cmd/upload-mp4-service/main.go
encoding:
	go run cmd/encoding-service/main.go
video-hls:
	go run cmd/video-hls-service/main.go

container:
	docker start service-postgres-1 
	docker start service-redis-1 
	docker start rabbitmq-ctn