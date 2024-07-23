package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// 	socketio "github.com/googollee/go-socket.io"
// 	"github.com/googollee/go-socket.io/engineio"
// )

// func main() {
// 	// Tạo ứng dụng Fiber
// 	app := fiber.New()

// 	// Tạo server Socket.IO
// 	server := socketio.NewServer(&engineio.Options{
// 		PingTimeout:  0,
// 		PingInterval: 0,
// 	})

// 	// Bản đồ lưu trữ các kết nối của người dùng theo user_id
// 	userConnections := make(map[string]socketio.Conn)

// 	// Xử lý sự kiện khi có kết nối mới
// 	server.OnConnect("/", func(s socketio.Conn) error {
// 		s.SetContext("")
// 		fmt.Println("connected:", s.ID())
// 		return nil
// 	})

// 	// Sự kiện để đăng ký một kết nối với user_id
// 	server.OnEvent("/", "register", func(s socketio.Conn, userID string) {
// 		fmt.Println("User registered:", userID)
// 		userConnections[userID] = s
// 		s.Emit("registered", "You are registered with user_id: "+userID)
// 	})

// 	// Sự kiện để gửi tin nhắn đến một người dùng cụ thể
// 	server.OnEvent("/", "send_message", func(s socketio.Conn, userID string, msg string) {
// 		if conn, ok := userConnections[userID]; ok {
// 			conn.Emit("receive_message", msg)
// 		} else {
// 			s.Emit("error", "User not found")
// 		}
// 	})

// 	// Sự kiện để xử lý khi người dùng rời đi
// 	server.OnEvent("/", "bye", func(s socketio.Conn) {
// 		last := s.Context().(string)
// 		s.Emit("bye", last)
// 		s.Close()
// 	})

// 	server.OnError("/", func(s socketio.Conn, e error) {
// 		fmt.Println("meet error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 		fmt.Println("closed", reason)
// 		// Dọn dẹp bản đồ userConnections
// 		for userID, conn := range userConnections {
// 			if conn.ID() == s.ID() {
// 				delete(userConnections, userID)
// 				break
// 			}
// 		}
// 	})

// 	// Khởi chạy server Socket.IO trong một goroutine
// 	go server.Serve()
// 	defer server.Close()

// 	// Tạo HTTP server và sử dụng Fiber để phục vụ các yêu cầu
// 	http.Handle("/socket.io/", server)
// 	http.Handle("/", app)

// 	// Khởi chạy server HTTP
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }
