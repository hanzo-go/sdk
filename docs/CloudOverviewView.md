# CloudOverviewView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **bool** |  | [optional] 
**Funnel** | Pointer to [**CloudFunnel**](CloudFunnel.md) |  | [optional] 
**Progress** | Pointer to [**CloudProgressView**](CloudProgressView.md) |  | [optional] 
**Steps** | Pointer to [**[]CloudStepView**](CloudStepView.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudOverviewView

`func NewCloudOverviewView() *CloudOverviewView`

NewCloudOverviewView instantiates a new CloudOverviewView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOverviewViewWithDefaults

`func NewCloudOverviewViewWithDefaults() *CloudOverviewView`

NewCloudOverviewViewWithDefaults instantiates a new CloudOverviewView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *CloudOverviewView) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CloudOverviewView) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CloudOverviewView) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *CloudOverviewView) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetFunnel

`func (o *CloudOverviewView) GetFunnel() CloudFunnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *CloudOverviewView) GetFunnelOk() (*CloudFunnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *CloudOverviewView) SetFunnel(v CloudFunnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *CloudOverviewView) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetProgress

`func (o *CloudOverviewView) GetProgress() CloudProgressView`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *CloudOverviewView) GetProgressOk() (*CloudProgressView, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *CloudOverviewView) SetProgress(v CloudProgressView)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *CloudOverviewView) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSteps

`func (o *CloudOverviewView) GetSteps() []CloudStepView`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CloudOverviewView) GetStepsOk() (*[]CloudStepView, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CloudOverviewView) SetSteps(v []CloudStepView)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CloudOverviewView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTitle

`func (o *CloudOverviewView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudOverviewView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudOverviewView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudOverviewView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *CloudOverviewView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudOverviewView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudOverviewView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudOverviewView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


