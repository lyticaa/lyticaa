package helpers

const (
	setupNav = "partials/nav/_setup"
	mainNav  = "partials/nav/_main"
)

func NavForSession(subscribed bool) string {
	if subscribed {
		return mainNav
	}

	return setupNav
}
