# MarketingSendIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At schedules it for a later time, unix seconds; zero sends now. | [optional] 
**Id** | Pointer to **string** | ID is the broadcast, from the path. | [optional] 

## Methods

### NewMarketingSendIn

`func NewMarketingSendIn() *MarketingSendIn`

NewMarketingSendIn instantiates a new MarketingSendIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSendInWithDefaults

`func NewMarketingSendInWithDefaults() *MarketingSendIn`

NewMarketingSendInWithDefaults instantiates a new MarketingSendIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketingSendIn) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketingSendIn) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketingSendIn) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketingSendIn) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetId

`func (o *MarketingSendIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingSendIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingSendIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingSendIn) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


