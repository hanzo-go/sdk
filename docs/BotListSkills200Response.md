# BotListSkills200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]BotSkill**](BotSkill.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewBotListSkills200Response

`func NewBotListSkills200Response() *BotListSkills200Response`

NewBotListSkills200Response instantiates a new BotListSkills200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotListSkills200ResponseWithDefaults

`func NewBotListSkills200ResponseWithDefaults() *BotListSkills200Response`

NewBotListSkills200ResponseWithDefaults instantiates a new BotListSkills200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BotListSkills200Response) GetItems() []BotSkill`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BotListSkills200Response) GetItemsOk() (*[]BotSkill, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BotListSkills200Response) SetItems(v []BotSkill)`

SetItems sets Items field to given value.

### HasItems

`func (o *BotListSkills200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextCursor

`func (o *BotListSkills200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BotListSkills200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BotListSkills200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *BotListSkills200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetHasMore

`func (o *BotListSkills200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *BotListSkills200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *BotListSkills200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *BotListSkills200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


