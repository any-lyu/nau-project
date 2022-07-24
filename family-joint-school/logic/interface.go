package logic

import (
	"context"
	"family-joint-school/common/errors/apierror"
	"family-joint-school/common/errors/errorcode"
	"family-joint-school/common/jwt"
	"family-joint-school/common/util"
	"family-joint-school/model"
	"family-joint-school/mysql/ent"
	"family-joint-school/mysql/ent/account"
	"family-joint-school/mysql/ent/homework"
	"family-joint-school/mysql/ent/notice"
	"family-joint-school/mysql/ent/user"
	"family-joint-school/mysql/ent/userhomework"
	"github.com/chanxuehong/log"
	"github.com/demdxx/gocast"
)

type Interface interface {
	Run(ctx context.Context)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginReply, error)
	NoticeList(ctx context.Context, req *model.NoticeListRequest) (*model.NoticeListReply, error)
	NoticePublish(ctx context.Context, req *model.NoticePublishRequest) (*model.NoticePublishReply, error)
	ClassList(ctx context.Context, req *model.ClassListRequest) (*model.ClassListReply, error)
	HomeWorkList(ctx context.Context, req *model.HomeWorkListRequest) (*model.HomeWorkListReply, error)
	HomeWorkPublish(ctx context.Context, req *model.HomeWorkPublishRequest) (*model.HomeWorkPublishReply, error)
	HomeWorkSubmit(ctx context.Context, req *model.HomeWorkSubmitRequest) (*model.HomeWorkSubmitReply, error)
	UserHomeWorkDetail(ctx context.Context, req *model.UserHomeWorkDetailRequest) (*model.UserHomeWorkDetailReply, error)
	UserHomeWorkList(ctx context.Context, req *model.UserHomeWorkRequest) (*model.UserHomeWorkListReply, error)
}

type logic struct {
	db *ent.Client
}

func (l *logic) Run(ctx context.Context) {
	_, err := l.HomeWorkPublish(ctx, &model.HomeWorkPublishRequest{
		Title:   "test-1",
		Content: "test24324234r3wjkljlk",
		Pics: []model.Pic{
			{
				Url: "http://baidu.com/11.image",
			},
			{
				Url: "http://baidu.com/22.image",
			},
		},
	})
	if err != nil {
		log.ErrorContext(ctx, "failed", "error", err.Error())
		return
	}
}

func (l *logic) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginReply, error) {
	_account, err := l.db.Account.Query().Where(account.Account(req.Account)).First(ctx)
	if err != nil {
		log.ErrorContext(ctx, "login-failed", "error", err.Error())
		return nil, apierror.New(errorcode.Code_INTERNAL, "账号错误")
	}

	if _account.Password != util.Md5(req.Password) {
		return nil, apierror.New(errorcode.Code_INTERNAL, "密码错误")
	}

	// 登录成功
	_user, err := l.db.User.Query().Where(user.ID(_account.UserID)).First(ctx)
	if err != nil {
		log.ErrorContext(ctx, "login-failed", "error", err.Error())
		return nil, apierror.New(errorcode.Code_INTERNAL, err.Error())
	}
	tokenGenerate, err := jwt.TokenGenerate(gocast.ToString(_user.ID), model.GetPermissionName(_account.Permission))
	if err != nil {
		log.ErrorContext(ctx, "login-failed", "error", err.Error())
		return nil, apierror.New(errorcode.Code_INTERNAL, err.Error())
	}
	var mobile = _user.Mobile
	if len(mobile) > 6 {
		_user.Mobile = mobile[:3] + "****" + mobile[len(mobile)-4:]
	}
	return &model.LoginReply{
		Token:          tokenGenerate,
		User:           _user,
		Permission:     _account.Permission,
		PermissionName: model.GetPermissionName(_account.Permission),
	}, nil
}

func (l *logic) NoticeList(ctx context.Context, req *model.NoticeListRequest) (*model.NoticeListReply, error) {
	id := gocast.ToUint64(req.Cursor)
	if id == 0 {
		id = model.MaxID
	}
	notices, err := l.db.Notice.Query().Where(notice.IDLT(id)).Order(ent.Desc(notice.FieldID)).Limit(10).All(ctx)
	if err != nil {
		log.ErrorContext(ctx, "NoticeList-failed", "error", err.Error())
		return nil, apierror.New(errorcode.Code_INTERNAL, err.Error())
	}
	var lastID uint64
	if len(notices) > 0 {
		lastID = notices[len(notices)-1].ID
	}
	return &model.NoticeListReply{
		List:   notices,
		Tips:   "level:1-普通通知 2-重要通知 3-紧急通知",
		Cursor: gocast.ToString(lastID),
	}, nil
}

func (l *logic) NoticePublish(ctx context.Context, req *model.NoticePublishRequest) (*model.NoticePublishReply, error) {
	_notice, err := l.db.Notice.Create().SetTitle(req.Title).SetContent(req.Content).SetLevel(uint8(req.Level)).Save(ctx)
	if err != nil {
		log.ErrorContext(ctx, "NoticePublish-failed", "error", err.Error())
		return nil, err
	}
	return &model.NoticePublishReply{ID: int64(_notice.ID)}, nil
}

