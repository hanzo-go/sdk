# ProvisioningProvisionedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to **string** | Database is the logical database, collection, index or bucket this resource resolves to on its backend. | [optional] 
**Host** | Pointer to **string** | Host is the address that actually routes to this resource — a dedicated instance&#39;s own in-cluster Service, or the public gateway for a shared one. | [optional] 
**Id** | Pointer to **string** | ID is the resource&#39;s server-minted handle, \&quot;rs_\&quot;-prefixed. | [optional] 
**Kind** | Pointer to **string** | Kind is the product: sql, vector, datastore, kv, search, s3 or docdb. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique slug the caller provisioned the resource under. | [optional] 
**Port** | Pointer to **int64** | Port is the port a client connects to on Host. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ready\&quot;, or \&quot;provisioning\&quot; while a dedicated instance is still being materialized. A dedicated resource&#39;s status is reconciled from the operator&#39;s live CR before this is answered, so it is never a stale ready. | [optional] 
**Username** | Pointer to **string** | Username is the credential&#39;s user, for the kinds that mint one per resource. Absent for a kind whose backend authenticates with a shared, out-of-band key. | [optional] 

## Methods

### NewProvisioningProvisionedResource

`func NewProvisioningProvisionedResource() *ProvisioningProvisionedResource`

NewProvisioningProvisionedResource instantiates a new ProvisioningProvisionedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningProvisionedResourceWithDefaults

`func NewProvisioningProvisionedResourceWithDefaults() *ProvisioningProvisionedResource`

NewProvisioningProvisionedResourceWithDefaults instantiates a new ProvisioningProvisionedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *ProvisioningProvisionedResource) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ProvisioningProvisionedResource) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ProvisioningProvisionedResource) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *ProvisioningProvisionedResource) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetHost

`func (o *ProvisioningProvisionedResource) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProvisioningProvisionedResource) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProvisioningProvisionedResource) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProvisioningProvisionedResource) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ProvisioningProvisionedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningProvisionedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningProvisionedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningProvisionedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ProvisioningProvisionedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProvisioningProvisionedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProvisioningProvisionedResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProvisioningProvisionedResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ProvisioningProvisionedResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningProvisionedResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningProvisionedResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningProvisionedResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ProvisioningProvisionedResource) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProvisioningProvisionedResource) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProvisioningProvisionedResource) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ProvisioningProvisionedResource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisioningProvisionedResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningProvisionedResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningProvisionedResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisioningProvisionedResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *ProvisioningProvisionedResource) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProvisioningProvisionedResource) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProvisioningProvisionedResource) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProvisioningProvisionedResource) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


