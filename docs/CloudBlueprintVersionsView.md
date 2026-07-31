# CloudBlueprintVersionsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand is the blueprint key the history belongs to — this deployment&#39;s brand, or \&quot;\&quot; (the base blueprint) when the brand has no row of its own. | [optional] 
**Versions** | Pointer to [**[]CloudVersionMeta**](CloudVersionMeta.md) | Versions are the stored versions, newest first: metadata only, never the documents. | [optional] 

## Methods

### NewCloudBlueprintVersionsView

`func NewCloudBlueprintVersionsView() *CloudBlueprintVersionsView`

NewCloudBlueprintVersionsView instantiates a new CloudBlueprintVersionsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBlueprintVersionsViewWithDefaults

`func NewCloudBlueprintVersionsViewWithDefaults() *CloudBlueprintVersionsView`

NewCloudBlueprintVersionsViewWithDefaults instantiates a new CloudBlueprintVersionsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *CloudBlueprintVersionsView) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CloudBlueprintVersionsView) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CloudBlueprintVersionsView) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CloudBlueprintVersionsView) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetVersions

`func (o *CloudBlueprintVersionsView) GetVersions() []CloudVersionMeta`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CloudBlueprintVersionsView) GetVersionsOk() (*[]CloudVersionMeta, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CloudBlueprintVersionsView) SetVersions(v []CloudVersionMeta)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CloudBlueprintVersionsView) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


