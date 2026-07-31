# CloudMirrorReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the local repo to mirror into, from the :name path segment. It is CREATED on first use. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope to land the repo in; empty uses the caller&#39;s own, exactly as a create would. | [optional] 
**Source** | Pointer to **string** | Source is the http(s) git URL to fetch from. The host is SSRF-guarded and the shared mirror credential is only sent to allowlisted hosts. | [optional] 

## Methods

### NewCloudMirrorReq

`func NewCloudMirrorReq() *CloudMirrorReq`

NewCloudMirrorReq instantiates a new CloudMirrorReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMirrorReqWithDefaults

`func NewCloudMirrorReqWithDefaults() *CloudMirrorReq`

NewCloudMirrorReqWithDefaults instantiates a new CloudMirrorReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudMirrorReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudMirrorReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudMirrorReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudMirrorReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *CloudMirrorReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudMirrorReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudMirrorReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudMirrorReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSource

`func (o *CloudMirrorReq) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudMirrorReq) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudMirrorReq) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudMirrorReq) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


