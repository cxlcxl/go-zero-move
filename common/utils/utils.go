package utils

import (
	"business/app/rbac/rpc/model"
	"business/common/vars"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func MD5(params string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// GeneratePassSecret 生成密码加密串
func GeneratePassSecret() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(15)
	if n < 3 {
		n = 8
	}
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@$&."
	bs := []byte(str)
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bs[rand.Intn(len(bs))]
	}
	return string(result)
}

// Password 登陆密码
func Password(pass, secret string) string {
	return strings.ToUpper(Base64Md5(secret + pass + secret))
}

// Base64Encode ...
func Base64Encode(params string) string {
	return base64.StdEncoding.EncodeToString([]byte(params))
}

// Base64Decode ...
func Base64Decode(params string) string {
	b, _ := base64.StdEncoding.DecodeString(params)
	return string(b)
}

// Sha256 ...
func Sha256(params string) string {
	h := sha256.New()
	h.Write([]byte(params))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Md5 先base64，然后MD5
func Base64Md5(params string) string {
	return MD5(base64.StdEncoding.EncodeToString([]byte(params)))
}

// Sha1 加密
func Sha1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// HttpBuildQuery 参数转化
func HttpBuildQuery(params map[string]interface{}) string {
	v := make(url.Values)
	for key, val := range params {
		v.Set(key, val.(string))
	}
	return v.Encode()
}

// JsonParamsQuery 参数转化
func JsonParamsQuery(params map[string]interface{}) string {
	jsonString, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonString)
}

// GetFileExt 获取文件后缀
func GetFileExt(fp multipart.File) string {
	buffer := make([]byte, 32)
	if _, err := fp.Read(buffer); err != nil {
		// 获取失败
		return ""
	}
	return http.DetectContentType(buffer)
}

// Round 保留小数位
func Round(f float64, n int) float64 {
	d, _ := decimal.NewFromFloat(f).Round(int32(n)).Float64()
	return d
}

// BufferConcat 字符串拼接
func BufferConcat(s []string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {
		buf.WriteString(s[i])
	}
	return buf.String()
}

// PostHttpRequest http 请求
func PostHttpRequest(host, data string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", host, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// RpcError 拆出 RPC 返回的错误
func RpcError(err error, emptyErrMsg string) error {
	if strings.Contains(err.Error(), "=") {
		errs := strings.Split(err.Error(), "=")
		e := strings.TrimSpace(errs[len(errs)-1])
		if strings.Contains(e, model.ErrNotFound.Error()) {
			return errors.New(emptyErrMsg)
		} else {
			return errors.New(e)
		}
	} else {
		return err
	}
}

// Pagination 分页信息
func Pagination(page, pageSize int64) (offset, limit uint64) {
	p := uint64(page)
	s := uint64(pageSize)
	if s == 0 {
		s = 10
	}
	if s > vars.MaxPageSize {
		s = vars.MaxPageSize
	}
	if p < 1 {
		p = 1
	}
	offset = (p - 1) * s
	return offset, s
}

func ArrayUnique(arr []int64) []int64 {
	size := len(arr)
	result := make([]int64, 0, size)
	temp := map[int64]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}
