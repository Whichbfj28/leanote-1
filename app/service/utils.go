package service

import (
	"context"

	"github.com/coocn-cn/leanote/app/note/model"
	"github.com/coocn-cn/leanote/app/note/repository"
)

func updateNote(repo repository.NoteRepository, ctx context.Context, predicate repository.Predicater, f func(*model.Note) error) error {
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

func updateNotes(repo repository.NoteRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.Note) error) error {
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

func updateContent(repo repository.ContentRepository, ctx context.Context, predicate repository.Predicater, f func(*model.Content) error) error {
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

func updateContents(repo repository.ContentRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.Content) error) error {
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

func updateHistory(repo repository.HistoryRepository, ctx context.Context, predicate repository.Predicater, f func(*model.History) error) error {
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

func updateHistorys(repo repository.HistoryRepository, ctx context.Context, predicate repository.Predicater, f func([]*model.History) error) error {
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
