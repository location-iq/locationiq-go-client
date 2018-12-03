/*
 * LocationIQ
 *
 * LocationIQ provides flexible enterprise-grade location based solutions. We work with developers, startups and enterprises worldwide serving billions of requests everyday. This page provides an overview of the technical aspects of our API and will help you get started.
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package locationiq

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ReverseApiService service

/*
ReverseApiService Reverse Geocoding
Reverse geocoding is the process of converting a coordinate or location (latitude, longitude) to a readable address or place name. This permits the identification of nearby street addresses, places, and/or area subdivisions such as a neighborhood, county, state, or country.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param lat Latitude of the location to generate an address for.
 * @param lon Longitude of the location to generate an address for.
 * @param format Format to geocode. Only JSON supported for SDKs
 * @param normalizecity Normalizes village to city level data to city
 * @param optional nil or *ReverseOpts - Optional Parameters:
 * @param "Addressdetails" (optional.Int32) -  Include a breakdown of the address into elements. Defaults to 1.
 * @param "AcceptLanguage" (optional.String) -  Preferred language order for showing search results, overrides the value specified in the Accept-Language HTTP header. Defaults to en. To use native language for the response when available, use accept-language=native
 * @param "Namedetails" (optional.Int32) -  Include a list of alternative names in the results. These may include language variants, references, operator and brand.
 * @param "Extratags" (optional.Int32) -  Include additional information in the result if available, e.g. wikipedia link, opening hours.
@return Location
*/

type ReverseOpts struct {
    Addressdetails optional.Int32
    AcceptLanguage optional.String
    Namedetails optional.Int32
    Extratags optional.Int32
}

func (a *ReverseApiService) Reverse(ctx context.Context, lat float32, lon float32, format string, normalizecity int32, localVarOptionals *ReverseOpts) (Location, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue Location
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reverse.php"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if lat < -90 {
		return localVarReturnValue, nil, reportError("lat must be greater than -90")
	}
	if lat > 90 {
		return localVarReturnValue, nil, reportError("lat must be less than 90")
	}
	if lon < -180 {
		return localVarReturnValue, nil, reportError("lon must be greater than -180")
	}
	if lon > 180 {
		return localVarReturnValue, nil, reportError("lon must be less than 180")
	}

	localVarQueryParams.Add("lat", parameterToString(lat, ""))
	localVarQueryParams.Add("lon", parameterToString(lon, ""))
	localVarQueryParams.Add("format", parameterToString(format, ""))
	localVarQueryParams.Add("normalizecity", parameterToString(normalizecity, ""))
	if localVarOptionals != nil && localVarOptionals.Addressdetails.IsSet() {
		localVarQueryParams.Add("addressdetails", parameterToString(localVarOptionals.Addressdetails.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AcceptLanguage.IsSet() {
		localVarQueryParams.Add("accept-language", parameterToString(localVarOptionals.AcceptLanguage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Namedetails.IsSet() {
		localVarQueryParams.Add("namedetails", parameterToString(localVarOptionals.Namedetails.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Extratags.IsSet() {
		localVarQueryParams.Add("extratags", parameterToString(localVarOptionals.Extratags.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("key", key)
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Location
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
