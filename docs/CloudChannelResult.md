# CloudChannelResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | the social integration id targeted | [optional] 
**Error** | Pointer to **string** | short reason, when it failed | [optional] 
**ExternalId** | Pointer to **string** | social post id, when it went out | [optional] 
**Provider** | Pointer to **string** | \&quot;x\&quot; | \&quot;instagram\&quot; | ... when known | [optional] 
**Status** | Pointer to **string** | \&quot;distributed\&quot; | \&quot;scheduled\&quot; | \&quot;failed\&quot; | [optional] 

## Methods

### NewCloudChannelResult

`func NewCloudChannelResult() *CloudChannelResult`

NewCloudChannelResult instantiates a new CloudChannelResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudChannelResultWithDefaults

`func NewCloudChannelResultWithDefaults() *CloudChannelResult`

NewCloudChannelResultWithDefaults instantiates a new CloudChannelResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CloudChannelResult) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudChannelResult) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudChannelResult) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudChannelResult) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetError

`func (o *CloudChannelResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudChannelResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudChannelResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudChannelResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudChannelResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudChannelResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudChannelResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudChannelResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvider

`func (o *CloudChannelResult) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudChannelResult) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudChannelResult) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudChannelResult) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *CloudChannelResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudChannelResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudChannelResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudChannelResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


