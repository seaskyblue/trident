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

// NewExportRuleCollectionGetParams creates a new ExportRuleCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleCollectionGetParams() *ExportRuleCollectionGetParams {
	return &ExportRuleCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleCollectionGetParamsWithTimeout creates a new ExportRuleCollectionGetParams object
// with the ability to set a timeout on a request.
func NewExportRuleCollectionGetParamsWithTimeout(timeout time.Duration) *ExportRuleCollectionGetParams {
	return &ExportRuleCollectionGetParams{
		timeout: timeout,
	}
}

// NewExportRuleCollectionGetParamsWithContext creates a new ExportRuleCollectionGetParams object
// with the ability to set a context for a request.
func NewExportRuleCollectionGetParamsWithContext(ctx context.Context) *ExportRuleCollectionGetParams {
	return &ExportRuleCollectionGetParams{
		Context: ctx,
	}
}

// NewExportRuleCollectionGetParamsWithHTTPClient creates a new ExportRuleCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleCollectionGetParamsWithHTTPClient(client *http.Client) *ExportRuleCollectionGetParams {
	return &ExportRuleCollectionGetParams{
		HTTPClient: client,
	}
}

/* ExportRuleCollectionGetParams contains all the parameters to send to the API endpoint
   for the export rule collection get operation.

   Typically these are written to a http.Request.
*/
type ExportRuleCollectionGetParams struct {

	/* AllowDeviceCreation.

	   Filter by allow_device_creation
	*/
	AllowDeviceCreationQueryParameter *bool

	/* AllowSuid.

	   Filter by allow_suid
	*/
	AllowSUIDQueryParameter *bool

	/* AnonymousUser.

	   Filter by anonymous_user
	*/
	AnonymousUserQueryParameter *string

	/* ChownMode.

	   Filter by chown_mode
	*/
	ChownModeQueryParameter *string

	/* ClientsMatch.

	   Filter by clients.match
	*/
	ClientsMatchQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NtfsUnixSecurity.

	   Filter by ntfs_unix_security
	*/
	NtfsUnixSecurityQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyIDPathParameter int64

	/* Protocols.

	   Filter by protocols
	*/
	ProtocolsQueryParameter *string

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

	/* RoRule.

	   Filter by ro_rule
	*/
	RoRuleQueryParameter *string

	/* RwRule.

	   Filter by rw_rule
	*/
	RwRuleQueryParameter *string

	/* Superuser.

	   Filter by superuser
	*/
	SuperuserQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleCollectionGetParams) WithDefaults() *ExportRuleCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := ExportRuleCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithTimeout(timeout time.Duration) *ExportRuleCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithContext(ctx context.Context) *ExportRuleCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithHTTPClient(client *http.Client) *ExportRuleCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeviceCreationQueryParameter adds the allowDeviceCreation to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithAllowDeviceCreationQueryParameter(allowDeviceCreation *bool) *ExportRuleCollectionGetParams {
	o.SetAllowDeviceCreationQueryParameter(allowDeviceCreation)
	return o
}

// SetAllowDeviceCreationQueryParameter adds the allowDeviceCreation to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetAllowDeviceCreationQueryParameter(allowDeviceCreation *bool) {
	o.AllowDeviceCreationQueryParameter = allowDeviceCreation
}

// WithAllowSUIDQueryParameter adds the allowSuid to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithAllowSUIDQueryParameter(allowSuid *bool) *ExportRuleCollectionGetParams {
	o.SetAllowSUIDQueryParameter(allowSuid)
	return o
}

// SetAllowSUIDQueryParameter adds the allowSuid to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetAllowSUIDQueryParameter(allowSuid *bool) {
	o.AllowSUIDQueryParameter = allowSuid
}

// WithAnonymousUserQueryParameter adds the anonymousUser to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithAnonymousUserQueryParameter(anonymousUser *string) *ExportRuleCollectionGetParams {
	o.SetAnonymousUserQueryParameter(anonymousUser)
	return o
}

// SetAnonymousUserQueryParameter adds the anonymousUser to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetAnonymousUserQueryParameter(anonymousUser *string) {
	o.AnonymousUserQueryParameter = anonymousUser
}

