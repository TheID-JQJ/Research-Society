package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hkc.ink/rss/database"
	"hkc.ink/rss/dto"
	"hkc.ink/rss/response"
	"hkc.ink/rss/service"
	"hkc.ink/rss/utils"
)

var userService service.UserService
var jwtUtil utils.JwtUtil
var valUtil utils.ValidatorUtil

// @Summary test
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} hello hkc
// @Router /test/hello [get]
func TestHello(ctx *gin.Context) {
	value, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"value": value}, "hello hkc")
}

// @Summary test register
// @Description test register
// @Tags test
// @Accept json
// @Produce json
// @Param	registerDto body string true "register"
// @Success 200 {string} success
// @Failure	400	{} fail
// @Failure	500	{} fail
// @Router /test/register [post]
func TestRegister(ctx *gin.Context) {
	//接收数据
	rm := dto.RegisterDto{}
	if err := ctx.ShouldBindJSON(&rm); err != nil {
		log.Println("数据绑定失败：" + err.Error())
		if s := valUtil.GetError(err, rm); s != "" {
			response.Fail(ctx, s)
			return
		}
		response.Response(ctx, http.StatusOK, 500, gin.H{}, "注册失败")
		return
	}

	db := database.GetDB()

	//判断是否存在账号
	if userService.ExistEmail(db, rm.Email) {
		response.Fail(ctx, "邮箱已被注册")
		return
	}

	//写入数据库并返回结果
	if id, err := userService.Register(db, &rm); err != nil {
		log.Println("sql执行错误：" + err.Error())
		response.Response(ctx, http.StatusOK, 500, gin.H{}, "注册失败")
	} else {
		authorization := jwtUtil.ReleaseToken(id)
		response.Success(ctx, gin.H{"Authorization": authorization}, "注册成功")
	}
}

// @Summary test insert
// @Description test insert
// @Tags test
// @Accept json
// @Produce json
// @Param	userDto body string true "user"
// @Success 200 {} success
// @Failure	400	{} fail
// @Failure	500	{} fail
// @Router /test/insert [post]
func TestInsert(ctx *gin.Context) {
	//接收数据
	ud := dto.UserDto{}
	if err := ctx.ShouldBindJSON(&ud); err != nil {
		log.Println("数据绑定失败：" + err.Error())
		if s := valUtil.GetError(err, ud); s != "" {
			response.Fail(ctx, s)
			return
		}
		response.Response(ctx, http.StatusOK, 500, gin.H{}, "创建失败")
		return
	}

	db := database.GetDB()

	//判断是否存在账号
	if userService.ExistEmail(db, ud.Email) {
		response.Fail(ctx, "邮箱已被使用")
		return
	}

	//写入数据库并返回结果
	if _, err := userService.Insert(db, ud.ToUserModel()); err != nil {
		log.Println("sql执行错误：" + err.Error())
		response.Response(ctx, http.StatusOK, 500, gin.H{}, "创建失败")
		return
	}

	response.Success(ctx, gin.H{}, "创建成功")

}

// @Summary test delete
// @Description test delete
// @Tags test
// @Accept json
// @Produce json
// @Param	id path string true "id"
// @Success 200 {} success
// @Failure	400	{} fail
// @Failure	500	{} fail
// @Router /test/delete/{id} [post]
func TestDelete(ctx *gin.Context) {
	//接收数据
	id := ctx.Param("id")
	db := database.GetDB()

	// 执行删除并返回结果
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id转换失败：" + err.Error())
		response.Fail(ctx, "用户id输入错误")
		return
	}

	if err := userService.DeleteByID(db, uint(intId)); err != nil {
		log.Println("sql执行错误：" + err.Error())
		response.Fail(ctx, "删除失败")
		return
	}
	response.Success(ctx, gin.H{}, "删除成功")
}

// @Summary test update
// @Description test update
// @Tags test
// @Accept json
// @Produce json
// @Param	userDto body string true "user"
// @Success 200 {} success
// @Failure	400	{} fail
// @Failure	500	{} fail
// @Router /test/update [post]
func TestUpdate(ctx *gin.Context) {
	//接收数据
	ud := dto.UserDto{}
	if err := ctx.ShouldBindJSON(&ud); err != nil {
		log.Println("数据绑定失败：" + err.Error())
		if s := valUtil.GetError(err, ud); s != "" {
			response.Fail(ctx, s)
			return
		}
		response.Response(ctx, http.StatusOK, 500, gin.H{}, "修改失败")
		return
	}

	db := database.GetDB()

	// 执行修改并返回结果
	if err := userService.Update(db, ud.ToUserModel()); err != nil {
		log.Println("sql执行错误：" + err.Error())
		response.Fail(ctx, "修改失败")
		return
	}
	response.Success(ctx, gin.H{}, "修改成功")
}

// @Summary test get
// @Description test get
// @Tags test
// @Accept json
// @Produce json
// @Param	id path string true "id"
// @Success 200 {User} success
// @Failure	400	{} fail
// @Failure	500	{} fail
// @Router /test/get/{id} [get]
func TestGet(ctx *gin.Context) {
	//接收数据
	id := ctx.Param("id")
	db := database.GetDB()

	// 执行删除并返回结果
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("id转换失败：" + err.Error())
		response.Fail(ctx, "用户id输入错误")
		return
	}

	// 执行查询并返回结果
	user, err := userService.GetByID(db, uint(intId))
	if err != nil {
		log.Println("sql执行错误：" + err.Error())
		response.Fail(ctx, "用户不存在")
		return
	}
	response.Success(ctx, gin.H{"user": user}, "查询成功")
}

// @Summary test getAll
// @Description test getAll
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {list} success
// @Failure	400	{} fail
// @Failure	500	{} fail
// @Router /test/get/all [get]
func TestGetAll(ctx *gin.Context) {
	db := database.GetDB()

	// 执行查询并返回结果
	users, err := userService.GetAll(db)
	if err != nil {
		log.Println("sql执行错误：" + err.Error())
		response.Fail(ctx, "查询失败")
		return
	}
	response.Success(ctx, gin.H{"users": users}, "查询成功")
}

var fileService service.FileService

func TestUpload(ctx *gin.Context) {
	fileName, err := fileService.Upload(ctx)
	if err != nil {
		response.Fail(ctx, "文件上传失败")
		return
	}
	response.Success(ctx, gin.H{"fileName": fileName}, "文件上传成功")
}

func TestDownload(ctx *gin.Context) {
	s := ctx.Param("fileName")
	path := fileService.GetFilePath(s)
	log.Println(path)
	if path == "" {
		response.Fail(ctx, "文件读取失败")
		return
	}
	ctx.File(path)
}
