package logic

import (
	"business/app/application/rpc/appcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type AssetUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

var (
	fileSize = map[string]int64{
		vars.AssetTypeVideo:   100 << 20,
		vars.AssetTypePicture: 10 << 20,
	}
)

func NewAssetUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *AssetUploadLogic {
	return &AssetUploadLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AssetUploadLogic) AssetUpload(req *types.AssetUploadReq) (resp *types.BaseResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	filFullPath, err := l.fileSave(req)
	if err != nil {
		return nil, err
	}
	baseName := filepath.Base(filFullPath)
	splitName := strings.Split(baseName, ".")
	assetName := strings.Join(splitName[:len(splitName)-1], "")
	params := map[string]string{
		"asset_name": assetName,
		"file_token": req.FileToken,
	}
	request, err := newUploadRequest(l.svcCtx.Config.MarketingApis.Asset.Create, params, filFullPath)
	if err != nil {
		return nil, errors.New("请求创建失败：" + err.Error())
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("接口调用失败：" + err.Error())
	}
	defer response.Body.Close()
	// defer os.Remove(filFullPath) // 删除本地文件
	var rs statements.AssetCreateResponse
	if err = json.NewDecoder(response.Body).Decode(&rs); err != nil {
		return nil, errors.New("接口返回数据解析失败：" + err.Error())
	}
	if rs.Code != "200" {
		return nil, errors.New(fmt.Sprintf("接口调用失败：%s %s", rs.Code, rs.Message))
	}
	app, err := l.svcCtx.AppRpcClient.GetAppInfoByAppId(l.ctx, &appcenter.GetByAppIdReq{AppId: req.AppId})
	if err != nil {
		return nil, utils.RpcError(err, "应用信息有误")
	}
	_, err = l.svcCtx.MarketingRpcClient.BatchInsertAsset(l.ctx, &marketingcenter.BatchInsertAssetReq{
		Assets: []*marketingcenter.Asset{
			{
				AppId:             req.AppId,
				AssetId:           rs.Data.AssetId,
				AssetName:         assetName,
				AssetType:         req.AssetType,
				FileUrl:           rs.Data.Url,
				Width:             req.Width,
				Height:            req.Height,
				VideoPlayDuration: req.Duration,
				AccountId:         app.AccountId,
				AdvertiserId:      app.AdvertiserId,
			},
		},
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}

	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}, nil
}

func (l *AssetUploadLogic) fileSave(req *types.AssetUploadReq) (fullPath string, err error) {
	v, ok := fileSize[req.AssetType]
	if !ok {
		return "", errors.New("文件大小超出限制")
	}
	if err = l.r.ParseMultipartForm(v); err != nil {
		return "", errors.New("文件上传出错：" + err.Error())
	}
	file, h, err := l.r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 校验尺寸
	if err = checkAssetDimension(req.Width, req.Height, req.AssetType); err != nil {
		return "", err
	}
	_path := filepath.Join(l.svcCtx.Config.MarketingApis.Asset.AssetTmpPath, req.AppId)
	if err = dirIsExists(_path); err != nil {
		return "", err
	}
	fullPath = filepath.Join(_path, h.Filename)
	tempFile, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()
	w, err := io.Copy(tempFile, file)
	fmt.Println(w, err)
	return
}

func checkAssetDimension(width, height int64, assetType string) error {
	var dimensions []statements.AssetDimension
	if assetType == vars.AssetTypePicture {
		dimensions = statements.PictureDimensions
		return nil
	} else if assetType == vars.AssetTypePicture {
		dimensions = statements.VideoDimensions
	} else {
		return errors.New("素材类型错误")
	}
	for _, dimension := range dimensions {
		if dimension.Width == width && dimension.Height == height {
			return nil
		}
	}
	return errors.New("素材不满足尺寸要求")
}

func dirIsExists(d string) (err error) {
	if _, err = os.Stat(d); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(d, os.ModePerm)
		} else {
			return err
		}
	}
	return nil
}

// 新建上传请求
func newUploadRequest(link string, params map[string]string, path string) (*http.Request, error) {
	fp, err := os.Open(path) // 打开文件句柄
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	body := &bytes.Buffer{}                                         // 初始化body参数
	writer := multipart.NewWriter(body)                             // 实例化multipart
	part, err := writer.CreateFormFile("file", filepath.Base(path)) // 创建multipart 文件字段
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fp) // 写入文件数据到multipart
	if err != nil {
		return nil, err
	}
	for key, val := range params {
		_ = writer.WriteField(key, val) // 写入body中额外参数，比如七牛上传时需要提供token
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", link, body) // 新建请求
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType()) // 设置请求头,!!!非常重要，否则远端无法识别请求
	return req, nil
}
