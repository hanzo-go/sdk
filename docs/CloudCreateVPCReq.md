# CloudCreateVPCReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRange** | Pointer to **string** | IPRange is the VPC&#39;s private CIDR. Empty lets DigitalOcean assign one. | [optional] 
**Name** | Pointer to **string** | Name is the FRIENDLY name, a DNS-safe slug of at most 40 characters. The physical DigitalOcean name is derived from it and the caller&#39;s org. | [optional] 
**Region** | Pointer to **string** | Region is the DigitalOcean region slug (nyc3, sfo3, …). Required. | [optional] 

## Methods

### NewCloudCreateVPCReq

`func NewCloudCreateVPCReq() *CloudCreateVPCReq`

NewCloudCreateVPCReq instantiates a new CloudCreateVPCReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateVPCReqWithDefaults

`func NewCloudCreateVPCReqWithDefaults() *CloudCreateVPCReq`

NewCloudCreateVPCReqWithDefaults instantiates a new CloudCreateVPCReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRange

`func (o *CloudCreateVPCReq) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *CloudCreateVPCReq) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *CloudCreateVPCReq) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *CloudCreateVPCReq) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetName

`func (o *CloudCreateVPCReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCreateVPCReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCreateVPCReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCreateVPCReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CloudCreateVPCReq) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudCreateVPCReq) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudCreateVPCReq) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudCreateVPCReq) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


