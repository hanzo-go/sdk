# ServiceModeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service is the slug to flip, taken from the path. | [optional] 
**WaitlistMode** | Pointer to **bool** | WaitlistMode is the new mode: true gates the service behind the waitlist, false opens it. This is the launch lever. | [optional] 

## Methods

### NewServiceModeIn

`func NewServiceModeIn() *ServiceModeIn`

NewServiceModeIn instantiates a new ServiceModeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceModeInWithDefaults

`func NewServiceModeInWithDefaults() *ServiceModeIn`

NewServiceModeInWithDefaults instantiates a new ServiceModeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceModeIn) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceModeIn) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceModeIn) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceModeIn) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWaitlistMode

`func (o *ServiceModeIn) GetWaitlistMode() bool`

GetWaitlistMode returns the WaitlistMode field if non-nil, zero value otherwise.

### GetWaitlistModeOk

`func (o *ServiceModeIn) GetWaitlistModeOk() (*bool, bool)`

GetWaitlistModeOk returns a tuple with the WaitlistMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistMode

`func (o *ServiceModeIn) SetWaitlistMode(v bool)`

SetWaitlistMode sets WaitlistMode field to given value.

### HasWaitlistMode

`func (o *ServiceModeIn) HasWaitlistMode() bool`

HasWaitlistMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


