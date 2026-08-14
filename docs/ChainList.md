# ChainList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chains** | Pointer to [**[]Chain**](Chain.md) | Chains is every chain this deployment is configured to reach, sorted by id. Empty when none are configured — never a fabricated entry. | [optional] 

## Methods

### NewChainList

`func NewChainList() *ChainList`

NewChainList instantiates a new ChainList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainListWithDefaults

`func NewChainListWithDefaults() *ChainList`

NewChainListWithDefaults instantiates a new ChainList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChains

`func (o *ChainList) GetChains() []Chain`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *ChainList) GetChainsOk() (*[]Chain, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *ChainList) SetChains(v []Chain)`

SetChains sets Chains field to given value.

### HasChains

`func (o *ChainList) HasChains() bool`

HasChains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


