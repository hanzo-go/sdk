# O11yNotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**O11yNotificationChannelConfig**](O11yNotificationChannelConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yNotificationChannel

`func NewO11yNotificationChannel() *O11yNotificationChannel`

NewO11yNotificationChannel instantiates a new O11yNotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNotificationChannelWithDefaults

`func NewO11yNotificationChannelWithDefaults() *O11yNotificationChannel`

NewO11yNotificationChannelWithDefaults instantiates a new O11yNotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yNotificationChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yNotificationChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yNotificationChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yNotificationChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yNotificationChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yNotificationChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yNotificationChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yNotificationChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *O11yNotificationChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yNotificationChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yNotificationChannel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yNotificationChannel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *O11yNotificationChannel) GetConfig() O11yNotificationChannelConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yNotificationChannel) GetConfigOk() (*O11yNotificationChannelConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yNotificationChannel) SetConfig(v O11yNotificationChannelConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yNotificationChannel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yNotificationChannel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yNotificationChannel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yNotificationChannel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yNotificationChannel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yNotificationChannel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yNotificationChannel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yNotificationChannel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yNotificationChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yNotificationChannel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yNotificationChannel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yNotificationChannel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yNotificationChannel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


