// Code generated by go-swagger; DO NOT EDIT.

package svm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSvmPeerInstanceGetParams creates a new SvmPeerInstanceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerInstanceGetParams() *SvmPeerInstanceGetParams {
	return &SvmPeerInstanceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerInstanceGetParamsWithTimeout creates a new SvmPeerInstanceGetParams object
// with the ability to set a timeout on a request.
func NewSvmPeerInstanceGetParamsWithTimeout(timeout time.Duration) *SvmPeerInstanceGetParams {
	return &SvmPeerInstanceGetParams{
		timeout: timeout,
	}
}

// NewSvmPeerInstanceGetParamsWithContext creates a new SvmPeerInstanceGetParams object
// with the ability to set a context for a request.
func NewSvmPeerInstanceGetParamsWithContext(ctx context.Context) *SvmPeerInstanceGetParams {
	return &SvmPeerInstanceGetParams{
		Context: ctx,
	}
}

// NewSvmPeerInstanceGetParamsWithHTTPClient creates a new SvmPeerInstanceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerInstanceGetParamsWithHTTPClient(client *http.Client) *SvmPeerInstanceGetParams {
	return &SvmPeerInstanceGetParams{
		HTTPClient: client,
	}
}

/* SvmPeerInstanceGetParams contains all the parameters to send to the API endpoint
   for the svm peer instance get operation.

   Typically these are written to a http.Request.
*/
type SvmPeerInstanceGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   SVM peer relation UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm peer instance get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerInstanceGetParams) WithDefaults() *SvmPeerInstanceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer instance get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerInstanceGetParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SvmPeerInstanceGetParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) WithTimeout(timeout time.Duration) *SvmPeerInstanceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) WithContext(ctx context.Context) *SvmPeerInstanceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) WithHTTPClient(client *http.Client) *SvmPeerInstanceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) WithFieldsQueryParameter(fields []string) *SvmPeerInstanceGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SvmPeerInstanceGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) WithUUIDPathParameter(uuid string) *SvmPeerInstanceGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the svm peer instance get params
func (o *SvmPeerInstanceGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerInstanceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSvmPeerInstanceGet binds the parameter fields
func (o *SvmPeerInstanceGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
