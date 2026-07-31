# CloudProjectPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the board&#39;s free-form blurb, at most 32768 characters. | [optional] 
**Key** | Pointer to **string** | Key is the project to update, from the path. | [optional] 
**Name** | Pointer to **string** | Name is the project&#39;s display name. Non-empty, at most 256 characters. | [optional] 

## Methods

### NewCloudProjectPatch

`func NewCloudProjectPatch() *CloudProjectPatch`

NewCloudProjectPatch instantiates a new CloudProjectPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProjectPatchWithDefaults

`func NewCloudProjectPatchWithDefaults() *CloudProjectPatch`

NewCloudProjectPatchWithDefaults instantiates a new CloudProjectPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudProjectPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudProjectPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudProjectPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudProjectPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *CloudProjectPatch) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudProjectPatch) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudProjectPatch) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudProjectPatch) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CloudProjectPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProjectPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProjectPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProjectPatch) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


