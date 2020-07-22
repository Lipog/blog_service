package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

//用于处理标签模块的DAO的操作

//d里面的gorm.DB实例通过tag.Count方法，将获得的数据放入到实例中去，即通过实例来查询
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name:name, State:state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name:name, State:state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Model: &model.Model{CreatedBy:createdBy},
		Name:  name,
		State: state,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID:id, ModifiedBy:modifiedBy},
		Name:  name,
		State: state,
	}
	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model:&model.Model{ID:id}}
	return tag.Delete(d.engine)
}


