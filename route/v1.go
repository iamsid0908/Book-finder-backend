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
	auth.POST("/google-login", h.Auth.GoogleLogin, middleware.VerifyGoogleToken())
	auth.GET("/logout", h.Auth.UserLogOut, middleware.JWTVerify())

	user := g.Group("/user", middleware.JWTVerify())
	user.GET("/get-user", h.User.GetUserName)

	book := g.Group("/books", middleware.JWTVerify())
	book.POST("/insert", h.Book.Insert)
	book.GET("/getall", h.Book.GellAllBook)
	book.POST("/bulk-insert", h.Book.BulkInsert)
	book.POST("/recommendation/books", h.Book.Recommend)

	booksummary := g.Group("/book-summary", middleware.JWTVerify())
	booksummary.POST("/insert", h.BookSummary.Insert)
	booksummary.GET("/book-detail/:book_id", h.BookSummary.GetBookDetails)

	cart := g.Group("/cart", middleware.JWTVerify())
	cart.POST("/insert", h.Cart.Insert)
	cart.GET("/get-cart", h.Cart.GetCartByUserId)
	cart.GET("/cart-size", h.Cart.GetSizeofCart)
	cart.DELETE("/cart-remove", h.Cart.RemoveFromCart)

}
