core:
	go run cmd/core-service/main.go
upload-mp4:
	go run cmd/upload-mp4-service/main.go
encoding:
	go run cmd/encoding-service/main.go
video-hls:
	go run cmd/video-hls-service/main.go
quizz:
	go run cmd/quizz-service/main.go
blob:
	go run cmd/blob-service/main.go
# push mess to queue with quantity
# job delete video
all:
	go run cmd/core-service/main.go &
	go run cmd/upload-mp4-service/main.go &
	go run cmd/encoding-service/main.go &
	go run cmd/video-hls-service/main.go &
	go run cmd/quizz-service/main.go &
	go run cmd/blob-service/main.go &
	wait
