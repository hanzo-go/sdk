# ProviderFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label is what the choice is called. | [optional] 
**Value** | Pointer to **string** | Value is what submitting it sends. | [optional] 

## Methods

### NewProviderFieldOption

`func NewProviderFieldOption() *ProviderFieldOption`

NewProviderFieldOption instantiates a new ProviderFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderFieldOptionWithDefaults

`func NewProviderFieldOptionWithDefaults() *ProviderFieldOption`

NewProviderFieldOptionWithDefaults instantiates a new ProviderFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ProviderFieldOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProviderFieldOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProviderFieldOption) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProviderFieldOption) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *ProviderFieldOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProviderFieldOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProviderFieldOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProviderFieldOption) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


