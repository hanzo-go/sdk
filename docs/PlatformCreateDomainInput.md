# PlatformCreateDomainInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] 
**Https** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**DomainType** | **string** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ComposeId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformCreateDomainInput

`func NewPlatformCreateDomainInput(host string, domainType string, ) *PlatformCreateDomainInput`

NewPlatformCreateDomainInput instantiates a new PlatformCreateDomainInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformCreateDomainInputWithDefaults

`func NewPlatformCreateDomainInputWithDefaults() *PlatformCreateDomainInput`

NewPlatformCreateDomainInputWithDefaults instantiates a new PlatformCreateDomainInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PlatformCreateDomainInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PlatformCreateDomainInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PlatformCreateDomainInput) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PlatformCreateDomainInput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlatformCreateDomainInput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlatformCreateDomainInput) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PlatformCreateDomainInput) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHttps

`func (o *PlatformCreateDomainInput) GetHttps() bool`

GetHttps returns the Https field if non-nil, zero value otherwise.

### GetHttpsOk

`func (o *PlatformCreateDomainInput) GetHttpsOk() (*bool, bool)`

GetHttpsOk returns a tuple with the Https field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttps

`func (o *PlatformCreateDomainInput) SetHttps(v bool)`

SetHttps sets Https field to given value.

### HasHttps

`func (o *PlatformCreateDomainInput) HasHttps() bool`

HasHttps returns a boolean if a field has been set.

### GetPath

`func (o *PlatformCreateDomainInput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformCreateDomainInput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformCreateDomainInput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformCreateDomainInput) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDomainType

`func (o *PlatformCreateDomainInput) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *PlatformCreateDomainInput) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *PlatformCreateDomainInput) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.


### GetApplicationId

`func (o *PlatformCreateDomainInput) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PlatformCreateDomainInput) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PlatformCreateDomainInput) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *PlatformCreateDomainInput) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetComposeId

`func (o *PlatformCreateDomainInput) GetComposeId() string`

GetComposeId returns the ComposeId field if non-nil, zero value otherwise.

### GetComposeIdOk

`func (o *PlatformCreateDomainInput) GetComposeIdOk() (*string, bool)`

GetComposeIdOk returns a tuple with the ComposeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeId

`func (o *PlatformCreateDomainInput) SetComposeId(v string)`

SetComposeId sets ComposeId field to given value.

### HasComposeId

`func (o *PlatformCreateDomainInput) HasComposeId() bool`

HasComposeId returns a boolean if a field has been set.

### GetServiceName

`func (o *PlatformCreateDomainInput) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PlatformCreateDomainInput) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PlatformCreateDomainInput) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PlatformCreateDomainInput) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetCertificateType

`func (o *PlatformCreateDomainInput) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *PlatformCreateDomainInput) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *PlatformCreateDomainInput) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *PlatformCreateDomainInput) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


