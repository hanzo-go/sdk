# MarketplaceVerifyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the expert listing, from the path. | [optional] 
**Verified** | Pointer to **bool** | Verified is whether staff vouch for the expert. | [optional] 

## Methods

### NewMarketplaceVerifyReq

`func NewMarketplaceVerifyReq() *MarketplaceVerifyReq`

NewMarketplaceVerifyReq instantiates a new MarketplaceVerifyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceVerifyReqWithDefaults

`func NewMarketplaceVerifyReqWithDefaults() *MarketplaceVerifyReq`

NewMarketplaceVerifyReqWithDefaults instantiates a new MarketplaceVerifyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketplaceVerifyReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceVerifyReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceVerifyReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceVerifyReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVerified

`func (o *MarketplaceVerifyReq) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MarketplaceVerifyReq) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MarketplaceVerifyReq) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MarketplaceVerifyReq) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


