# IamSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**ExclusiveSignin** | Pointer to **bool** | ExclusiveSignin is a transient control flag (v1 xorm:\&quot;-\&quot;): a caller sets it on a create to collapse SessionId down to the single incoming cookie instead of appending. It is never stored — a persisted session always carries it false, so orm:\&quot;-\&quot; keeps it off the column backends and omitempty keeps it out of the SQLite JSON blob. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIamSession

`func NewIamSession() *IamSession`

NewIamSession instantiates a new IamSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSessionWithDefaults

`func NewIamSessionWithDefaults() *IamSession`

NewIamSessionWithDefaults instantiates a new IamSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamSession) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamSession) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamSession) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IamSession) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamSession) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamSession) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamSession) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamSession) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamSession) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamSession) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamSession) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamSession) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetExclusiveSignin

`func (o *IamSession) GetExclusiveSignin() bool`

GetExclusiveSignin returns the ExclusiveSignin field if non-nil, zero value otherwise.

### GetExclusiveSigninOk

`func (o *IamSession) GetExclusiveSigninOk() (*bool, bool)`

GetExclusiveSigninOk returns a tuple with the ExclusiveSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveSignin

`func (o *IamSession) SetExclusiveSignin(v bool)`

SetExclusiveSignin sets ExclusiveSignin field to given value.

### HasExclusiveSignin

`func (o *IamSession) HasExclusiveSignin() bool`

HasExclusiveSignin returns a boolean if a field has been set.

### GetId

`func (o *IamSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IamSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamSession) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamSession) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamSession) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamSession) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamSession) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamSession) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSessionId

`func (o *IamSession) GetSessionId() []string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *IamSession) GetSessionIdOk() (*[]string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *IamSession) SetSessionId(v []string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *IamSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamSession) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamSession) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamSession) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamSession) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


