# IamObjectTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]IamObjectTicketMessage**](IamObjectTicketMessage.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectTicket

`func NewIamObjectTicket() *IamObjectTicket`

NewIamObjectTicket instantiates a new IamObjectTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectTicketWithDefaults

`func NewIamObjectTicketWithDefaults() *IamObjectTicket`

NewIamObjectTicketWithDefaults instantiates a new IamObjectTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *IamObjectTicket) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IamObjectTicket) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IamObjectTicket) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IamObjectTicket) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectTicket) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectTicket) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectTicket) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectTicket) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectTicket) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectTicket) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectTicket) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectTicket) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMessages

`func (o *IamObjectTicket) GetMessages() []IamObjectTicketMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *IamObjectTicket) GetMessagesOk() (*[]IamObjectTicketMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *IamObjectTicket) SetMessages(v []IamObjectTicketMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *IamObjectTicket) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *IamObjectTicket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectTicket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectTicket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectTicket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectTicket) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectTicket) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectTicket) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectTicket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetState

`func (o *IamObjectTicket) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamObjectTicket) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamObjectTicket) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamObjectTicket) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *IamObjectTicket) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamObjectTicket) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamObjectTicket) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamObjectTicket) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *IamObjectTicket) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *IamObjectTicket) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *IamObjectTicket) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *IamObjectTicket) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUser

`func (o *IamObjectTicket) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamObjectTicket) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamObjectTicket) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamObjectTicket) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


