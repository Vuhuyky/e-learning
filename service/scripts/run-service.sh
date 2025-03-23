#!/bin/bash

# Chạy các service trong nền và lưu PID của chúng
go run cmd/core-service/main.go & pid1=$!
go run cmd/upload-mp4-service/main.go & pid2=$!
go run cmd/encoding-service/main.go & pid3=$!
go run cmd/video-hls-service/main.go & pid4=$!
go run cmd/blob-service/main.go & pid5=$!

# Định nghĩa hàm để dừng tất cả các process khi script bị dừng
cleanup() {
  echo "Stopping all services..."
  kill $pid1 $pid2 $pid3 $pid4 $pid5
  wait
  echo "All services stopped."
}

# Bắt tín hiệu Ctrl+C (SIGINT) và gọi cleanup
trap cleanup SIGINT

# Chờ các process hoàn thành
wait
