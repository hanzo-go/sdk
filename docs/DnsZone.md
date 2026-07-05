# DnsZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** | Zone name (e.g. example.com.) | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Nameservers** | Pointer to **[]string** | Assigned authoritative nameservers | [optional] 
**Soa** | Pointer to [**DnsSOARecord**](DnsSOARecord.md) |  | [optional] 
**RecordCount** | Pointer to **int32** |  | [optional] 
**DnssecEnabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDnsZone

`func NewDnsZone() *DnsZone`

NewDnsZone instantiates a new DnsZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneWithDefaults

`func NewDnsZoneWithDefaults() *DnsZone`

NewDnsZoneWithDefaults instantiates a new DnsZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetZone

`func (o *DnsZone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DnsZone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DnsZone) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *DnsZone) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetStatus

`func (o *DnsZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DnsZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DnsZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DnsZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNameservers

`func (o *DnsZone) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *DnsZone) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *DnsZone) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *DnsZone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetSoa

`func (o *DnsZone) GetSoa() DnsSOARecord`

GetSoa returns the Soa field if non-nil, zero value otherwise.

### GetSoaOk

`func (o *DnsZone) GetSoaOk() (*DnsSOARecord, bool)`

GetSoaOk returns a tuple with the Soa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoa

`func (o *DnsZone) SetSoa(v DnsSOARecord)`

SetSoa sets Soa field to given value.

### HasSoa

`func (o *DnsZone) HasSoa() bool`

HasSoa returns a boolean if a field has been set.

### GetRecordCount

`func (o *DnsZone) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *DnsZone) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *DnsZone) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *DnsZone) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *DnsZone) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *DnsZone) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *DnsZone) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *DnsZone) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DnsZone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DnsZone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DnsZone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DnsZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DnsZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DnsZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DnsZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DnsZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


