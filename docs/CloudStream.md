# CloudStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**CloudConfig**](CloudConfig.md) | Config is the stream&#39;s configuration. | [optional] 
**Created** | Pointer to **time.Time** | Created is when the stream was created. | [optional] 
**Name** | Pointer to **string** | Name is the stream name within the org. | [optional] 
**State** | Pointer to [**CloudState**](CloudState.md) | State is the stream&#39;s current state. | [optional] 

## Methods

### NewCloudStream

`func NewCloudStream() *CloudStream`

NewCloudStream instantiates a new CloudStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStreamWithDefaults

`func NewCloudStreamWithDefaults() *CloudStream`

NewCloudStreamWithDefaults instantiates a new CloudStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CloudStream) GetConfig() CloudConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CloudStream) GetConfigOk() (*CloudConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CloudStream) SetConfig(v CloudConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CloudStream) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *CloudStream) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudStream) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudStream) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudStream) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *CloudStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudStream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudStream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *CloudStream) GetState() CloudState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudStream) GetStateOk() (*CloudState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudStream) SetState(v CloudState)`

SetState sets State field to given value.

### HasState

`func (o *CloudStream) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


