# TelNumberList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TelNumber**](TelNumber.md) | Data is the numbers, and which numbers depends on the route: a search answers what the carrier has available, a list answers what this org already holds. | [optional] 

## Methods

### NewTelNumberList

`func NewTelNumberList() *TelNumberList`

NewTelNumberList instantiates a new TelNumberList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelNumberListWithDefaults

`func NewTelNumberListWithDefaults() *TelNumberList`

NewTelNumberListWithDefaults instantiates a new TelNumberList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TelNumberList) GetData() []TelNumber`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TelNumberList) GetDataOk() (*[]TelNumber, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TelNumberList) SetData(v []TelNumber)`

SetData sets Data field to given value.

### HasData

`func (o *TelNumberList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


