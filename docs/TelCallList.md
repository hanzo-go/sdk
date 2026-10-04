# TelCallList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TelCall**](TelCall.md) | Data is this org&#39;s own calls, newest first — what this platform placed or received on its behalf, which is our record rather than the carrier&#39;s. | [optional] 

## Methods

### NewTelCallList

`func NewTelCallList() *TelCallList`

NewTelCallList instantiates a new TelCallList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelCallListWithDefaults

`func NewTelCallListWithDefaults() *TelCallList`

NewTelCallListWithDefaults instantiates a new TelCallList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TelCallList) GetData() []TelCall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TelCallList) GetDataOk() (*[]TelCall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TelCallList) SetData(v []TelCall)`

SetData sets Data field to given value.

### HasData

`func (o *TelCallList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


