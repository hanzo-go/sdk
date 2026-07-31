# CloudPushEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** |  | [optional] 
**Pusher** | Pointer to [**CloudPushEventPusher**](CloudPushEventPusher.md) |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**CloudPushEventRepository**](CloudPushEventRepository.md) |  | [optional] 

## Methods

### NewCloudPushEvent

`func NewCloudPushEvent() *CloudPushEvent`

NewCloudPushEvent instantiates a new CloudPushEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPushEventWithDefaults

`func NewCloudPushEventWithDefaults() *CloudPushEvent`

NewCloudPushEventWithDefaults instantiates a new CloudPushEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CloudPushEvent) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CloudPushEvent) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CloudPushEvent) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CloudPushEvent) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetPusher

`func (o *CloudPushEvent) GetPusher() CloudPushEventPusher`

GetPusher returns the Pusher field if non-nil, zero value otherwise.

### GetPusherOk

`func (o *CloudPushEvent) GetPusherOk() (*CloudPushEventPusher, bool)`

GetPusherOk returns a tuple with the Pusher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPusher

`func (o *CloudPushEvent) SetPusher(v CloudPushEventPusher)`

SetPusher sets Pusher field to given value.

### HasPusher

`func (o *CloudPushEvent) HasPusher() bool`

HasPusher returns a boolean if a field has been set.

### GetRef

`func (o *CloudPushEvent) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudPushEvent) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudPushEvent) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudPushEvent) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepository

`func (o *CloudPushEvent) GetRepository() CloudPushEventRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CloudPushEvent) GetRepositoryOk() (*CloudPushEventRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CloudPushEvent) SetRepository(v CloudPushEventRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CloudPushEvent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


