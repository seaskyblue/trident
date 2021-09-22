// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CloudTargetCollectionGet(params *CloudTargetCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetCollectionGetOK, error)

	CloudTargetCreate(params *CloudTargetCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetCreateAccepted, error)

	CloudTargetDelete(params *CloudTargetDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetDeleteAccepted, error)

	CloudTargetGet(params *CloudTargetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetGetOK, error)

	CloudTargetModify(params *CloudTargetModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetModifyAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CloudTargetCollectionGet Retrieves the collection of cloud targets in the cluster.
### Related ONTAP commands
* `storage aggregate object-store config show`

*/
func (a *Client) CloudTargetCollectionGet(params *CloudTargetCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetCollectionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudTargetCollectionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud_target_collection_get",
		Method:             "GET",
		PathPattern:        "/cloud/targets",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudTargetCollectionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CloudTargetCollectionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CloudTargetCollectionGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CloudTargetCreate Creates a cloud target.
### Required properties
* `name` - Name for the cloud target.
* `owner` - Owner of the target: _fabricpool_, _snapmirror_.
* `provider_type` - Type of cloud provider: _AWS_S3_, _Azure_Cloud_, _SGWS_, _IBM_COS_, _AliCloud_, _GoogleCloud_, _ONTAP_S3_.
* `server` - Fully qualified domain name of the object store server. Required when `provider_type` is one of the following: _SGWS_, _IBM_COS_, _AliCloud_.
* `container` - Data bucket/container name.
* `access_key` - Access key ID if `provider_type` is not _Azure_Cloud_ and `authentication_type` is _key_.
* `secret_password` - Secret access key if `provider_type` is not _Azure_Cloud_ and `authentication_type` is _key_.
* `azure_account` - Azure account if `provider_type` is _Azure_Cloud_.
* `azure_private_key` - Azure access key if `provider_type` is _Azure_Cloud_.
* `cap_url` - Full URL of the request to a CAP server for retrieving temporary credentials if `authentication_type` is _cap_.
* `svm.name` or `svm.uuid` - Name or UUID of SVM if `owner` is _snapmirror_.
* `snapmirror_use` - Use of the cloud target if `owner` is _snapmirror_: data, metadata.
### Recommended optional properties
* `authentication_type` - Authentication used to access the target: _key_, _cap_, _ec2_iam_, _gcp_sa_, _azure_msi_.
* `ssl_enabled` - SSL/HTTPS enabled or disabled.
* `port` - Port number of the object store that ONTAP uses when establishing a connection.
* `ipspace` - IPspace to use in order to reach the cloud target.
* `use_http_proxy` - Use the HTTP proxy when connecting to the object store server.
### Default property values
* `authentication_type`
  - _ec2_iam_ - if running in Cloud Volumes ONTAP in AWS
  - _gcp_sa_ - if running in Cloud Volumes ONTAP in GCP
  - _azure_msi_ - if running in Cloud Volumes ONTAP in Azure
  - _key_  - in all other cases.
* `server`
  - _s3.amazonaws.com_ - if `provider_type` is _AWS_S3_
  - _blob.core.windows.net_ - if `provider_type` is _Azure_Cloud_
  - _storage.googleapis.com_ - if `provider_type` is _GoogleCloud_
* `ssl_enabled` - _true_
* `port`
  - _443_ if `ssl_enabled` is _true_ and `provider_type` is not _SGWS_
  - _8082_ if `ssl_enabled` is _true_ and `provider_type` is _SGWS_
  - _80_ if `ssl_enabled` is _false_ and `provider_type` is not _SGWS_
  - _8084_ if `ssl_enabled` is _false_ and `provider_type` is _SGWS_
* `ipspace` - _Default_
* `certificate_validation_enabled` - _true_
* `ignore_warnings` - _false_
* `check_only` - _false_
* `use_http_proxy` - _false_
* `server_side_encryption`
  - _none_ - if `provider_type` is _ONTAP_S3_
  - _sse_s3_ - if `provider_type` is not _ONTAP_S3_
* `url_style`
  - _path_style_ - if `provider_type` is neither _AWS_S3_ nor _AliCloud_
  - _virtual_hosted_style_ - if `provider_type` is either _AWS_S3 or _AliCloud__
### Related ONTAP commands
* `storage aggregate object-store config create`

*/
func (a *Client) CloudTargetCreate(params *CloudTargetCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetCreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudTargetCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud_target_create",
		Method:             "POST",
		PathPattern:        "/cloud/targets",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudTargetCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CloudTargetCreateAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CloudTargetCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CloudTargetDelete Deletes the cloud target specified by the UUID. This request starts a job and returns a link to that job.
### Related ONTAP commands
* `storage aggregate object-store config delete`

*/
func (a *Client) CloudTargetDelete(params *CloudTargetDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudTargetDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud_target_delete",
		Method:             "DELETE",
		PathPattern:        "/cloud/targets/{uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudTargetDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CloudTargetDeleteAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CloudTargetDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CloudTargetGet Retrieves the cloud target specified by the UUID.
### Related ONTAP commands
* `storage aggregate object-store config show`

*/
func (a *Client) CloudTargetGet(params *CloudTargetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudTargetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud_target_get",
		Method:             "GET",
		PathPattern:        "/cloud/targets/{uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudTargetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CloudTargetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CloudTargetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CloudTargetModify Updates the cloud target specified by the UUID with the fields in the body. This request starts a job and returns a link to that job.
### Related ONTAP commands
* `storage aggregate object-store config modify`

*/
func (a *Client) CloudTargetModify(params *CloudTargetModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloudTargetModifyAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloudTargetModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cloud_target_modify",
		Method:             "PATCH",
		PathPattern:        "/cloud/targets/{uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloudTargetModifyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CloudTargetModifyAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CloudTargetModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}