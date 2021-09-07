package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"github.com/sjsbjyck/common"
)

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//获取token
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		tokenVal := meta["Token"]
		fmt.Println("==>", common.AppConf().GetString("jwt.sign"))
		jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.AppConf().GetString("jwt.sign")), nil
		})

		return fn(ctx, req, rsp)
	}
}
