# CloudLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**[]CloudLogLine**](CloudLogLine.md) |  | [optional] 
**NextCursor** | Pointer to **int32** | pass back as ?sinceNs for the next tail poll | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** | \&quot;infra\&quot; (admin) | \&quot;request\&quot; (per-org) | [optional] 

## Methods

### NewCloudLogsResponse

`func NewCloudLogsResponse() *CloudLogsResponse`

NewCloudLogsResponse instantiates a new CloudLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLogsResponseWithDefaults

`func NewCloudLogsResponseWithDefaults() *CloudLogsResponse`

NewCloudLogsResponseWithDefaults instantiates a new CloudLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *CloudLogsResponse) GetLines() []CloudLogLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *CloudLogsResponse) GetLinesOk() (*[]CloudLogLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *CloudLogsResponse) SetLines(v []CloudLogLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *CloudLogsResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNextCursor

`func (o *CloudLogsResponse) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CloudLogsResponse) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CloudLogsResponse) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *CloudLogsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetProduct

`func (o *CloudLogsResponse) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudLogsResponse) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudLogsResponse) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudLogsResponse) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetView

`func (o *CloudLogsResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *CloudLogsResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *CloudLogsResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *CloudLogsResponse) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


