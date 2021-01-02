package helpers

var (
	PartialsNav               = "partials/_nav"
	PartialsNavSetup          = "partials/nav/_setup"
	PartialsNavMain           = "partials/nav/_main"
	PartialsNavAccountAccount = "partials/nav/account/_account"
	PartialsNavAccountSetup   = "partials/nav/account/_setup"
	PartialsNavAccountMain    = "partials/nav/account/_main"

	SetupNav = []string{
		PartialsNav,
		PartialsNavSetup,
		PartialsNavAccountAccount,
		PartialsNavAccountSetup,
	}

	DefaultNav = []string{
		PartialsNav,
		PartialsNavMain,
		PartialsNavAccountAccount,
		PartialsNavAccountMain,
	}
)
