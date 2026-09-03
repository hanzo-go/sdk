# ReviewQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is how many founders are waiting. | [optional] 
**Queue** | Pointer to [**[]Waiting**](Waiting.md) | Queue is one entry per unsettled founder, oldest formation first. | [optional] 

## Methods

### NewReviewQueue

`func NewReviewQueue() *ReviewQueue`

NewReviewQueue instantiates a new ReviewQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewQueueWithDefaults

`func NewReviewQueueWithDefaults() *ReviewQueue`

NewReviewQueueWithDefaults instantiates a new ReviewQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ReviewQueue) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReviewQueue) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReviewQueue) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ReviewQueue) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetQueue

`func (o *ReviewQueue) GetQueue() []Waiting`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *ReviewQueue) GetQueueOk() (*[]Waiting, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *ReviewQueue) SetQueue(v []Waiting)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *ReviewQueue) HasQueue() bool`

HasQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


