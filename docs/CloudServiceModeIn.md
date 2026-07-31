# CloudServiceModeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service is the slug to flip, taken from the path. | [optional] 
**WaitlistMode** | Pointer to **bool** | WaitlistMode is the new mode: true gates the service behind the waitlist, false opens it. This is the launch lever. | [optional] 

## Methods

### NewCloudServiceModeIn

`func NewCloudServiceModeIn() *CloudServiceModeIn`

NewCloudServiceModeIn instantiates a new CloudServiceModeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceModeInWithDefaults

`func NewCloudServiceModeInWithDefaults() *CloudServiceModeIn`

NewCloudServiceModeInWithDefaults instantiates a new CloudServiceModeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudServiceModeIn) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudServiceModeIn) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudServiceModeIn) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudServiceModeIn) HasService() bool`

HasService returns a boolean if a field has been set.

### GetWaitlistMode

`func (o *CloudServiceModeIn) GetWaitlistMode() bool`

GetWaitlistMode returns the WaitlistMode field if non-nil, zero value otherwise.

### GetWaitlistModeOk

`func (o *CloudServiceModeIn) GetWaitlistModeOk() (*bool, bool)`

GetWaitlistModeOk returns a tuple with the WaitlistMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistMode

`func (o *CloudServiceModeIn) SetWaitlistMode(v bool)`

SetWaitlistMode sets WaitlistMode field to given value.

### HasWaitlistMode

`func (o *CloudServiceModeIn) HasWaitlistMode() bool`

HasWaitlistMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


