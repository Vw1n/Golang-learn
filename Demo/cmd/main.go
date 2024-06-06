package main

import (
	"Demo/internal/handlers"
	"Demo/internal/models"
	"Demo/internal/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 创建一个 gorm.DB 类型的变量
	var db *gorm.DB
	// 调用 Open 方法，传入驱动名和连接字符串
	db, err := gorm.Open(mysql.Open("root:Wen20040620@(localhost:3306)/test?parseTime=true"), &gorm.Config{})
	// 检查是否有错误
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	// 打印成功信息
	fmt.Println("连接数据库成功")
	// 自动迁移数据库
	db.AutoMigrate(&models.User{})

	// 初始化用户存储库
	userRepo := repositories.NewUserRepository(db)

	// 初始化用户处理程序
	userHandler := handlers.NewUserHandler(userRepo)

	// 创建 Fiber 应用
	app := fiber.New()

	// 设置路由
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/users", userHandler.CreateUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)

	// 启动 Fiber 服务
	app.Listen(":3000")
}
