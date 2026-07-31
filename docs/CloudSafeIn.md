# CloudSafeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentIds** | Pointer to **[]string** | DocumentIDs are data room document ids to raise a signature request over. Required. | [optional] 
**Signers** | Pointer to [**[]CloudSigner**](CloudSigner.md) | Signers are the recipients, each a name and an email. Required. | [optional] 

## Methods

### NewCloudSafeIn

`func NewCloudSafeIn() *CloudSafeIn`

NewCloudSafeIn instantiates a new CloudSafeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSafeInWithDefaults

`func NewCloudSafeInWithDefaults() *CloudSafeIn`

NewCloudSafeInWithDefaults instantiates a new CloudSafeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentIds

`func (o *CloudSafeIn) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *CloudSafeIn) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *CloudSafeIn) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *CloudSafeIn) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetSigners

`func (o *CloudSafeIn) GetSigners() []CloudSigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *CloudSafeIn) GetSignersOk() (*[]CloudSigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *CloudSafeIn) SetSigners(v []CloudSigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *CloudSafeIn) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


