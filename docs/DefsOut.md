# DefsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DefRow**](DefRow.md) | Data is every definition in the caller&#39;s (org, project) store, by key. | [optional] 

## Methods

### NewDefsOut

`func NewDefsOut() *DefsOut`

NewDefsOut instantiates a new DefsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefsOutWithDefaults

`func NewDefsOutWithDefaults() *DefsOut`

NewDefsOutWithDefaults instantiates a new DefsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DefsOut) GetData() []DefRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DefsOut) GetDataOk() (*[]DefRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DefsOut) SetData(v []DefRow)`

SetData sets Data field to given value.

### HasData

`func (o *DefsOut) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


