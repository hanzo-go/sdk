# NetworkServiceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is where the HOSTING identity forwards a connection — an address the host device itself can reach, \&quot;127.0.0.1\&quot; for a server on the device. | [optional] 
**Name** | Pointer to **string** | Name is the service&#39;s name within the org — a DNS label. The fabric knows the service as \&quot;&lt;name&gt;.&lt;org&gt;\&quot; and answers for it at \&quot;&lt;name&gt;.&lt;org&gt;.zt\&quot;. | [optional] 
**Port** | Pointer to **int64** | Port is the port beside Host, and the one the DNS name intercepts. | [optional] 

## Methods

### NewNetworkServiceIn

`func NewNetworkServiceIn() *NetworkServiceIn`

NewNetworkServiceIn instantiates a new NetworkServiceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceInWithDefaults

`func NewNetworkServiceInWithDefaults() *NetworkServiceIn`

NewNetworkServiceInWithDefaults instantiates a new NetworkServiceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NetworkServiceIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NetworkServiceIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NetworkServiceIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NetworkServiceIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetName

`func (o *NetworkServiceIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkServiceIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkServiceIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkServiceIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *NetworkServiceIn) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworkServiceIn) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworkServiceIn) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworkServiceIn) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


