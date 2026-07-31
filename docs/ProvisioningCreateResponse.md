# ProvisioningCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Server-generated resource id. | 
**Kind** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** | Resource status (e.g. \&quot;ready\&quot;, or \&quot;provisioning\&quot; for a dedicated instance still coming up). | 
**Host** | **string** | Customer-facing host (public gateway host, or dedicated instance service). | 
**Port** | **int32** |  | 
**Username** | Pointer to **string** | Present only for secretful kinds. | [optional] 
**Database** | **string** |  | 
**ConnectionString** | **string** | Public, routable DSN (internal admin host is remapped out). | 
**Password** | Pointer to **string** | Present only for secretful kinds; returned ONCE. | [optional] 

## Methods

### NewProvisioningCreateResponse

`func NewProvisioningCreateResponse(id string, kind string, name string, status string, host string, port int32, database string, connectionString string, ) *ProvisioningCreateResponse`

NewProvisioningCreateResponse instantiates a new ProvisioningCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningCreateResponseWithDefaults

`func NewProvisioningCreateResponseWithDefaults() *ProvisioningCreateResponse`

NewProvisioningCreateResponseWithDefaults instantiates a new ProvisioningCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningCreateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ProvisioningCreateResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProvisioningCreateResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProvisioningCreateResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ProvisioningCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningCreateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ProvisioningCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHost

`func (o *ProvisioningCreateResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProvisioningCreateResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProvisioningCreateResponse) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ProvisioningCreateResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProvisioningCreateResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProvisioningCreateResponse) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *ProvisioningCreateResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProvisioningCreateResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProvisioningCreateResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProvisioningCreateResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDatabase

`func (o *ProvisioningCreateResponse) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ProvisioningCreateResponse) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ProvisioningCreateResponse) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetConnectionString

`func (o *ProvisioningCreateResponse) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *ProvisioningCreateResponse) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *ProvisioningCreateResponse) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetPassword

`func (o *ProvisioningCreateResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProvisioningCreateResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProvisioningCreateResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProvisioningCreateResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


