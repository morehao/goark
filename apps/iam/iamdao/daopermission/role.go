package daopermission

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

type RoleCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type RoleDao struct {
	iamdao.Base
}

func NewRoleDao() *RoleDao {
	return &RoleDao{}
}

func (d *RoleDao) TableName() string {
	return iammodel.TableNameRole
}

func (d *RoleDao) WithTx(db *gorm.DB) *RoleDao {
	return &RoleDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *RoleDao) Insert(ctx context.Context, entity *iammodel.RoleEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[RoleDao] Insert fail, entity:%s", gutil.ToJsonString(entity))
	}
	return nil
}

func (d *RoleDao) BatchInsert(ctx context.Context, entityList iammodel.RoleEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(gerror.DBInsertErr).Wrapf(nil, "[RoleDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[RoleDao] BatchInsert fail, entityList:%s", gutil.ToJsonString(entityList))
	}
	return nil
}

func (d *RoleDao) UpdateByID(ctx context.Context, id uint, entity *iammodel.RoleEntity) error {
	db := d.DB(ctx).Model(&iammodel.RoleEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[RoleDao] UpdateByID fail, id:%d entity:%s", id, gutil.ToJsonString(entity))
	}
	return nil
}

func (d *RoleDao) UpdateMap(ctx context.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Model(&iammodel.RoleEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[RoleDao] UpdateMap fail, id:%d, updateMap:%s", id, gutil.ToJsonString(updateMap))
	}
	return nil
}

func (d *RoleDao) Delete(ctx context.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Model(&iammodel.RoleEntity{}).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(gerror.DBDeleteErr).Wrapf(err, "[RoleDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *RoleDao) GetById(ctx context.Context, id uint) (*iammodel.RoleEntity, error) {
	var entity iammodel.RoleEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *RoleDao) GetByCond(ctx context.Context, cond *RoleCond) (*iammodel.RoleEntity, error) {
	var entity iammodel.RoleEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleDao] GetById fail, cond:%s", gutil.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *RoleDao) GetListByCond(ctx context.Context, cond *RoleCond) (iammodel.RoleEntityList, error) {
	var entityList iammodel.RoleEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleDao] GetListByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *RoleDao) GetPageListByCond(ctx context.Context, cond *RoleCond) (iammodel.RoleEntityList, int64, error) {
	db := d.DB(ctx).Model(&iammodel.RoleEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleDao] GetPageListByCond count fail, cond:%s", gutil.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.RoleEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleDao] GetPageListByCond find fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *RoleDao) CountByCond(ctx context.Context, cond *RoleCond) (int64, error) {
	db := d.DB(ctx).Model(&iammodel.RoleEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleDao] CountByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return count, nil
}

func (d *RoleDao) BuildCondition(db *gorm.DB, cond *RoleCond) {
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
