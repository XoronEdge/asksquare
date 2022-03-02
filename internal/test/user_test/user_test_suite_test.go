package user_test_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/XoronEdge/asksquare/initial"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	os.Setenv("ENV", "TEST")
}

var (
	db  *gorm.DB
	err error
)

func TestUserTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserTest Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Before suit invoked")
	db = initial.GetDB()
})

var _ = AfterSuite(func() {
	fmt.Println("After suit invoked")
	defer db.Close()
})
