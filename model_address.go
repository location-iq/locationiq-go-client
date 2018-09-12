/*
 * LocationIQ
 *
 * LocationIQ provides flexible enterprise-grade location based solutions. We work with developers, startups and enterprises worldwide serving billions of requests everyday. This page provides an overview of the technical aspects of our API and will help you get started.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package locationiq

type Address struct {
	HouseNumber string `json:"house_number,omitempty"`
	Road string `json:"road,omitempty"`
	Residential string `json:"residential,omitempty"`
	Village string `json:"village,omitempty"`
	County string `json:"county,omitempty"`
	State string `json:"state,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	Country string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City string `json:"city,omitempty"`
}
