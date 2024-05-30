package main

import (
	"Demo/internal/handlers"
	"Demo/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 初始化用户存储库
	userRepo := repositories.NewUserRepository()

	// 初始化用户处理程序
	userHandler := handlers.NewUserHandler(userRepo)

	// 创建 Fiber 应用
	app := fiber.New()

	// 设置路由
	app.Get("/users", userHandler.GetAllUsers)

	// 启动 Fiber 服务
	app.Listen(":3000")
}
