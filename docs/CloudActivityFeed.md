# CloudActivityFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**[]CloudActivityView**](CloudActivityView.md) | Activity is the merged run/create/update events, newest first, capped at 50. | [optional] 

## Methods

### NewCloudActivityFeed

`func NewCloudActivityFeed() *CloudActivityFeed`

NewCloudActivityFeed instantiates a new CloudActivityFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudActivityFeedWithDefaults

`func NewCloudActivityFeedWithDefaults() *CloudActivityFeed`

NewCloudActivityFeedWithDefaults instantiates a new CloudActivityFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *CloudActivityFeed) GetActivity() []CloudActivityView`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *CloudActivityFeed) GetActivityOk() (*[]CloudActivityView, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *CloudActivityFeed) SetActivity(v []CloudActivityView)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *CloudActivityFeed) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


