package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/orangebottle/blockchain-grant/application/routers/api/v1"
	"net/http"
	"strings"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	//swagger文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//创建路由组
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/createWyuUser", v1.CreateWyuUser)
		apiV1.POST("/queryWyuUser", v1.QueryWyuUser)
		apiV1.POST("/createProofMaterial", v1.CreateProofMaterial)
		apiV1.POST("/queryProofMaterial", v1.QueryProofMaterial)
		apiV1.POST("/queryProofMaterialOnly", v1.QueryProofMaterialOnly)
		apiV1.POST("/queryProofCertify", v1.QueryProofCertify)
		apiV1.POST("/createPhotoMaterial", v1.CreatePhotoMaterial)
		apiV1.POST("/queryPhotoMaterial", v1.QueryPhotoMaterial)
		apiV1.POST("/updateProofMaterial", v1.UpdateProofMaterial)
		apiV1.POST("/updatePower", v1.UpdatePower)
		apiV1.POST("/setPower", v1.SetPower)
		apiV1.POST("/queryPower", v1.QueryPower)
		apiV1.POST("/createQueryStipendRanking", v1.CreateQueryStipendRanking)
		apiV1.POST("/queryAwardList", v1.QueryAwardList)
		apiV1.POST("createVote", v1.CreateVote)
		apiV1.POST("/queryVote", v1.QueryVote)
		apiV1.POST("/queryVoteOnly", v1.QueryVoteOnly)
	}
	// 静态文件路由
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}

//gin通过跨域中间件cors
// Cors 允许跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 这是允许访问所有域
			c.Header("Access-Control-Allow-Origin", "*")
			// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			// 缓存请求信息 单位为秒
			c.Header("Access-Control-Max-Age", "172800")
			// 跨域请求是否需要带cookie信息 默认设置为true
			c.Header("Access-Control-Allow-Credentials", "false")
			// 设置返回格式是json
			c.Set("content-type", "application/json")
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
