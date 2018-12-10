// +build mock

package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** TYPES **********/

type (
	// Request represents the http request client
	Request struct {
		authKey, clientID, clientSecret, fingerprint, ipAddress string
	}
)

/********** GLOBAL VARIABLES **********/
var goreq = gorequest.New()

/********** METHODS **********/

func (req *Request) updateRequest(clientID, clientSecret, fingerprint, ipAddress string, authKey ...string) Request {
	var aKey string

	if len(authKey) > 0 {
		aKey = authKey[0]
	}

	return Request{
		authKey:      aKey,
		clientID:     clientID,
		clientSecret: clientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}
}

/********** REQUEST **********/

// Get performs a GET request
func (req *Request) Get(url string, queryParams []string, result interface{}) ([]byte, error) {
	var params string
	if len(queryParams) > 0 {
		params = queryParams[0]
	}

	res, body, errs := goreq.
		Get(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(params).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var params string
	if len(queryParams) > 0 {
		params = queryParams[0]
	}

	res, body, errs := goreq.
		Post(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var params string
	if len(queryParams) > 0 {
		params = queryParams[0]
	}

	res, body, errs := goreq.
		Patch(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string, result interface{}) ([]byte, error) {
	res, body, errs := goreq.
		Delete(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}