package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// 定义结构体
type Student struct {
	ID      string
	Name    string
	Sayname string
}

// 定义接口
type StudentRepo interface {
	Save(context.Context, *Student) (*Student, error)
	Get(context.Context, *Student) (*Student, error)
}

// 加上日志
type StudentUsercase struct {
	repo StudentRepo
	log  *log.Helper
}

// 初始化StudentUserCase
func NewStudentUserCase(repo StudentRepo, logger log.Logger) *StudentUsercase {
	return &StudentUsercase{repo: repo, log: log.NewHelper(logger)}
}

// CreateStudent
// 这里是一些业务逻辑
// 有点类似与三层架构里的logic
// biz主要完成业务逻辑组装，数据处理
func (uc *StudentUsercase) CreateStudent(ctx context.Context, stu *Student) (*Student, error) {
	uc.log.WithContext(ctx).Info("CreateStudent:%v", stu.ID)
	return uc.repo.Save(ctx, stu)
}
