# IamInvitationsListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | Pointer to [**[]IamInvitation**](IamInvitation.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewIamInvitationsListOutput

`func NewIamInvitationsListOutput() *IamInvitationsListOutput`

NewIamInvitationsListOutput instantiates a new IamInvitationsListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamInvitationsListOutputWithDefaults

`func NewIamInvitationsListOutputWithDefaults() *IamInvitationsListOutput`

NewIamInvitationsListOutputWithDefaults instantiates a new IamInvitationsListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *IamInvitationsListOutput) GetInvitations() []IamInvitation`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *IamInvitationsListOutput) GetInvitationsOk() (*[]IamInvitation, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *IamInvitationsListOutput) SetInvitations(v []IamInvitation)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *IamInvitationsListOutput) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetTotal

`func (o *IamInvitationsListOutput) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamInvitationsListOutput) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamInvitationsListOutput) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamInvitationsListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


