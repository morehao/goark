package daouser

import (
	"context"
	"fmt"
	"time"

	"github.com/morehao/goark/apps/iam/iamdao"
	"github.com/morehao/goark/apps/iam/iammodel"
	"github.com/morehao/goark/pkg/code"
	"github.com/morehao/golib/gerror"
	"github.com/morehao/golib/gutil"
	"gorm.io/gorm"
)

type UserCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type UserDao struct {
	iamdao.Base
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) TableName() string {
	return iammodel.TableNameUser
}

func (d *UserDao) WithTx(db *gorm.DB) *UserDao {
	return &UserDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *UserDao) Insert(ctx context.Context, entity *iammodel.UserEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[UserDao] Insert fail, entity:%s", gutil.ToJsonString(entity))
	}
	return nil
}

func (d *UserDao) BatchInsert(ctx context.Context, entityList iammodel.UserEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(gerror.DBInsertErr).Wrapf(nil, "[UserDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[UserDao] BatchInsert fail, entityList:%s", gutil.ToJsonString(entityList))
	}
	return nil
}

func (d *UserDao) UpdateByID(ctx context.Context, id uint, entity *iammodel.UserEntity) error {
	db := d.DB(ctx).Model(&iammodel.UserEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[UserDao] UpdateByID fail, id:%d entity:%s", id, gutil.ToJsonString(entity))
	}
	return nil
}

func (d *UserDao) UpdateMap(ctx context.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Model(&iammodel.UserEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[UserDao] UpdateMap fail, id:%d, updateMap:%s", id, gutil.ToJsonString(updateMap))
	}
	return nil
}

func (d *UserDao) Delete(ctx context.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Model(&iammodel.UserEntity{}).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(gerror.DBDeleteErr).Wrapf(err, "[UserDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *UserDao) GetById(ctx context.Context, id uint) (*iammodel.UserEntity, error) {
	var entity iammodel.UserEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[UserDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *UserDao) GetByCond(ctx context.Context, cond *UserCond) (*iammodel.UserEntity, error) {
	var entity iammodel.UserEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[UserDao] GetById fail, cond:%s", gutil.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *UserDao) GetListByCond(ctx context.Context, cond *UserCond) (iammodel.UserEntityList, error) {
	var entityList iammodel.UserEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[UserDao] GetListByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *UserDao) GetPageListByCond(ctx context.Context, cond *UserCond) (iammodel.UserEntityList, int64, error) {
	db := d.DB(ctx).Model(&iammodel.UserEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[UserDao] GetPageListByCond count fail, cond:%s", gutil.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.UserEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[UserDao] GetPageListByCond find fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *UserDao) CountByCond(ctx context.Context, cond *UserCond) (int64, error) {
	db := d.DB(ctx).Model(&iammodel.UserEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[UserDao] CountByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return count, nil
}

func (d *UserDao) BuildCondition(db *gorm.DB, cond *UserCond) {
	if cond.ID > 0 {
		query := fmt.Sprintf("%s.id = ?", d.TableName())
		db.Where(query, cond.ID)
	}
	if len(cond.IDs) > 0 {
		query := fmt.Sprintf("%s.id in (?)", d.TableName())
		db.Where(query, cond.IDs)
	}
	if cond.CreatedAtStart > 0 {
		query := fmt.Sprintf("%s.created_at >= ?", d.TableName())
		db.Where(query, time.Unix(cond.CreatedAtStart, 0))
	}
	if cond.CreatedAtEnd > 0 {
		query := fmt.Sprintf("%s.created_at <= ?", d.TableName())
		db.Where(query, time.Unix(cond.CreatedAtEnd, 0))
	}
	if cond.IsDelete {
		db.Unscoped()
	}

	if cond.OrderField != "" {
		db.Order(cond.OrderField)
	}
}
