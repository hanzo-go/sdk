# CloudBlueprintView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | Pointer to [**CloudBlueprint**](CloudBlueprint.md) | Blueprint is the whole authored document, including items disabled for the org-facing reads, with every enabled flag written out explicitly. | [optional] 
**Brand** | Pointer to **string** | Brand is the key this blueprint is stored under — the deployment&#39;s brand, or \&quot;\&quot; for the shared base blueprint it falls back to. | [optional] 
**Counts** | Pointer to [**CloudBlueprintCounts**](CloudBlueprintCounts.md) | Counts summarises how many items each collection holds. | [optional] 
**Version** | Pointer to **int32** | Version is the active stored version number (1 is the seed). Each edit appends a new one; nothing is ever overwritten. | [optional] 

## Methods

### NewCloudBlueprintView

`func NewCloudBlueprintView() *CloudBlueprintView`

NewCloudBlueprintView instantiates a new CloudBlueprintView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBlueprintViewWithDefaults

`func NewCloudBlueprintViewWithDefaults() *CloudBlueprintView`

NewCloudBlueprintViewWithDefaults instantiates a new CloudBlueprintView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *CloudBlueprintView) GetBlueprint() CloudBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *CloudBlueprintView) GetBlueprintOk() (*CloudBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *CloudBlueprintView) SetBlueprint(v CloudBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *CloudBlueprintView) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetBrand

`func (o *CloudBlueprintView) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CloudBlueprintView) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CloudBlueprintView) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CloudBlueprintView) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCounts

`func (o *CloudBlueprintView) GetCounts() CloudBlueprintCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *CloudBlueprintView) GetCountsOk() (*CloudBlueprintCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *CloudBlueprintView) SetCounts(v CloudBlueprintCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *CloudBlueprintView) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetVersion

`func (o *CloudBlueprintView) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudBlueprintView) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudBlueprintView) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudBlueprintView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


