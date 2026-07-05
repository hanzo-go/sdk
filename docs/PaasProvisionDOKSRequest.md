# PaasProvisionDOKSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Name** | **string** |  | 
**Region** | **string** |  | 
**Version** | **string** |  | 
**Ha** | Pointer to **bool** |  | [optional] [default to false]
**NodePool** | [**PaasProvisionDOKSRequestNodePool**](PaasProvisionDOKSRequestNodePool.md) |  | 

## Methods

### NewPaasProvisionDOKSRequest

`func NewPaasProvisionDOKSRequest(orgId string, name string, region string, version string, nodePool PaasProvisionDOKSRequestNodePool, ) *PaasProvisionDOKSRequest`

NewPaasProvisionDOKSRequest instantiates a new PaasProvisionDOKSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasProvisionDOKSRequestWithDefaults

`func NewPaasProvisionDOKSRequestWithDefaults() *PaasProvisionDOKSRequest`

NewPaasProvisionDOKSRequestWithDefaults instantiates a new PaasProvisionDOKSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *PaasProvisionDOKSRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PaasProvisionDOKSRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PaasProvisionDOKSRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetName

`func (o *PaasProvisionDOKSRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasProvisionDOKSRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasProvisionDOKSRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *PaasProvisionDOKSRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PaasProvisionDOKSRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PaasProvisionDOKSRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVersion

`func (o *PaasProvisionDOKSRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PaasProvisionDOKSRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PaasProvisionDOKSRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetHa

`func (o *PaasProvisionDOKSRequest) GetHa() bool`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *PaasProvisionDOKSRequest) GetHaOk() (*bool, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *PaasProvisionDOKSRequest) SetHa(v bool)`

SetHa sets Ha field to given value.

### HasHa

`func (o *PaasProvisionDOKSRequest) HasHa() bool`

HasHa returns a boolean if a field has been set.

### GetNodePool

`func (o *PaasProvisionDOKSRequest) GetNodePool() PaasProvisionDOKSRequestNodePool`

GetNodePool returns the NodePool field if non-nil, zero value otherwise.

### GetNodePoolOk

`func (o *PaasProvisionDOKSRequest) GetNodePoolOk() (*PaasProvisionDOKSRequestNodePool, bool)`

GetNodePoolOk returns a tuple with the NodePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePool

`func (o *PaasProvisionDOKSRequest) SetNodePool(v PaasProvisionDOKSRequestNodePool)`

SetNodePool sets NodePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


