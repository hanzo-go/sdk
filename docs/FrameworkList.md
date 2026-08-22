# FrameworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frameworks** | Pointer to [**[]FrameworkRow**](FrameworkRow.md) | Frameworks is each framework and how many clauses it publishes. | [optional] 

## Methods

### NewFrameworkList

`func NewFrameworkList() *FrameworkList`

NewFrameworkList instantiates a new FrameworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkListWithDefaults

`func NewFrameworkListWithDefaults() *FrameworkList`

NewFrameworkListWithDefaults instantiates a new FrameworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrameworks

`func (o *FrameworkList) GetFrameworks() []FrameworkRow`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *FrameworkList) GetFrameworksOk() (*[]FrameworkRow, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *FrameworkList) SetFrameworks(v []FrameworkRow)`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *FrameworkList) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


