# MirrorReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the local repo to mirror into, from the :name path segment. It is CREATED on first use. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope to land the repo in; empty uses the caller&#39;s own, exactly as a create would. | [optional] 
**Source** | Pointer to **string** | Source is the http(s) git URL to fetch from. The host is SSRF-guarded, and a credential is sent only if we hold one NAMED FOR that host — so a tenant-supplied URL to anywhere else fetches anonymously. | [optional] 

## Methods

### NewMirrorReq

`func NewMirrorReq() *MirrorReq`

NewMirrorReq instantiates a new MirrorReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMirrorReqWithDefaults

`func NewMirrorReqWithDefaults() *MirrorReq`

NewMirrorReqWithDefaults instantiates a new MirrorReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MirrorReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MirrorReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MirrorReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MirrorReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *MirrorReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MirrorReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MirrorReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *MirrorReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSource

`func (o *MirrorReq) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MirrorReq) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MirrorReq) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MirrorReq) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


