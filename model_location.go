/*
 * LocationIQ
 *
 * LocationIQ provides flexible enterprise-grade location based solutions. We work with developers, startups and enterprises worldwide serving billions of requests everyday. This page provides an overview of the technical aspects of our API and will help you get started.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package locationiq
// Location struct for Location
type Location struct {
	Distance float32 `json:"distance,omitempty"`
	PlaceId string `json:"place_id,omitempty"`
	Licence string `json:"licence,omitempty"`
	OsmType string `json:"osm_type,omitempty"`
	OsmId string `json:"osm_id,omitempty"`
	Boundingbox []string `json:"boundingbox,omitempty"`
	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Class string `json:"class,omitempty"`
	Type string `json:"type,omitempty"`
	Importance float32 `json:"importance,omitempty"`
	Address Address `json:"address,omitempty"`
	Namedetails Namedetails `json:"namedetails,omitempty"`
	Matchquality Matchquality `json:"matchquality,omitempty"`
}
