# BotMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is whether the agent projects as a LIVE space member, derived from its registry status: empty, \&quot;active\&quot; and \&quot;ready\&quot; are live, anything else (archived/retired) is not. An inactive bot drops out of the Team list while its past authorship survives. | [optional] 
**Id** | Pointer to **string** | the agent id | [optional] 
**Name** | Pointer to **string** | display name | [optional] 
**PersonRef** | Pointer to **string** | the projected Person _id | [optional] 
**UserId** | Pointer to **string** | derived member account uuid (personUuid) | [optional] 

## Methods

### NewBotMember

`func NewBotMember() *BotMember`

NewBotMember instantiates a new BotMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotMemberWithDefaults

`func NewBotMemberWithDefaults() *BotMember`

NewBotMemberWithDefaults instantiates a new BotMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BotMember) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BotMember) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BotMember) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BotMember) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *BotMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BotMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BotMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonRef

`func (o *BotMember) GetPersonRef() string`

GetPersonRef returns the PersonRef field if non-nil, zero value otherwise.

### GetPersonRefOk

`func (o *BotMember) GetPersonRefOk() (*string, bool)`

GetPersonRefOk returns a tuple with the PersonRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonRef

`func (o *BotMember) SetPersonRef(v string)`

SetPersonRef sets PersonRef field to given value.

### HasPersonRef

`func (o *BotMember) HasPersonRef() bool`

HasPersonRef returns a boolean if a field has been set.

### GetUserId

`func (o *BotMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BotMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BotMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BotMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


