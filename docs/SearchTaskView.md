# SearchTaskView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** |  | [optional] 
**IndexUid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CanceledBy** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to [**SearchResponseError**](SearchResponseError.md) |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**EnqueuedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSearchTaskView

`func NewSearchTaskView() *SearchTaskView`

NewSearchTaskView instantiates a new SearchTaskView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTaskViewWithDefaults

`func NewSearchTaskViewWithDefaults() *SearchTaskView`

NewSearchTaskViewWithDefaults instantiates a new SearchTaskView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *SearchTaskView) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SearchTaskView) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SearchTaskView) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SearchTaskView) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetIndexUid

`func (o *SearchTaskView) GetIndexUid() string`

GetIndexUid returns the IndexUid field if non-nil, zero value otherwise.

### GetIndexUidOk

`func (o *SearchTaskView) GetIndexUidOk() (*string, bool)`

GetIndexUidOk returns a tuple with the IndexUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUid

`func (o *SearchTaskView) SetIndexUid(v string)`

SetIndexUid sets IndexUid field to given value.

### HasIndexUid

`func (o *SearchTaskView) HasIndexUid() bool`

HasIndexUid returns a boolean if a field has been set.

### GetStatus

`func (o *SearchTaskView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchTaskView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchTaskView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchTaskView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SearchTaskView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchTaskView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchTaskView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchTaskView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCanceledBy

`func (o *SearchTaskView) GetCanceledBy() int32`

GetCanceledBy returns the CanceledBy field if non-nil, zero value otherwise.

### GetCanceledByOk

`func (o *SearchTaskView) GetCanceledByOk() (*int32, bool)`

GetCanceledByOk returns a tuple with the CanceledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledBy

`func (o *SearchTaskView) SetCanceledBy(v int32)`

SetCanceledBy sets CanceledBy field to given value.

### HasCanceledBy

`func (o *SearchTaskView) HasCanceledBy() bool`

HasCanceledBy returns a boolean if a field has been set.

### GetDetails

`func (o *SearchTaskView) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SearchTaskView) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SearchTaskView) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SearchTaskView) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetError

`func (o *SearchTaskView) GetError() SearchResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchTaskView) GetErrorOk() (*SearchResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchTaskView) SetError(v SearchResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SearchTaskView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDuration

`func (o *SearchTaskView) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SearchTaskView) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SearchTaskView) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SearchTaskView) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnqueuedAt

`func (o *SearchTaskView) GetEnqueuedAt() time.Time`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *SearchTaskView) GetEnqueuedAtOk() (*time.Time, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *SearchTaskView) SetEnqueuedAt(v time.Time)`

SetEnqueuedAt sets EnqueuedAt field to given value.

### HasEnqueuedAt

`func (o *SearchTaskView) HasEnqueuedAt() bool`

HasEnqueuedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *SearchTaskView) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SearchTaskView) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SearchTaskView) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SearchTaskView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *SearchTaskView) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SearchTaskView) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SearchTaskView) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *SearchTaskView) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


