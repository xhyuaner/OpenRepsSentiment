package router

import (
	"SDDS/api"
	"github.com/gin-gonic/gin"
)

func InitSentimentRouter(Router *gin.RouterGroup) {
	SentimentRouter := Router.Group("api")
	{
		// 获取某个项目在某个周月标识情况下全球情感的总数，正向情感数量，负向情感数
		SentimentRouter.GET("getDeveopSentimentB", api.GetOverall)
		// 获取时间轴时间
		SentimentRouter.GET("getDevelopSentimentTime", api.GetTimeAxis)
		// 获取某个项目在某个周月标识情况下情感的数量和经纬度数据
		SentimentRouter.GET("getDevelopSentimentDefault", api.GetDetail)
		SentimentRouter.GET("getDatas", api.GetDatas)
		//SentimentRouter.GET("getGptResponse", api.GetGPTRes)

		//SentimentRouter.GET("getDeveopSentimentB", api.GetSentimentByRepoName)
		//SentimentRouter.GET("getDevelopSentimentTime", api.GetSentimentTime)
		//SentimentRouter.GET("getDevelopSentimentDefault", api.GetSentimentDefault)
		//SentimentRouter.GET("getDevelopTrend", api.GetDevelopTrend)
	}

	//SentimentRouter := Router.Group("sentiment")
	//{
	//	SentimentRouter.GET("byRepoName", api.GetSentimentByRepoName)
	//	SentimentRouter.GET("sentimentTime", api.GetSentimentTime)
	//	SentimentRouter.GET("default", api.GetSentimentDefault)
	//	SentimentRouter.GET("developTrend", api.GetDevelopTrend)
	//}
}
