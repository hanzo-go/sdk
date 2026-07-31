# CloudDeviceStartIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label names this connection so one user can hold several per provider (\&quot;work\&quot;, \&quot;personal\&quot;). Empty means \&quot;default\&quot;. 1-64 of [A-Za-z0-9._-]. | [optional] 
**Provider** | Pointer to **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | [optional] 

## Methods

### NewCloudDeviceStartIn

`func NewCloudDeviceStartIn() *CloudDeviceStartIn`

NewCloudDeviceStartIn instantiates a new CloudDeviceStartIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeviceStartInWithDefaults

`func NewCloudDeviceStartInWithDefaults() *CloudDeviceStartIn`

NewCloudDeviceStartInWithDefaults instantiates a new CloudDeviceStartIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CloudDeviceStartIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudDeviceStartIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudDeviceStartIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudDeviceStartIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProvider

`func (o *CloudDeviceStartIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudDeviceStartIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudDeviceStartIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudDeviceStartIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


