# KmsAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to [**KmsAuditLogActor**](KmsAuditLogActor.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Event** | Pointer to [**KmsAuditLogActor**](KmsAuditLogActor.md) |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsAuditLog

`func NewKmsAuditLog() *KmsAuditLog`

NewKmsAuditLog instantiates a new KmsAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsAuditLogWithDefaults

`func NewKmsAuditLogWithDefaults() *KmsAuditLog`

NewKmsAuditLogWithDefaults instantiates a new KmsAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsAuditLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsAuditLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsAuditLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsAuditLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActor

`func (o *KmsAuditLog) GetActor() KmsAuditLogActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *KmsAuditLog) GetActorOk() (*KmsAuditLogActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *KmsAuditLog) SetActor(v KmsAuditLogActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *KmsAuditLog) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetOrgId

`func (o *KmsAuditLog) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *KmsAuditLog) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *KmsAuditLog) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *KmsAuditLog) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *KmsAuditLog) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KmsAuditLog) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KmsAuditLog) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *KmsAuditLog) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetEvent

`func (o *KmsAuditLog) GetEvent() KmsAuditLogActor`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *KmsAuditLog) GetEventOk() (*KmsAuditLogActor, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *KmsAuditLog) SetEvent(v KmsAuditLogActor)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *KmsAuditLog) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetIpAddress

`func (o *KmsAuditLog) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *KmsAuditLog) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *KmsAuditLog) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *KmsAuditLog) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *KmsAuditLog) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *KmsAuditLog) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *KmsAuditLog) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *KmsAuditLog) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsAuditLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsAuditLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsAuditLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsAuditLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


