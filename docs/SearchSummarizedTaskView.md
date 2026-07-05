# SearchSummarizedTaskView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUid** | Pointer to **int32** |  | [optional] 
**IndexUid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**EnqueuedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSearchSummarizedTaskView

`func NewSearchSummarizedTaskView() *SearchSummarizedTaskView`

NewSearchSummarizedTaskView instantiates a new SearchSummarizedTaskView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSummarizedTaskViewWithDefaults

`func NewSearchSummarizedTaskViewWithDefaults() *SearchSummarizedTaskView`

NewSearchSummarizedTaskViewWithDefaults instantiates a new SearchSummarizedTaskView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUid

`func (o *SearchSummarizedTaskView) GetTaskUid() int32`

GetTaskUid returns the TaskUid field if non-nil, zero value otherwise.

### GetTaskUidOk

`func (o *SearchSummarizedTaskView) GetTaskUidOk() (*int32, bool)`

GetTaskUidOk returns a tuple with the TaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUid

`func (o *SearchSummarizedTaskView) SetTaskUid(v int32)`

SetTaskUid sets TaskUid field to given value.

### HasTaskUid

`func (o *SearchSummarizedTaskView) HasTaskUid() bool`

HasTaskUid returns a boolean if a field has been set.

### GetIndexUid

`func (o *SearchSummarizedTaskView) GetIndexUid() string`

GetIndexUid returns the IndexUid field if non-nil, zero value otherwise.

### GetIndexUidOk

`func (o *SearchSummarizedTaskView) GetIndexUidOk() (*string, bool)`

GetIndexUidOk returns a tuple with the IndexUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUid

`func (o *SearchSummarizedTaskView) SetIndexUid(v string)`

SetIndexUid sets IndexUid field to given value.

### HasIndexUid

`func (o *SearchSummarizedTaskView) HasIndexUid() bool`

HasIndexUid returns a boolean if a field has been set.

### GetStatus

`func (o *SearchSummarizedTaskView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchSummarizedTaskView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchSummarizedTaskView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchSummarizedTaskView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SearchSummarizedTaskView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchSummarizedTaskView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchSummarizedTaskView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchSummarizedTaskView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnqueuedAt

`func (o *SearchSummarizedTaskView) GetEnqueuedAt() time.Time`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *SearchSummarizedTaskView) GetEnqueuedAtOk() (*time.Time, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *SearchSummarizedTaskView) SetEnqueuedAt(v time.Time)`

SetEnqueuedAt sets EnqueuedAt field to given value.

### HasEnqueuedAt

`func (o *SearchSummarizedTaskView) HasEnqueuedAt() bool`

HasEnqueuedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


