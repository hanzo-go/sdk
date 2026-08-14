# LicensingVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App overrides the app_id the token is expected to carry. Leave it empty and the token&#39;s own app_id is used — an online verify is informational, and it is the ENGINE that enforces the app at boot. | [optional] 
**Token** | **string** | Token is the license token to check. | 

## Methods

### NewLicensingVerifyRequest

`func NewLicensingVerifyRequest(token string, ) *LicensingVerifyRequest`

NewLicensingVerifyRequest instantiates a new LicensingVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingVerifyRequestWithDefaults

`func NewLicensingVerifyRequestWithDefaults() *LicensingVerifyRequest`

NewLicensingVerifyRequestWithDefaults instantiates a new LicensingVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *LicensingVerifyRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *LicensingVerifyRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *LicensingVerifyRequest) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *LicensingVerifyRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetToken

`func (o *LicensingVerifyRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LicensingVerifyRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LicensingVerifyRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


