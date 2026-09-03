# IamInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DefaultCode** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsRegexp** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Quota** | Pointer to **int64** |  | [optional] 
**SignupGroup** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**UsedCount** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewIamInvitation

`func NewIamInvitation() *IamInvitation`

NewIamInvitation instantiates a new IamInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamInvitationWithDefaults

`func NewIamInvitationWithDefaults() *IamInvitation`

NewIamInvitationWithDefaults instantiates a new IamInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamInvitation) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamInvitation) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamInvitation) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamInvitation) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCode

`func (o *IamInvitation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamInvitation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamInvitation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamInvitation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamInvitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamInvitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamInvitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamInvitation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamInvitation) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamInvitation) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamInvitation) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamInvitation) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDefaultCode

`func (o *IamInvitation) GetDefaultCode() string`

GetDefaultCode returns the DefaultCode field if non-nil, zero value otherwise.

### GetDefaultCodeOk

`func (o *IamInvitation) GetDefaultCodeOk() (*string, bool)`

GetDefaultCodeOk returns a tuple with the DefaultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCode

`func (o *IamInvitation) SetDefaultCode(v string)`

SetDefaultCode sets DefaultCode field to given value.

### HasDefaultCode

`func (o *IamInvitation) HasDefaultCode() bool`

HasDefaultCode returns a boolean if a field has been set.

### GetDeleted

`func (o *IamInvitation) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamInvitation) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamInvitation) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamInvitation) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamInvitation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamInvitation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamInvitation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamInvitation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *IamInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IamInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *IamInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRegexp

`func (o *IamInvitation) GetIsRegexp() bool`

GetIsRegexp returns the IsRegexp field if non-nil, zero value otherwise.

### GetIsRegexpOk

`func (o *IamInvitation) GetIsRegexpOk() (*bool, bool)`

GetIsRegexpOk returns a tuple with the IsRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegexp

`func (o *IamInvitation) SetIsRegexp(v bool)`

SetIsRegexp sets IsRegexp field to given value.

### HasIsRegexp

`func (o *IamInvitation) HasIsRegexp() bool`

HasIsRegexp returns a boolean if a field has been set.

### GetName

`func (o *IamInvitation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamInvitation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamInvitation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamInvitation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamInvitation) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamInvitation) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamInvitation) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamInvitation) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPhone

`func (o *IamInvitation) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *IamInvitation) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *IamInvitation) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *IamInvitation) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetQuota

`func (o *IamInvitation) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *IamInvitation) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *IamInvitation) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *IamInvitation) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetSignupGroup

`func (o *IamInvitation) GetSignupGroup() string`

GetSignupGroup returns the SignupGroup field if non-nil, zero value otherwise.

### GetSignupGroupOk

`func (o *IamInvitation) GetSignupGroupOk() (*string, bool)`

GetSignupGroupOk returns a tuple with the SignupGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupGroup

`func (o *IamInvitation) SetSignupGroup(v string)`

SetSignupGroup sets SignupGroup field to given value.

### HasSignupGroup

`func (o *IamInvitation) HasSignupGroup() bool`

HasSignupGroup returns a boolean if a field has been set.

### GetState

`func (o *IamInvitation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamInvitation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamInvitation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IamInvitation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamInvitation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamInvitation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamInvitation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamInvitation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *IamInvitation) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *IamInvitation) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *IamInvitation) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *IamInvitation) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUsedCount

`func (o *IamInvitation) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *IamInvitation) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *IamInvitation) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *IamInvitation) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetUsername

`func (o *IamInvitation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamInvitation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamInvitation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamInvitation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


