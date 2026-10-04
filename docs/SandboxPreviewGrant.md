# SandboxPreviewGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int64** | ExpiresIn is how long the ticket in URL is good for, in seconds. | [optional] 
**Host** | Pointer to **string** | Host is the preview&#39;s origin host, the same for every ticket for this port of this sandbox, before and after the sandbox is parked and resumed. | [optional] 
**Port** | Pointer to **int64** | Port is the port the preview serves. | [optional] 
**Url** | Pointer to **string** | URL opens the preview: its own origin, with a single-use ticket that sets the preview&#39;s cookie and redirects to its root. Open it in a frame or a tab within ExpiresIn seconds, and mint another to reopen. | [optional] 

## Methods

### NewSandboxPreviewGrant

`func NewSandboxPreviewGrant() *SandboxPreviewGrant`

NewSandboxPreviewGrant instantiates a new SandboxPreviewGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPreviewGrantWithDefaults

`func NewSandboxPreviewGrantWithDefaults() *SandboxPreviewGrant`

NewSandboxPreviewGrantWithDefaults instantiates a new SandboxPreviewGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *SandboxPreviewGrant) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *SandboxPreviewGrant) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *SandboxPreviewGrant) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *SandboxPreviewGrant) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetHost

`func (o *SandboxPreviewGrant) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SandboxPreviewGrant) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SandboxPreviewGrant) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SandboxPreviewGrant) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *SandboxPreviewGrant) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SandboxPreviewGrant) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SandboxPreviewGrant) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SandboxPreviewGrant) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUrl

`func (o *SandboxPreviewGrant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SandboxPreviewGrant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SandboxPreviewGrant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SandboxPreviewGrant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


