package daoorg

import (
	"fmt"
	"time"

	"github.com/morehao/goark/apps/iam/iamdao"
	"github.com/morehao/goark/apps/iam/iammodel"
	"github.com/morehao/goark/pkg/code"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gutils"
	"gorm.io/gorm"
)

type DepartmentCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type DepartmentDao struct {
	iamdao.Base
}

func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{}
}

func (d *DepartmentDao) TableName() string {
	return iammodel.TableNameDepartment
}

func (d *DepartmentDao) WithTx(db *gorm.DB) *DepartmentDao {
	return &DepartmentDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *DepartmentDao) Insert(ctx *gin.Context, entity *iammodel.DepartmentEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(code.DBInsertErr).Wrapf(err, "[DepartmentDao] Insert fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (d *DepartmentDao) BatchInsert(ctx *gin.Context, entityList iammodel.DepartmentEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(code.DBInsertErr).Wrapf(nil, "[DepartmentDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(code.DBInsertErr).Wrapf(err, "[DepartmentDao] BatchInsert fail, entityList:%s", gutils.ToJsonString(entityList))
	}
	return nil
}

func (d *DepartmentDao) UpdateByID(ctx *gin.Context, id uint, entity *iammodel.DepartmentEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[DepartmentDao] UpdateByID fail, id:%d entity:%s", id, gutils.ToJsonString(entity))
	}
	return nil
}

func (d *DepartmentDao) UpdateMap(ctx *gin.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[DepartmentDao] UpdateMap fail, id:%d, updateMap:%s", id, gutils.ToJsonString(updateMap))
	}
	return nil
}

func (d *DepartmentDao) Delete(ctx *gin.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[DepartmentDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *DepartmentDao) GetById(ctx *gin.Context, id uint) (*iammodel.DepartmentEntity, error) {
	var entity iammodel.DepartmentEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[DepartmentDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *DepartmentDao) GetByCond(ctx *gin.Context, cond *DepartmentCond) (*iammodel.DepartmentEntity, error) {
	var entity iammodel.DepartmentEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[DepartmentDao] GetById fail, cond:%s", gutils.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *DepartmentDao) GetListByCond(ctx *gin.Context, cond *DepartmentCond) (iammodel.DepartmentEntityList, error) {
	var entityList iammodel.DepartmentEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[DepartmentDao] GetListByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *DepartmentDao) GetPageListByCond(ctx *gin.Context, cond *DepartmentCond) (iammodel.DepartmentEntityList, int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(code.DBFindErr).Wrapf(err, "[DepartmentDao] GetPageListByCond count fail, cond:%s", gutils.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.DepartmentEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(code.DBFindErr).Wrapf(err, "[DepartmentDao] GetPageListByCond find fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *DepartmentDao) CountByCond(ctx *gin.Context, cond *DepartmentCond) (int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(code.DBFindErr).Wrapf(err, "[DepartmentDao] CountByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return count, nil
}

func (d *DepartmentDao) BuildCondition(db *gorm.DB, cond *DepartmentCond) {
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
