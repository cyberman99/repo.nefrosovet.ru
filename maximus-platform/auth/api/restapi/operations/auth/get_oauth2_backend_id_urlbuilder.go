// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetOauth2BackendIDURL generates an URL for the get oauth2 backend ID operation
type GetOauth2BackendIDURL struct {
	BackendID string

	RedirectURI string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetOauth2BackendIDURL) WithBasePath(bp string) *GetOauth2BackendIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetOauth2BackendIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetOauth2BackendIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/oauth2/{backendID}"

	backendID := o.BackendID
	if backendID != "" {
		_path = strings.Replace(_path, "{backendID}", backendID, -1)
	} else {
		return nil, errors.New("BackendID is required on GetOauth2BackendIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	redirectURI := o.RedirectURI
	if redirectURI != "" {
		qs.Set("redirectURI", redirectURI)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetOauth2BackendIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetOauth2BackendIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetOauth2BackendIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetOauth2BackendIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetOauth2BackendIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetOauth2BackendIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
