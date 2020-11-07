package graph

import (
	"github.com/XoronEdge/asksquare/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ...
type Resolver struct {
	Uc  domain.UserUsecase
	QRc domain.QaReportUsecase
}
