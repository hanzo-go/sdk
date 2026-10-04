# MarketplaceDisputeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the job, from the path. | [optional] 
**Reason** | Pointer to **string** | Reason is the disputing party&#39;s words. Required, at most 4096 characters. | [optional] 

## Methods

### NewMarketplaceDisputeIn

`func NewMarketplaceDisputeIn() *MarketplaceDisputeIn`

NewMarketplaceDisputeIn instantiates a new MarketplaceDisputeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDisputeInWithDefaults

`func NewMarketplaceDisputeInWithDefaults() *MarketplaceDisputeIn`

NewMarketplaceDisputeInWithDefaults instantiates a new MarketplaceDisputeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketplaceDisputeIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceDisputeIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceDisputeIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceDisputeIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *MarketplaceDisputeIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketplaceDisputeIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketplaceDisputeIn) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketplaceDisputeIn) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


