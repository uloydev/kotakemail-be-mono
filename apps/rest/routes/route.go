package routes

import "kotakemail.id/pkg/rest"

func GetRoutes() []*rest.RestRoute {
	routes := []*rest.RestRoute{
		HeathCheckRoute(),
	}

	return routes
}
