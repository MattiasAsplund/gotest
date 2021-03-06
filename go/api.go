/*
 * webapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// WeatherForecastApiRouter defines the required methods for binding the api requests to a responses for the WeatherForecastApi
// The WeatherForecastApiRouter implementation should parse necessary information from the http request,
// pass the data to a WeatherForecastApiServicer to perform the required actions, then write the service results to the http response.
type WeatherForecastApiRouter interface { 
	GetWeatherForecast(http.ResponseWriter, *http.Request)
}


// WeatherForecastApiServicer defines the api actions for the WeatherForecastApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type WeatherForecastApiServicer interface { 
	GetWeatherForecast(context.Context) (ImplResponse, error)
}
