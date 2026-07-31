# CloudReviewQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is how many founders are waiting. | [optional] 
**Queue** | Pointer to [**[]CloudWaiting**](CloudWaiting.md) | Queue is one entry per unsettled founder, oldest formation first. | [optional] 

## Methods

### NewCloudReviewQueue

`func NewCloudReviewQueue() *CloudReviewQueue`

NewCloudReviewQueue instantiates a new CloudReviewQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudReviewQueueWithDefaults

`func NewCloudReviewQueueWithDefaults() *CloudReviewQueue`

NewCloudReviewQueueWithDefaults instantiates a new CloudReviewQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CloudReviewQueue) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudReviewQueue) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudReviewQueue) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudReviewQueue) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetQueue

`func (o *CloudReviewQueue) GetQueue() []CloudWaiting`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CloudReviewQueue) GetQueueOk() (*[]CloudWaiting, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CloudReviewQueue) SetQueue(v []CloudWaiting)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CloudReviewQueue) HasQueue() bool`

HasQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


