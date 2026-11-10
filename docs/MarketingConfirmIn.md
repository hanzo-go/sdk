# MarketingConfirmIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Token is the t value in the fragment of the link the confirmation mail carried (&lt;brand site&gt;/newsletter/confirm#t&#x3D;&lt;token&gt;). It is the request&#39;s only authority: it names the address, and only this deployment signs one. | [optional] 

## Methods

### NewMarketingConfirmIn

`func NewMarketingConfirmIn() *MarketingConfirmIn`

NewMarketingConfirmIn instantiates a new MarketingConfirmIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingConfirmInWithDefaults

`func NewMarketingConfirmInWithDefaults() *MarketingConfirmIn`

NewMarketingConfirmInWithDefaults instantiates a new MarketingConfirmIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *MarketingConfirmIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MarketingConfirmIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MarketingConfirmIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MarketingConfirmIn) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


