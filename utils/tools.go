package utils

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/viadroid/xpanel-go/global"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!#"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetIpLocation(ip string) string {
	data := "GeoIP2 服务未配置"
	// TODO
	return data
}

/**
 * 根据流量值自动转换单位输出
 */
func AutoBytes(size int64) string {
	// precision := 2
	if size <= 0 {
		return "0B"
	}

	if size > 1208925819614629174 {
		return "∞"
	}

	//  base := math.Log(float64(size))
	base := math.Log(float64(size)) / math.Log(1024)
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	res := fmt.Sprintf("%d%s", int(math.Round(math.Pow(1024, base-math.Floor(base)))), units[int(math.Floor(float64(base)))])
	return res
}

/**
 * 根据含单位的流量值转换 B 输出
 */
func AutoBytesR(sizeStr string) int {

	if v, err := strconv.ParseFloat(sizeStr[0:len(sizeStr)-1], 64); err == nil {
		return int(v)
	}

	suffix := sizeStr[len(sizeStr)-2:]
	base, _ := strconv.ParseFloat(sizeStr[0:len(sizeStr)-2], 64)
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	if base > 999 && suffix == "EB" {
		return -1
	}

	return (int)(base * math.Pow(1024, float64(slices.Index(units, suffix))))
}

/**
 * 根据速率值自动转换单位输出
 */
func AutoMbps(size int) string {
	if size <= 0 {
		return "0Bps"
	}

	if size > 1000000000 {
		return "∞"
	}

	base := math.Log(float64(size)) / math.Log(1000)

	units := []string{"Mbps", "Gbps", "Tbps", "Pbps"}
	res := fmt.Sprintf("%d%s", int(math.Round(math.Pow(1000, base-math.Floor(base)))), units[int(math.Floor(float64(base)))])
	return res
}

/**
 * 虽然名字是toMB，但是实际上功能是from MB to B
 */
func ToMB(traffic int) int64 {
	return int64(traffic) * 1024 * 1024
}

/**
 * 虽然名字是toGB，但是实际上功能是from GB to B
 */
func ToGB(traffic int) int64 {
	return int64(traffic) * 1024 * 1024 * 1024
}

func FlowToMB(traffic float64) float64 {
	return math.Round(traffic / (1024 * 1024))
}

func FlowToGB(traffic float64) float64 {
	return math.Round(traffic / (1024 * 1024 * 1024))
}

func GetSsMethod(typ string) []string {
	var res []string
	switch typ {
	case "ss_obfs":
		res = []string{
			"simple_obfs_http",
			"simple_obfs_http_compatible",
			"simple_obfs_tls",
			"simple_obfs_tls_compatible",
		}
	default:
		res = []string{
			"aes-128-gcm",
			"aes-192-gcm",
			"aes-256-gcm",
			"chacha20-ietf-poly1305",
			"xchacha20-ietf-poly1305",
		}
	}
	return res
}

func IsEmailLegal(email string) map[string]any {
	var res = make(map[string]any)
	res["ret"] = 0

	if !IsEmail(email) {
		res["msg"] = "邮箱不规范"
		return res
	}

	mail_suffix := strings.Split(email, "@")[1]
	mail_filter_list, _ := global.ConfMap["mail_filter_list"].([]string)

	switch global.ConfMap["mail_filter"] {
	case 1:
		// 白名单
		if slices.Contains(mail_filter_list, mail_suffix) {
			res["ret"] = 1
		} else {
			res["msg"] = "邮箱域名 " + mail_suffix + " 无效，请更换邮件地址"
		}

		return res
	case 2:
		// 黑名单
		if !slices.Contains(mail_filter_list, mail_suffix) {
			res["ret"] = 1
		} else {
			res["msg"] = "邮箱域名 " + mail_suffix + " 无效，请更换邮件地址"
		}

		return res
	default:
		res["ret"] = 1
		return res
	}
}

func IsEmail(input string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(input)
}

func GenRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetSsPort() int {
	var maxP, minP orm.Params
	global.DB.Raw("select * from config where item =?", "max_port").RowsToMap(&maxP, "item", "value")
	global.DB.Raw("select * from config where item =?", "min_port").RowsToMap(&minP, "item", "value")

	// conf := models.NewConfig()
	// maxPort := conf.Obtain("max_port").ValueToInt()
	// minPort := conf.Obtain("min_port").ValueToInt()

	maxPort, _ := strconv.Atoi(maxP["max_port"].(string))
	minPort, _ := strconv.Atoi(minP["min_port"].(string))

	// user := models.NewUser()

	userCount, _ := global.DB.QueryTable("user").Count()

	if minPort >= 65535 || minPort <= 0 ||
		maxPort > 65535 || maxPort <= 0 ||
		minPort > maxPort || int(userCount) >= maxPort-minPort+1 {
		return 0
	}

	// var ports []int
	// det := user.PortArray()
	var det []int
	global.DB.Raw("select port from user").QueryRows(&det)

	arr := []int{}
	for i := minPort; i < maxPort; i++ {
		arr = append(arr, i)
	}

	diffPorts := DiffArray(arr, det)
	randIndex := seededRand.Intn(len(diffPorts))

	return diffPorts[randIndex]
}

// random_group a string saparated with comma
func GetRandomGroup(random_group string) int {
	groups := strings.Split(random_group, ",")
	n := seededRand.Intn(len(groups))

	group, err := strconv.Atoi(groups[n])
	if err != nil {
		return 0
	}
	return group
}

// DiffArray 求两个切片的差集
func DiffArray(a []int, b []int) []int {
	var diffArray []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}
