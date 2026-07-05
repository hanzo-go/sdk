# DnsListZones200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zones** | Pointer to [**[]DnsZone**](DnsZone.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnsListZones200Response

`func NewDnsListZones200Response() *DnsListZones200Response`

NewDnsListZones200Response instantiates a new DnsListZones200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsListZones200ResponseWithDefaults

`func NewDnsListZones200ResponseWithDefaults() *DnsListZones200Response`

NewDnsListZones200ResponseWithDefaults instantiates a new DnsListZones200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZones

`func (o *DnsListZones200Response) GetZones() []DnsZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *DnsListZones200Response) GetZonesOk() (*[]DnsZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *DnsListZones200Response) SetZones(v []DnsZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *DnsListZones200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetTotal

`func (o *DnsListZones200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DnsListZones200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DnsListZones200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DnsListZones200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


