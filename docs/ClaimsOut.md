# ClaimsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ClaimRow**](ClaimRow.md) | Data is one row per (benchmark, model, SOURCE) — every independent claim, not one per model. Effective values only: the row that wins after layering for each source, never the superseded readings behind it. | [optional] 
**Total** | Pointer to **int64** | Total is how many rows Data holds. | [optional] 

## Methods

### NewClaimsOut

`func NewClaimsOut() *ClaimsOut`

NewClaimsOut instantiates a new ClaimsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimsOutWithDefaults

`func NewClaimsOutWithDefaults() *ClaimsOut`

NewClaimsOutWithDefaults instantiates a new ClaimsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ClaimsOut) GetData() []ClaimRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClaimsOut) GetDataOk() (*[]ClaimRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClaimsOut) SetData(v []ClaimRow)`

SetData sets Data field to given value.

### HasData

`func (o *ClaimsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *ClaimsOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ClaimsOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ClaimsOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ClaimsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


