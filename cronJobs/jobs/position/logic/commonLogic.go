package logic

import (
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"fmt"
	"time"
)

var (
	positionElementTypes = []string{
		"image",
		"icon",
		"video",
		"title",
		"description",
		"corprate_name",
		"landing_page_url",
		"impression_tracking_url",
		"click_tracking_url",
	}
)

type requestPriceChan struct {
	Params statements.CreativeSizePriceRequest
	Token  string
}

type requestElementChan struct {
	Params statements.PositionElementRequest
	Token  string
}

func getTokenMap(ctx context.Context, svcCtx *svc.ServiceContext) (tokens map[int64]string, validAccountIds []int64, err error) {
	list, err := svcCtx.TokenModel.GetAccessTokenList(ctx)
	if err != nil {
		return nil, nil, err
	}
	tokens = make(map[int64]string)
	for _, token := range list {
		if token.ExpiredAt.After(time.Now()) {
			tokens[token.AccountId] = fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
			validAccountIds = append(validAccountIds, token.AccountId)
		}
	}
	return
}
