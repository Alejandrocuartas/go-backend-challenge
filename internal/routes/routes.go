package router

import (
	"fmt"
	"go-backend-challenge/environment"
	"go-backend-challenge/internal/controller"
	"net/http"

	utils "github.com/Alejandrocuartas/core-utils-private-library"

	"github.com/gorilla/mux"
)

func ApiRouter(c controller.CustomControllerStruct) *mux.Router {
	r := mux.NewRouter()

	basicAuth := utils.BasicAuth(
		environment.BasicAuthUsername,
		environment.BasicAuthPassword,
	)

	r.Use(basicAuth)

	r.NotFoundHandler = http.HandlerFunc(utils.NotFoundHandler)

	r.HandleFunc(
		"/v1/actions",
		c.Actions.CreateActionControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/useragencyrelation",
		c.UserAgencyRelations.CreateUserAgencyRelationsControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/users",
		c.Users.CreateUserControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/agencies",
		c.Agencies.CreateAgencyControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/campaigns",
		c.Campaigns.CreateCampaignControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/campaigns/list",
		c.Campaigns.ListCampaignsControllerMethod,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/v1/campaigns/{campaign_id:[0-9]+}",
		c.Campaigns.GetCampaignByIdControllerMethod,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/v1/campaigns/{campaign_id:[0-9]+}",
		c.Campaigns.UpdateCampaignControllerMethod,
	).Methods(http.MethodPut)

	fmt.Println("Available Routes:")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, e := route.GetPathTemplate()
		if e != nil {
			return e
		}

		methods, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println(t, methods)
		return nil
	})

	return r
}
