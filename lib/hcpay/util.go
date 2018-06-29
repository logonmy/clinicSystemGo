package hcpay

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

//GetRSA1Sign 获取签名
/**
 * 获取签名
 * @param formData 待签名对象
 * @param privateKeyPath 私钥地址
 */
// func GetRSA1Sign(formData map[string]string, privateKeyPath string) string {
// 	sortData := sortArg(formData)
// 	var s []string
// 	for key := range sortData {
// 		s = append(s, key+"="+sortData[key])
// 	}
// 	queryString := strings.Join(s, "&")
// 	if contents, err := ioutil.ReadFile(privateKeyPath); err == nil {
// 		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
// 		prvKey := strings.Replace(string(contents), "\n", "", 1)
// 		keyByts, err := hex.DecodeString(prvKey)
// 		if err != nil {
// 			fmt.Println(err)
// 			return ""
// 		}
// 		privateKey, err := x509.ParsePKCS8PrivateKey(keyByts)
// 		if err != nil {
// 			fmt.Println("ParsePKCS8PrivateKey err", err)
// 			return ""
// 		}
// 		h := sha1.New()
// 		h.Write([]byte([]byte(queryString)))
// 		hash := h.Sum(nil)
// 		signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
// 		if err != nil {
// 			fmt.Printf("Error from signing: %s\n", err)
// 			return ""
// 		}
// 		out := hex.EncodeToString(signature)
// 		return out
// 	}
// 	return ""
// }

//GetSign 计算签名
func GetSign(mReq map[string]string, key string) (sign string) {

	mReq = filteParam(mReq)
	//STEP 1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k := range mReq {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		value := mReq[k]
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

//SignVeryfy 验证签名
/**
 * @param notifyData 支付宝通知数据（需 json 格式）
 * @param publicKeyPath 公钥地址
 * @param key 安全检验码，以数字和字母组成的32位字符（config 配置中的key）
 * @param signType 签名类型 （config 配置中的sign_type）
 */
func SignVeryfy(notifyData map[string]string, publicKeyPath string, key string, signType string) bool {
	sign := notifyData["sign"]
	paraFilter := filteParam(notifyData)
	paraSort := sortArg(paraFilter)
	prestr := createLinkstring(paraSort)
	isSgin := false
	signType = strings.ToUpper(signType)
	if signType == "MD5" {
		isSgin = md5Verify(prestr, sign, key)
	} else if signType == "RSA" {
		if contents, err := ioutil.ReadFile(publicKeyPath); err == nil {
			pubKey := strings.Replace(string(contents), "\n", "", 1)
			error := rsaVerySignWithSha1Base64(prestr, sign, pubKey)
			if error == nil {
				isSgin = true
			}
		}
	}
	return isSgin
}

//sortArg 对象排序
func sortArg(para map[string]string) map[string]string {
	var keys []string
	sortArg := make(map[string]string)
	for k := range para {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", para[k])
		sortArg[k] = para[k]
	}
	return sortArg
}

//createLinkstring 把所有元素，按照“参数=参数值”的模式用“&”字符拼接成字符串
func createLinkstring(para map[string]string) string {
	var ls string
	for k := range para {
		ls = ls + k + "=" + para[k] + "&"
	}
	ls = substr(ls, 0, 0)
	return ls
}

//filteParam 取出需要验证的属性
func filteParam(para map[string]string) map[string]string {
	paraFilter := make(map[string]string)
	for key := range para {
		if key == "sign" || key == "signType" || para[key] == "" {
			continue
		} else {
			paraFilter[key] = para[key]
		}
	}
	return paraFilter
}

//Md5Verify md5验证
func md5Verify(prestr string, sign string, key string) bool {
	prestr = prestr + key
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(prestr))
	cipherStr := md5Ctx.Sum(nil)
	mysgin := hex.EncodeToString(cipherStr)
	if mysgin == sign {
		return true
	}
	return false
}

//rsaVerySignWithSha1Base64 rsa1 验证
func rsaVerySignWithSha1Base64(originalData string, signData string, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}

//substr 截取字符串 start 起点下标 length 需要截取的长度
func substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	if length == 0 {
		length = rl
	}
	end = start + length - 1

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//GenerateNonceString 生成随机字符串
func GenerateNonceString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//CreateTradeNo 生成系统交易号
//06123xxxxx
//sum 最少10位,sum 表示全部单号位数
func CreateTradeNo(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(getDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	rand.Seed(time.Now().UnixNano())
	pow := math.Pow(10, float64(sum)) - 1

	result := strconv.Itoa(rand.Intn(int(pow)))
	count = len(result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}

//getDaysInYearByThisYear 年中的第几天
func getDaysInYearByThisYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1

	} else {
		total = total + d
	}
	return total
}
