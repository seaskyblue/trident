// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewPortsetCreateParams creates a new PortsetCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortsetCreateParams() *PortsetCreateParams {
	return &PortsetCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortsetCreateParamsWithTimeout creates a new PortsetCreateParams object
// with the ability to set a timeout on a request.
func NewPortsetCreateParamsWithTimeout(timeout time.Duration) *PortsetCreateParams {
	return &PortsetCreateParams{
		timeout: timeout,
	}
}

// NewPortsetCreateParamsWithContext creates a new PortsetCreateParams object
// with the ability to set a context for a request.
func NewPortsetCreateParamsWithContext(ctx context.Context) *PortsetCreateParams {
	return &PortsetCreateParams{
		Context: ctx,
	}
}

// NewPortsetCreateParamsWithHTTPClient creates a new PortsetCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortsetCreateParamsWithHTTPClient(client *http.Client) *PortsetCreateParams {
	return &PortsetCreateParams{
		HTTPClient: client,
	}
}

/* PortsetCreateParams contains all the parameters to send to the API endpoint
   for the portset create operation.

   Typically these are written to a http.Request.
*/
type PortsetCreateParams struct {

	/* Info.

	   The property values for the new portset.

	*/
	Info *models.Portset

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the portset create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetCreateParams) WithDefaults() *PortsetCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the portset create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := PortsetCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the portset create params
func (o *PortsetCreateParams) WithTimeout(timeout time.Duration) *PortsetCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portset create params
func (o *PortsetCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portset create params
func (o *PortsetCreateParams) WithContext(ctx context.Context) *PortsetCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portset create params
func (o *PortsetCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portset create params
func (o *PortsetCreateParams) WithHTTPClient(client *http.Client) *PortsetCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portset create params
func (o *PortsetCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the portset create params
func (o *PortsetCreateParams) WithInfo(info *models.Portset) *PortsetCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the portset create params
func (o *PortsetCreateParams) SetInfo(info *models.Portset) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the portset create params
func (o *PortsetCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *PortsetCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the portset create params
func (o *PortsetCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *PortsetCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
