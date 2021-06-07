package handlers

import (
	"context"
	"log"
	"strings"

	"github.com/go-kit/kit/endpoint"

	accountPb "github.com/eedkevin/go-kit-boilerplate/account"
)

func AuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			log.Printf("received request in auth middleware: %+v", request)
			log.Printf("Authorization: %v", ctx.Value("Authorization"))
			authHeader := ctx.Value("Authorization").(string)
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			_, err := grpcClients.Account.AuthTokenValidate(ctx, &accountPb.AuthTokenValidateRequest{Token: token})
			if err != nil {
				log.Println("permission deny")
				return nil, err
			} else {
				return next(ctx, request)
			}
		}
	}
}
