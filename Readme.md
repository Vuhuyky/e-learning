# Online Learning Platform

## Giới thiệu

Đây là một ứng dụng web học online, nơi người dùng có thể:

- Đăng ký và đăng nhập tài khoản.
- Đăng tải các khóa học (dành cho giảng viên hoặc người tạo nội dung).
- Tìm kiếm và đăng ký học các khóa học (dành cho học viên).
- Xem nội dung khóa học và theo dõi tiến trình học tập.

## Tính năng chính

### Đối với người tạo khóa học:
- Đăng khóa học với các thông tin chi tiết: tiêu đề, mô tả, giá cả, nội dung, v.v.
- Quản lý các khóa học đã đăng.
- Theo dõi số lượng học viên tham gia.

### Đối với học viên:
- Duyệt và tìm kiếm khóa học theo từ khóa, danh mục, hoặc đánh giá.
- Đăng ký học các khóa học yêu thích.
- Theo dõi tiến trình học tập của bản thân.
- Để lại đánh giá và nhận xét cho khóa học.

### Tính năng chung:
- Quản lý người dùng (đăng ký, đăng nhập, quên mật khẩu).
- Hỗ trợ thanh toán (nếu cần thiết) cho các khóa học trả phí.
- Hệ thống đánh giá và bình luận cho khóa học.
- Responsive UI, hỗ trợ mọi thiết bị.

## Công nghệ sử dụng

### Backend
- **Ngôn ngữ:** Golang.
- **Framework:** Gin hoặc Fiber.
- **Database:** PostgreSQL hoặc MongoDB.
- **Authentication:** JWT (JSON Web Token).

### Frontend
- **Ngôn ngữ:** TypeScript.
- **Framework:** ReactJS.
- **UI Library:** Material-UI hoặc Tailwind CSS.

### DevOps
- **Containerization:** Docker.
- **CI/CD:** GitHub Actions.
- **Hosting:** AWS hoặc Vercel (cho frontend), Heroku (cho backend).

## Cài đặt và chạy dự án

### Yêu cầu:
- **Go:** >=1.20
- **Docker:** >=20.x