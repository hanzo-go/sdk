# PlatformPreviewView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the preview application&#39;s own slug, &#x60;&lt;app&gt;-&lt;branch&gt;&#x60;. | [optional] 
**Branch** | Pointer to **string** | Branch is the branch this preview maps. | [optional] 
**Deployment** | Pointer to [**PlatformDeploymentView**](PlatformDeploymentView.md) | Deployment is the deployment the preview recorded. | [optional] 
**Url** | Pointer to **string** | URL is the preview&#39;s live HTTPS address. | [optional] 

## Methods

### NewPlatformPreviewView

`func NewPlatformPreviewView() *PlatformPreviewView`

NewPlatformPreviewView instantiates a new PlatformPreviewView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformPreviewViewWithDefaults

`func NewPlatformPreviewViewWithDefaults() *PlatformPreviewView`

NewPlatformPreviewViewWithDefaults instantiates a new PlatformPreviewView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PlatformPreviewView) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PlatformPreviewView) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PlatformPreviewView) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PlatformPreviewView) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBranch

`func (o *PlatformPreviewView) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PlatformPreviewView) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PlatformPreviewView) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PlatformPreviewView) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetDeployment

`func (o *PlatformPreviewView) GetDeployment() PlatformDeploymentView`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *PlatformPreviewView) GetDeploymentOk() (*PlatformDeploymentView, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *PlatformPreviewView) SetDeployment(v PlatformDeploymentView)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *PlatformPreviewView) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetUrl

`func (o *PlatformPreviewView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlatformPreviewView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlatformPreviewView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PlatformPreviewView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


