package api

import (
	"family-joint-school/common/errors/errorcode"
	ginx "family-joint-school/common/http/gin"
	"family-joint-school/logic"
	"family-joint-school/model"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/chanxuehong/log"
	"github.com/demdxx/gocast"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Handler 外部调用接口
type Handler interface {
	Login(ctx ginx.Context) (interface{}, error)
	NoticeList(ctx ginx.Context) (interface{}, error)
	NoticePublish(ctx ginx.Context) (interface{}, error)
	ClassList(ctx ginx.Context) (interface{}, error)
	HomeWorkList(ctx ginx.Context) (interface{}, error)
	HomeWorkPublish(ctx ginx.Context) (interface{}, error)
	HomeWorkSubmit(ctx ginx.Context) (interface{}, error)
	UserHomeWorkDetail(ctx ginx.Context) (interface{}, error)
	UserHomeWorkList(ctx ginx.Context) (interface{}, error)
	Upload(c *gin.Context)
}

type handler struct {
	logic  logic.Interface
	bucket *oss.Bucket
}

func (h *handler) NoticeList(ctx ginx.Context) (interface{}, error) {
	params := &model.NoticeListRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.NoticeList(ctx, params)
}

func (h *handler) NoticePublish(ctx ginx.Context) (interface{}, error) {
	params := &model.NoticePublishRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.NoticePublish(ctx, params)
}

func (h *handler) ClassList(ctx ginx.Context) (interface{}, error) {
	params := &model.ClassListRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.ClassList(ctx, params)
}

func (h *handler) HomeWorkList(ctx ginx.Context) (interface{}, error) {
	params := &model.HomeWorkListRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.HomeWorkList(ctx, params)
}

func (h *handler) HomeWorkPublish(ctx ginx.Context) (interface{}, error) {
	params := &model.HomeWorkPublishRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.HomeWorkPublish(ctx, params)
}

func (h *handler) HomeWorkSubmit(ctx ginx.Context) (interface{}, error) {
	params := &model.HomeWorkSubmitRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	params.UserID = gocast.ToUint64(ctx.GetString("uid"))
	return h.logic.HomeWorkSubmit(ctx, params)
}

func (h *handler) UserHomeWorkDetail(ctx ginx.Context) (interface{}, error) {
	params := &model.UserHomeWorkDetailRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	params.UserID = gocast.ToUint64(ctx.GetString("uid"))
	return h.logic.UserHomeWorkDetail(ctx, params)
}

func (h *handler) UserHomeWorkList(ctx ginx.Context) (interface{}, error) {
	params := &model.UserHomeWorkRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.UserHomeWorkList(ctx, params)
}

func (h *handler) Login(ctx ginx.Context) (interface{}, error) {
	params := &model.LoginRequest{}
	err := ctx.BindJSON(params)
	if err != nil {
		log.ErrorContext(ctx, "handler-failed", "error", err.Error())
		return nil, err
	}
	return h.logic.Login(ctx, params)
}

func (h *handler) Upload(c *gin.Context) {
	type response struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}
	resp := response{
		Code: int(errorcode.Code_OK),
	}
	file, err := c.FormFile("file")
	if err != nil {
		resp.Code = int(errorcode.Code_INTERNAL)
		resp.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	// c.JSON(200, gin.H{"message": file.Header.Context})
	//c.SaveUploadedFile(file, file.Filename) // TODO 改为OSS
	//c.String(http.StatusOK, file.Filename)

	reader, err := file.Open()
	if err != nil {
		resp.Code = int(errorcode.Code_INTERNAL)
		resp.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	defer reader.Close()
	salt := gocast.ToString(time.Now().In(model.ShanghaiLocation).Nanosecond())
	filePath := "tmp/test/" + salt + file.Filename
	log.Info("filePath", "path", filePath)
	if err := h.bucket.PutObject(filePath, reader); err != nil {
		resp.Code = int(errorcode.Code_INTERNAL)
		resp.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	resp.Data = struct {
		URL string `json:"url"`
	}{
		URL: "https://" + model.OssImageDomain + "/" + filePath,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func NewHandler(l logic.Interface) Handler {
	client, err := oss.New(model.OssEndPoint, model.OssAccessKey, model.OssAccessSecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(model.OssBucket)
	if err != nil {
		panic(err)
	}
	return &handler{
		logic:  l,
		bucket: bucket,
	}
}
