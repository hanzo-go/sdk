# FlagActivityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is one of created, updated, deleted. | [optional] 
**Actor** | Pointer to **string** | Actor is the email of the principal who made the change. Empty for a write by an in-process composer; a project key can never appear here, because evaluating flags is all a key may do. | [optional] 
**At** | Pointer to **string** | At is when the change was made, RFC 3339 UTC. | [optional] 
**Detail** | Pointer to **string** | Detail is free-form context about the change. Nothing writes it today, so it is absent from every row the store serves. | [optional] 
**Id** | Pointer to **int64** | ID is the log&#39;s own sequence number, rising with each entry. The log is served newest-first, which is this descending. | [optional] 
**Key** | Pointer to **string** | Key is the flag that changed. It survives a delete, so the log still names flags the definition store no longer holds. | [optional] 

## Methods

### NewFlagActivityRow

`func NewFlagActivityRow() *FlagActivityRow`

NewFlagActivityRow instantiates a new FlagActivityRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagActivityRowWithDefaults

`func NewFlagActivityRowWithDefaults() *FlagActivityRow`

NewFlagActivityRowWithDefaults instantiates a new FlagActivityRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FlagActivityRow) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FlagActivityRow) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FlagActivityRow) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FlagActivityRow) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActor

`func (o *FlagActivityRow) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *FlagActivityRow) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *FlagActivityRow) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *FlagActivityRow) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAt

`func (o *FlagActivityRow) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *FlagActivityRow) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *FlagActivityRow) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *FlagActivityRow) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDetail

`func (o *FlagActivityRow) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *FlagActivityRow) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *FlagActivityRow) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *FlagActivityRow) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetId

`func (o *FlagActivityRow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagActivityRow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagActivityRow) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FlagActivityRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *FlagActivityRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagActivityRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagActivityRow) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlagActivityRow) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


