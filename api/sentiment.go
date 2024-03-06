package api

import (
	"SDDS/forms"
	"SDDS/global"
	"SDDS/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//var client = openai.NewClient("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")

// HandleValidatorError 处理参数校验错误
func HandleValidatorError(c *gin.Context, err error) {
	//var errs validator.ValidationErrors
	//ok := errors.As(err, &errs)
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "参数校验失败",
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

//func GetGPTRes(ctx *gin.Context) {
//	userQuestionForm := forms.UserQuestionForm{}
//	if err := ctx.ShouldBind(&userQuestionForm); err != nil {
//		HandleValidatorError(ctx, err)
//		return
//	}
//	resp, err := client.CreateChatCompletion(
//		context.Background(),
//		openai.ChatCompletionRequest{
//			Model: openai.GPT3Dot5Turbo,
//			Messages: []openai.ChatCompletionMessage{
//				{
//					Role:    openai.ChatMessageRoleUser,
//					Content: userQuestionForm.UserQuestion,
//				},
//			},
//		},
//	)
//
//	if err != nil {
//		fmt.Printf("ChatCompletion error: %v\n", err)
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"status":  http.StatusOK,
//		"message": "回复成功！",
//	})
//
//	fmt.Println(resp.Choices[0].Message.Content)
//}

// GetOverall 获取某个项目在某个周月标识情况下情感的总数，正向情感数量，负向情感数
func GetOverall(ctx *gin.Context) {
	sentimentBForm := forms.SentimentBForm{}
	if err := ctx.ShouldBind(&sentimentBForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	var sentiment []*models.Sentiment
	res := global.DB.Model(&models.Sentiment{}).Select("SUM(mblog_total) as mblog_total,SUM(pos_num) as pos_num,SUM(neg_num) as neg_num").
		Where("repo_name=? and time_flag=? and start_date=? and end_date=?",
			sentimentBForm.RepoName, sentimentBForm.TimeFlag, sentimentBForm.StartDate, sentimentBForm.EndDate).
		Find(&sentiment)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "未查询到数据",
		})
		return
	}

	//var response []interface{}
	//response = append(response, sentiment)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": sentiment,
	})
}

// GetTimeAxis 获取时间轴时间
func GetTimeAxis(ctx *gin.Context) {
	sentimentBForm := forms.SentimentBForm{}
	if err := ctx.ShouldBind(&sentimentBForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	var sentiments []*models.Sentiment
	// SELECT DISTINCT(start_date),end_date from developsentiment WHERE repo_name='flutter' and time_flag=0;
	res := global.DB.Model(&models.Sentiment{}).Select("start_date, end_date").
		Where("repo_name=? and time_flag=?",
			sentimentBForm.RepoName, sentimentBForm.TimeFlag).Distinct("start_date", "end_date").
		Order("start_date").Find(&sentiments)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "未查询到数据",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": sentiments,
	})
}

// GetDetail 获取某个项目在某个周月标识情况下情感的数量和经纬度数据
func GetDetail(ctx *gin.Context) {
	sentimentBForm := forms.SentimentBForm{}
	if err := ctx.ShouldBind(&sentimentBForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	var sentiments []*models.Sentiment
	res := global.DB.Model(&models.Sentiment{}).Where("repo_name=? AND time_flag=?",
		sentimentBForm.RepoName, sentimentBForm.TimeFlag).Preload("Area").Find(&sentiments)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "未查询到数据",
		})
		return
	}
	// 封装返回数据格式
	// 生成随机数
	//rand.Seed(time.Now().UnixNano())

	//fmt.Printf("随机浮点数: %f\n", randomFloat)
	var response []interface{}
	var features []gin.H
	for _, sentiment := range sentiments {
		// 生成0到100之间的随机浮点数
		//randomFloat := rand.Float64() * 100.0
		features = append(features, gin.H{
			"type": "Feature",
			"properties": gin.H{
				"adcode":      "United States",
				"mblog_total": sentiment.MblogTotal,
				"pos_num":     sentiment.PosNum,
				"neg_num":     sentiment.NegNum,
				"startdate":   sentiment.StartDate,
				"enddate":     sentiment.EndDate,
			},
			"geometry": gin.H{
				"type": "Point",
				"coordinates": []float64{
					sentiment.Area.Longitude, sentiment.Area.Latitude,
				},
			},
		})
	}

	response = append(response, gin.H{
		"type":     "FeatureCollection",
		"features": features,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": response,
	})
}

