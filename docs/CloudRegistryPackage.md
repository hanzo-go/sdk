# CloudRegistryPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description says what the package is, as published. | [optional] 
**Name** | Pointer to **string** | Name is the package name (&#x60;&lt;org&gt;&#x60; or &#x60;@&lt;org&gt;/…&#x60;). | [optional] 
**Updated** | Pointer to **string** | Updated is when the package last changed, as the registry reports it. | [optional] 
**Version** | Pointer to **string** | Version is the latest published version. | [optional] 

## Methods

### NewCloudRegistryPackage

`func NewCloudRegistryPackage() *CloudRegistryPackage`

NewCloudRegistryPackage instantiates a new CloudRegistryPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegistryPackageWithDefaults

`func NewCloudRegistryPackageWithDefaults() *CloudRegistryPackage`

NewCloudRegistryPackageWithDefaults instantiates a new CloudRegistryPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudRegistryPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudRegistryPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudRegistryPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudRegistryPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CloudRegistryPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRegistryPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRegistryPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRegistryPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudRegistryPackage) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudRegistryPackage) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudRegistryPackage) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudRegistryPackage) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *CloudRegistryPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudRegistryPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudRegistryPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudRegistryPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


