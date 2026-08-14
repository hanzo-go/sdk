# RegistryPackageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RegistryPackage**](RegistryPackage.md) | Data is the packages in the org&#39;s scope. | [optional] 

## Methods

### NewRegistryPackageList

`func NewRegistryPackageList() *RegistryPackageList`

NewRegistryPackageList instantiates a new RegistryPackageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryPackageListWithDefaults

`func NewRegistryPackageListWithDefaults() *RegistryPackageList`

NewRegistryPackageListWithDefaults instantiates a new RegistryPackageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RegistryPackageList) GetData() []RegistryPackage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RegistryPackageList) GetDataOk() (*[]RegistryPackage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RegistryPackageList) SetData(v []RegistryPackage)`

SetData sets Data field to given value.

### HasData

`func (o *RegistryPackageList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


