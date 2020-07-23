package model
//该包对文章的模型操作进行封装，封装创建，更新，获取等方法
import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

type ArticleRow struct {
	ArticleID uint32
	TagID uint32
	TagName string
	ArticleTitle string
	ArticleDesc string
	CoverImagesUrl string
	Content string
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	err := db.Create(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

//文章的更新操作，values存放要更新的数据
func (a Article) Update(db *gorm.DB, values interface{}) error {
	err := db.Update(values).Where("id = ? and is_del = ?", a.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? and state = ? and is_del = ?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	err := db.Where("id = ? and is_del = ?", a.Model.ID, 0).Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}

//获取文章列表时，要进行关联查询
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title",
		"ar.desc AS article_desc", "ar.cover_image_url", "ar.content",
	"t.id AS tag_id", "t.name AS tag_name"}

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	//Select指定要检索的数据库字段，不指定则默认检索全部字段
	//joins指定关联查询的语句
	//利用blog_article_tag表的文章ID和标签ID进行了两次左关联查询，并且为防止有重合的字段名，
	//在select子句中通过as分别设置了字段的别名
	//使用rows扫描后，结果与原生的model结构体不同，因此重新定义了ArticleRow结构体,
	//用于处理和返回文章列表的结果集
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName() + "AS at").
		Joins("left join `" + Tag{}.TableName() + "` as t on at.tag_id = t.id").
		Joins("left join `" + Article{}.TableName() + "` as ar on at.article_id = ar.id").
		Where("at.tag_id = ? and ar.state = ? and ar.is_del = ?", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImagesUrl,
			&r.Content, &r.TagID, &r.TagName)
		if err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

//该查询方法用来获取文章列表总数
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName() + " as at").
		Joins("left join `" + Tag{}.TableName() + "` as t on at.tag_id = t.id").
		Joins("left join `" + Article{}.TableName() + "` as ar on at.article_id = ar.id").
		Where("at.tag_id = ? and ar.state = ? and ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}