# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Disabled is true for a channel the org switched off at the social edge. It is still listed — this is what the org has CONNECTED, not what it can post to — but a publish never targets it, neither by name nor as part of the \&quot;every channel\&quot; default. | [optional] 
**Id** | Pointer to **string** | ID is the social integration id a post targets. It is the exact value to put in a content item&#39;s &#x60;channels&#x60; list to reach this one connected account. | [optional] 
**Name** | Pointer to **string** | Name is the account label as the org connected it — the handle a human recognises. It is never an address: a publish resolves channels by ID or by Provider and never by this. | [optional] 
**Provider** | Pointer to **string** | Provider is the network behind the integration: \&quot;x\&quot;, \&quot;instagram\&quot;, \&quot;tiktok\&quot; and the rest of what the org connected. Naming a provider in a publish targets EVERY connected account of it, so it is the coarse handle where ID is the precise one. | [optional] 

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


