package forms

//PageForm
// RepoNameForm 获取指定仓库整体数据---请求参数结构体
type PageForm struct {
	PageIndex int `form:"pageindex" json:"pageindex"`
	Nums      int `form:"nums" json:"nums"`
}

// SentimentBForm 获取某个项目在某个周月标识情况下情感的总数，正向情感数，负向情感数---请求参数结构体
// /api/getDeveopSentimentB?reponame=tensorflow&startdate=2020-01-02&enddate=2020-01-07&timeflag=0
type SentimentBForm struct {
	RepoName  string `form:"reponame" json:"reponame" binding:"required"`
	TimeFlag  string `form:"timeflag" json:"timeflag" binding:"required"`
	StartDate string `form:"startdate" json:"startdate"`
	EndDate   string `form:"enddate" json:"enddate"`
}

// RepoNameForm 获取指定仓库整体数据---请求参数结构体
type RepoNameForm struct {
	RepoName string `form:"repo_name" json:"repo_name" binding:"required"`
}

// RepoDetailForm 获取指定仓库详细情感数据---请求参数结构体
type RepoDetailForm struct {
	RepoName  string `form:"repo_name" json:"repo_name" binding:"required"`
	TimeFlag  string `form:"time_flag" json:"time_flag" binding:"required"`
	StartDate string `form:"start_date" json:"start_date" binding:"required"`
}

// SentiByRepoForm 返回数据结构体
type SentiByRepoForm struct {
	RepoName     string `form:"repo_name" json:"repo_name"`
	TotalComment string `form:"total_comment" json:"total_comment"`
	TotalPos     string `form:"total_pos" json:"total_pos"`
	TotalNeg     string `form:"total_neg" json:"total_neg"`
}
