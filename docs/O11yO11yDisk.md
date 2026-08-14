# O11yO11yDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the disk&#39;s name. | [optional] 
**Type** | Pointer to **string** | Type is the disk&#39;s type, e.g. local or s3. | [optional] 

## Methods

### NewO11yO11yDisk

`func NewO11yO11yDisk() *O11yO11yDisk`

NewO11yO11yDisk instantiates a new O11yO11yDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDiskWithDefaults

`func NewO11yO11yDiskWithDefaults() *O11yO11yDisk`

NewO11yO11yDiskWithDefaults instantiates a new O11yO11yDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *O11yO11yDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yDisk) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


