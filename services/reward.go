package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type RewardService struct{}

func NewRewardService() *RewardService {
	return &RewardService{}
}

func (s *RewardService) IssuePaybackReward(userId, refUserId int) {
	// TODO
}

func (s *RewardService) IssueRegReward(userId, refUserId int) {

	conf := models.NewConfig()

	invite_reg_money_reward := conf.Obtain("invite_reg_money_reward").ValueToInt()
	invite_reg_traffic_reward := conf.Obtain("invite_reg_traffic_reward").ValueToInt()

	user := models.NewUser().FindById(userId)

	var refUser models.User
	err := global.DB.QueryTable(models.User{}).
		Filter("id", refUserId).
		Filter("is_banned", 0).
		Filter("is_shadow_banned", 0).One(&refUser)

	if err != nil {
		return
	}

	// TODO change to transaction flow

	if user != nil && refUser.Id != 0 {

		if invite_reg_money_reward != 0 {
			money_before := user.Money
			user.Money += float64(invite_reg_money_reward)
			user.Update()

			// 添加余额记录
			models.NewUserMoneyLog().Add(
				userId,
				money_before,
				user.Money,
				float64(invite_reg_money_reward),
				fmt.Sprintf("被用户 #%d 邀请注册奖励", refUserId),
			)

		}

		if invite_reg_traffic_reward != 0 {
			refUser.TransferEnable += utils.ToGB(invite_reg_traffic_reward)
			refUser.Update()
		}

	}

}

func (s *RewardService) IssueCheckinReward(userId int) (int, error) {

	user := models.NewUser().FindById(userId)

	if user == nil {
		return 0, errors.New("user not found")
	}
	conf := models.NewConfig()

	checkin_min := conf.Obtain("checkin_min").ValueToInt()
	checkin_max := conf.Obtain("checkin_max").ValueToInt()

	traffic := 0

	if checkin_min == checkin_max {
		traffic = checkin_max
	} else {
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		traffic = seededRand.Intn(checkin_max-checkin_min) + checkin_min
	}

	if traffic != 0 {
		user.TransferEnable += utils.ToMB(traffic)
		user.LastCheckInTime = time.Now().UnixMilli()
		if _, err := user.Update(); err != nil {
			return 0, err
		}
		
	}

	return traffic, nil
}
