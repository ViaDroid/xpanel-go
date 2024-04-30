package models

import (
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/viadroid/xpanel-go/global"
)

type Config struct {
	FakeId   *int   `orm:"-" json:"id"`
	Id       int    `orm:"description(配置ID)" json:"-"`
	Item     string `orm:"description(配置项)"`
	Value    string `orm:"description(配置值)"`
	Class    string `orm:"description(配置类别)"`
	IsPublic int    `orm:"description(是否为公共参数)"`
	Type     string `orm:"description()"`
	Default  string `orm:"description(默认值)"`
	Mark     string `orm:"description(备注)"`
}

func NewConfig() Config {
	return Config{}
}

// Save
func (c Config) Save() {
	global.DB.Insert(&c)
}

// Update
func (c Config) Update() (int64, error) {
	return global.DB.Update(&c)
}

// update by item
func (c Config) UpdateByItem(item string, value string) (int64, error) {
	err := global.DB.QueryTable(c).Filter("item", item).One(&c)
	if err != nil {
		return 0, err
	}
	// if c.Type == "array" {
	// } else {
	// }
	c.Value = value

	id, err := global.DB.Update(&c)
	if err == nil {
		// Update global conf map
		c.updateSettingMap()
	}
	return id, err
}

func (c Config) Obtain(item string) (conf Config) {
	orm.NewOrm().QueryTable(Config{}).Filter("item", item).One(&conf)
	return conf
}

func (c Config) ObtainValue(item string) string {
	return c.Obtain(item).Value
}

func (c Config) IsExist(item string) bool {
	return global.DB.QueryTable(Config{}).Filter("item", item).Exist()
}

func (c Config) ValueToInt() int {
	i, err := strconv.Atoi(c.Value)
	if err != nil {
		return 0
	}
	return i
}

func (c Config) ValueToBool() bool {
	return c.ValueIsOne()
}

func (c Config) ValueIsOne() bool {
	return c.Value == "1"
}

func (c Config) GetPublicConfig() (m map[string]any) {
	var list []Config
	orm.NewOrm().QueryTable(Config{}).Filter("is_public", 1).All(&list)

	return listToMap(list)
}

func (c Config) FindByClass(class string) map[string]any {
	var list []Config
	global.DB.QueryTable(c).Filter("class", class).All(&list)

	return listToMap(list)
}

// fetch list by item
func (c Config) FetchListByItem(item string) []Config {
	var list []Config
	global.DB.QueryTable(c).Filter("item", item).All(&list)
	return list
}

func listToMap(list []Config) map[string]any {
	m := make(map[string]any)

	for _, v := range list {
		if v.Type == "bool" {
			m[v.Item] = v.Value == "1"
		} else if v.Type == "int" {
			i, _ := strconv.Atoi(v.Value)
			m[v.Item] = i
		} else {
			m[v.Item] = v.Value
		}
	}

	return m
}

func (c Config) updateSettingMap() {
	if c.Type == "bool" {
		global.SettingMap[c.Item] = c.Value == "1"
	} else if c.Type == "int" {
		i, _ := strconv.Atoi(c.Value)
		global.SettingMap[c.Item] = i
	} else {
		global.SettingMap[c.Item] = c.Value
	}
}