func GetDatas(ctx *gin.Context) {
	pageForm := forms.PageForm{}
	if err := ctx.ShouldBind(&pageForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	var sentiments []*models.Sentiment
	res := global.DB.Model(&models.Sentiment{}).Preload("Area").Limit(pageForm.Nums).
		Offset(pageForm.Nums * (pageForm.PageIndex - 1)).Find(&sentiments)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "未查询到数据",
		})
		return
	}
	// 封装返回数据格式
	// 生成随机数
	//rand.Seed(time.Now().UnixNano())

	//fmt.Printf("随机浮点数: %f\n", randomFloat)
	var response []interface{}
	var features []gin.H
	for _, sentiment := range sentiments {
		// 生成0到100之间的随机浮点数
		//randomFloat := rand.Float64() * 100.0
		features = append(features, gin.H{
			"type": "Feature",
			"properties": gin.H{
				"adcode":      "United States",
				"mblog_total": sentiment.MblogTotal,
				"pos_num":     sentiment.PosNum,
				"neg_num":     sentiment.NegNum,
				"startdate":   sentiment.StartDate,
				"enddate":     sentiment.EndDate,
			},
			"geometry": gin.H{
				"type": "Point",
				"coordinates": []float64{
					sentiment.Area.Longitude, sentiment.Area.Latitude,
				},
			},
		})
	}

	response = append(response, gin.H{
		"type":     "FeatureCollection",
		"features": features,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": response,
	})
}

//// GetSentimentByRepoName 获取指定仓库整体数据
//func GetSentimentByRepoName(ctx *gin.Context) {
//
//	repoNameForm := forms.RepoNameForm{}
//	if err := ctx.ShouldBind(&repoNameForm); err != nil {
//		HandleValidatorError(ctx, err)
//		return
//	}
//
//	result := forms.SentiByRepoForm{}
//	res := global.DB.Model(&models.Sentiment{}).Select(
//		"repo_name, sum(mblog_total) as total_comment, sum(pos_num) as total_pos, sum(neg_num) as total_neg").
//		Where("repo_name=?", repoNameForm.RepoName).Group("repo_name").Find(&result)
//	//global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
//	if res.RowsAffected == 0 {
//		ctx.JSON(http.StatusNotFound, gin.H{
//			"msg": "未查询到数据",
//		})
//		return
//	}
//
//	fmt.Println("通过仓库名称查询到情感数据")
//	ctx.JSON(http.StatusOK, gin.H{
//		"data": result,
//	})
//}
//
//// GetSentimentTime 获取指定仓库详细情感数据
//func GetSentimentTime(ctx *gin.Context) {
//	repoDetail := forms.RepoDetailForm{}
//	if err := ctx.ShouldBind(&repoDetail); err != nil {
//		HandleValidatorError(ctx, err)
//		return
//	}
//
//	var result []*models.Sentiment
//	res := global.DB.Model(&models.Sentiment{}).Where("repo_name=? AND time_flag=? AND start_date=?",
//		repoDetail.RepoName, repoDetail.TimeFlag, repoDetail.StartDate).Preload("Area").Find(&result)
//	//global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
//	if res.RowsAffected == 0 {
//		ctx.JSON(http.StatusNotFound, gin.H{
//			"msg": "未查询到数据",
//		})
//		return
//	}
//
//	fmt.Println("查询到仓库详细情感数据")
//	ctx.JSON(http.StatusOK, gin.H{
//		"data": result,
//	})
//}
//
//// GetSentimentDefault 按照默认参数（repo_name=flutter&time_flag=0&start_date=2020-01-01）
//// 获取指定仓库详细情感数据，并写入json文件
//func GetSentimentDefault(ctx *gin.Context) {
//	var result []*models.Sentiment
//	global.DB.Model(&models.Sentiment{}).Where("repo_name='flutter' AND time_flag='0' AND start_date='2020-01-01'").
//		Preload("Area").Find(&result)
//
//	// 调用 json.MarshalIndent 将结构体序列化为 JSON 格式的字节切片
//	jsonData, err := json.MarshalIndent(result, "", "  ")
//	if err != nil {
//		fmt.Println("Error marshalling JSON:", err)
//		return
//	}
//
//	// 将 JSON 数据写入文件
//	file, err := os.Create("./developer-sentiment/data/default_sentiment_data.json")
//	if err != nil {
//		fmt.Println("Error creating file:", err)
//		return
//	}
//	defer file.Close()
//
//	file.Write(jsonData)
//	fmt.Println("文件写入完成！！")
//
//}

//func GetDevelopTrend(ctx *gin.Context) {}
