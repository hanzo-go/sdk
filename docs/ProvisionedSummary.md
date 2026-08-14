# ProvisionedSummary

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

### NewProvisionedSummary

`func NewProvisionedSummary() *ProvisionedSummary`

NewProvisionedSummary instantiates a new ProvisionedSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedSummaryWithDefaults

`func NewProvisionedSummaryWithDefaults() *ProvisionedSummary`

NewProvisionedSummaryWithDefaults instantiates a new ProvisionedSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ProvisionedSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProvisionedSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProvisionedSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProvisionedSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *ProvisionedSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProvisionedSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProvisionedSummary) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProvisionedSummary) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ProvisionedSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionedSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionedSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisionedSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ProvisionedSummary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProvisionedSummary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProvisionedSummary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProvisionedSummary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ProvisionedSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionedSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionedSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionedSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ProvisionedSummary) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProvisionedSummary) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProvisionedSummary) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ProvisionedSummary) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisionedSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisionedSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisionedSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisionedSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


