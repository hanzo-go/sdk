# ActivityFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**[]ActivityView**](ActivityView.md) | Activity is the merged run/create/update events, newest first, capped at 50. | [optional] 

## Methods

### NewActivityFeed

`func NewActivityFeed() *ActivityFeed`

NewActivityFeed instantiates a new ActivityFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityFeedWithDefaults

`func NewActivityFeedWithDefaults() *ActivityFeed`

NewActivityFeedWithDefaults instantiates a new ActivityFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *ActivityFeed) GetActivity() []ActivityView`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ActivityFeed) GetActivityOk() (*[]ActivityView, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ActivityFeed) SetActivity(v []ActivityView)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ActivityFeed) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


