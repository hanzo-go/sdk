# NotifyNotifyStored

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the credential that was set. | [optional] 
**Provider** | Pointer to **string** | Provider is the provider the credential was set for. | [optional] 
**Stored** | Pointer to **bool** | Stored is true: the value is sealed and the next send reads it. | [optional] 

## Methods

### NewNotifyNotifyStored

`func NewNotifyNotifyStored() *NotifyNotifyStored`

NewNotifyNotifyStored instantiates a new NotifyNotifyStored object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyNotifyStoredWithDefaults

`func NewNotifyNotifyStoredWithDefaults() *NotifyNotifyStored`

NewNotifyNotifyStoredWithDefaults instantiates a new NotifyNotifyStored object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *NotifyNotifyStored) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NotifyNotifyStored) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NotifyNotifyStored) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *NotifyNotifyStored) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetProvider

`func (o *NotifyNotifyStored) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotifyNotifyStored) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotifyNotifyStored) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotifyNotifyStored) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStored

`func (o *NotifyNotifyStored) GetStored() bool`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *NotifyNotifyStored) GetStoredOk() (*bool, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *NotifyNotifyStored) SetStored(v bool)`

SetStored sets Stored field to given value.

### HasStored

`func (o *NotifyNotifyStored) HasStored() bool`

HasStored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


