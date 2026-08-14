# PreviewView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the preview application&#39;s own slug, &#x60;&lt;app&gt;-&lt;branch&gt;&#x60;. | [optional] 
**Branch** | Pointer to **string** | Branch is the branch this preview maps. | [optional] 
**Deployment** | Pointer to [**DeploymentView**](DeploymentView.md) | Deployment is the deployment the preview recorded. | [optional] 
**Url** | Pointer to **string** | URL is the preview&#39;s live HTTPS address. | [optional] 

## Methods

### NewPreviewView

`func NewPreviewView() *PreviewView`

NewPreviewView instantiates a new PreviewView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewViewWithDefaults

`func NewPreviewViewWithDefaults() *PreviewView`

NewPreviewViewWithDefaults instantiates a new PreviewView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PreviewView) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PreviewView) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PreviewView) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PreviewView) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBranch

`func (o *PreviewView) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PreviewView) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PreviewView) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PreviewView) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetDeployment

`func (o *PreviewView) GetDeployment() DeploymentView`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *PreviewView) GetDeploymentOk() (*DeploymentView, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *PreviewView) SetDeployment(v DeploymentView)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *PreviewView) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetUrl

`func (o *PreviewView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PreviewView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PreviewView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PreviewView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


