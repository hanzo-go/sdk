# CaptableInvestments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CaptableInvestment**](CaptableInvestment.md) | Data is every investment across every round, newest first. | [optional] 

## Methods

### NewCaptableInvestments

`func NewCaptableInvestments() *CaptableInvestments`

NewCaptableInvestments instantiates a new CaptableInvestments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableInvestmentsWithDefaults

`func NewCaptableInvestmentsWithDefaults() *CaptableInvestments`

NewCaptableInvestmentsWithDefaults instantiates a new CaptableInvestments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CaptableInvestments) GetData() []CaptableInvestment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CaptableInvestments) GetDataOk() (*[]CaptableInvestment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CaptableInvestments) SetData(v []CaptableInvestment)`

SetData sets Data field to given value.

### HasData

`func (o *CaptableInvestments) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


