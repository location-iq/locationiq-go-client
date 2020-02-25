/*
 * LocationIQ
 *
 * LocationIQ provides flexible enterprise-grade location based solutions. We work with developers, startups and enterprises worldwide serving billions of requests everyday. This page provides an overview of the technical aspects of our API and will help you get started.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package locationiq

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ReverseApiService ReverseApi service
type ReverseApiService service

// ReverseOpts Optional parameters for the method 'Reverse'
type ReverseOpts struct {
    Addressdetails optional.Int32
    AcceptLanguage optional.String
    Namedetails optional.Int32
    Extratags optional.Int32
    Statecode optional.Int32
    Showdistance optional.Int32
    Postaladdress optional.Int32
}

/*
Reverse Reverse Geocoding
Reverse geocoding is the process of converting a coordinate or location (latitude, longitude) to a readable address or place name. This permits the identification of nearby street addresses, places, and/or area subdivisions such as a neighborhood, county, state, or country.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param lat Latitude of the location to generate an address for.
 * @param lon Longitude of the location to generate an address for.
 * @param format Format to geocode. Only JSON supported for SDKs
 * @param normalizecity Normalizes village to city level data to city
 * @param optional nil or *ReverseOpts - Optional Parameters:
 * @param "Addressdetails" (optional.Int32) -  Include a breakdown of the address into elements. Defaults to 1.
 * @param "AcceptLanguage" (optional.String) -  Preferred language order for showing search results, overrides the value specified in the Accept-Language HTTP header. Defaults to en. To use native language for the response when available, use accept-language=native
 * @param "Namedetails" (optional.Int32) -  Include a list of alternative names in the results. These may include language variants, references, operator and brand.
 * @param "Extratags" (optional.Int32) -  Include additional information in the result if available, e.g. wikipedia link, opening hours.
 * @param "Statecode" (optional.Int32) -  Adds state or province code when available to the statecode key inside the address element. Currently supported for addresses in the USA, Canada and Australia. Defaults to 0
 * @param "Showdistance" (optional.Int32) -  Returns the straight line distance (meters) between the input location and the result's location. Value is set in the distance key of the response. Defaults to 0 [0,1]
 * @param "Postaladdress" (optional.Int32) -  Returns address inside the postaladdress key, that is specifically formatted for each country. Currently supported for addresses in Germany. Defaults to 0 [0,1]
@return Location
*/
func (a *ReverseApiService) Reverse(ctx _context.Context, lat float32, lon float32, format string, normalizecity int32, localVarOptionals *ReverseOpts) (Location, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Location
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reverse.php"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
	if localVarOptionals != nil && localVarOptionals.Statecode.IsSet() {
		localVarQueryParams.Add("statecode", parameterToString(localVarOptionals.Statecode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Showdistance.IsSet() {
		localVarQueryParams.Add("showdistance", parameterToString(localVarOptionals.Showdistance.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Postaladdress.IsSet() {
		localVarQueryParams.Add("postaladdress", parameterToString(localVarOptionals.Postaladdress.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Location
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
