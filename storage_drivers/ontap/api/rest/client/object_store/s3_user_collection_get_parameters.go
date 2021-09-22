// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3UserCollectionGetParams creates a new S3UserCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3UserCollectionGetParams() *S3UserCollectionGetParams {
	return &S3UserCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3UserCollectionGetParamsWithTimeout creates a new S3UserCollectionGetParams object
// with the ability to set a timeout on a request.
func NewS3UserCollectionGetParamsWithTimeout(timeout time.Duration) *S3UserCollectionGetParams {
	return &S3UserCollectionGetParams{
		timeout: timeout,
	}
}

// NewS3UserCollectionGetParamsWithContext creates a new S3UserCollectionGetParams object
// with the ability to set a context for a request.
func NewS3UserCollectionGetParamsWithContext(ctx context.Context) *S3UserCollectionGetParams {
	return &S3UserCollectionGetParams{
		Context: ctx,
	}
}

// NewS3UserCollectionGetParamsWithHTTPClient creates a new S3UserCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3UserCollectionGetParamsWithHTTPClient(client *http.Client) *S3UserCollectionGetParams {
	return &S3UserCollectionGetParams{
		HTTPClient: client,
	}
}

/* S3UserCollectionGetParams contains all the parameters to send to the API endpoint
   for the s3 user collection get operation.

   Typically these are written to a http.Request.
*/
type S3UserCollectionGetParams struct {

	/* AccessKey.

	   Filter by access_key
	*/
	AccessKeyQueryParameter *string

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 user collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3UserCollectionGetParams) WithDefaults() *S3UserCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 user collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3UserCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := S3UserCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithTimeout(timeout time.Duration) *S3UserCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithContext(ctx context.Context) *S3UserCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithHTTPClient(client *http.Client) *S3UserCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyQueryParameter adds the accessKey to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithAccessKeyQueryParameter(accessKey *string) *S3UserCollectionGetParams {
	o.SetAccessKeyQueryParameter(accessKey)
	return o
}

// SetAccessKeyQueryParameter adds the accessKey to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetAccessKeyQueryParameter(accessKey *string) {
	o.AccessKeyQueryParameter = accessKey
}

// WithCommentQueryParameter adds the comment to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithCommentQueryParameter(comment *string) *S3UserCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithFieldsQueryParameter adds the fields to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithFieldsQueryParameter(fields []string) *S3UserCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *S3UserCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithNameQueryParameter(name *string) *S3UserCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *S3UserCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *S3UserCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *S3UserCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *S3UserCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *S3UserCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithSVMUUIDPathParameter adds the svmUUID to the s3 user collection get params
func (o *S3UserCollectionGetParams) WithSVMUUIDPathParameter(svmUUID string) *S3UserCollectionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the s3 user collection get params
func (o *S3UserCollectionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3UserCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyQueryParameter != nil {

		// query param access_key
		var qrAccessKey string

		if o.AccessKeyQueryParameter != nil {
			qrAccessKey = *o.AccessKeyQueryParameter
		}
		qAccessKey := qrAccessKey
		if qAccessKey != "" {

			if err := r.SetQueryParam("access_key", qAccessKey); err != nil {
				return err
			}
		}
	}

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
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

// bindParamS3UserCollectionGet binds the parameter fields
func (o *S3UserCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamS3UserCollectionGet binds the parameter order_by
func (o *S3UserCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}