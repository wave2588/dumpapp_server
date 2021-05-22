package impl

type AuthController struct{}

var DefaultAuthController *AuthController

func init() {
	DefaultAuthController = NewAuthController()
}

func NewAuthController() *AuthController {
	return &AuthController{}
}
