# ProvisionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | Pointer to **string** | ConnectionString is the ready-to-use DSN, credential included. RETURNED HERE ONCE: no read beside this one carries it, so a caller that does not keep it must provision again. | [optional] 
**Database** | Pointer to **string** | Database is the logical database, collection, index or bucket this resource resolves to on its backend. It is derived from Name under an org-namespacing hash, so it is not Name and two orgs cannot land on one. | [optional] 
**Host** | Pointer to **string** | Host is the address that routes to this resource — a dedicated instance&#39;s own in-cluster Service, or the public gateway for a shared one. Never the internal admin address of a shared backend. | [optional] 
**Id** | Pointer to **string** | ID is the resource&#39;s server-minted handle, \&quot;rs_\&quot;-prefixed. The caller does not choose it, and it is what every read and the delete address. | [optional] 
**Kind** | Pointer to **string** | Kind is the product provisioned: sql, vector, datastore, kv, search, s3 or docdb. It is the route that was called, not a body field. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique slug the caller asked for, lower-cased. Every physical name on the backend derives from it. | [optional] 
**Password** | Pointer to **string** | Password is the minted credential, in plaintext, for the kinds that have one. RETURNED HERE ONCE — where KMS is configured it is sealed there and only a reference is persisted; where it is not, it is stored nowhere at all. It is never held in plaintext on either side. | [optional] 
**Port** | Pointer to **int64** | Port is the port a client connects to on Host. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ready\&quot;, or \&quot;provisioning\&quot; while a dedicated instance is still being materialized by the operator. A shared-backend create is \&quot;ready\&quot; here; a dedicated one answers 201 still launching, and reaches ready only when a later read reconciles it against the operator&#39;s live CR — never fabricated. | [optional] 
**Username** | Pointer to **string** | Username is the credential&#39;s user, for the kinds that mint one per resource. Absent for a kind whose backend authenticates with a shared, out-of-band key. | [optional] 

## Methods

### NewProvisionResult

`func NewProvisionResult() *ProvisionResult`

NewProvisionResult instantiates a new ProvisionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionResultWithDefaults

`func NewProvisionResultWithDefaults() *ProvisionResult`

NewProvisionResultWithDefaults instantiates a new ProvisionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *ProvisionResult) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *ProvisionResult) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *ProvisionResult) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *ProvisionResult) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### GetDatabase

`func (o *ProvisionResult) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ProvisionResult) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ProvisionResult) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *ProvisionResult) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetHost

`func (o *ProvisionResult) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProvisionResult) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProvisionResult) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProvisionResult) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ProvisionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisionResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ProvisionResult) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProvisionResult) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProvisionResult) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProvisionResult) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ProvisionResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *ProvisionResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProvisionResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProvisionResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProvisionResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *ProvisionResult) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProvisionResult) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProvisionResult) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ProvisionResult) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisionResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisionResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *ProvisionResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProvisionResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProvisionResult) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProvisionResult) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


