# MarketplaceDeclineIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the job, from the path. | [optional] 
**Reason** | Pointer to **string** | Reason is the seller&#39;s words to the buyer, at most 4096 characters. | [optional] 

## Methods

### NewMarketplaceDeclineIn

`func NewMarketplaceDeclineIn() *MarketplaceDeclineIn`

NewMarketplaceDeclineIn instantiates a new MarketplaceDeclineIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDeclineInWithDefaults

`func NewMarketplaceDeclineInWithDefaults() *MarketplaceDeclineIn`

NewMarketplaceDeclineInWithDefaults instantiates a new MarketplaceDeclineIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketplaceDeclineIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceDeclineIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceDeclineIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceDeclineIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *MarketplaceDeclineIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketplaceDeclineIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketplaceDeclineIn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketplaceDeclineIn) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


