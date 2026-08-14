# IamAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsTriggered** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**RequestUri** | Pointer to **string** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewIamAuditLog

`func NewIamAuditLog() *IamAuditLog`

NewIamAuditLog instantiates a new IamAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAuditLogWithDefaults

`func NewIamAuditLogWithDefaults() *IamAuditLog`

NewIamAuditLogWithDefaults instantiates a new IamAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IamAuditLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IamAuditLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IamAuditLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IamAuditLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClientIp

`func (o *IamAuditLog) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *IamAuditLog) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *IamAuditLog) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *IamAuditLog) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamAuditLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamAuditLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamAuditLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamAuditLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamAuditLog) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamAuditLog) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamAuditLog) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamAuditLog) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamAuditLog) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamAuditLog) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamAuditLog) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamAuditLog) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *IamAuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamAuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamAuditLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamAuditLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsTriggered

`func (o *IamAuditLog) GetIsTriggered() bool`

GetIsTriggered returns the IsTriggered field if non-nil, zero value otherwise.

### GetIsTriggeredOk

`func (o *IamAuditLog) GetIsTriggeredOk() (*bool, bool)`

GetIsTriggeredOk returns a tuple with the IsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTriggered

`func (o *IamAuditLog) SetIsTriggered(v bool)`

SetIsTriggered sets IsTriggered field to given value.

### HasIsTriggered

`func (o *IamAuditLog) HasIsTriggered() bool`

HasIsTriggered returns a boolean if a field has been set.

### GetLanguage

`func (o *IamAuditLog) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IamAuditLog) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IamAuditLog) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IamAuditLog) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMethod

`func (o *IamAuditLog) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IamAuditLog) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IamAuditLog) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IamAuditLog) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetName

`func (o *IamAuditLog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamAuditLog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamAuditLog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamAuditLog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *IamAuditLog) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *IamAuditLog) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *IamAuditLog) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *IamAuditLog) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOrganization

`func (o *IamAuditLog) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamAuditLog) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamAuditLog) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamAuditLog) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamAuditLog) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamAuditLog) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamAuditLog) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamAuditLog) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRequestUri

`func (o *IamAuditLog) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *IamAuditLog) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *IamAuditLog) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *IamAuditLog) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetResponse

`func (o *IamAuditLog) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *IamAuditLog) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *IamAuditLog) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *IamAuditLog) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatusCode

`func (o *IamAuditLog) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *IamAuditLog) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *IamAuditLog) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *IamAuditLog) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamAuditLog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamAuditLog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamAuditLog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamAuditLog) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *IamAuditLog) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamAuditLog) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamAuditLog) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IamAuditLog) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


