package user

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type ProductController struct {
	controllers.BaseController
}

func (c *ProductController) Index() {

	tabps := models.NewProduct().FetchList("tabp", 1)
	bandwidths := models.NewProduct().FetchList("bandwidth", 1)
	times := models.NewProduct().FetchList("time", 1)

	unmarshalContent := func(list []models.Product) {
		for i := range list {
			list[i].Parse()
		}
	}
	unmarshalContent(tabps)
	unmarshalContent(bandwidths)
	unmarshalContent(times)

	c.Data["tabps"] = tabps
	c.Data["bandwidths"] = bandwidths
	c.Data["times"] = times
	c.TplName = "views/tabler/user/product.tpl"
}
