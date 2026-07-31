# ObserveLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** | \&quot;infra\&quot; (admin) or \&quot;request\&quot; (per-org). | [optional] 
**Lines** | Pointer to [**[]ObserveLogLine**](ObserveLogLine.md) |  | [optional] 
**NextCursor** | Pointer to **int64** | Pass back as &#x60;sinceNs&#x60; for the next tail poll. | [optional] 

## Methods

### NewObserveLogsResponse

`func NewObserveLogsResponse() *ObserveLogsResponse`

NewObserveLogsResponse instantiates a new ObserveLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObserveLogsResponseWithDefaults

`func NewObserveLogsResponseWithDefaults() *ObserveLogsResponse`

NewObserveLogsResponseWithDefaults instantiates a new ObserveLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ObserveLogsResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ObserveLogsResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ObserveLogsResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ObserveLogsResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetView

`func (o *ObserveLogsResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ObserveLogsResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ObserveLogsResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ObserveLogsResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetLines

`func (o *ObserveLogsResponse) GetLines() []ObserveLogLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ObserveLogsResponse) GetLinesOk() (*[]ObserveLogLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ObserveLogsResponse) SetLines(v []ObserveLogLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *ObserveLogsResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNextCursor

`func (o *ObserveLogsResponse) GetNextCursor() int64`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ObserveLogsResponse) GetNextCursorOk() (*int64, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ObserveLogsResponse) SetNextCursor(v int64)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ObserveLogsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


