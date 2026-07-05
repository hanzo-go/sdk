# BotListPersonas200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]BotPersona**](BotPersona.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewBotListPersonas200Response

`func NewBotListPersonas200Response() *BotListPersonas200Response`

NewBotListPersonas200Response instantiates a new BotListPersonas200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotListPersonas200ResponseWithDefaults

`func NewBotListPersonas200ResponseWithDefaults() *BotListPersonas200Response`

NewBotListPersonas200ResponseWithDefaults instantiates a new BotListPersonas200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BotListPersonas200Response) GetItems() []BotPersona`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BotListPersonas200Response) GetItemsOk() (*[]BotPersona, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BotListPersonas200Response) SetItems(v []BotPersona)`

SetItems sets Items field to given value.

### HasItems

`func (o *BotListPersonas200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetHasMore

`func (o *BotListPersonas200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *BotListPersonas200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *BotListPersonas200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *BotListPersonas200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


