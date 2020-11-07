package postgres

import (
	"context"
	"fmt"

	"github.com/XoronEdge/asksquare/domain"
	"github.com/jinzhu/gorm"
)

type qaReportRepo struct {
	db *gorm.DB
}

// NewQaReportRepo  ...
func NewQaReportRepo(db *gorm.DB) domain.QaReportRepo {
	return &qaReportRepo{db}
}

//Fetch ...
func (qr *qaReportRepo) Fetch(ctx context.Context) ([]*domain.QaReport, error) {
	qar := []*domain.QaReport{}
	err := qr.db.Find(&qar).Error
	if err != nil {
		return nil, err
	}
	return qar, nil
}

//Store ...
func (qr *qaReportRepo) Store(ctx context.Context, qar *domain.QaReport) (err error) {
	fmt.Printf("%+v\n", qar)
	err = qr.db.Create(&qar).Error
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qr *qaReportRepo) GetByID(ctx context.Context, id uint) (qar *domain.QaReport, err error) {
	qar = &domain.QaReport{}
	err = qr.db.First(&qar, id).Error
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qr *qaReportRepo) Update(ctx context.Context, qar *domain.QaReport) (err error) {
	err = qr.db.Save(&qar).Error
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qr *qaReportRepo) Delete(ctx context.Context, id uint) (err error) {
	err = qr.db.Delete(&domain.QaReport{}, id).Error
	return
}
