# CreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is a free-form blurb, max 4KiB. | [optional] 
**Name** | Pointer to **string** | Name is the repo&#39;s handle, unique within the scope, and the last segment of both clone URLs. Must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$; a trailing \&quot;.git\&quot; is stripped first. Required. | [optional] 
**Project** | Pointer to **string** | Project narrows the repo to a sub-scope of the org. Omit it to use the caller&#39;s own X-Project-Id scope; it can never widen past the caller&#39;s org. | [optional] 
**Public** | Pointer to **bool** | Public grants ANONYMOUS read (fetch) only; push and the whole control plane stay org-authed. Defaults to false. | [optional] 

## Methods

### NewCreateReq

`func NewCreateReq() *CreateReq`

NewCreateReq instantiates a new CreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReqWithDefaults

`func NewCreateReqWithDefaults() *CreateReq`

NewCreateReqWithDefaults instantiates a new CreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *CreateReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublic

`func (o *CreateReq) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreateReq) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreateReq) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreateReq) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


