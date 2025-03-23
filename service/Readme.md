# encoding:
	cmd := exec.Command("ffmpeg",
		"-f", "webm", // Định dạng đầu vào
		"-i", "pipe:0", // Nhận từ stdin
		"-vf", "scale=-2:360", // Giảm độ phân giải
		"-preset", "ultrafast", // Cấu hình preset nhanh
		"-vcodec", "libx264", // Bộ mã hóa video H.264
		"-acodec", "aac", // Bộ mã hóa âm thanh AAC
		"-fflags", "+genpts", // Đảm bảo timestamp chính xác
		"-movflags", "+frag_keyframe+empty_moov", // Đảm bảo phát trực tiếp
		"-f", "mpegts", // Định dạng đầu ra là MPEG-TS
		"pipe:1", // Ghi ra stdout
	)

  cmd := exec.Command("ffmpeg",
		"-f", "webm", // Định dạng đầu vào
		"-i", "pipe:0", // Nhận từ stdin
		"-c:v", "copy", // Copy codec video, không mã hóa lại
		"-c:a", "copy", // Copy codec âm thanh, không mã hóa lại
		"-fflags", "+genpts", // Đảm bảo timestamp chính xác
		"-f", "mpegts", // Định dạng đầu ra là MPEG-TS
		"pipe:1", // Ghi ra stdout
	)

# note:
  - Bỏ service: stream-service, proxy-service, quantity-blob-service, merge-blob-service

# build docker:
  - docker build -t core-service -f docker/Dockerfile.core . && docker tag core-service:latest nguyen040904/core-service:latest && docker push nguyen040904/core-service:latest

  - docker build -t encoding-service -f docker/Dockerfile.encoding . && docker tag encoding-service:latest nguyen040904/encoding-service:latest && docker push nguyen040904/encoding-service:latest

  - docker build -t blob-service -f docker/Dockerfile.blob . && docker tag blob-service:latest nguyen040904/blob-service:latest && docker push nguyen040904/blob-service:latest

  - docker build -t upload-mp4-service -f docker/Dockerfile.upload-mp4 . && docker tag upload-mp4-service:latest nguyen040904/upload-mp4-service:latest && docker push nguyen040904/upload-mp4-service:latest

  - docker build -t video-hls-service -f docker/Dockerfile.video-hls . && docker tag video-hls-service:latest nguyen040904/video-hls-service:latest && docker push nguyen040904/video-hls-service:latest