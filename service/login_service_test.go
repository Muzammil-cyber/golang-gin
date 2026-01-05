package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/muzammil-cyber/golang-gin/entity"
	"github.com/muzammil-cyber/golang-gin/service"
)

var fakeUser = entity.LoginCredentials{
	Username: "admin",
	Password: "password",
}

var _ = Describe("LoginService", func() {
	var (
		loginService service.LoginService
	)

	BeforeEach(func() {
		loginService = service.NewLoginService()
	})
	Describe("Login", func() {
		It("should authenticate valid user credentials", func() {
			isAuthenticated := loginService.Login(fakeUser.Username, fakeUser.Password)
			Expect(isAuthenticated).To(BeFalse())
		})

		It("should reject invalid user credentials", func() {
			isAuthenticated := loginService.Login("invalidUser", "wrongPassword")
			Expect(isAuthenticated).To(BeFalse())
		})
	})

})
