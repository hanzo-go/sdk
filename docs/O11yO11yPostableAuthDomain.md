# O11yO11yPostableAuthDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**O11yO11yAuthDomainConfig**](O11yO11yAuthDomainConfig.md) | Config is the domain&#39;s SSO configuration. | [optional] 
**Name** | Pointer to **string** | Name is the email domain being claimed, e.g. example.com. | [optional] 

## Methods

### NewO11yO11yPostableAuthDomain

`func NewO11yO11yPostableAuthDomain() *O11yO11yPostableAuthDomain`

NewO11yO11yPostableAuthDomain instantiates a new O11yO11yPostableAuthDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPostableAuthDomainWithDefaults

`func NewO11yO11yPostableAuthDomainWithDefaults() *O11yO11yPostableAuthDomain`

NewO11yO11yPostableAuthDomainWithDefaults instantiates a new O11yO11yPostableAuthDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *O11yO11yPostableAuthDomain) GetConfig() O11yO11yAuthDomainConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yO11yPostableAuthDomain) GetConfigOk() (*O11yO11yAuthDomainConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yO11yPostableAuthDomain) SetConfig(v O11yO11yAuthDomainConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yO11yPostableAuthDomain) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yPostableAuthDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yPostableAuthDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yPostableAuthDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yPostableAuthDomain) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