func (l *logic) ClassList(ctx context.Context, req *model.ClassListRequest) (*model.ClassListReply, error) {
	classes, err := l.db.Class.Query().All(ctx)
	if err != nil {
		log.ErrorContext(ctx, "ClassList-failed", "error", err.Error())
		return nil, err
	}
	return &model.ClassListReply{List: classes}, nil
}

func (l *logic) HomeWorkList(ctx context.Context, req *model.HomeWorkListRequest) (*model.HomeWorkListReply, error) {
	id := gocast.ToUint64(req.Cursor)
	if id == 0 {
		id = model.MaxID
	}
	homeworks, err := l.db.Homework.Query().Where(homework.IDLT(id)).Order(ent.Desc(homework.FieldID)).All(ctx)
	if err != nil {
		log.ErrorContext(ctx, "HomeWorkList-failed", "error", err.Error())
		return nil, err
	}
	var (
		list   []*model.Homework
		lastID uint64
	)
	for _, h := range homeworks {
		list = append(list, &model.Homework{
			ID:        h.ID,
			Title:     h.Title,
			Content:   h.Content,
			Pics:      h.Pics,
			UpdatedAt: h.UpdatedAt,
			CreatedAt: h.CreatedAt,
		})
		lastID = h.ID
	}
	return &model.HomeWorkListReply{
		List:   list,
		Cursor: gocast.ToString(lastID),
	}, nil
}

func (l *logic) HomeWorkPublish(ctx context.Context, req *model.HomeWorkPublishRequest) (*model.HomeWorkPublishReply, error) {
	_homework, err := l.db.Homework.Create().SetTitle(req.Title).SetClassID(req.ClassID).
		SetContent(req.Content).SetPics(req.Pics).Save(ctx)
	if err != nil {
		log.ErrorContext(ctx, "HomeWorkPublish-failed", "error", err.Error())
		return nil, err
	}
	return &model.HomeWorkPublishReply{
		ID: int64(_homework.ID),
	}, nil
}

func (l *logic) HomeWorkSubmit(ctx context.Context, req *model.HomeWorkSubmitRequest) (*model.HomeWorkSubmitReply, error) {
	userHomeWork, err := l.db.UserHomeWork.Create().SetHomeworkID(uint64(req.HomeworkID)).SetTitle(req.Title).
		SetContent(req.Content).SetPics(req.Pics).SetUserID(req.UserID).Save(ctx)
	if err != nil {
		log.ErrorContext(ctx, "HomeWorkSubmit-failed", "error", err.Error())
		return nil, err
	}
	return &model.HomeWorkSubmitReply{ID: int64(userHomeWork.ID)}, nil
}

func (l *logic) UserHomeWorkDetail(ctx context.Context, req *model.UserHomeWorkDetailRequest) (*model.UserHomeWorkDetailReply, error) {
	_homework, err := l.db.Homework.Query().Where(homework.ID(uint64(req.HomeworkID))).First(ctx)
	if err != nil {
		log.ErrorContext(ctx, "UserHomeWorkDetail-failed", "error", err.Error())
		return nil, err
	}
	reply := &model.UserHomeWorkDetailReply{
		Status: 0,
		Homework: &model.Homework{
			ID:        _homework.ID,
			Title:     _homework.Title,
			Content:   _homework.Content,
			Pics:      _homework.Pics,
			UpdatedAt: _homework.UpdatedAt,
			CreatedAt: _homework.CreatedAt,
		},
		UserHomework: nil,
	}
	userHomework, err := l.db.UserHomeWork.Query().
		Where(userhomework.HomeworkID(uint64(req.HomeworkID))).Where(userhomework.UserID(req.UserID)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return reply, nil
		}
		log.ErrorContext(ctx, "UserHomeWorkDetail-failed", "error", err.Error())
		return nil, err
	}
	reply.Status = 1
	reply.UserHomework = &model.Homework{
		ID:        userHomework.ID,
		Title:     userHomework.Title,
		Content:   userHomework.Content,
		Pics:      userHomework.Pics,
		UpdatedAt: userHomework.UpdatedAt,
		CreatedAt: userHomework.CreatedAt,
	}
	return reply, nil

}

func (l *logic) UserHomeWorkList(ctx context.Context, req *model.UserHomeWorkRequest) (*model.UserHomeWorkListReply, error) {
	id := gocast.ToUint64(req.Cursor)
	if id == 0 {
		id = model.MaxID
	}
	homeworks, err := l.db.UserHomeWork.Query().Where(userhomework.IDLT(id)).Order(ent.Desc(userhomework.FieldID)).All(ctx)
	if err != nil {
		log.ErrorContext(ctx, "UserHomeWorkList-failed", "error", err.Error())
		return nil, err
	}
	var list []*model.Homework
	var lastID uint64 = 0
	for _, h := range homeworks {
		list = append(list, &model.Homework{
			ID:        h.ID,
			Title:     h.Title,
			Content:   h.Content,
			Pics:      h.Pics,
			UpdatedAt: h.UpdatedAt,
			CreatedAt: h.CreatedAt,
		})
		lastID = h.ID
	}
	return &model.UserHomeWorkListReply{
		List:   list,
		Cursor: gocast.ToString(lastID),
	}, nil
}

func NewLogic(db *ent.Client) Interface {
	return &logic{
		db: db,
	}
}
