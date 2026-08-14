# Principle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Change** | Pointer to **string** | the Book of Changes reading | [optional] 
**Domain** | Pointer to **string** | the growth / go-to-market domain it governs | [optional] 
**Hexagram** | Pointer to **string** | the I-Ching hexagram (pinyin + gloss) | [optional] 
**N** | Pointer to **int32** | 1..64, the hexagram number + canonical order | [optional] 
**Name** | Pointer to **string** | the principle&#39;s short name | [optional] 
**Principle** | Pointer to **string** | the actionable growth law | [optional] 
**Slug** | Pointer to **string** | stable identifier a tactic files under | [optional] 
**SunTzu** | Pointer to **string** | the Art of War teaching | [optional] 

## Methods

### NewPrinciple

`func NewPrinciple() *Principle`

NewPrinciple instantiates a new Principle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipleWithDefaults

`func NewPrincipleWithDefaults() *Principle`

NewPrincipleWithDefaults instantiates a new Principle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChange

`func (o *Principle) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Principle) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Principle) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *Principle) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetDomain

`func (o *Principle) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Principle) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Principle) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Principle) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHexagram

`func (o *Principle) GetHexagram() string`

GetHexagram returns the Hexagram field if non-nil, zero value otherwise.

### GetHexagramOk

`func (o *Principle) GetHexagramOk() (*string, bool)`

GetHexagramOk returns a tuple with the Hexagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexagram

`func (o *Principle) SetHexagram(v string)`

SetHexagram sets Hexagram field to given value.

### HasHexagram

`func (o *Principle) HasHexagram() bool`

HasHexagram returns a boolean if a field has been set.

### GetN

`func (o *Principle) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *Principle) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *Principle) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *Principle) HasN() bool`

HasN returns a boolean if a field has been set.

### GetName

`func (o *Principle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Principle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Principle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Principle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrinciple

`func (o *Principle) GetPrinciple() string`

GetPrinciple returns the Principle field if non-nil, zero value otherwise.

### GetPrincipleOk

`func (o *Principle) GetPrincipleOk() (*string, bool)`

GetPrincipleOk returns a tuple with the Principle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinciple

`func (o *Principle) SetPrinciple(v string)`

SetPrinciple sets Principle field to given value.

### HasPrinciple

`func (o *Principle) HasPrinciple() bool`

HasPrinciple returns a boolean if a field has been set.

### GetSlug

`func (o *Principle) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Principle) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Principle) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Principle) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSunTzu

`func (o *Principle) GetSunTzu() string`

GetSunTzu returns the SunTzu field if non-nil, zero value otherwise.

### GetSunTzuOk

`func (o *Principle) GetSunTzuOk() (*string, bool)`

GetSunTzuOk returns a tuple with the SunTzu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunTzu

`func (o *Principle) SetSunTzu(v string)`

SetSunTzu sets SunTzu field to given value.

### HasSunTzu

`func (o *Principle) HasSunTzu() bool`

HasSunTzu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


