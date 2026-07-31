# CloudMirrorTargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is an optional assertion of the target&#39;s hostname. The authoritative host is the one in URL; a value that disagrees with it is refused. | [optional] 
**Name** | Pointer to **string** | Name is the repo whose advanced refs are pushed downstream, from the :name path segment. | [optional] 
**Url** | Pointer to **string** | URL is the downstream https git remote. Must be https to an allowlisted host (github.com / gitlab.com); any embedded credentials are stripped. Required. | [optional] 

## Methods

### NewCloudMirrorTargetReq

`func NewCloudMirrorTargetReq() *CloudMirrorTargetReq`

NewCloudMirrorTargetReq instantiates a new CloudMirrorTargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMirrorTargetReqWithDefaults

`func NewCloudMirrorTargetReqWithDefaults() *CloudMirrorTargetReq`

NewCloudMirrorTargetReqWithDefaults instantiates a new CloudMirrorTargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CloudMirrorTargetReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudMirrorTargetReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudMirrorTargetReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudMirrorTargetReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *CloudMirrorTargetReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMirrorTargetReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMirrorTargetReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMirrorTargetReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *CloudMirrorTargetReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudMirrorTargetReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudMirrorTargetReq) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudMirrorTargetReq) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


