package daoorg

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

type CompanyCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type CompanyDao struct {
	iamdao.Base
}

func NewCompanyDao() *CompanyDao {
	return &CompanyDao{}
}

func (d *CompanyDao) TableName() string {
	return iammodel.TableNameCompany
}

func (d *CompanyDao) WithTx(db *gorm.DB) *CompanyDao {
	return &CompanyDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *CompanyDao) Insert(ctx context.Context, entity *iammodel.CompanyEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[CompanyDao] Insert fail, entity:%s", gutil.ToJsonString(entity))
	}
	return nil
}

func (d *CompanyDao) BatchInsert(ctx context.Context, entityList iammodel.CompanyEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(gerror.DBInsertErr).Wrapf(nil, "[CompanyDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[CompanyDao] BatchInsert fail, entityList:%s", gutil.ToJsonString(entityList))
	}
	return nil
}

func (d *CompanyDao) UpdateByID(ctx context.Context, id uint, entity *iammodel.CompanyEntity) error {
	db := d.DB(ctx).Model(&iammodel.CompanyEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[CompanyDao] UpdateByID fail, id:%d entity:%s", id, gutil.ToJsonString(entity))
	}
	return nil
}

func (d *CompanyDao) UpdateMap(ctx context.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Model(&iammodel.CompanyEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[CompanyDao] UpdateMap fail, id:%d, updateMap:%s", id, gutil.ToJsonString(updateMap))
	}
	return nil
}

func (d *CompanyDao) Delete(ctx context.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Model(&iammodel.CompanyEntity{}).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(gerror.DBDeleteErr).Wrapf(err, "[CompanyDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *CompanyDao) GetById(ctx context.Context, id uint) (*iammodel.CompanyEntity, error) {
	var entity iammodel.CompanyEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[CompanyDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *CompanyDao) GetByCond(ctx context.Context, cond *CompanyCond) (*iammodel.CompanyEntity, error) {
	var entity iammodel.CompanyEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[CompanyDao] GetById fail, cond:%s", gutil.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *CompanyDao) GetListByCond(ctx context.Context, cond *CompanyCond) (iammodel.CompanyEntityList, error) {
	var entityList iammodel.CompanyEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[CompanyDao] GetListByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *CompanyDao) GetPageListByCond(ctx context.Context, cond *CompanyCond) (iammodel.CompanyEntityList, int64, error) {
	db := d.DB(ctx).Model(&iammodel.CompanyEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[CompanyDao] GetPageListByCond count fail, cond:%s", gutil.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.CompanyEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[CompanyDao] GetPageListByCond find fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *CompanyDao) CountByCond(ctx context.Context, cond *CompanyCond) (int64, error) {
	db := d.DB(ctx).Model(&iammodel.CompanyEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[CompanyDao] CountByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return count, nil
}

func (d *CompanyDao) BuildCondition(db *gorm.DB, cond *CompanyCond) {
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
