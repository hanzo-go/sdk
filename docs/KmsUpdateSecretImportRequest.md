# KmsUpdateSecretImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **int32** |  | [optional] 
**Import** | Pointer to [**KmsUpdateSecretImportRequestImport**](KmsUpdateSecretImportRequestImport.md) |  | [optional] 

## Methods

### NewKmsUpdateSecretImportRequest

`func NewKmsUpdateSecretImportRequest() *KmsUpdateSecretImportRequest`

NewKmsUpdateSecretImportRequest instantiates a new KmsUpdateSecretImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsUpdateSecretImportRequestWithDefaults

`func NewKmsUpdateSecretImportRequestWithDefaults() *KmsUpdateSecretImportRequest`

NewKmsUpdateSecretImportRequestWithDefaults instantiates a new KmsUpdateSecretImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *KmsUpdateSecretImportRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *KmsUpdateSecretImportRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *KmsUpdateSecretImportRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *KmsUpdateSecretImportRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetImport

`func (o *KmsUpdateSecretImportRequest) GetImport() KmsUpdateSecretImportRequestImport`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *KmsUpdateSecretImportRequest) GetImportOk() (*KmsUpdateSecretImportRequestImport, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *KmsUpdateSecretImportRequest) SetImport(v KmsUpdateSecretImportRequestImport)`

SetImport sets Import field to given value.

### HasImport

`func (o *KmsUpdateSecretImportRequest) HasImport() bool`

HasImport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


