package utils

import (
	"business/common/vars"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
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

// HttpRequest http 请求
func HttpRequest(host, data, method string, headers map[string]string) ([]byte, error) {
	m := strings.ToUpper(method)
	if m != "POST" && m != "GET" {
		return nil, errors.New("不允许的请求方法")
	}
	req, err := http.NewRequest(m, host, strings.NewReader(data))
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
	if IsErrNotFound(err) {
		return errors.New(emptyErrMsg)
	} else {
		errs := strings.Split(err.Error(), "=")
		return errors.New(strings.TrimSpace(errs[len(errs)-1]))
	}
}

// IsErrNotFound 拆出 RPC 返回的错误
func IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), sqlx.ErrNotFound.Error())
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

// GetDsn mysql dsn 连接
func GetDsn(user, pass, host, name, charset string, port int64) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user, pass, host, port, name, charset,
	)
}

// WhereIn where in 条件组合
func WhereIn(s interface{}) (w string, args []interface{}, err error) {
	query := make([]string, 0)
	switch s.(type) {
	case []string:
		for i := range s.([]string) {
			query = append(query, "?")
			args = append(args, s.([]string)[i])
		}
	case []int64:
		for i := range s.([]int64) {
			query = append(query, "?")
			args = append(args, s.([]int64)[i])
		}
	default:
		return "", nil, errors.New("不支持的类型")
	}

	return "(" + strings.Join(query, ",") + ")", args, nil
}
