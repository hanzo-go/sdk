# TeamTeamReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | Accounts are the account uuids that reacted with it, earliest first. | [optional] 
**Count** | Pointer to **int64** | Count is how many people reacted with it. | [optional] 
**Emoji** | Pointer to **string** | Emoji is the reaction itself, as the person picked it. | [optional] 

## Methods

### NewTeamTeamReaction

`func NewTeamTeamReaction() *TeamTeamReaction`

NewTeamTeamReaction instantiates a new TeamTeamReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamReactionWithDefaults

`func NewTeamTeamReactionWithDefaults() *TeamTeamReaction`

NewTeamTeamReactionWithDefaults instantiates a new TeamTeamReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *TeamTeamReaction) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *TeamTeamReaction) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *TeamTeamReaction) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *TeamTeamReaction) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCount

`func (o *TeamTeamReaction) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TeamTeamReaction) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TeamTeamReaction) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TeamTeamReaction) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEmoji

`func (o *TeamTeamReaction) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *TeamTeamReaction) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *TeamTeamReaction) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *TeamTeamReaction) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


