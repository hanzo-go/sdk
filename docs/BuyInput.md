# BuyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E164** | Pointer to **string** | E164 is the number to buy, in E.164 (a leading + and digits), exactly as the search returned it. This is the number itself, not the id from a search result. | [optional] 

## Methods

### NewBuyInput

`func NewBuyInput() *BuyInput`

NewBuyInput instantiates a new BuyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyInputWithDefaults

`func NewBuyInputWithDefaults() *BuyInput`

NewBuyInputWithDefaults instantiates a new BuyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE164

`func (o *BuyInput) GetE164() string`

GetE164 returns the E164 field if non-nil, zero value otherwise.

### GetE164Ok

`func (o *BuyInput) GetE164Ok() (*string, bool)`

GetE164Ok returns a tuple with the E164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE164

`func (o *BuyInput) SetE164(v string)`

SetE164 sets E164 field to given value.

### HasE164

`func (o *BuyInput) HasE164() bool`

HasE164 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


