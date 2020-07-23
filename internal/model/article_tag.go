package model

import "github.com/jinzhu/gorm"

//该包对文章和标签的关联模型操作进行封装
type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (a ArticleTag) GetByID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? and is_del = ?", a.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}
	return articleTag, nil
}

func (a ArticleTag) ListByID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("tag_id = ? and is_del = ?", a.TagID, 0).Find(&articleTags).Error
	if err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (a ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id in (?) and is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (a ArticleTag) Create(db *gorm.DB) error {
	err := db.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	err := db.Model(&a).Where("article_id = ? and is_del = ?", a.ArticleID, 0).
		Limit(1).Update(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	err := db.Where("id = ? and is_del = ?", a.Model.ID, 0).Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	err := db.Where("article_id = ? and is_del = ?", a.ArticleID, 0).
		Delete(&a).Limit(1).Error
	if err != nil {
		return err
	}
	return nil
}