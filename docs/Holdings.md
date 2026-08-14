# Holdings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**[]Holding**](Holding.md) | Domains is the caller org&#39;s domains, newest registration first. | [optional] 

## Methods

### NewHoldings

`func NewHoldings() *Holdings`

NewHoldings instantiates a new Holdings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingsWithDefaults

`func NewHoldingsWithDefaults() *Holdings`

NewHoldingsWithDefaults instantiates a new Holdings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *Holdings) GetDomains() []Holding`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Holdings) GetDomainsOk() (*[]Holding, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Holdings) SetDomains(v []Holding)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Holdings) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


