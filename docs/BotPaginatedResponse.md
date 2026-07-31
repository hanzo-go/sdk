# BotPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]interface{}** |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewBotPaginatedResponse

`func NewBotPaginatedResponse() *BotPaginatedResponse`

NewBotPaginatedResponse instantiates a new BotPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotPaginatedResponseWithDefaults

`func NewBotPaginatedResponseWithDefaults() *BotPaginatedResponse`

NewBotPaginatedResponseWithDefaults instantiates a new BotPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BotPaginatedResponse) GetItems() []interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BotPaginatedResponse) GetItemsOk() (*[]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BotPaginatedResponse) SetItems(v []interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *BotPaginatedResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextCursor

`func (o *BotPaginatedResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BotPaginatedResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BotPaginatedResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *BotPaginatedResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetHasMore

`func (o *BotPaginatedResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *BotPaginatedResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *BotPaginatedResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *BotPaginatedResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


