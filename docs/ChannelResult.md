# ChannelResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | the social integration id targeted | [optional] 
**Error** | Pointer to **string** | short reason, when it failed | [optional] 
**ExternalId** | Pointer to **string** | social post id, when it went out | [optional] 
**Provider** | Pointer to **string** | \&quot;x\&quot; | \&quot;instagram\&quot; | ... when known | [optional] 
**Status** | Pointer to **string** | \&quot;distributed\&quot; | \&quot;scheduled\&quot; | \&quot;failed\&quot; | [optional] 

## Methods

### NewChannelResult

`func NewChannelResult() *ChannelResult`

NewChannelResult instantiates a new ChannelResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelResultWithDefaults

`func NewChannelResultWithDefaults() *ChannelResult`

NewChannelResultWithDefaults instantiates a new ChannelResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelResult) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelResult) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelResult) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelResult) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetError

`func (o *ChannelResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChannelResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChannelResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ChannelResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExternalId

`func (o *ChannelResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ChannelResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ChannelResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ChannelResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvider

`func (o *ChannelResult) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ChannelResult) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ChannelResult) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ChannelResult) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ChannelResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChannelResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChannelResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChannelResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


