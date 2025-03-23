package osapp

func GetHostIP() (string, error) {
	// Gửi yêu cầu HTTP tới checkip.amazonaws.com để lấy IP công cộng
	// resp, err := http.Get("http://checkip.amazonaws.com")
	// if err != nil {
	// 	return "", fmt.Errorf("could not get public IP: %v", err)
	// }
	// defer resp.Body.Close()

	// // Đọc và trả về IP từ response body
	// ip, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", fmt.Errorf("could not read response: %v", err)
	// }

	// // Loại bỏ ký tự xuống dòng và khoảng trắng thừa
	// return strings.TrimSpace(string(ip)), nil
	return "localhost", nil
}
