package postgres

import (
	"context"

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

type qaHideRepo struct {
	db *gorm.DB
}

// NewQaHideRepo  ...
func NewQaHideRepo(db *gorm.DB) domain.QaHideRepo {
	return &qaHideRepo{db}
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

////Question Hide--------------------------------
//Fetch ...
func (qhr *qaHideRepo) Fetch(ctx context.Context) ([]*domain.QaHide, error) {
	qh := []*domain.QaHide{}
	err := qhr.db.Find(&qh).Error
	if err != nil {
		return nil, err
	}
	return qh, nil
}

//Store ...
func (qhr *qaHideRepo) Store(ctx context.Context, qh *domain.QaHide) (err error) {
	err = qhr.db.Create(&qh).Error
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qhr *qaHideRepo) GetByID(ctx context.Context, id uint) (qh *domain.QaHide, err error) {
	qh = &domain.QaHide{}
	err = qhr.db.First(&qh, id).Error
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qhr *qaHideRepo) Update(ctx context.Context, qh *domain.QaHide) (err error) {
	err = qhr.db.Save(&qh).Error
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qhr *qaHideRepo) Delete(ctx context.Context, id uint) (err error) {
	err = qhr.db.Delete(&domain.QaHide{}, id).Error
	return
}
