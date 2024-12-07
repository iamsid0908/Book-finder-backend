package route

import (
	"core/domain"
	"core/handler"
	"core/service"
)

type AppModel struct {
	Health      handler.HealthHandler
	User        handler.UserHandler
	Auth        handler.AuthHandler
	Role        handler.RoleHandler
	Book        handler.BookHandler
	BookSummary handler.BookSummaryHandler
	Cart        handler.CartHandler
}

func App() AppModel {
	//domain
	healthDomain := &domain.HealthDomainCtx{}
	authDomain := &domain.AuthDomainCtx{}
	userDomain := &domain.UserDomainCtx{}
	roleDomain := &domain.RoleDomainCtx{}
	bookDomain := &domain.BookDomainCtx{}
	bookSummaryDomain := &domain.BookSummaryDomainCtx{}
	cartDomain := &domain.CartDomainCtx{}

	//service
	healthService := service.HealthService{
		HealthDomain: healthDomain,
	}
	userService := service.UserService{
		UserDomain: userDomain,
	}
	authService := service.AuthService{
		AuthDomain: authDomain,
		UserDomain: userDomain,
	}
	roleService := service.RoleService{
		RoleDomain: roleDomain,
	}
	bookService := service.BookService{
		BookDomain:        bookDomain,
		BookSummaryDomain: bookSummaryDomain,
	}
	bookSummaryService := service.BookSummaryService{
		BookSummaryDomain: bookSummaryDomain,
	}
	cartService := service.CartService{
		CartDomain: cartDomain,
	}

	//handler
	healthHandler := handler.HealthHandler{
		HealthService: healthService,
	}
	userHandler := handler.UserHandler{
		UserService: userService,
	}
	authHandler := handler.AuthHandler{
		AuthService: authService,
	}
	roleHandler := handler.RoleHandler{
		RoleService: roleService,
	}
	bookHandler := handler.BookHandler{
		BookService: bookService,
	}
	bookSummaryHandler := handler.BookSummaryHandler{
		BookSummaryService: bookSummaryService,
	}
	cartHandler := handler.CartHandler{
		CartService: cartService,
	}

	return AppModel{
		Health:      healthHandler,
		User:        userHandler,
		Auth:        authHandler,
		Role:        roleHandler,
		Book:        bookHandler,
		BookSummary: bookSummaryHandler,
		Cart:        cartHandler,
	}
}
