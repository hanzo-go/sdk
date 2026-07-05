# DnsSOARecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mname** | Pointer to **string** | Primary nameserver | [optional] 
**Rname** | Pointer to **string** | Admin email (DNS format) | [optional] 
**Serial** | Pointer to **int64** |  | [optional] 
**Refresh** | Pointer to **int32** |  | [optional] [default to 3600]
**Retry** | Pointer to **int32** |  | [optional] [default to 600]
**Expire** | Pointer to **int32** |  | [optional] [default to 604800]
**Minimum** | Pointer to **int32** | Negative cache TTL | [optional] [default to 300]

## Methods

### NewDnsSOARecord

`func NewDnsSOARecord() *DnsSOARecord`

NewDnsSOARecord instantiates a new DnsSOARecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSOARecordWithDefaults

`func NewDnsSOARecordWithDefaults() *DnsSOARecord`

NewDnsSOARecordWithDefaults instantiates a new DnsSOARecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMname

`func (o *DnsSOARecord) GetMname() string`

GetMname returns the Mname field if non-nil, zero value otherwise.

### GetMnameOk

`func (o *DnsSOARecord) GetMnameOk() (*string, bool)`

GetMnameOk returns a tuple with the Mname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMname

`func (o *DnsSOARecord) SetMname(v string)`

SetMname sets Mname field to given value.

### HasMname

`func (o *DnsSOARecord) HasMname() bool`

HasMname returns a boolean if a field has been set.

### GetRname

`func (o *DnsSOARecord) GetRname() string`

GetRname returns the Rname field if non-nil, zero value otherwise.

### GetRnameOk

`func (o *DnsSOARecord) GetRnameOk() (*string, bool)`

GetRnameOk returns a tuple with the Rname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRname

`func (o *DnsSOARecord) SetRname(v string)`

SetRname sets Rname field to given value.

### HasRname

`func (o *DnsSOARecord) HasRname() bool`

HasRname returns a boolean if a field has been set.

### GetSerial

`func (o *DnsSOARecord) GetSerial() int64`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DnsSOARecord) GetSerialOk() (*int64, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DnsSOARecord) SetSerial(v int64)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DnsSOARecord) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRefresh

`func (o *DnsSOARecord) GetRefresh() int32`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *DnsSOARecord) GetRefreshOk() (*int32, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *DnsSOARecord) SetRefresh(v int32)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *DnsSOARecord) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRetry

`func (o *DnsSOARecord) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *DnsSOARecord) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *DnsSOARecord) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *DnsSOARecord) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetExpire

`func (o *DnsSOARecord) GetExpire() int32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *DnsSOARecord) GetExpireOk() (*int32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *DnsSOARecord) SetExpire(v int32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *DnsSOARecord) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetMinimum

`func (o *DnsSOARecord) GetMinimum() int32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *DnsSOARecord) GetMinimumOk() (*int32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *DnsSOARecord) SetMinimum(v int32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *DnsSOARecord) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