// WithChownModeQueryParameter adds the chownMode to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithChownModeQueryParameter(chownMode *string) *ExportRuleCollectionGetParams {
	o.SetChownModeQueryParameter(chownMode)
	return o
}

// SetChownModeQueryParameter adds the chownMode to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetChownModeQueryParameter(chownMode *string) {
	o.ChownModeQueryParameter = chownMode
}

// WithClientsMatchQueryParameter adds the clientsMatch to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithClientsMatchQueryParameter(clientsMatch *string) *ExportRuleCollectionGetParams {
	o.SetClientsMatchQueryParameter(clientsMatch)
	return o
}

// SetClientsMatchQueryParameter adds the clientsMatch to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetClientsMatchQueryParameter(clientsMatch *string) {
	o.ClientsMatchQueryParameter = clientsMatch
}

// WithFieldsQueryParameter adds the fields to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithFieldsQueryParameter(fields []string) *ExportRuleCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexQueryParameter adds the index to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithIndexQueryParameter(index *int64) *ExportRuleCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *ExportRuleCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNtfsUnixSecurityQueryParameter adds the ntfsUnixSecurity to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithNtfsUnixSecurityQueryParameter(ntfsUnixSecurity *string) *ExportRuleCollectionGetParams {
	o.SetNtfsUnixSecurityQueryParameter(ntfsUnixSecurity)
	return o
}

// SetNtfsUnixSecurityQueryParameter adds the ntfsUnixSecurity to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetNtfsUnixSecurityQueryParameter(ntfsUnixSecurity *string) {
	o.NtfsUnixSecurityQueryParameter = ntfsUnixSecurity
}

// WithOrderByQueryParameter adds the orderBy to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *ExportRuleCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPolicyIDPathParameter adds the policyID to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithPolicyIDPathParameter(policyID int64) *ExportRuleCollectionGetParams {
	o.SetPolicyIDPathParameter(policyID)
	return o
}

// SetPolicyIDPathParameter adds the policyId to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetPolicyIDPathParameter(policyID int64) {
	o.PolicyIDPathParameter = policyID
}

// WithProtocolsQueryParameter adds the protocols to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithProtocolsQueryParameter(protocols *string) *ExportRuleCollectionGetParams {
	o.SetProtocolsQueryParameter(protocols)
	return o
}

// SetProtocolsQueryParameter adds the protocols to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetProtocolsQueryParameter(protocols *string) {
	o.ProtocolsQueryParameter = protocols
}

// WithReturnRecordsQueryParameter adds the returnRecords to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ExportRuleCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *ExportRuleCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithRoRuleQueryParameter adds the roRule to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithRoRuleQueryParameter(roRule *string) *ExportRuleCollectionGetParams {
	o.SetRoRuleQueryParameter(roRule)
	return o
}

// SetRoRuleQueryParameter adds the roRule to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetRoRuleQueryParameter(roRule *string) {
	o.RoRuleQueryParameter = roRule
}

// WithRwRuleQueryParameter adds the rwRule to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithRwRuleQueryParameter(rwRule *string) *ExportRuleCollectionGetParams {
	o.SetRwRuleQueryParameter(rwRule)
	return o
}

// SetRwRuleQueryParameter adds the rwRule to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetRwRuleQueryParameter(rwRule *string) {
	o.RwRuleQueryParameter = rwRule
}

// WithSuperuserQueryParameter adds the superuser to the export rule collection get params
func (o *ExportRuleCollectionGetParams) WithSuperuserQueryParameter(superuser *string) *ExportRuleCollectionGetParams {
	o.SetSuperuserQueryParameter(superuser)
	return o
}

