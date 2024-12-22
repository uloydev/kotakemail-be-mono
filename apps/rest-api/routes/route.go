package routes

import "kotakemail.id/shared/base/rest"

func GetRoutes() []*rest.RestRoute {
	routes := []*rest.RestRoute{
		rest.SwaggerRoute(),
		HeathCheckRoute(),
	}

	return routes
}
