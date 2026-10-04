# CloudflarePagesEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type is \&quot;plain_text\&quot; or \&quot;secret_text\&quot; and decides that: plain text is readable afterwards, secret text is write-only. Empty is Cloudflare&#39;s default, plain_text — so a secret with no type set is stored in the clear. | [optional] 
**Value** | Pointer to **string** | Value is the variable&#39;s value. Under type \&quot;secret_text\&quot; Cloudflare encrypts it on arrival and never reads it back, so a later read of the project shows the variable without this. | [optional] 

## Methods

### NewCloudflarePagesEnvVar

`func NewCloudflarePagesEnvVar() *CloudflarePagesEnvVar`

NewCloudflarePagesEnvVar instantiates a new CloudflarePagesEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePagesEnvVarWithDefaults

`func NewCloudflarePagesEnvVarWithDefaults() *CloudflarePagesEnvVar`

NewCloudflarePagesEnvVarWithDefaults instantiates a new CloudflarePagesEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudflarePagesEnvVar) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudflarePagesEnvVar) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudflarePagesEnvVar) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudflarePagesEnvVar) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CloudflarePagesEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudflarePagesEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudflarePagesEnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudflarePagesEnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


