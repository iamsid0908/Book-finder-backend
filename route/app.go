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
	Customer    handler.CustomerHandler
	Billing     handler.BillingHandler
	Payment     handler.PaymentHandler
	Email       handler.EmailHandler
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
	customerDomain := &domain.CustomerDomainCtx{}
	billingDomain := &domain.BillingDomainCtx{}
	paymentDomain := &domain.PaymentDomainCtx{}
	emailDomain := &domain.EmailDomainCtx{}
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
	customerService := service.CustomerService{
		CustomerDomain: customerDomain,
	}
	billingService := service.BillingService{
		BillingDomain: billingDomain,
	}
	paymentService := service.PaymentService{
		PaymentDomain: paymentDomain,
	}
	emailService := service.EmailService{
		EmailDomain: emailDomain,
	}
	bookService := service.BookService{
		BookDomain: bookDomain,
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
	customerHandler := handler.CustomerHandler{
		CustomerService: customerService,
	}
	billingHandler := handler.BillingHandler{
		BillingService: billingService,
	}
	paymentHandler := handler.PaymentHandler{
		PaymentService: paymentService,
	}
	emailHandler := handler.EmailHandler{
		EmailService: emailService,
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
		Customer:    customerHandler,
		Billing:     billingHandler,
		Payment:     paymentHandler,
		Email:       emailHandler,
		Book:        bookHandler,
		BookSummary: bookSummaryHandler,
		Cart:        cartHandler,
	}
}
