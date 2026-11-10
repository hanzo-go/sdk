# PatentDiscsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PatentPatentDisclosure**](PatentPatentDisclosure.md) | Data are the canonical disclosures, oldest append first. | [optional] 
**Total** | Pointer to **int64** | Total is len(data). | [optional] 

## Methods

### NewPatentDiscsOut

`func NewPatentDiscsOut() *PatentDiscsOut`

NewPatentDiscsOut instantiates a new PatentDiscsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentDiscsOutWithDefaults

`func NewPatentDiscsOutWithDefaults() *PatentDiscsOut`

NewPatentDiscsOutWithDefaults instantiates a new PatentDiscsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatentDiscsOut) GetData() []PatentPatentDisclosure`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatentDiscsOut) GetDataOk() (*[]PatentPatentDisclosure, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatentDiscsOut) SetData(v []PatentPatentDisclosure)`

SetData sets Data field to given value.

### HasData

`func (o *PatentDiscsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *PatentDiscsOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PatentDiscsOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PatentDiscsOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PatentDiscsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


