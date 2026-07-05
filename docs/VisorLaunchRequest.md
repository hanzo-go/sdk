# VisorLaunchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | Machine size slug (or use instanceType) | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** | When true, return a price quote and spend nothing | [optional] 

## Methods

### NewVisorLaunchRequest

`func NewVisorLaunchRequest() *VisorLaunchRequest`

NewVisorLaunchRequest instantiates a new VisorLaunchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorLaunchRequestWithDefaults

`func NewVisorLaunchRequestWithDefaults() *VisorLaunchRequest`

NewVisorLaunchRequestWithDefaults instantiates a new VisorLaunchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VisorLaunchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisorLaunchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisorLaunchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VisorLaunchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *VisorLaunchRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VisorLaunchRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VisorLaunchRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VisorLaunchRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetInstanceType

`func (o *VisorLaunchRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *VisorLaunchRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *VisorLaunchRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *VisorLaunchRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetRegion

`func (o *VisorLaunchRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VisorLaunchRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VisorLaunchRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VisorLaunchRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDryRun

`func (o *VisorLaunchRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *VisorLaunchRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *VisorLaunchRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *VisorLaunchRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


