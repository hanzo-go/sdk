# AuthorsApproveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareBps** | Pointer to **int32** | Royalty share in basis points (500 &#x3D; 5%). | [optional] 

## Methods

### NewAuthorsApproveRequest

`func NewAuthorsApproveRequest() *AuthorsApproveRequest`

NewAuthorsApproveRequest instantiates a new AuthorsApproveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsApproveRequestWithDefaults

`func NewAuthorsApproveRequestWithDefaults() *AuthorsApproveRequest`

NewAuthorsApproveRequestWithDefaults instantiates a new AuthorsApproveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShareBps

`func (o *AuthorsApproveRequest) GetShareBps() int32`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *AuthorsApproveRequest) GetShareBpsOk() (*int32, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *AuthorsApproveRequest) SetShareBps(v int32)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *AuthorsApproveRequest) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


