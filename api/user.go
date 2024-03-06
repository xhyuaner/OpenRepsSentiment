package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

// 去除校验返回错误提示信息中Key的结构体名称
// 例如：将 `"LoginForm.password":"password为必填字段",` 转换为 `"password":"password为必填字段",`
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, value := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = value
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

//func HandleValidatorError(c *gin.Context, err error) {
//	//var errs validator.ValidationErrors
//	//ok := errors.As(err, &errs)
//	errs, ok := err.(validator.ValidationErrors)
//	if !ok {
//		c.JSON(http.StatusOK, gin.H{
//			"msg": err.Error(),
//		})
//	}
//	c.JSON(http.StatusBadRequest, gin.H{
//		"error": removeTopStruct(errs.Translate(global.Trans)),
//	})
//	return
//}

//func GetUserList(ctx *gin.Context) {
//	claims, _ := ctx.Get("claims")
//	currentUser := claims.(*models.CustomClaims)
//	zap.S().Infof("访问用户: %d", currentUser.ID)
//
//	pn := ctx.DefaultQuery("pn", "0")
//	pnInt, _ := strconv.Atoi(pn)
//	pSize := ctx.DefaultQuery("psize", "10")
//	pSizeInt, _ := strconv.Atoi(pSize)
//	//使用UserSrvClient调用grpc用户服务接口
//	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
//		Pn:    uint32(pnInt),
//		PSize: uint32(pSizeInt),
//	})
//	if err != nil {
//		zap.S().Errorw("[GetUserList] 查询 【用户列表】失败")
//		HandleGrpcErrorToHttp(err, ctx)
//		return
//	}
//
//	reMap := gin.H{
//		"total": rsp.Total,
//	}
//	result := make([]interface{}, 0)
//	for _, value := range rsp.Data {
//		user := reponse.UserResponse{
//			Id:       value.Id,
//			NickName: value.NickName,
//			//Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
//			Birthday: reponse.JsonTime(time.Unix(int64(value.BirthDay), 0)),
//			Gender:   value.Gender,
//			Mobile:   value.Mobile,
//		}
//		result = append(result, user)
//	}
//
//	reMap["data"] = result
//	ctx.JSON(http.StatusOK, reMap)
//}
//
//func PassWordLogin(c *gin.Context) {
//	//表单验证
//	passwordLoginForm := forms.PassWordLoginForm{}
//	if err := c.ShouldBind(&passwordLoginForm); err != nil {
//		HandleValidatorError(c, err)
//		return
//	}
//
//	//图形验证码校验
//	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"captcha": "验证码错误",
//		})
//		return
//	}
//
//	//登录的逻辑
//	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
//		Mobile: passwordLoginForm.Mobile,
//	}); err != nil {
//		if e, ok := status.FromError(err); ok {
//			switch e.Code() {
//			case codes.NotFound:
//				c.JSON(http.StatusBadRequest, gin.H{
//					"mobile": "用户不存在",
//				})
//			default:
//				c.JSON(http.StatusInternalServerError, gin.H{
//					"mobile": "登录失败，内部错误",
//				})
//			}
//			return
//		}
//	} else {
//		//只是查询到用户了而已，并没有检查密码
//		if passRsp, pasErr := global.UserSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
//			Password:          passwordLoginForm.PassWord,
//			EncryptedPassword: rsp.PassWord,
//		}); pasErr != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"password": "登录失败，内部错误",
//			})
//		} else {
//			if passRsp.Success {
//				//生成token
//				j := middlewares.NewJWT()
//				claims := models.CustomClaims{
//					ID:          uint(rsp.Id),
//					NickName:    rsp.NickName,
//					AuthorityId: uint(rsp.Role),
//					StandardClaims: jwt.StandardClaims{
//						NotBefore: time.Now().Unix(),               //签名的生效时间
//						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
//						Issuer:    "xhyuaner",                      //签发者
//					},
//				}
//				token, err := j.CreateToken(claims)
//				if err != nil {
//					c.JSON(http.StatusInternalServerError, gin.H{
//						"msg": "生成token失败",
//					})
//					return
//				}
//
//				c.JSON(http.StatusOK, gin.H{
//					"id":         rsp.Id,
//					"nick_name":  rsp.NickName,
//					"token":      token,
//					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
//				})
//			} else {
//				c.JSON(http.StatusBadRequest, gin.H{
//					"msg": "登录失败，密码错误",
//				})
//			}
//		}
//	}
//}
//
//func Register(c *gin.Context) {
//	//用户注册
//	registerForm := forms.RegisterForm{}
//	if err := c.ShouldBind(&registerForm); err != nil {
//		HandleValidatorError(c, err)
//		return
//	}
//
//	//验证码
//	rdb := redis.NewClient(&redis.Options{
//		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
//	})
//	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
//	//错误判断
//	if err == redis.Nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code": "验证码错误，手机号不存在",
//		})
//		return
//	} else {
//		if value != registerForm.Code {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"code": "验证码错误",
//			})
//			return
//		}
//	}
//
//	// TODO:优化代码
//	//错误判断
//	//if err != nil {
//	//	if err == redis.Nil {
//	//		c.JSON(http.StatusBadRequest, gin.H{
//	//			"code": "验证码错误，手机号不存在",
//	//		})
//	//		return
//	//	}
//	//	zap.S().Errorf("[Register] 查询 redis 失败: %s", err.Error())
//	//	c.JSON(http.StatusInternalServerError, gin.H{
//	//        "code": "服务不可用",
//	//    })
//	//	return
//	//}
//	//if value != registerForm.Code {
//	//	c.JSON(http.StatusBadRequest, gin.H{
//	//		"code": "验证码错误",
//	//	})
//	//	return
//	//}
//
//	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
//		NickName: registerForm.Mobile,
//		PassWord: registerForm.PassWord,
//		Mobile:   registerForm.Mobile,
//	})
//
//	if err != nil {
//		zap.S().Errorf("[Register] 查询 【新建用户失败】失败: %s", err.Error())
//		HandleGrpcErrorToHttp(err, c)
//		return
//	}
//
//	j := middlewares.NewJWT()
//	claims := models.CustomClaims{
//		ID:          uint(user.Id),
//		NickName:    user.NickName,
//		AuthorityId: uint(user.Role),
//		StandardClaims: jwt.StandardClaims{
//			NotBefore: time.Now().Unix(),               //签名的生效时间
//			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
//			Issuer:    "xhyuaner",
//		},
//	}
//	token, err := j.CreateToken(claims)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "生成token失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"id":         user.Id,
//		"nick_name":  user.NickName,
//		"token":      token,
//		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
//	})
//}
//
//func GetUserDetail(ctx *gin.Context) {
//	claims, _ := ctx.Get("claims")
//	currentUser := claims.(*models.CustomClaims)
//	zap.S().Infof("访问用户: %d", currentUser.ID)
//
//	rsp, err := global.UserSrvClient.GetUserById(context.Background(), &proto.IdRequest{
//		Id: int32(currentUser.ID),
//	})
//	if err != nil {
//		HandleGrpcErrorToHttp(err, ctx)
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"name":     rsp.NickName,
//		"birthday": time.Unix(int64(rsp.BirthDay), 0).Format("2006-01-02"),
//		"gender":   rsp.Gender,
//		"mobile":   rsp.Mobile,
//	})
//}
//
//func UpdateUser(ctx *gin.Context) {
//	updateUserForm := forms.UpdateUserForm{}
//	if err := ctx.ShouldBind(&updateUserForm); err != nil {
//		HandleValidatorError(ctx, err)
//		return
//	}
//
//	claims, _ := ctx.Get("claims")
//	currentUser := claims.(*models.CustomClaims)
//	zap.S().Infof("访问用户: %d", currentUser.ID)
//
//	//将前端传递过来的日期格式转换成int
//	loc, _ := time.LoadLocation("Local") //local的L必须大写
//	birthDay, _ := time.ParseInLocation("2006-01-02", updateUserForm.Birthday, loc)
//	_, err := global.UserSrvClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
//		Id:       int32(currentUser.ID),
//		NickName: updateUserForm.Name,
//		Gender:   updateUserForm.Gender,
//		BirthDay: uint64(birthDay.Unix()),
//	})
//	if err != nil {
//		HandleGrpcErrorToHttp(err, ctx)
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{})
//}
