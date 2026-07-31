# ProvisioningGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Kind** | **string** |  | 
**Status** | **string** |  | 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Username** | Pointer to **string** | Present only for secretful kinds. | [optional] 
**Database** | **string** |  | 

## Methods

### NewProvisioningGetResponse

`func NewProvisioningGetResponse(id string, name string, kind string, status string, host string, port int32, database string, ) *ProvisioningGetResponse`

NewProvisioningGetResponse instantiates a new ProvisioningGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningGetResponseWithDefaults

`func NewProvisioningGetResponseWithDefaults() *ProvisioningGetResponse`

NewProvisioningGetResponseWithDefaults instantiates a new ProvisioningGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningGetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningGetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningGetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProvisioningGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningGetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *ProvisioningGetResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProvisioningGetResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProvisioningGetResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *ProvisioningGetResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningGetResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningGetResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHost

`func (o *ProvisioningGetResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProvisioningGetResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProvisioningGetResponse) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ProvisioningGetResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProvisioningGetResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProvisioningGetResponse) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *ProvisioningGetResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProvisioningGetResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProvisioningGetResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProvisioningGetResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDatabase

`func (o *ProvisioningGetResponse) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ProvisioningGetResponse) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ProvisioningGetResponse) SetDatabase(v string)`

SetDatabase sets Database field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


