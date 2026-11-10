# PatentIdeasOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PatentPatentIdea**](PatentPatentIdea.md) | Data are the canonical entries, oldest append first. | [optional] 
**Total** | Pointer to **int64** | Total is len(data). | [optional] 

## Methods

### NewPatentIdeasOut

`func NewPatentIdeasOut() *PatentIdeasOut`

NewPatentIdeasOut instantiates a new PatentIdeasOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentIdeasOutWithDefaults

`func NewPatentIdeasOutWithDefaults() *PatentIdeasOut`

NewPatentIdeasOutWithDefaults instantiates a new PatentIdeasOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatentIdeasOut) GetData() []PatentPatentIdea`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatentIdeasOut) GetDataOk() (*[]PatentPatentIdea, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatentIdeasOut) SetData(v []PatentPatentIdea)`

SetData sets Data field to given value.

### HasData

`func (o *PatentIdeasOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *PatentIdeasOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PatentIdeasOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PatentIdeasOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PatentIdeasOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


