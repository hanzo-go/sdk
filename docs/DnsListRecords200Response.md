# DnsListRecords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]DnsRecord**](DnsRecord.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnsListRecords200Response

`func NewDnsListRecords200Response() *DnsListRecords200Response`

NewDnsListRecords200Response instantiates a new DnsListRecords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsListRecords200ResponseWithDefaults

`func NewDnsListRecords200ResponseWithDefaults() *DnsListRecords200Response`

NewDnsListRecords200ResponseWithDefaults instantiates a new DnsListRecords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *DnsListRecords200Response) GetRecords() []DnsRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *DnsListRecords200Response) GetRecordsOk() (*[]DnsRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *DnsListRecords200Response) SetRecords(v []DnsRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *DnsListRecords200Response) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *DnsListRecords200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DnsListRecords200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DnsListRecords200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DnsListRecords200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


