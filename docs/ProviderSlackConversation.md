# ProviderSlackConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the channel ID to pass to message reads and sends. | [optional] 
**IsIm** | Pointer to **bool** | IsIM identifies a direct message. | [optional] 
**IsMember** | Pointer to **bool** | IsMember reports whether Hanzo has joined this channel. | [optional] 
**IsMpim** | Pointer to **bool** | IsMPIM identifies a group direct message. | [optional] 
**IsPrivate** | Pointer to **bool** | IsPrivate identifies a private channel or conversation. | [optional] 
**Name** | Pointer to **string** | Name is the channel name, without #. DMs may have no name. | [optional] 

## Methods

### NewProviderSlackConversation

`func NewProviderSlackConversation() *ProviderSlackConversation`

NewProviderSlackConversation instantiates a new ProviderSlackConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackConversationWithDefaults

`func NewProviderSlackConversationWithDefaults() *ProviderSlackConversation`

NewProviderSlackConversationWithDefaults instantiates a new ProviderSlackConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderSlackConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderSlackConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderSlackConversation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderSlackConversation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsIm

`func (o *ProviderSlackConversation) GetIsIm() bool`

GetIsIm returns the IsIm field if non-nil, zero value otherwise.

### GetIsImOk

`func (o *ProviderSlackConversation) GetIsImOk() (*bool, bool)`

GetIsImOk returns a tuple with the IsIm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIm

`func (o *ProviderSlackConversation) SetIsIm(v bool)`

SetIsIm sets IsIm field to given value.

### HasIsIm

`func (o *ProviderSlackConversation) HasIsIm() bool`

HasIsIm returns a boolean if a field has been set.

### GetIsMember

`func (o *ProviderSlackConversation) GetIsMember() bool`

GetIsMember returns the IsMember field if non-nil, zero value otherwise.

### GetIsMemberOk

`func (o *ProviderSlackConversation) GetIsMemberOk() (*bool, bool)`

GetIsMemberOk returns a tuple with the IsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMember

`func (o *ProviderSlackConversation) SetIsMember(v bool)`

SetIsMember sets IsMember field to given value.

### HasIsMember

`func (o *ProviderSlackConversation) HasIsMember() bool`

HasIsMember returns a boolean if a field has been set.

### GetIsMpim

`func (o *ProviderSlackConversation) GetIsMpim() bool`

GetIsMpim returns the IsMpim field if non-nil, zero value otherwise.

### GetIsMpimOk

`func (o *ProviderSlackConversation) GetIsMpimOk() (*bool, bool)`

GetIsMpimOk returns a tuple with the IsMpim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMpim

`func (o *ProviderSlackConversation) SetIsMpim(v bool)`

SetIsMpim sets IsMpim field to given value.

### HasIsMpim

`func (o *ProviderSlackConversation) HasIsMpim() bool`

HasIsMpim returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ProviderSlackConversation) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ProviderSlackConversation) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ProviderSlackConversation) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ProviderSlackConversation) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetName

`func (o *ProviderSlackConversation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderSlackConversation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderSlackConversation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderSlackConversation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


