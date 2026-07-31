# CloudVersionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildDate** | Pointer to **string** | BuildDate is the time THIS RESPONSE was generated, in RFC 3339 — not a build timestamp. There is no argocd build here to report one for. | [optional] 
**Compiler** | Pointer to **string** | Compiler is the constant \&quot;gc\&quot; the SPA expects; it is not read from this process. | [optional] 
**GoVersion** | Pointer to **string** | GoVersion is always empty. | [optional] 
**Platform** | Pointer to **string** | Platform is the constant \&quot;linux/amd64\&quot; the SPA expects; it is not this process&#39;s own GOOS/GOARCH. | [optional] 
**Version** | Pointer to **string** | Version names the projection, \&quot;hanzo-cd (projection)\&quot;. | [optional] 

## Methods

### NewCloudVersionMessage

`func NewCloudVersionMessage() *CloudVersionMessage`

NewCloudVersionMessage instantiates a new CloudVersionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVersionMessageWithDefaults

`func NewCloudVersionMessageWithDefaults() *CloudVersionMessage`

NewCloudVersionMessageWithDefaults instantiates a new CloudVersionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildDate

`func (o *CloudVersionMessage) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *CloudVersionMessage) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *CloudVersionMessage) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *CloudVersionMessage) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetCompiler

`func (o *CloudVersionMessage) GetCompiler() string`

GetCompiler returns the Compiler field if non-nil, zero value otherwise.

### GetCompilerOk

`func (o *CloudVersionMessage) GetCompilerOk() (*string, bool)`

GetCompilerOk returns a tuple with the Compiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiler

`func (o *CloudVersionMessage) SetCompiler(v string)`

SetCompiler sets Compiler field to given value.

### HasCompiler

`func (o *CloudVersionMessage) HasCompiler() bool`

HasCompiler returns a boolean if a field has been set.

### GetGoVersion

`func (o *CloudVersionMessage) GetGoVersion() string`

GetGoVersion returns the GoVersion field if non-nil, zero value otherwise.

### GetGoVersionOk

`func (o *CloudVersionMessage) GetGoVersionOk() (*string, bool)`

GetGoVersionOk returns a tuple with the GoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoVersion

`func (o *CloudVersionMessage) SetGoVersion(v string)`

SetGoVersion sets GoVersion field to given value.

### HasGoVersion

`func (o *CloudVersionMessage) HasGoVersion() bool`

HasGoVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudVersionMessage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudVersionMessage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudVersionMessage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudVersionMessage) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetVersion

`func (o *CloudVersionMessage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudVersionMessage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudVersionMessage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudVersionMessage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


