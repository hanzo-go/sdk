# PatrolPatrolCamera

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is the camera type. | [optional] 
**Label** | Pointer to **string** | Label is what the camera is called on screen. | [optional] 
**Name** | Pointer to **string** | Name is the camera&#39;s document name and the segment /v1/patrol/camera/{name}/stream addresses it by. | [optional] 
**Online** | Pointer to **bool** | Online says whether the camera last reported healthy. | [optional] 
**Site** | Pointer to **string** | Site is the site the camera watches. | [optional] 

## Methods

### NewPatrolPatrolCamera

`func NewPatrolPatrolCamera() *PatrolPatrolCamera`

NewPatrolPatrolCamera instantiates a new PatrolPatrolCamera object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolCameraWithDefaults

`func NewPatrolPatrolCameraWithDefaults() *PatrolPatrolCamera`

NewPatrolPatrolCameraWithDefaults instantiates a new PatrolPatrolCamera object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PatrolPatrolCamera) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatrolPatrolCamera) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatrolPatrolCamera) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatrolPatrolCamera) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *PatrolPatrolCamera) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatrolPatrolCamera) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatrolPatrolCamera) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatrolPatrolCamera) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolCamera) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolCamera) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolCamera) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolCamera) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *PatrolPatrolCamera) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *PatrolPatrolCamera) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *PatrolPatrolCamera) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *PatrolPatrolCamera) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetSite

`func (o *PatrolPatrolCamera) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatrolPatrolCamera) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatrolPatrolCamera) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatrolPatrolCamera) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


