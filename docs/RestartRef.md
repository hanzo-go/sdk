# RestartRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the service&#39;s CR name, from the path. It must be a DNS-1123 label. | [optional] 
**Env** | Pointer to **string** | Env is REQUIRED and must be main, test or dev. A bare call does not default to production, which is what closes the fat-finger and confused-deputy hazard.  It carries no &#x60;validate:\&quot;required\&quot;&#x60;: the handler already refuses an empty env with the sentence that names the three values, and a validator tag would replace that sentence with a generic one. The requirement is stated here and enforced there, once. | [optional] 

## Methods

### NewRestartRef

`func NewRestartRef() *RestartRef`

NewRestartRef instantiates a new RestartRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartRefWithDefaults

`func NewRestartRefWithDefaults() *RestartRef`

NewRestartRefWithDefaults instantiates a new RestartRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *RestartRef) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *RestartRef) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *RestartRef) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *RestartRef) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetEnv

`func (o *RestartRef) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *RestartRef) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *RestartRef) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *RestartRef) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


