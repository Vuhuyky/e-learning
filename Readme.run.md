# Chạy backend: docker start $(docker ps -a -q)

# Chạy từng service:
    // backend xử lí mp4
    docker start service-upload-mp4-service-1 => Lệnh này dùng để chạy backend (đơn lẻ)
    docker stop service-upload-mp4-service-1 => Lệnh này dùng để dừng backend (đơn lẻ)

    // backend xử lí nghiệp vụ chính
    docker start service-core-service-1 => Lệnh này dùng để chạy backend (đơn lẻ)
    docker stop service-core-service-1 => Lệnh này dùng để dừng backend (đơn lẻ)

    // backend xử lí chất lượng video
    docker start service-encoding-service-1 => Lệnh này dùng để chạy backend (đơn lẻ)
    docker stop service-encoding-service-1 => Lệnh này dùng để dừng backend (đơn lẻ)

    // backend xử lí phát video
    docker start service-video-hls-service-1 => Lệnh này dùng để chạy backend (đơn lẻ)
    docker stop service-video-hls-service-1 => Lệnh này dùng để dừng backend (đơn lẻ)

# Các service:
    docker start service-postgres-1
    docker stop service-postgres-1

    docker start service-redis-1
    docker stop service-redis-1

    docker start rabbitmq-ctn
    docker stop rabbitmq-ctn
