# DnsUpdateRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Proxied** | Pointer to **bool** |  | [optional] 

## Methods

### NewDnsUpdateRecordRequest

`func NewDnsUpdateRecordRequest() *DnsUpdateRecordRequest`

NewDnsUpdateRecordRequest instantiates a new DnsUpdateRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsUpdateRecordRequestWithDefaults

`func NewDnsUpdateRecordRequestWithDefaults() *DnsUpdateRecordRequest`

NewDnsUpdateRecordRequestWithDefaults instantiates a new DnsUpdateRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DnsUpdateRecordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsUpdateRecordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsUpdateRecordRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsUpdateRecordRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DnsUpdateRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsUpdateRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsUpdateRecordRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DnsUpdateRecordRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *DnsUpdateRecordRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DnsUpdateRecordRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DnsUpdateRecordRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DnsUpdateRecordRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTtl

`func (o *DnsUpdateRecordRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsUpdateRecordRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsUpdateRecordRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsUpdateRecordRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPriority

`func (o *DnsUpdateRecordRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DnsUpdateRecordRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DnsUpdateRecordRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DnsUpdateRecordRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProxied

`func (o *DnsUpdateRecordRequest) GetProxied() bool`

GetProxied returns the Proxied field if non-nil, zero value otherwise.

### GetProxiedOk

`func (o *DnsUpdateRecordRequest) GetProxiedOk() (*bool, bool)`

GetProxiedOk returns a tuple with the Proxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxied

`func (o *DnsUpdateRecordRequest) SetProxied(v bool)`

SetProxied sets Proxied field to given value.

### HasProxied

`func (o *DnsUpdateRecordRequest) HasProxied() bool`

HasProxied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


