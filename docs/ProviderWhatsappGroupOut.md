# ProviderWhatsappGroupOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the group id. A reply addresses it with recipient_type group. | [optional] 
**InviteLink** | Pointer to **string** | InviteLink is the link a person opens to join. Anyone holding it can join, so it is answered only to an org admin. | [optional] 

## Methods

### NewProviderWhatsappGroupOut

`func NewProviderWhatsappGroupOut() *ProviderWhatsappGroupOut`

NewProviderWhatsappGroupOut instantiates a new ProviderWhatsappGroupOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWhatsappGroupOutWithDefaults

`func NewProviderWhatsappGroupOutWithDefaults() *ProviderWhatsappGroupOut`

NewProviderWhatsappGroupOutWithDefaults instantiates a new ProviderWhatsappGroupOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderWhatsappGroupOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderWhatsappGroupOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderWhatsappGroupOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderWhatsappGroupOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInviteLink

`func (o *ProviderWhatsappGroupOut) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *ProviderWhatsappGroupOut) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *ProviderWhatsappGroupOut) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *ProviderWhatsappGroupOut) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


