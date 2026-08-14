# LicensingHealthView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Env is the deployment environment (\&quot;dev\&quot; | \&quot;staging\&quot; | \&quot;prod\&quot;). | [optional] 
**Service** | Pointer to **string** | Service is always \&quot;licensing\&quot;. | [optional] 
**Signer** | Pointer to **string** | Signer names the KMS provider signing licenses here. \&quot;local\&quot; means a development key: tokens it mints are not production credentials. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; whenever the process is up — this is not a dependency probe. | [optional] 

## Methods

### NewLicensingHealthView

`func NewLicensingHealthView() *LicensingHealthView`

NewLicensingHealthView instantiates a new LicensingHealthView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingHealthViewWithDefaults

`func NewLicensingHealthViewWithDefaults() *LicensingHealthView`

NewLicensingHealthViewWithDefaults instantiates a new LicensingHealthView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *LicensingHealthView) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *LicensingHealthView) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *LicensingHealthView) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *LicensingHealthView) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetService

`func (o *LicensingHealthView) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LicensingHealthView) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *LicensingHealthView) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *LicensingHealthView) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSigner

`func (o *LicensingHealthView) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *LicensingHealthView) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *LicensingHealthView) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *LicensingHealthView) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStatus

`func (o *LicensingHealthView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LicensingHealthView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LicensingHealthView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LicensingHealthView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


