package route

import (
	"core/middleware"

	"github.com/labstack/echo"
)

func v1Routes(g *echo.Group, h AppModel) {
	g.GET("/health", h.Health.Check)
	auth := g.Group("/auth")
	auth.POST("/register", h.Auth.RegisterUser)
	auth.POST("/login", h.Auth.LoginUser)
	auth.GET("/test", h.Auth.Test, middleware.JWTVerify())

	role := g.Group("/role")
	role.POST("/insert", h.Role.Insert)
	role.GET("/list", h.Role.FindAll)

	customer := g.Group("/customer")
	customer.POST("/insert", h.Customer.Insert, middleware.JWTVerify())
	customer.GET("/get", h.Customer.GetAll, middleware.JWTVerify(), middleware.VerifyRoles("Sales"))
	customer.PUT("/update", h.Customer.UpdateCustomer, middleware.JWTVerify(), middleware.VerifyRoles("Sales"))

	billing := g.Group("/billing")
	billing.POST("/insert", h.Billing.Insert, middleware.JWTVerify())
	billing.PUT("/update", h.Billing.Update, middleware.JWTVerify(), middleware.VerifyRoles("Sales", "Accountant"))
	billing.GET("/get", h.Billing.List, middleware.JWTVerify(), middleware.VerifyRoles("Sales", "Accountant"))

	payment := g.Group("/payment")
	payment.POST("/insert", h.Payment.Insert, middleware.JWTVerify())
	payment.GET("/get", h.Payment.GetAll, middleware.JWTVerify(), middleware.VerifyRoles("HR", "Accountant"))
	payment.POST("/update", h.Payment.Update, middleware.JWTVerify(), middleware.VerifyRoles("HR", "Accountant"))

	user := g.Group("/user")
	user.PUT("/update", h.User.Update, middleware.JWTVerify(), middleware.VerifyRoles("Administrator"))
	user.GET("/list", h.User.ListUser, middleware.JWTVerify(), middleware.VerifyRoles("Administrator"))
}
