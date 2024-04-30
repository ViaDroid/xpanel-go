package models

import "github.com/viadroid/xpanel-go/global"

type GiftCard struct {
	Id         int     `orm:"description(礼品卡ID)"`
	Card       string  `orm:"description(卡号)"`
	Balance    float64 `orm:"description(余额)"`
	CreateTime int64   `orm:"description(创建时间)"`
	Status     int     `orm:"description(使用状态)"`
	UseTime    int64   `orm:"description(使用时间)"`
	UseUser    int     `orm:"description(使用用户)"`

	Op            string `orm:"-"`
	StatusStr     string `orm:"-"`
	CreateTimeStr string `orm:"-"`
	UseTimeStr    string `orm:"-"`
}

func NewGiftCard() *GiftCard {
	return &GiftCard{}
}

// find by card
func (item *GiftCard) FindByCard(card string) (*GiftCard, error) {
	item.Card = card
	err := global.DB.QueryTable(item).Filter("card", card).One(item)
	return item, err
}

// Delete
func (item *GiftCard) Delete(id int) (int64, error) {
	item.Id = id
	return global.DB.Delete(item)
}

// Fetch all order by id
func (item *GiftCard) FetchAll() []GiftCard {
	var items []GiftCard
	global.DB.QueryTable(item).OrderBy("-id").All(&items)
	return items
}

func (item *GiftCard) ParseStatus() string {
	if item.Status == 0 {
		return "未使用"
	} else if item.Status == 1 {
		return "已使用"
	} else {
		return "已过期"
	}
}
