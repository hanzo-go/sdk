# Web3Chain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **int64** | ChainID is the EIP-155 id, so a caller can check it matches the wallet they are about to sign with. | [optional] 
**Id** | Pointer to **string** | ID is the URL name: the value of :chain. | [optional] 
**Name** | Pointer to **string** | Name is for humans. | [optional] 

## Methods

### NewWeb3Chain

`func NewWeb3Chain() *Web3Chain`

NewWeb3Chain instantiates a new Web3Chain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeb3ChainWithDefaults

`func NewWeb3ChainWithDefaults() *Web3Chain`

NewWeb3ChainWithDefaults instantiates a new Web3Chain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *Web3Chain) GetChainId() int64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Web3Chain) GetChainIdOk() (*int64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Web3Chain) SetChainId(v int64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Web3Chain) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetId

`func (o *Web3Chain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Web3Chain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Web3Chain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Web3Chain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Web3Chain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Web3Chain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Web3Chain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Web3Chain) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


