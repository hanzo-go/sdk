# PreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the parent application&#39;s slug, from the path. | [optional] 
**Branch** | Pointer to **string** | Branch is the branch to preview; defaults to the parent app&#39;s branch. | [optional] 
**Image** | Pointer to **string** | Image is the already-built image ref to deploy. Required — a preview never builds. | [optional] 
**Project** | Pointer to **string** | Project is the project the parent application lives under, from the path. | [optional] 

## Methods

### NewPreviewReq

`func NewPreviewReq() *PreviewReq`

NewPreviewReq instantiates a new PreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewReqWithDefaults

`func NewPreviewReqWithDefaults() *PreviewReq`

NewPreviewReqWithDefaults instantiates a new PreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PreviewReq) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PreviewReq) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PreviewReq) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PreviewReq) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBranch

`func (o *PreviewReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PreviewReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PreviewReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PreviewReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetImage

`func (o *PreviewReq) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PreviewReq) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PreviewReq) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PreviewReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetProject

`func (o *PreviewReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PreviewReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PreviewReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PreviewReq) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


