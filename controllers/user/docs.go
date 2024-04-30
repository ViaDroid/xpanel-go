package user

import (
	"github.com/viadroid/xpanel-go/controllers"
	"github.com/viadroid/xpanel-go/models"
)

type DocsController struct {
	controllers.BaseController
}

func (c *DocsController) Index() {
	conf := models.NewConfig()
	if !conf.Obtain("display_docs").ValueToBool() ||
		(conf.Obtain("display_docs_only_for_paid_user").ValueToBool() && c.User.Class == 0) {

		c.Redirect("/user", 302)
		return
	}

	docs := models.NewDocs().FetchDocs()

	c.Data["docs"] = docs
	c.TplName = "views/tabler/user/docs/index.tpl"
}

func (c *DocsController) Detail() {
	conf := models.NewConfig()
	if !conf.Obtain("display_docs").ValueToBool() ||
		(conf.Obtain("display_docs_only_for_paid_user").ValueToBool() && c.User.Class == 0) {

		c.Redirect("/user", 302)
		return
	}

	docId, err := c.GetInt("id")
	if err != nil {
		c.Redirect("/404", 302)
		return
	}

	doc := models.NewDocs().FindDocsById(docId)

	c.Data["doc"] = doc
	c.TplName = "views/tabler/user/docs/view.tpl"
}
