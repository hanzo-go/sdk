# S3ListServiceAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]S3ServiceAccount**](S3ServiceAccount.md) |  | [optional] 

## Methods

### NewS3ListServiceAccounts200Response

`func NewS3ListServiceAccounts200Response() *S3ListServiceAccounts200Response`

NewS3ListServiceAccounts200Response instantiates a new S3ListServiceAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ListServiceAccounts200ResponseWithDefaults

`func NewS3ListServiceAccounts200ResponseWithDefaults() *S3ListServiceAccounts200Response`

NewS3ListServiceAccounts200ResponseWithDefaults instantiates a new S3ListServiceAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *S3ListServiceAccounts200Response) GetAccounts() []S3ServiceAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *S3ListServiceAccounts200Response) GetAccountsOk() (*[]S3ServiceAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *S3ListServiceAccounts200Response) SetAccounts(v []S3ServiceAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *S3ListServiceAccounts200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


