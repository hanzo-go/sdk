# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectedSource** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) |  | [optional] 
**Usage** | Pointer to [**Tally**](Tally.md) |  | [optional] 

## Methods

### NewDelivery

`func NewDelivery() *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectedSource

`func (o *Delivery) GetDetectedSource() string`

GetDetectedSource returns the DetectedSource field if non-nil, zero value otherwise.

### GetDetectedSourceOk

`func (o *Delivery) GetDetectedSourceOk() (*string, bool)`

GetDetectedSourceOk returns a tuple with the DetectedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedSource

`func (o *Delivery) SetDetectedSource(v string)`

SetDetectedSource sets DetectedSource field to given value.

### HasDetectedSource

`func (o *Delivery) HasDetectedSource() bool`

HasDetectedSource returns a boolean if a field has been set.

### GetModel

`func (o *Delivery) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Delivery) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Delivery) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Delivery) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTier

`func (o *Delivery) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Delivery) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Delivery) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Delivery) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTranslations

`func (o *Delivery) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *Delivery) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *Delivery) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *Delivery) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetUsage

`func (o *Delivery) GetUsage() Tally`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Delivery) GetUsageOk() (*Tally, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Delivery) SetUsage(v Tally)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Delivery) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


