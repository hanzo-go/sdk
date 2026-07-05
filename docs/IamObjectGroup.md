# IamObjectGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]IamObjectGroup**](IamObjectGroup.md) |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**HaveChildren** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsTopGroup** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamObjectGroup

`func NewIamObjectGroup() *IamObjectGroup`

NewIamObjectGroup instantiates a new IamObjectGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectGroupWithDefaults

`func NewIamObjectGroupWithDefaults() *IamObjectGroup`

NewIamObjectGroupWithDefaults instantiates a new IamObjectGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *IamObjectGroup) GetChildren() []IamObjectGroup`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *IamObjectGroup) GetChildrenOk() (*[]IamObjectGroup, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *IamObjectGroup) SetChildren(v []IamObjectGroup)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *IamObjectGroup) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetContactEmail

`func (o *IamObjectGroup) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *IamObjectGroup) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *IamObjectGroup) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *IamObjectGroup) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamObjectGroup) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamObjectGroup) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamObjectGroup) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamObjectGroup) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamObjectGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamObjectGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamObjectGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamObjectGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHaveChildren

`func (o *IamObjectGroup) GetHaveChildren() bool`

GetHaveChildren returns the HaveChildren field if non-nil, zero value otherwise.

### GetHaveChildrenOk

`func (o *IamObjectGroup) GetHaveChildrenOk() (*bool, bool)`

GetHaveChildrenOk returns a tuple with the HaveChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaveChildren

`func (o *IamObjectGroup) SetHaveChildren(v bool)`

SetHaveChildren sets HaveChildren field to given value.

### HasHaveChildren

`func (o *IamObjectGroup) HasHaveChildren() bool`

HasHaveChildren returns a boolean if a field has been set.

### GetIsEnabled

`func (o *IamObjectGroup) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *IamObjectGroup) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *IamObjectGroup) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *IamObjectGroup) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsTopGroup

`func (o *IamObjectGroup) GetIsTopGroup() bool`

GetIsTopGroup returns the IsTopGroup field if non-nil, zero value otherwise.

### GetIsTopGroupOk

`func (o *IamObjectGroup) GetIsTopGroupOk() (*bool, bool)`

GetIsTopGroupOk returns a tuple with the IsTopGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTopGroup

`func (o *IamObjectGroup) SetIsTopGroup(v bool)`

SetIsTopGroup sets IsTopGroup field to given value.

### HasIsTopGroup

`func (o *IamObjectGroup) HasIsTopGroup() bool`

HasIsTopGroup returns a boolean if a field has been set.

### GetKey

`func (o *IamObjectGroup) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IamObjectGroup) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IamObjectGroup) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IamObjectGroup) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetManager

`func (o *IamObjectGroup) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *IamObjectGroup) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *IamObjectGroup) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *IamObjectGroup) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *IamObjectGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectGroup) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectGroup) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectGroup) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectGroup) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentId

`func (o *IamObjectGroup) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *IamObjectGroup) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *IamObjectGroup) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *IamObjectGroup) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentName

`func (o *IamObjectGroup) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *IamObjectGroup) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *IamObjectGroup) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *IamObjectGroup) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetTitle

`func (o *IamObjectGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IamObjectGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IamObjectGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IamObjectGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *IamObjectGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IamObjectGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IamObjectGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IamObjectGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *IamObjectGroup) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *IamObjectGroup) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *IamObjectGroup) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *IamObjectGroup) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUsers

`func (o *IamObjectGroup) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IamObjectGroup) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IamObjectGroup) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IamObjectGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


