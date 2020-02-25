/*
 * LocationIQ
 *
 * LocationIQ provides flexible enterprise-grade location based solutions. We work with developers, startups and enterprises worldwide serving billions of requests everyday. This page provides an overview of the technical aspects of our API and will help you get started.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package locationiq
// DirectionsMatrixSources struct for DirectionsMatrixSources
type DirectionsMatrixSources struct {
	Distance float32 `json:"distance,omitempty"`
	Location []float32 `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
}
