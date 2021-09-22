// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewApplicationComponentSnapshotRestoreParams creates a new ApplicationComponentSnapshotRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationComponentSnapshotRestoreParams() *ApplicationComponentSnapshotRestoreParams {
	return &ApplicationComponentSnapshotRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationComponentSnapshotRestoreParamsWithTimeout creates a new ApplicationComponentSnapshotRestoreParams object
// with the ability to set a timeout on a request.
func NewApplicationComponentSnapshotRestoreParamsWithTimeout(timeout time.Duration) *ApplicationComponentSnapshotRestoreParams {
	return &ApplicationComponentSnapshotRestoreParams{
		timeout: timeout,
	}
}

// NewApplicationComponentSnapshotRestoreParamsWithContext creates a new ApplicationComponentSnapshotRestoreParams object
// with the ability to set a context for a request.
func NewApplicationComponentSnapshotRestoreParamsWithContext(ctx context.Context) *ApplicationComponentSnapshotRestoreParams {
	return &ApplicationComponentSnapshotRestoreParams{
		Context: ctx,
	}
}

// NewApplicationComponentSnapshotRestoreParamsWithHTTPClient creates a new ApplicationComponentSnapshotRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationComponentSnapshotRestoreParamsWithHTTPClient(client *http.Client) *ApplicationComponentSnapshotRestoreParams {
	return &ApplicationComponentSnapshotRestoreParams{
		HTTPClient: client,
	}
}

/* ApplicationComponentSnapshotRestoreParams contains all the parameters to send to the API endpoint
   for the application component snapshot restore operation.

   Typically these are written to a http.Request.
*/
type ApplicationComponentSnapshotRestoreParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUIDPathParameter string

	/* ComponentUUID.

	   Application Component UUID
	*/
	ComponentUUIDPathParameter string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   Snapshot copy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application component snapshot restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentSnapshotRestoreParams) WithDefaults() *ApplicationComponentSnapshotRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application component snapshot restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationComponentSnapshotRestoreParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := ApplicationComponentSnapshotRestoreParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithTimeout(timeout time.Duration) *ApplicationComponentSnapshotRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithContext(ctx context.Context) *ApplicationComponentSnapshotRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithHTTPClient(client *http.Client) *ApplicationComponentSnapshotRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUIDPathParameter adds the applicationUUID to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithApplicationUUIDPathParameter(applicationUUID string) *ApplicationComponentSnapshotRestoreParams {
	o.SetApplicationUUIDPathParameter(applicationUUID)
	return o
}

// SetApplicationUUIDPathParameter adds the applicationUuid to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetApplicationUUIDPathParameter(applicationUUID string) {
	o.ApplicationUUIDPathParameter = applicationUUID
}

// WithComponentUUIDPathParameter adds the componentUUID to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithComponentUUIDPathParameter(componentUUID string) *ApplicationComponentSnapshotRestoreParams {
	o.SetComponentUUIDPathParameter(componentUUID)
	return o
}

// SetComponentUUIDPathParameter adds the componentUuid to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetComponentUUIDPathParameter(componentUUID string) {
	o.ComponentUUIDPathParameter = componentUUID
}

// WithReturnRecordsQueryParameter adds the returnRecords to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ApplicationComponentSnapshotRestoreParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *ApplicationComponentSnapshotRestoreParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) WithUUIDPathParameter(uuid string) *ApplicationComponentSnapshotRestoreParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the application component snapshot restore params
func (o *ApplicationComponentSnapshotRestoreParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationComponentSnapshotRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUIDPathParameter); err != nil {
		return err
	}

	// path param component.uuid
	if err := r.SetPathParam("component.uuid", o.ComponentUUIDPathParameter); err != nil {
		return err
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