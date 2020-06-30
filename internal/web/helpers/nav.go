package helpers

const (
	setupPrimaryNav = "partials/nav/_setup"
	mainPrimaryNav  = "partials/nav/_main"
	setupAccountNav = "partials/nav/account/_setup"
	mainAccountNav  = "partials/nav/account/_main"
)

func PrimaryNavForSession(subscribed bool) string {
	if subscribed {
		return mainPrimaryNav
	}

	return setupPrimaryNav
}

func AccountNavForSession(subscribed bool) string {
	if subscribed {
		return mainAccountNav
	}

	return setupAccountNav
}
