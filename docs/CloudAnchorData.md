# CloudAnchorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | Pointer to [**CloudAnchorStatus**](CloudAnchorStatus.md) | Anchor is the Hanzo L1 anchoring status of the ledger root after this call. | [optional] 

## Methods

### NewCloudAnchorData

`func NewCloudAnchorData() *CloudAnchorData`

NewCloudAnchorData instantiates a new CloudAnchorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnchorDataWithDefaults

`func NewCloudAnchorDataWithDefaults() *CloudAnchorData`

NewCloudAnchorDataWithDefaults instantiates a new CloudAnchorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *CloudAnchorData) GetAnchor() CloudAnchorStatus`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *CloudAnchorData) GetAnchorOk() (*CloudAnchorStatus, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *CloudAnchorData) SetAnchor(v CloudAnchorStatus)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *CloudAnchorData) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


