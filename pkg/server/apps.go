package server

import (
	"rest/pkg/apps/companies"
	"rest/pkg/apps/user"
	"rest/pkg/shared"
)

var apps = shared.Apps{
	shared.Route("/users", user.App),
	shared.Route("/companies", companies.App),
}
