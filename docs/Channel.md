# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | the social integration id to target in a post | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | \&quot;x\&quot; | \&quot;instagram\&quot; | \&quot;tiktok\&quot; | ... | [optional] 

## Methods

### NewChannel

`func NewChannel() *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *Channel) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Channel) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Channel) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Channel) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *Channel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Channel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Channel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Channel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Channel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *Channel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Channel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Channel) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Channel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


