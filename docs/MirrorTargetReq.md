# MirrorTargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is an optional assertion of the target&#39;s hostname. The authoritative host is the one in URL; a value that disagrees with it is refused. | [optional] 
**Name** | Pointer to **string** | Name is the repo whose advanced refs are pushed downstream, from the :name path segment. | [optional] 
**Url** | Pointer to **string** | URL is the downstream https git remote. Must be https to an allowlisted host (github.com / gitlab.com); any embedded credentials are stripped. Required. | [optional] 

## Methods

### NewMirrorTargetReq

`func NewMirrorTargetReq() *MirrorTargetReq`

NewMirrorTargetReq instantiates a new MirrorTargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirrorTargetReqWithDefaults

`func NewMirrorTargetReqWithDefaults() *MirrorTargetReq`

NewMirrorTargetReqWithDefaults instantiates a new MirrorTargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MirrorTargetReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MirrorTargetReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MirrorTargetReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MirrorTargetReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *MirrorTargetReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MirrorTargetReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MirrorTargetReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MirrorTargetReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *MirrorTargetReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MirrorTargetReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MirrorTargetReq) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MirrorTargetReq) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


