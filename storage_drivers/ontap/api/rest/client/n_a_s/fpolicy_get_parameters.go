// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewFpolicyGetParams creates a new FpolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyGetParams() *FpolicyGetParams {
	return &FpolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyGetParamsWithTimeout creates a new FpolicyGetParams object
// with the ability to set a timeout on a request.
func NewFpolicyGetParamsWithTimeout(timeout time.Duration) *FpolicyGetParams {
	return &FpolicyGetParams{
		timeout: timeout,
	}
}

// NewFpolicyGetParamsWithContext creates a new FpolicyGetParams object
// with the ability to set a context for a request.
func NewFpolicyGetParamsWithContext(ctx context.Context) *FpolicyGetParams {
	return &FpolicyGetParams{
		Context: ctx,
	}
}

// NewFpolicyGetParamsWithHTTPClient creates a new FpolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyGetParamsWithHTTPClient(client *http.Client) *FpolicyGetParams {
	return &FpolicyGetParams{
		HTTPClient: client,
	}
}

/* FpolicyGetParams contains all the parameters to send to the API endpoint
   for the fpolicy get operation.

   Typically these are written to a http.Request.
*/
type FpolicyGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyGetParams) WithDefaults() *FpolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fpolicy get params
func (o *FpolicyGetParams) WithTimeout(timeout time.Duration) *FpolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy get params
func (o *FpolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy get params
func (o *FpolicyGetParams) WithContext(ctx context.Context) *FpolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy get params
func (o *FpolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy get params
func (o *FpolicyGetParams) WithHTTPClient(client *http.Client) *FpolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy get params
func (o *FpolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the fpolicy get params
func (o *FpolicyGetParams) WithFieldsQueryParameter(fields []string) *FpolicyGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the fpolicy get params
func (o *FpolicyGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the fpolicy get params
func (o *FpolicyGetParams) WithSVMUUIDPathParameter(svmUUID string) *FpolicyGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the fpolicy get params
func (o *FpolicyGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFpolicyGet binds the parameter fields
func (o *FpolicyGetParams) bindParamFields(formats strfmt.Registry) []string {
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