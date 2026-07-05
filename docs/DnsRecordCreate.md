# DnsRecordCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Ttl** | Pointer to **int32** |  | [optional] [default to 300]
**Content** | **string** | Record content varies by type: - A: IPv4 address (e.g. 1.2.3.4) - AAAA: IPv6 address - CNAME: Target hostname - MX: Mail server hostname - TXT: Text value (auto-quoted) - SRV: weight port target (priority set separately) - NS: Nameserver hostname - CAA: flags tag value (e.g. 0 issue letsencrypt.org)  | 
**Priority** | Pointer to **int32** | Required for MX and SRV records | [optional] 
**Proxied** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDnsRecordCreate

`func NewDnsRecordCreate(name string, type_ string, content string, ) *DnsRecordCreate`

NewDnsRecordCreate instantiates a new DnsRecordCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordCreateWithDefaults

`func NewDnsRecordCreateWithDefaults() *DnsRecordCreate`

NewDnsRecordCreateWithDefaults instantiates a new DnsRecordCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DnsRecordCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DnsRecordCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecordCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecordCreate) SetType(v string)`

SetType sets Type field to given value.


### GetTtl

`func (o *DnsRecordCreate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordCreate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordCreate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordCreate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetContent

`func (o *DnsRecordCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DnsRecordCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DnsRecordCreate) SetContent(v string)`

SetContent sets Content field to given value.


### GetPriority

`func (o *DnsRecordCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsRecordCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsRecordCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsRecordCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProxied

`func (o *DnsRecordCreate) GetProxied() bool`

GetProxied returns the Proxied field if non-nil, zero value otherwise.

### GetProxiedOk

`func (o *DnsRecordCreate) GetProxiedOk() (*bool, bool)`

GetProxiedOk returns a tuple with the Proxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxied

`func (o *DnsRecordCreate) SetProxied(v bool)`

SetProxied sets Proxied field to given value.

### HasProxied

`func (o *DnsRecordCreate) HasProxied() bool`

HasProxied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


