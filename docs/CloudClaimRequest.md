# CloudClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the referrer&#39;s referral code, as it appeared in their ?ref&#x3D; link. Case and surrounding whitespace do not matter. | [optional] 

## Methods

### NewCloudClaimRequest

`func NewCloudClaimRequest() *CloudClaimRequest`

NewCloudClaimRequest instantiates a new CloudClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClaimRequestWithDefaults

`func NewCloudClaimRequestWithDefaults() *CloudClaimRequest`

NewCloudClaimRequestWithDefaults instantiates a new CloudClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CloudClaimRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudClaimRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudClaimRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudClaimRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


