# O11yLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** | infra for a validated SuperAdmin, request for every other org. | [optional] 
**Lines** | Pointer to [**[]O11yLogLine**](O11yLogLine.md) |  | [optional] 
**NextCursor** | Pointer to **int64** | Max nanosecond cursor. Pass back as sinceNs for the next tail poll. | [optional] 

## Methods

### NewO11yLogsResponse

`func NewO11yLogsResponse() *O11yLogsResponse`

NewO11yLogsResponse instantiates a new O11yLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yLogsResponseWithDefaults

`func NewO11yLogsResponseWithDefaults() *O11yLogsResponse`

NewO11yLogsResponseWithDefaults instantiates a new O11yLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *O11yLogsResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *O11yLogsResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *O11yLogsResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *O11yLogsResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetView

`func (o *O11yLogsResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *O11yLogsResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *O11yLogsResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *O11yLogsResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetLines

`func (o *O11yLogsResponse) GetLines() []O11yLogLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *O11yLogsResponse) GetLinesOk() (*[]O11yLogLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *O11yLogsResponse) SetLines(v []O11yLogLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *O11yLogsResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNextCursor

`func (o *O11yLogsResponse) GetNextCursor() int64`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *O11yLogsResponse) GetNextCursorOk() (*int64, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *O11yLogsResponse) SetNextCursor(v int64)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *O11yLogsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


