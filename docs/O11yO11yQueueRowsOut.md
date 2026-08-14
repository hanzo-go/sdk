# O11yO11yQueueRowsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yO11yQueueRow**](O11yO11yQueueRow.md) | Data holds one row per messaging destination. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yQueueRowsOut

`func NewO11yO11yQueueRowsOut() *O11yO11yQueueRowsOut`

NewO11yO11yQueueRowsOut instantiates a new O11yO11yQueueRowsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueRowsOutWithDefaults

`func NewO11yO11yQueueRowsOutWithDefaults() *O11yO11yQueueRowsOut`

NewO11yO11yQueueRowsOutWithDefaults instantiates a new O11yO11yQueueRowsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yQueueRowsOut) GetData() []O11yO11yQueueRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yQueueRowsOut) GetDataOk() (*[]O11yO11yQueueRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yQueueRowsOut) SetData(v []O11yO11yQueueRow)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yQueueRowsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yQueueRowsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yQueueRowsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yQueueRowsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yQueueRowsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


