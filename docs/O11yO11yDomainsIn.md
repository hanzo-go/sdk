# O11yO11yDomainsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain narrows the read to one external domain (the domain view requires it). | [optional] 
**End** | Pointer to **int32** | End is the window&#39;s end, epoch milliseconds. | [optional] 
**Endpoint** | Pointer to **string** | Endpoint narrows the domain view to one endpoint. | [optional] 
**Filter** | Pointer to [**O11yO11yDomainFilter**](O11yO11yDomainFilter.md) | Filter is an additional predicate in the query-builder filter syntax. | [optional] 
**GroupBy** | Pointer to [**[]O11yO11yDomainGroupBy**](O11yO11yDomainGroupBy.md) | GroupBy adds grouping columns to the result. | [optional] 
**ShowIp** | Pointer to **bool** | ShowIP keeps rows whose domain is a bare IP address; they are dropped otherwise. | [optional] 
**Start** | Pointer to **int32** | Start is the window&#39;s start, epoch milliseconds. | [optional] 

## Methods

### NewO11yO11yDomainsIn

`func NewO11yO11yDomainsIn() *O11yO11yDomainsIn`

NewO11yO11yDomainsIn instantiates a new O11yO11yDomainsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDomainsInWithDefaults

`func NewO11yO11yDomainsInWithDefaults() *O11yO11yDomainsIn`

NewO11yO11yDomainsInWithDefaults instantiates a new O11yO11yDomainsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *O11yO11yDomainsIn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *O11yO11yDomainsIn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *O11yO11yDomainsIn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *O11yO11yDomainsIn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnd

`func (o *O11yO11yDomainsIn) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yDomainsIn) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yDomainsIn) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yDomainsIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEndpoint

`func (o *O11yO11yDomainsIn) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *O11yO11yDomainsIn) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *O11yO11yDomainsIn) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *O11yO11yDomainsIn) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetFilter

`func (o *O11yO11yDomainsIn) GetFilter() O11yO11yDomainFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yO11yDomainsIn) GetFilterOk() (*O11yO11yDomainFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yO11yDomainsIn) SetFilter(v O11yO11yDomainFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yO11yDomainsIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *O11yO11yDomainsIn) GetGroupBy() []O11yO11yDomainGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yO11yDomainsIn) GetGroupByOk() (*[]O11yO11yDomainGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yO11yDomainsIn) SetGroupBy(v []O11yO11yDomainGroupBy)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yO11yDomainsIn) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetShowIp

`func (o *O11yO11yDomainsIn) GetShowIp() bool`

GetShowIp returns the ShowIp field if non-nil, zero value otherwise.

### GetShowIpOk

`func (o *O11yO11yDomainsIn) GetShowIpOk() (*bool, bool)`

GetShowIpOk returns a tuple with the ShowIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIp

`func (o *O11yO11yDomainsIn) SetShowIp(v bool)`

SetShowIp sets ShowIp field to given value.

### HasShowIp

`func (o *O11yO11yDomainsIn) HasShowIp() bool`

HasShowIp returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yDomainsIn) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yDomainsIn) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yDomainsIn) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yDomainsIn) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


