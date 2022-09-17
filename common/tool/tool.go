package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

func IntInArray(item int, arr []int) (dx int) {
	for idx, v := range arr {
		if v == item {
			dx = idx
			return
		}
	}
	return -1
}

func StringInArray(item string, arr []string) (dx int) {
	for idx, v := range arr {
		if v == item {
			dx = idx
			return
		}
	}
	return -1
}

func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetUid(c *gin.Context) int64 {
	uid, ok := c.Get("uid")
	if !ok {
		return 0
	}
	if uidInt64, ok := uid.(int64); ok {
		return uidInt64
	}
	return 0
}

func DataToBytes(d interface{}) (databyte []byte) {
	da, _ := json.Marshal(d)
	databyte = da
	return
}

func DataToString(d interface{}) (datastr string) {
	da, _ := json.Marshal(d)
	datastr = string(da)
	return
}

func BindProto(c *gin.Context, message proto.Message) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	defer c.Request.Body.Close()
	err = proto.Unmarshal(body, message)
	if err != nil {
		return err
	}
	return nil
}

func MergeMap(map1, map2 map[string]string) (ret map[string]string) {
	for k1, v1 := range map1 {
		map2[k1] = v1
	}
	return map2
}

type OrderFabric struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func GetOrderFabric(url string) (ret []OrderFabric, err error) {
	_, _, errs := GetRequest().Get(url).EndStruct(&ret)
	if len(errs) != 0 {
		err = errs[0]
		return
	}
	return
}

func JoinIntArr(arr []int, sep string) string {
	strArr := []string{}
	for _, v := range arr {
		strArr = append(strArr, fmt.Sprintf("%d", v))
	}
	return strings.Join(strArr, sep)
}

func GetTimeShow(ctime time.Time) string {
	sec := time.Now().Unix() - ctime.Unix()
	if sec < 30 {
		return "刚刚"
	}
	if sec < 60 {
		return fmt.Sprintf("%d秒前", sec)
	}
	if sec >= 60 && sec < 3600 {
		return fmt.Sprintf("%d分钟前", sec/60)
	}
	if sec >= 3600 && sec < 86400 {
		return fmt.Sprintf("%d小时前", sec/3600)
	}
	if sec >= 86400 && sec < 604800 {
		return fmt.Sprintf("%d天前", sec/86400)
	}
	return ctime.Format("2006-01-02 15:04:05")
}
