package route

import (
	"core/domain"
	"core/handler"
	"core/service"
)

type AppModel struct {
	Health          handler.HealthHandler
	User            handler.UserHandler
	Auth            handler.AuthHandler
	Role            handler.RoleHandler
	Book            handler.BookHandler
	BookSummary     handler.BookSummaryHandler
	Cart            handler.CartHandler
	Workspace       handler.WorkspaceHandler
	Channel         handler.ChannelHandler
	Document        handler.DocumentHandler
	ManageWorkspace handler.ManageWorkspaceHandler
	ManageChannels  handler.ManageChannelsHandler
	Credentials     handler.CredentialsHandler
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
	workspaceDomain := &domain.WorkspaceDomainCtx{}
	channelDomain := &domain.ChannelDomainCtx{}
	documentDomain := &domain.DocumentDomainCtx{}
	manageWorkspaceDomain := &domain.ManageWorkspaceDomainCtx{}
	manageChannelsDomain := &domain.ManageChannelsDomainCtx{}
	credentialsDomain := &domain.CredentialsDomainCtx{}

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
	workspaceService := service.WorkspaceService{
		WorkspaceDomain:       workspaceDomain,
		ManageWorkspaceDomain: manageWorkspaceDomain,
		ChannelDomain:         channelDomain,
		DocumentDomain:        documentDomain,
		CredentialsDomain:     credentialsDomain,
		ManageChannelsDomain:  manageChannelsDomain,
		UserDomain:            userDomain,
	}
	channelService := service.ChannelService{
		ChannelDomain:         channelDomain,
		ManageChannelsDomain:  manageChannelsDomain,
		UserDomain:            userDomain,
		ManageWorkspaceDomain: manageWorkspaceDomain,
	}
	documentService := service.DocumentService{
		DocumentDomain: documentDomain,
	}
	manageWorkspaceService := service.ManageWorkspaceService{
		ManageWorkspaceDomain: manageWorkspaceDomain,
	}
	manageChannelsService := service.ManageChannelsService{
		ManageChannelsDomain: manageChannelsDomain,
	}
	credentialsService := service.CredentialsService{
		CredentialsDomain: credentialsDomain,
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
	workspaceHandler := handler.WorkspaceHandler{
		WorkspaceService: workspaceService,
	}
	channelHandler := handler.ChannelHandler{
		ChannelService: channelService,
	}
	documentHandler := handler.DocumentHandler{
		DocumentService: documentService,
	}
	manageWorkspaceHandler := handler.ManageWorkspaceHandler{
		ManageWorkspaceService: manageWorkspaceService,
	}
	manageChannelsHandler := handler.ManageChannelsHandler{
		ManageChannelsService: manageChannelsService,
	}
	credentialsHandler := handler.CredentialsHandler{
		CredentialsService: credentialsService,
	}

	return AppModel{
		Health:          healthHandler,
		User:            userHandler,
		Auth:            authHandler,
		Role:            roleHandler,
		Book:            bookHandler,
		BookSummary:     bookSummaryHandler,
		Cart:            cartHandler,
		Workspace:       workspaceHandler,
		Channel:         channelHandler,
		Document:        documentHandler,
		ManageWorkspace: manageWorkspaceHandler,
		ManageChannels:  manageChannelsHandler,
		Credentials:     credentialsHandler,
	}
}
