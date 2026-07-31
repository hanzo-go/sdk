# CloudProvisionedSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the resource was provisioned, in unix seconds. | [optional] 
**Host** | Pointer to **string** | Host is the address that actually routes to this resource — a dedicated instance&#39;s own in-cluster Service, or the public gateway for a shared one. Never the internal admin address of a shared backend. | [optional] 
**Id** | Pointer to **string** | ID is the resource&#39;s server-minted handle, \&quot;rs_\&quot;-prefixed. | [optional] 
**Kind** | Pointer to **string** | Kind is the product: sql, vector, datastore, kv, search, s3 or docdb. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique slug the caller provisioned the resource under. | [optional] 
**Port** | Pointer to **int32** | Port is the port a client connects to on Host. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ready\&quot;, or \&quot;provisioning\&quot; while a dedicated instance is still being materialized by the operator. | [optional] 

## Methods

### NewCloudProvisionedSummary

`func NewCloudProvisionedSummary() *CloudProvisionedSummary`

NewCloudProvisionedSummary instantiates a new CloudProvisionedSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProvisionedSummaryWithDefaults

`func NewCloudProvisionedSummaryWithDefaults() *CloudProvisionedSummary`

NewCloudProvisionedSummaryWithDefaults instantiates a new CloudProvisionedSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudProvisionedSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudProvisionedSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudProvisionedSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudProvisionedSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *CloudProvisionedSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudProvisionedSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudProvisionedSummary) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudProvisionedSummary) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *CloudProvisionedSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProvisionedSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProvisionedSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProvisionedSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudProvisionedSummary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudProvisionedSummary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudProvisionedSummary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudProvisionedSummary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CloudProvisionedSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProvisionedSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProvisionedSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProvisionedSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *CloudProvisionedSummary) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudProvisionedSummary) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudProvisionedSummary) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CloudProvisionedSummary) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *CloudProvisionedSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudProvisionedSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudProvisionedSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudProvisionedSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


