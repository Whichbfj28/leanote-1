package service

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/repository"
	tag_model "github.com/coocn-cn/leanote/app/tag/model"
	tag_repo "github.com/coocn-cn/leanote/app/tag/repository"
)

func updateTag(repo tag_repo.TagRepository, ctx context.Context, predicate repository.Predicater, f func(*tag_model.Tag) error) error {
	model, err := repo.Find(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(model); err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	return repo.Save(ctx, model)
}

func updateTags(repo tag_repo.TagRepository, ctx context.Context, predicate repository.Predicater, f func([]*tag_model.Tag) error) error {
	models, err := repo.FindAll(ctx, predicate)
	if err != nil {
		return err
	}

	if err := f(models); err != nil {
		return err
	}

	if models == nil {
		return nil
	}

	return repo.Save(ctx, models...)
}