// SetSuperuserQueryParameter adds the superuser to the export rule collection get params
func (o *ExportRuleCollectionGetParams) SetSuperuserQueryParameter(superuser *string) {
	o.SuperuserQueryParameter = superuser
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeviceCreationQueryParameter != nil {

		// query param allow_device_creation
		var qrAllowDeviceCreation bool

		if o.AllowDeviceCreationQueryParameter != nil {
			qrAllowDeviceCreation = *o.AllowDeviceCreationQueryParameter
		}
		qAllowDeviceCreation := swag.FormatBool(qrAllowDeviceCreation)
		if qAllowDeviceCreation != "" {

			if err := r.SetQueryParam("allow_device_creation", qAllowDeviceCreation); err != nil {
				return err
			}
		}
	}

	if o.AllowSUIDQueryParameter != nil {

		// query param allow_suid
		var qrAllowSuid bool

		if o.AllowSUIDQueryParameter != nil {
			qrAllowSuid = *o.AllowSUIDQueryParameter
		}
		qAllowSuid := swag.FormatBool(qrAllowSuid)
		if qAllowSuid != "" {

			if err := r.SetQueryParam("allow_suid", qAllowSuid); err != nil {
				return err
			}
		}
	}

	if o.AnonymousUserQueryParameter != nil {

		// query param anonymous_user
		var qrAnonymousUser string

		if o.AnonymousUserQueryParameter != nil {
			qrAnonymousUser = *o.AnonymousUserQueryParameter
		}
		qAnonymousUser := qrAnonymousUser
		if qAnonymousUser != "" {

			if err := r.SetQueryParam("anonymous_user", qAnonymousUser); err != nil {
				return err
			}
		}
	}

	if o.ChownModeQueryParameter != nil {

		// query param chown_mode
		var qrChownMode string

		if o.ChownModeQueryParameter != nil {
			qrChownMode = *o.ChownModeQueryParameter
		}
		qChownMode := qrChownMode
		if qChownMode != "" {

			if err := r.SetQueryParam("chown_mode", qChownMode); err != nil {
				return err
			}
		}
	}

	if o.ClientsMatchQueryParameter != nil {

		// query param clients.match
		var qrClientsMatch string

		if o.ClientsMatchQueryParameter != nil {
			qrClientsMatch = *o.ClientsMatchQueryParameter
		}
		qClientsMatch := qrClientsMatch
		if qClientsMatch != "" {

			if err := r.SetQueryParam("clients.match", qClientsMatch); err != nil {
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

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
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

	if o.NtfsUnixSecurityQueryParameter != nil {

		// query param ntfs_unix_security
		var qrNtfsUnixSecurity string

		if o.NtfsUnixSecurityQueryParameter != nil {
			qrNtfsUnixSecurity = *o.NtfsUnixSecurityQueryParameter
		}
		qNtfsUnixSecurity := qrNtfsUnixSecurity
		if qNtfsUnixSecurity != "" {

			if err := r.SetQueryParam("ntfs_unix_security", qNtfsUnixSecurity); err != nil {
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

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyIDPathParameter)); err != nil {
		return err
	}

	if o.ProtocolsQueryParameter != nil {

		// query param protocols
		var qrProtocols string

		if o.ProtocolsQueryParameter != nil {
			qrProtocols = *o.ProtocolsQueryParameter
		}
		qProtocols := qrProtocols
		if qProtocols != "" {

			if err := r.SetQueryParam("protocols", qProtocols); err != nil {
				return err
			}
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

	if o.RoRuleQueryParameter != nil {

		// query param ro_rule
		var qrRoRule string

		if o.RoRuleQueryParameter != nil {
			qrRoRule = *o.RoRuleQueryParameter
		}
		qRoRule := qrRoRule
		if qRoRule != "" {

			if err := r.SetQueryParam("ro_rule", qRoRule); err != nil {
				return err
			}
		}
	}

	if o.RwRuleQueryParameter != nil {

		// query param rw_rule
		var qrRwRule string

		if o.RwRuleQueryParameter != nil {
			qrRwRule = *o.RwRuleQueryParameter
		}
		qRwRule := qrRwRule
		if qRwRule != "" {

			if err := r.SetQueryParam("rw_rule", qRwRule); err != nil {
				return err
			}
		}
	}

	if o.SuperuserQueryParameter != nil {

		// query param superuser
		var qrSuperuser string

		if o.SuperuserQueryParameter != nil {
			qrSuperuser = *o.SuperuserQueryParameter
		}
		qSuperuser := qrSuperuser
		if qSuperuser != "" {

			if err := r.SetQueryParam("superuser", qSuperuser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamExportRuleCollectionGet binds the parameter fields
func (o *ExportRuleCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamExportRuleCollectionGet binds the parameter order_by
func (o *ExportRuleCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
