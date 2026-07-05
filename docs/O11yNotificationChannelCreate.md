# O11yNotificationChannelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Config** | [**O11yNotificationChannelCreateConfig**](O11yNotificationChannelCreateConfig.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewO11yNotificationChannelCreate

`func NewO11yNotificationChannelCreate(name string, type_ string, config O11yNotificationChannelCreateConfig, ) *O11yNotificationChannelCreate`

NewO11yNotificationChannelCreate instantiates a new O11yNotificationChannelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNotificationChannelCreateWithDefaults

`func NewO11yNotificationChannelCreateWithDefaults() *O11yNotificationChannelCreate`

NewO11yNotificationChannelCreateWithDefaults instantiates a new O11yNotificationChannelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *O11yNotificationChannelCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yNotificationChannelCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yNotificationChannelCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *O11yNotificationChannelCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yNotificationChannelCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yNotificationChannelCreate) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *O11yNotificationChannelCreate) GetConfig() O11yNotificationChannelCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yNotificationChannelCreate) GetConfigOk() (*O11yNotificationChannelCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yNotificationChannelCreate) SetConfig(v O11yNotificationChannelCreateConfig)`

SetConfig sets Config field to given value.


### GetEnabled

`func (o *O11yNotificationChannelCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yNotificationChannelCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yNotificationChannelCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yNotificationChannelCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


