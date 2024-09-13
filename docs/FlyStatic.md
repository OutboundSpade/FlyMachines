# FlyStatic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestPath** | **string** |  | 
**IndexDocument** | Pointer to **string** |  | [optional] 
**TigrisBucket** | Pointer to **string** |  | [optional] 
**UrlPrefix** | **string** |  | 

## Methods

### NewFlyStatic

`func NewFlyStatic(guestPath string, urlPrefix string, ) *FlyStatic`

NewFlyStatic instantiates a new FlyStatic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlyStaticWithDefaults

`func NewFlyStaticWithDefaults() *FlyStatic`

NewFlyStaticWithDefaults instantiates a new FlyStatic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestPath

`func (o *FlyStatic) GetGuestPath() string`

GetGuestPath returns the GuestPath field if non-nil, zero value otherwise.

### GetGuestPathOk

`func (o *FlyStatic) GetGuestPathOk() (*string, bool)`

GetGuestPathOk returns a tuple with the GuestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPath

`func (o *FlyStatic) SetGuestPath(v string)`

SetGuestPath sets GuestPath field to given value.


### GetIndexDocument

`func (o *FlyStatic) GetIndexDocument() string`

GetIndexDocument returns the IndexDocument field if non-nil, zero value otherwise.

### GetIndexDocumentOk

`func (o *FlyStatic) GetIndexDocumentOk() (*string, bool)`

GetIndexDocumentOk returns a tuple with the IndexDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDocument

`func (o *FlyStatic) SetIndexDocument(v string)`

SetIndexDocument sets IndexDocument field to given value.

### HasIndexDocument

`func (o *FlyStatic) HasIndexDocument() bool`

HasIndexDocument returns a boolean if a field has been set.

### GetTigrisBucket

`func (o *FlyStatic) GetTigrisBucket() string`

GetTigrisBucket returns the TigrisBucket field if non-nil, zero value otherwise.

### GetTigrisBucketOk

`func (o *FlyStatic) GetTigrisBucketOk() (*string, bool)`

GetTigrisBucketOk returns a tuple with the TigrisBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTigrisBucket

`func (o *FlyStatic) SetTigrisBucket(v string)`

SetTigrisBucket sets TigrisBucket field to given value.

### HasTigrisBucket

`func (o *FlyStatic) HasTigrisBucket() bool`

HasTigrisBucket returns a boolean if a field has been set.

### GetUrlPrefix

`func (o *FlyStatic) GetUrlPrefix() string`

GetUrlPrefix returns the UrlPrefix field if non-nil, zero value otherwise.

### GetUrlPrefixOk

`func (o *FlyStatic) GetUrlPrefixOk() (*string, bool)`

GetUrlPrefixOk returns a tuple with the UrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrefix

`func (o *FlyStatic) SetUrlPrefix(v string)`

SetUrlPrefix sets UrlPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


