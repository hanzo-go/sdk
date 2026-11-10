# MarketingField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to **[]string** | Choices are the values a choice field may take; at most 100. | [optional] 
**Key** | Pointer to **string** | Key names the field in a subscriber&#39;s fields and in a filter: a lower-case letter, then letters, digits and underscores, at most 40. | [optional] 
**Label** | Pointer to **string** | Label is how the field is shown; at most 100 characters. | [optional] 
**Type** | Pointer to **string** | Type is text, number, date (YYYY-MM-DD), boolean (true or false) or choice. | [optional] 

## Methods

### NewMarketingField

`func NewMarketingField() *MarketingField`

NewMarketingField instantiates a new MarketingField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingFieldWithDefaults

`func NewMarketingFieldWithDefaults() *MarketingField`

NewMarketingFieldWithDefaults instantiates a new MarketingField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *MarketingField) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *MarketingField) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *MarketingField) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *MarketingField) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetKey

`func (o *MarketingField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MarketingField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MarketingField) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MarketingField) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *MarketingField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MarketingField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MarketingField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MarketingField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *MarketingField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketingField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketingField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarketingField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


