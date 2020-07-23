package service

import (
	"blog-service/internal/dao"
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

type Article struct {
	ID uint32 `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
	Tag *model.Tag `json:"tag"`
}

type ArticleRequest struct {
	ID uint32
	State uint8
}

type CountArticleRequest struct {
	Title string
	Desc string
	Content string
	CoverImageUrl string
	CreatedBy string
	State uint8
	TagID uint32
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	Desc string `form:"desc" binding:"min=3,max=100"`
	Content string `form:"content" binding:"min=3"`
	CoverImageUrl string `form:"cover_image_url"`
	TagID uint32 `form:"tag_id"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID uint32 `form:"id"`
	Title string `form:"title" binding:"min=3,max=100"`
	State uint8 `form:"state"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	Desc string `form:"desc" binding:"min=3,max=100"`
	Content string `form:"content" binding:"min=3"`
	CoverImageUrl string `form:"cover_image_url"`
	TagID uint32 `form:"tag_id"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	articleTag, err := svc.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}
	tag, err := svc.dao.GetTag(articleTag.TagID)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           tag,
	}, nil
}

func (svc *Service) GetArticleList(param *ArticleListRequest,
	pager *app.Pager) ([]*Article, int, error) {
		articleCount, err := svc.dao.CountArticleListByTagID(param.TagID, param.State)
		if err != nil {
			return nil, 0, err
		}

		articles, err := svc.dao.GetArticleListByTagID(param.TagID, param.State,
			pager.Page, pager.PageSize)
		if err != nil {
			return nil, 0, err
		}
		var articleList []*Article
		for _, article := range articles {
			articleList = append(articleList, &Article{
				ID:            article.ArticleID,
				Title:         article.ArticleTitle,
				Desc:          article.ArticleDesc,
				Content:       article.Content,
				CoverImageUrl: article.CoverImagesUrl,
				Tag:           &model.Tag{Model: &model.Model{ID:article.TagID},
					Name:article.TagName},
			})
		}

		return articleList, articleCount, nil
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		CreatedBy:     param.CreatedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.ID, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		ModifiedBy:    param.ModifiedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}
	err = svc.dao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}
	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}
	return nil
}
