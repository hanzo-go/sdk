# O11yO11yAffectedAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the asset&#39;s id. | [optional] 
**ImpactedLabels** | Pointer to **[]string** | ImpactedLabels are the rule labels the asset uses. | [optional] 
**Name** | Pointer to **string** | Name is the asset&#39;s name. | [optional] 
**Type** | Pointer to **string** | Type is dashboard or alert_rule. | [optional] 
**Widget** | Pointer to [**O11yO11yAffectedWidget**](O11yO11yAffectedWidget.md) | Widget is the affected panel, for a dashboard. | [optional] 

## Methods

### NewO11yO11yAffectedAsset

`func NewO11yO11yAffectedAsset() *O11yO11yAffectedAsset`

NewO11yO11yAffectedAsset instantiates a new O11yO11yAffectedAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAffectedAssetWithDefaults

`func NewO11yO11yAffectedAssetWithDefaults() *O11yO11yAffectedAsset`

NewO11yO11yAffectedAssetWithDefaults instantiates a new O11yO11yAffectedAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yO11yAffectedAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yAffectedAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yAffectedAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yAffectedAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpactedLabels

`func (o *O11yO11yAffectedAsset) GetImpactedLabels() []string`

GetImpactedLabels returns the ImpactedLabels field if non-nil, zero value otherwise.

### GetImpactedLabelsOk

`func (o *O11yO11yAffectedAsset) GetImpactedLabelsOk() (*[]string, bool)`

GetImpactedLabelsOk returns a tuple with the ImpactedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedLabels

`func (o *O11yO11yAffectedAsset) SetImpactedLabels(v []string)`

SetImpactedLabels sets ImpactedLabels field to given value.

### HasImpactedLabels

`func (o *O11yO11yAffectedAsset) HasImpactedLabels() bool`

HasImpactedLabels returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yAffectedAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yAffectedAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yAffectedAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yAffectedAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yAffectedAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yAffectedAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yAffectedAsset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yAffectedAsset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidget

`func (o *O11yO11yAffectedAsset) GetWidget() O11yO11yAffectedWidget`

GetWidget returns the Widget field if non-nil, zero value otherwise.

### GetWidgetOk

`func (o *O11yO11yAffectedAsset) GetWidgetOk() (*O11yO11yAffectedWidget, bool)`

GetWidgetOk returns a tuple with the Widget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidget

`func (o *O11yO11yAffectedAsset) SetWidget(v O11yO11yAffectedWidget)`

SetWidget sets Widget field to given value.

### HasWidget

`func (o *O11yO11yAffectedAsset) HasWidget() bool`

HasWidget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


