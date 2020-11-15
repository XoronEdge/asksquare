package wira

// import (
// 	"fmt"
// 	"time"

// 	"github.com/XoronEdge/asksquare/domain"
// 	"github.com/XoronEdge/asksquare/initial"
// 	qaActionRepo "github.com/XoronEdge/asksquare/internal/questionAction/repo/postgres"
// 	qaActionUsecase "github.com/XoronEdge/asksquare/internal/questionAction/usecase"
// 	userRepo "github.com/XoronEdge/asksquare/internal/user/repo/postgres"
// 	userUsecase "github.com/XoronEdge/asksquare/internal/user/usecase"
// 	"github.com/google/wire"
// )

// func init() {
// 	fmt.Println("----------Its Working")
// }

// //Di ...
// type Di struct {
// 	Uc  domain.UserUsecase
// 	QRc domain.QaReportUsecase
// 	QHc domain.QaHideUsecase
// }

// //NewDi return
// func NewDi(Uc domain.UserUsecase, QRc domain.QaReportUsecase, QHc domain.QaHideUsecase) Di {
// 	return Di{Uc: Uc, QRc: QRc, QHc: QHc}
// }

// //Set1 ...
// var Set1 = wire.NewSet(userUsecase.NewUserUsecase,
// 	wire.Bind(new(domain.UserUsecase), new(*userUsecase.UserUsecase)),
// )

// //Set2 ...
// var Set2 = wire.NewSet(qaActionUsecase.NewQaReportUsecase,
// 	wire.Bind(new(domain.QaReportUsecase), new(*qaActionUsecase.QaReportUsecase)),
// )

// //Set3 ...
// var Set3 = wire.NewSet(qaActionUsecase.NewQaHideUsecase,
// 	wire.Bind(new(domain.QaHideUsecase), new(*qaActionUsecase.QaHideUsecase)),
// )

// //NewTime ...
// func NewTime() time.Duration {
// 	return time.Minute
// }

// //InitializeDi ...
// func InitializeDi() Di {
// 	wire.Build(initial.GetDB, NewTime, userRepo.NewUserRepo,
// 		qaActionRepo.NewQaReportRepo,
// 		qaActionRepo.NewQaHideRepo,
// 		Set1, Set2, Set3,
// 		NewDi)
// 	return Di{}
// }
