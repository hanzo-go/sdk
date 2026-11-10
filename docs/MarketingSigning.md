# MarketingSigning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is where the SAFE stands: submitted, investor signed, countersigned, completed, declined or voided. Empty when the document is no investor&#39;s of this deployment. | [optional] 

## Methods

### NewMarketingSigning

`func NewMarketingSigning() *MarketingSigning`

NewMarketingSigning instantiates a new MarketingSigning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSigningWithDefaults

`func NewMarketingSigningWithDefaults() *MarketingSigning`

NewMarketingSigningWithDefaults instantiates a new MarketingSigning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MarketingSigning) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingSigning) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingSigning) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingSigning) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


