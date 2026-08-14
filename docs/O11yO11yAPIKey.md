# O11yO11yAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the key was minted. | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt is when the key stops working, as a unix timestamp in seconds; zero means never. | [optional] 
**Id** | Pointer to **string** | ID is the key id. | [optional] 
**LastObservedAt** | Pointer to **time.Time** | LastObservedAt is when the key was last seen authenticating. | [optional] 
**Name** | Pointer to **string** | Name is the key&#39;s name. | [optional] 
**ServiceAccountId** | Pointer to **string** | ServiceAccountID is the account the key belongs to. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewO11yO11yAPIKey

`func NewO11yO11yAPIKey() *O11yO11yAPIKey`

NewO11yO11yAPIKey instantiates a new O11yO11yAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAPIKeyWithDefaults

`func NewO11yO11yAPIKeyWithDefaults() *O11yO11yAPIKey`

NewO11yO11yAPIKeyWithDefaults instantiates a new O11yO11yAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yAPIKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yAPIKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yAPIKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yAPIKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *O11yO11yAPIKey) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *O11yO11yAPIKey) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *O11yO11yAPIKey) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *O11yO11yAPIKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastObservedAt

`func (o *O11yO11yAPIKey) GetLastObservedAt() time.Time`

GetLastObservedAt returns the LastObservedAt field if non-nil, zero value otherwise.

### GetLastObservedAtOk

`func (o *O11yO11yAPIKey) GetLastObservedAtOk() (*time.Time, bool)`

GetLastObservedAtOk returns a tuple with the LastObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedAt

`func (o *O11yO11yAPIKey) SetLastObservedAt(v time.Time)`

SetLastObservedAt sets LastObservedAt field to given value.

### HasLastObservedAt

`func (o *O11yO11yAPIKey) HasLastObservedAt() bool`

HasLastObservedAt returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yAPIKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yAPIKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yAPIKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yAPIKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *O11yO11yAPIKey) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *O11yO11yAPIKey) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *O11yO11yAPIKey) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *O11yO11yAPIKey) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yAPIKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yAPIKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yAPIKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yAPIKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


