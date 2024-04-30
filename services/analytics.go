package services

import (
	"math"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type AnalyticsService struct{}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{}
}

func (s *AnalyticsService) GetIncome(req string) float64 {
	var number float64

	now := time.Now()
	today := utils.GetTodayZeroTime()
	var startTime, endTime time.Time

	switch req {
	case "today":
		startTime = today
		endTime = now
	case "yesterday":
		startTime = utils.GetTimeBeforeDays(today, 1)
		endTime = today
	case "this month":
		startTime = utils.FirstDayOfMonth()
		endTime = now
	default:
		global.DB.Raw("select sum(?) from paylist where status=?", "total", 1).QueryRow(&number)
		return math.Round(number)
	}

	global.DB.Raw("select sum(?) from paylist where status=? and datetime between ? and ?", "total", 1, startTime, endTime).QueryRow(&number)
	return math.Round(number)
}

func (s *AnalyticsService) GetTotalUser() int64 {
	count, _ := global.DB.QueryTable(models.User{}).Count()
	return count
}

func (s *AnalyticsService) GetCheckinUser() int64 {
	count, _ := global.DB.QueryTable(models.User{}).Filter("last_check_in_time__gt", 0).Count()
	return count
}

func (s *AnalyticsService) GetTodayCheckinUser() int64 {
	count, _ := global.DB.QueryTable(models.NewUser()).Filter("last_check_in_time__gt", time.Now().UnixMilli()).Count()
	return count
}

func (s *AnalyticsService) GetTrafficUsage() string {
	user := models.NewUser()
	sum := user.SumByField("u") + user.SumByField("d")
	return utils.AutoBytes(int64(sum))
}

func (s *AnalyticsService) GetTodayTrafficUsage() string {
	var sum = models.NewUser().SumByField("transfer_today")
	return utils.AutoBytes(int64(sum))
}

func (s *AnalyticsService) GetRawTodayTrafficUsage() float64 {
	var sum = models.NewUser().SumByField("transfer_today")
	return sum
}

func (s *AnalyticsService) GetRawGbTodayTrafficUsage() float64 {
	var sum = models.NewUser().SumByField("transfer_today")
	return utils.FlowToGB(sum)
}

func (s *AnalyticsService) GetLastTrafficUsage() string {
	user := models.NewUser()
	sum := user.SumByField("u") + user.SumByField("d") - user.SumByField("transfer_today")
	return utils.AutoBytes(int64(sum))
}

func (s *AnalyticsService) GetRawLastTrafficUsage() int {
	user := models.NewUser()
	sum := user.SumByField("u") + user.SumByField("d") + user.SumByField("transfer_today")
	return int(sum)
}

func (s *AnalyticsService) GetRawGbLastTrafficUsage() float64 {
	user := models.NewUser()
	sum := user.SumByField("u") + user.SumByField("d") - user.SumByField("transfer_today")
	return utils.FlowToGB(sum)
}

func (s *AnalyticsService) GetUnusedTrafficUsage() string {
	user := models.NewUser()
	sum := user.SumByField("transfer_enable") - user.SumByField("u") - user.SumByField("d")
	return utils.AutoBytes(int64(sum))
}

func (s *AnalyticsService) GetRawUnusedTrafficUsage() int {
	user := models.NewUser()
	sum := user.SumByField("transfer_enable") - user.SumByField("u") - user.SumByField("d")
	return int(sum)
}

func (s *AnalyticsService) GetRawGbUnusedTrafficUsage() float64 {
	user := models.NewUser()
	sum := user.SumByField("transfer_enable") - user.SumByField("u") - user.SumByField("d")
	return utils.FlowToGB(sum)
}

func (s *AnalyticsService) GetTotalTraffic() string {
	var sum = models.NewUser().SumByField("transfer_enable")
	return utils.AutoBytes(int64(sum))
}

func (s *AnalyticsService) GetRawTotalTraffic() float64 {
	var sum float64
	global.DB.Raw("select sum(transfer_enable) from user").QueryRow(&sum)
	return sum
}

func (s *AnalyticsService) GetRawGbTotalTraffic() float64 {
	var sum float64
	global.DB.Raw("select sum(transfer_enable) from user").QueryRow(&sum)
	return utils.FlowToGB(sum)
}

func (s *AnalyticsService) GetTotalNode() int64 {
	count, _ := global.DB.QueryTable(models.Node{}).Filter("node_heartbeat__gte", 0).Count()
	return count
}

func (s *AnalyticsService) GetAliveNode() int64 {
	ninetyScds, _ := time.ParseDuration("-90s")
	count, _ := global.DB.QueryTable(models.Node{}).Filter("node_heartbeat__gt", time.Now().Add(ninetyScds)).Count()
	return count
}

func (s *AnalyticsService) GetInactiveUser() int64 {
	count, _ := global.DB.QueryTable(models.User{}).Filter("is_inactive", 1).Count()
	return count
}

func (s *AnalyticsService) GetActiveUser() int64 {
	count, _ := global.DB.QueryTable(models.User{}).Filter("is_inactive", 0).Count()
	return count
}

func (s *AnalyticsService) GetUserHourlyUsage(user_id int, date string) map[string]any {
	var hourlyUsage models.HourlyUsage
	err := global.DB.QueryTable(models.HourlyUsage{}).Filter("user_id", user_id).Filter("date", date).One(&hourlyUsage)

	m := make(map[string]any)
	if err != nil {

		for i := range 24 {
			m[strconv.Itoa(i)] = 0
		}
	} else {
		m[hourlyUsage.Usage] = true
	}
	return m
}

func (s *AnalyticsService) GetUserTodayHourlyUsage(user_id int) map[string]any {
	date := time.Now().Format(time.DateOnly)

	return s.GetUserHourlyUsage(user_id, date)
}
