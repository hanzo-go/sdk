# BlueprintView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | Pointer to [**Blueprint**](Blueprint.md) | Blueprint is the whole authored document, including items disabled for the org-facing reads, with every enabled flag written out explicitly. | [optional] 
**Brand** | Pointer to **string** | Brand is the key this blueprint is stored under — the deployment&#39;s brand, or \&quot;\&quot; for the shared base blueprint it falls back to. | [optional] 
**Counts** | Pointer to [**BlueprintCounts**](BlueprintCounts.md) | Counts summarises how many items each collection holds. | [optional] 
**Version** | Pointer to **int32** | Version is the active stored version number (1 is the seed). Each edit appends a new one; nothing is ever overwritten. | [optional] 

## Methods

### NewBlueprintView

`func NewBlueprintView() *BlueprintView`

NewBlueprintView instantiates a new BlueprintView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintViewWithDefaults

`func NewBlueprintViewWithDefaults() *BlueprintView`

NewBlueprintViewWithDefaults instantiates a new BlueprintView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *BlueprintView) GetBlueprint() Blueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *BlueprintView) GetBlueprintOk() (*Blueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *BlueprintView) SetBlueprint(v Blueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *BlueprintView) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetBrand

`func (o *BlueprintView) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *BlueprintView) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *BlueprintView) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *BlueprintView) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCounts

`func (o *BlueprintView) GetCounts() BlueprintCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *BlueprintView) GetCountsOk() (*BlueprintCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *BlueprintView) SetCounts(v BlueprintCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *BlueprintView) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetVersion

`func (o *BlueprintView) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintView) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintView) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlueprintView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


