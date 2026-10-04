# MarketplaceDeliverIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | Hash is the seller&#39;s own commitment to what was delivered, 32 bytes as 0x hex; the SHA-256 of the note and URL when empty. | [optional] 
**Id** | Pointer to **string** | ID is the job, from the path. | [optional] 
**Note** | Pointer to **string** | Note says what was delivered. Required, at most 4096 characters. | [optional] 
**Url** | Pointer to **string** | URL is where it is, https. | [optional] 

## Methods

### NewMarketplaceDeliverIn

`func NewMarketplaceDeliverIn() *MarketplaceDeliverIn`

NewMarketplaceDeliverIn instantiates a new MarketplaceDeliverIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDeliverInWithDefaults

`func NewMarketplaceDeliverInWithDefaults() *MarketplaceDeliverIn`

NewMarketplaceDeliverInWithDefaults instantiates a new MarketplaceDeliverIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *MarketplaceDeliverIn) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *MarketplaceDeliverIn) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *MarketplaceDeliverIn) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *MarketplaceDeliverIn) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceDeliverIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceDeliverIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceDeliverIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceDeliverIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNote

`func (o *MarketplaceDeliverIn) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *MarketplaceDeliverIn) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *MarketplaceDeliverIn) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *MarketplaceDeliverIn) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUrl

`func (o *MarketplaceDeliverIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketplaceDeliverIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketplaceDeliverIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MarketplaceDeliverIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


