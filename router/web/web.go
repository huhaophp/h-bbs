package web

import (
	"bbs/app/controllers/web"
	"bbs/app/controllers/web/captcha"
	"bbs/app/controllers/web/comment"
	"bbs/app/controllers/web/file"
	"bbs/app/controllers/web/node"
	"bbs/app/controllers/web/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// init 初始化web路由
func init() {
	webController := new(web.Controller)
	fileController := new(file.Controller)
	userController := new(user.Controller)
	nodeController := new(node.Controller)
	commentController := new(comment.Controller)
	captchaController := new(captcha.Controller)
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", webController.Home)
		group.GET("/captcha", captchaController.Get)
		group.GET("/posts/{postsId}", webController.PostDetail)
		group.GET("/user/login", userController.Login)
		group.POST("/user/login", userController.Login)
		group.POST("/user/logout", userController.Logout)
		group.GET("/user/register", userController.Register)
		group.POST("/user/register", userController.Register)
		group.GET("/node/{nodeId}", nodeController.Index)
		group.POST("/comment", commentController.Add)
		group.GET("/user/edit", userController.Edit)
		group.POST("/user/edit", userController.Edit)
		group.POST("/markdown/file", fileController.MdFileStore)
	})
}
