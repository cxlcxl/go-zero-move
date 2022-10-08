package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"context"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionElementLogic {
	return &PositionElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionElementLogic) PositionElement(req *types.PositionElementReq) (resp *types.PositionElementResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	size := strings.Split(req.CreativeSize, "*")
	w, _ := strconv.ParseInt(size[0], 0, 64)
	h, _ := strconv.ParseInt(size[1], 0, 64)
	elements, err := l.svcCtx.MarketingRpcClient.GetPositionElement(l.ctx, &marketingcenter.PositionElementReq{
		CreativeSizeId: req.CreativeSizeId,
		SubType:        req.SubType,
		Width:          uint64(w),
		Height:         uint64(h),
	})
	if err != nil {
		return nil, utils.RpcError(err, "未查询到版位元素信息")
	}
	var __elements []*types.ElementItem
	if err = copier.Copy(&__elements, elements.Elements); err != nil {
		return nil, err
	}
	return &types.PositionElementResp{
		BaseResp: successFul(),
		Data:     __elements,
	}, nil
}
