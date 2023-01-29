package repository

import (
	"context"
	"self-payrol/config"
	"self-payrol/model"
)

type positionRepository struct {
	Cfg config.Config
}

func NewPositionRepository(cfg config.Config) model.PositionRepository {
	return &positionRepository{Cfg: cfg}
}

func (p *positionRepository) FindByID(ctx context.Context, id int) (*model.Position, error) {
	position := new(model.Position)

	if err := p.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		First(position).Error; err != nil {
		return nil, err
	}
	return position, nil
}

func (p *positionRepository) Create(ctx context.Context, position *model.Position) (*model.Position, error) {
	if err := p.Cfg.Database().WithContext(ctx).Create(&position).Error; err != nil {
		return nil, err
	}
	return position, nil
}

func (p *positionRepository) UpdateByID(ctx context.Context, id int, position *model.Position) (*model.Position, error) {
	if err := p.Cfg.Database().WithContext(ctx).
		Model(&model.Position{ID: id}).Updates(position).Find(position).Error; err != nil {
		return nil, err
	}

	return position, nil
}

func (p *positionRepository) Delete(ctx context.Context, id int) error {

	_, err := p.FindByID(ctx, id)

	if err != nil {
		return err
	}

	res := p.Cfg.Database().WithContext(ctx).
		Delete(&model.Position{}, id)
	if res.Error != nil {

		return res.Error
	}
	return nil
}

func (p *positionRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Position, error) {
	var data []*model.Position

	if err := p.Cfg.Database().WithContext(ctx).
		Select("id", "name", "salary", "created_at", "updated_at").
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
