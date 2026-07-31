# CloudSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address the signature request is sent to. | [optional] 
**Name** | Pointer to **string** | Name is the recipient&#39;s name, as it appears on the signature request. | [optional] 

## Methods

### NewCloudSigner

`func NewCloudSigner() *CloudSigner`

NewCloudSigner instantiates a new CloudSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSignerWithDefaults

`func NewCloudSignerWithDefaults() *CloudSigner`

NewCloudSignerWithDefaults instantiates a new CloudSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CloudSigner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudSigner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudSigner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudSigner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CloudSigner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSigner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSigner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSigner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


