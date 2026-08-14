# BlueprintVersionsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Brand is the blueprint key the history belongs to — this deployment&#39;s brand, or \&quot;\&quot; (the base blueprint) when the brand has no row of its own. | [optional] 
**Versions** | Pointer to [**[]VersionMeta**](VersionMeta.md) | Versions are the stored versions, newest first: metadata only, never the documents. | [optional] 

## Methods

### NewBlueprintVersionsView

`func NewBlueprintVersionsView() *BlueprintVersionsView`

NewBlueprintVersionsView instantiates a new BlueprintVersionsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintVersionsViewWithDefaults

`func NewBlueprintVersionsViewWithDefaults() *BlueprintVersionsView`

NewBlueprintVersionsViewWithDefaults instantiates a new BlueprintVersionsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *BlueprintVersionsView) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *BlueprintVersionsView) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *BlueprintVersionsView) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *BlueprintVersionsView) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetVersions

`func (o *BlueprintVersionsView) GetVersions() []VersionMeta`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BlueprintVersionsView) GetVersionsOk() (*[]VersionMeta, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BlueprintVersionsView) SetVersions(v []VersionMeta)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *BlueprintVersionsView) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


