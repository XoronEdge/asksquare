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

type qaAnswerLaterRepo struct {
	db *gorm.DB
}

// NewQaAnswerLaterRepo  ...
func NewQaAnswerLaterRepo(db *gorm.DB) domain.QaAnswerLaterRepo {
	return &qaAnswerLaterRepo{db}
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
func (qalr *qaHideRepo) Fetch(ctx context.Context) ([]*domain.QaHide, error) {
	qh := []*domain.QaHide{}
	err := qalr.db.Find(&qh).Error
	if err != nil {
		return nil, err
	}
	return qh, nil
}

//Store ...
func (qalr *qaHideRepo) Store(ctx context.Context, qh *domain.QaHide) (err error) {
	err = qalr.db.Create(&qh).Error
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qalr *qaHideRepo) GetByID(ctx context.Context, id uint) (qh *domain.QaHide, err error) {
	qh = &domain.QaHide{}
	err = qalr.db.First(&qh, id).Error
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qalr *qaHideRepo) Update(ctx context.Context, qh *domain.QaHide) (err error) {
	err = qalr.db.Save(&qh).Error
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qalr *qaHideRepo) Delete(ctx context.Context, id uint) (err error) {
	err = qalr.db.Delete(&domain.QaHide{}, id).Error
	return
}

////Question Answer Later--------------------------------
//Fetch ...
func (qalr *qaAnswerLaterRepo) Fetch(ctx context.Context) ([]*domain.QaAnswerLater, error) {
	qal := []*domain.QaAnswerLater{}
	err := qalr.db.Find(&qal).Error
	if err != nil {
		return nil, err
	}
	return qal, nil
}

//Store ...
func (qalr *qaAnswerLaterRepo) Store(ctx context.Context, qal *domain.QaAnswerLater) (err error) {
	err = qalr.db.Create(&qal).Error
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qalr *qaAnswerLaterRepo) GetByID(ctx context.Context, id uint) (qal *domain.QaAnswerLater, err error) {
	qal = &domain.QaAnswerLater{}
	err = qalr.db.First(&qal, id).Error
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qalr *qaAnswerLaterRepo) Update(ctx context.Context, qal *domain.QaAnswerLater) (err error) {
	err = qalr.db.Save(&qal).Error
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qalr *qaAnswerLaterRepo) Delete(ctx context.Context, id uint) (err error) {
	err = qalr.db.Delete(&domain.QaAnswerLater{}, id).Error
	return
}
