# PatentApprovalsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PatentPatentApproval**](PatentPatentApproval.md) | Data are the current decisions of record, oldest append first. | [optional] 
**Total** | Pointer to **int64** | Total is len(data). | [optional] 

## Methods

### NewPatentApprovalsOut

`func NewPatentApprovalsOut() *PatentApprovalsOut`

NewPatentApprovalsOut instantiates a new PatentApprovalsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentApprovalsOutWithDefaults

`func NewPatentApprovalsOutWithDefaults() *PatentApprovalsOut`

NewPatentApprovalsOutWithDefaults instantiates a new PatentApprovalsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatentApprovalsOut) GetData() []PatentPatentApproval`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatentApprovalsOut) GetDataOk() (*[]PatentPatentApproval, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatentApprovalsOut) SetData(v []PatentPatentApproval)`

SetData sets Data field to given value.

### HasData

`func (o *PatentApprovalsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *PatentApprovalsOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PatentApprovalsOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PatentApprovalsOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PatentApprovalsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


