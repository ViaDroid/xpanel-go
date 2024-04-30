package admin

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type CouponController struct {
	controllers.BaseController
}

var coupon_menu_details = map[string]any{
	"field": []any{
		NewField("Op", "操作"),
		NewField("Id", "ID"),
		NewField("Code", "优惠码"),
		NewField("Type", "类型"),
		NewField("Value", "额度"),
		NewField("ProductId", "可用商品ID"),
		NewField("UseTime", "使用次数（每用户）"),
		NewField("TotalUseTime", "使用次数（累计）"),
		NewField("NewUser", "仅限新用户使用"),
		NewField("Disabled", "已禁用"),
		NewField("UseCount", "总使用次数"),
		NewField("CreateTimeStr", "创建时间"),
		NewField("ExpireTimeStr", "过期时间"),
	},
	"create_dialog": []any{
		map[string]any{
			"Id":          "code",
			"Info":        "优惠码",
			"Type":        "input",
			"Placeholder": "",
		},
		map[string]any{
			"Id":   "type",
			"Info": "优惠码类型",
			"Type": "select",
			"select": map[string]any{
				"percentage": "百分比",
				"fixed":      "固定金额",
			},
		},
		map[string]any{
			"Id":          "value",
			"Info":        "优惠码额度",
			"Type":        "input",
			"Placeholder": "",
		},
		map[string]any{
			"Id":          "product_id",
			"Info":        "可用商品ID（多个ID以英文半角逗号分隔）",
			"Type":        "input",
			"Placeholder": "",
		},
		map[string]any{
			"Id":          "use_time",
			"Info":        "每个用户可使用次数限制（小于0为不限）",
			"Type":        "input",
			"Placeholder": "",
		},
		map[string]any{
			"Id":          "total_use_time",
			"Info":        "总使用次数限制（小于0为不限）",
			"Type":        "input",
			"Placeholder": "",
		},
		map[string]any{
			"Id":   "new_user",
			"Info": "仅限新用户使用",
			"Type": "select",
			"select": map[string]any{
				"1": "启用",
				"0": "禁用",
			},
		},
		map[string]any{
			"Id":   "generate_method",
			"Info": "生成方式",
			"Type": "select",
			"select": map[string]any{
				"char":        "指定字符",
				"random":      "随机字符（无视优惠码参数）",
				"char_random": "指定字符+随机字符",
			},
		},
	},
}

func (c *CouponController) Index() {
	c.Data["details"] = coupon_menu_details
	c.TplName = "views/tabler/admin/coupon.tpl"
}

func (c *CouponController) Add() {
	code := c.GetString("code")
	typ := c.GetString("type")
	value, _ := c.GetInt("value", 0)
	productID := c.GetString("product_id")
	useTime, _ := c.GetInt("use_time", 0)
	totalTime, _ := c.GetInt("total_use_time", 0)
	newUser, _ := c.GetInt("new_user", 0)
	generateMethod := c.GetString("generate_method")
	expireTime := c.GetString("expire_time")

	var parseExpireTime int64

	methods := []string{"char", "random", "char_random"}
	if code == "" && slices.Contains(methods, generateMethod) {
		c.Error(&c.Controller, "优惠码不能为空")
		return
	}
	// 无效的优惠码参数
	if typ == "" || value <= 0 {
		c.Error(&c.Controller, "无效的优惠码参数")
		return
	}
	// handle expiretime
	if expireTime != "" {
		t, err := time.ParseInLocation("2006-01-02 15:04", expireTime, time.Local)
		if err != nil {
			c.Error(&c.Controller, "无效的过期时间")
			return
		}
		parseExpireTime = t.UnixMilli()
		if parseExpireTime < time.Now().UnixMilli() {
			c.Error(&c.Controller, "过期时间不能早于当前时间")
			return
		}
	}

	coupon := models.NewUserCoupon()
	if generateMethod == methods[0] && coupon.CountByCode(code) > 0 {
		c.Error(&c.Controller, "优惠码已存在")
		return
	}

	if generateMethod == methods[1] {
		code = utils.GenRandomString(8)
		if coupon.CountByCode(code) > 0 {
			c.Error(&c.Controller, "出现了一些问题，请稍后重试")
		}
	}

	if generateMethod == methods[2] {
		code += utils.GenRandomString(8)
		if coupon.CountByCode(code) > 0 {
			c.Error(&c.Controller, "出现了一些问题，请稍后重试")
		}
	}

	contentMap := map[string]any{
		"type":  typ,
		"value": value,
	}

	limitMap := map[string]any{
		"product_id":     productID,
		"use_time":       useTime,
		"total_use_time": totalTime,
		"new_user":       newUser,
		"disabled":       0,
	}

	contentJsonBs, _ := json.Marshal(contentMap)
	limitJsonBs, _ := json.Marshal(limitMap)

	coupon.Code = code
	coupon.Content = string(contentJsonBs)
	coupon.Limit = string(limitJsonBs)
	coupon.CreateTime = time.Now().UnixMilli()

	coupon.ExpireTime = parseExpireTime

	if _, err := coupon.Save(); err != nil {
		c.Error(&c.Controller, "优惠码添加失败")
		return
	}
	c.Success(&c.Controller, "优惠码添加成功")
}

func (c *CouponController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "无效的优惠码ID")
		return
	}

	if _, err := models.NewUserCoupon().Delete(id); err != nil {
		c.Error(&c.Controller, "优惠码删除失败")
		return
	}
	c.Success(&c.Controller, "优惠码删除成功")

}

func (c *CouponController) Disable() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(&c.Controller, "无效的优惠码ID")
		return
	}

	coupon, err := models.NewUserCoupon().FindById(id)
	if err != nil {
		c.Error(&c.Controller, "优惠码不存在")
		return
	}

	var limitMap map[string]any
	json.Unmarshal([]byte(coupon.Limit), &limitMap)

	limitMap["disabled"] = 1
	limitJsonBs, _ := json.Marshal(limitMap)
	coupon.Limit = string(limitJsonBs)

	if _, err := coupon.Update(); err != nil {
		c.Error(&c.Controller, "优惠码禁用失败")
		return
	}
	c.Success(&c.Controller, "优惠码禁用成功")
}

func (c *CouponController) Ajax() {
	coupons := models.NewUserCoupon().FetchAll()

	for i, v := range coupons {
		var content, limit map[string]any
		json.Unmarshal([]byte(v.Content), &content)
		json.Unmarshal([]byte(v.Limit), &limit)

		coupons[i].Parse()

		op := fmt.Sprintf(`
		<button type="button" class="btn btn-red" id="delete-coupon-%d"
                onclick="deleteCoupon(%d)">删除</button>
		`, v.Id, v.Id)

		if limit["disabled"].(float64) != 1 {
			op += fmt.Sprintf(`
			<button type="button" class="btn btn-primary" id="disable-coupon-%d"
                onclick="disableCoupon(%d)">禁用</button>
			`, v.Id, v.Id)
		}
		coupons[i].Op = op

	}

	c.JSONResp(map[string]any{
		"coupons": coupons,
	})
}
