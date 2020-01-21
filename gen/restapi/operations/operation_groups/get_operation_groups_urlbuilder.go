// Code generated by go-swagger; DO NOT EDIT.

package operation_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetOperationGroupsURL generates an URL for the get operation groups operation
type GetOperationGroupsURL struct {
	Network  string
	Platform string

	AccountDelegate      []string
	AccountID            []string
	AccountManager       []string
	BlockID              []string
	BlockLevel           []int64
	BlockNetid           []string
	BlockProtocol        []string
	Limit                *int64
	Offset               *int64
	OperationDestination []string
	OperationID          []string
	OperationKind        []string
	OperationParticipant []string
	OperationSource      []string
	Order                *string
	SortBy               *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetOperationGroupsURL) WithBasePath(bp string) *GetOperationGroupsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetOperationGroupsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetOperationGroupsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/data/{platform}/{network}/operation_groups"

	network := o.Network
	if network != "" {
		_path = strings.Replace(_path, "{network}", network, -1)
	} else {
		return nil, errors.New("network is required on GetOperationGroupsURL")
	}

	platform := o.Platform
	if platform != "" {
		_path = strings.Replace(_path, "{platform}", platform, -1)
	} else {
		return nil, errors.New("platform is required on GetOperationGroupsURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var accountDelegateIR []string
	for _, accountDelegateI := range o.AccountDelegate {
		accountDelegateIS := accountDelegateI
		if accountDelegateIS != "" {
			accountDelegateIR = append(accountDelegateIR, accountDelegateIS)
		}
	}

	accountDelegate := swag.JoinByFormat(accountDelegateIR, "multi")

	for _, qsv := range accountDelegate {
		qs.Add("account_delegate", qsv)
	}

	var accountIDIR []string
	for _, accountIDI := range o.AccountID {
		accountIDIS := accountIDI
		if accountIDIS != "" {
			accountIDIR = append(accountIDIR, accountIDIS)
		}
	}

	accountID := swag.JoinByFormat(accountIDIR, "multi")

	for _, qsv := range accountID {
		qs.Add("account_id", qsv)
	}

	var accountManagerIR []string
	for _, accountManagerI := range o.AccountManager {
		accountManagerIS := accountManagerI
		if accountManagerIS != "" {
			accountManagerIR = append(accountManagerIR, accountManagerIS)
		}
	}

	accountManager := swag.JoinByFormat(accountManagerIR, "multi")

	for _, qsv := range accountManager {
		qs.Add("account_manager", qsv)
	}

	var blockIDIR []string
	for _, blockIDI := range o.BlockID {
		blockIDIS := blockIDI
		if blockIDIS != "" {
			blockIDIR = append(blockIDIR, blockIDIS)
		}
	}

	blockID := swag.JoinByFormat(blockIDIR, "multi")

	for _, qsv := range blockID {
		qs.Add("block_id", qsv)
	}

	var blockLevelIR []string
	for _, blockLevelI := range o.BlockLevel {
		blockLevelIS := swag.FormatInt64(blockLevelI)
		if blockLevelIS != "" {
			blockLevelIR = append(blockLevelIR, blockLevelIS)
		}
	}

	blockLevel := swag.JoinByFormat(blockLevelIR, "multi")

	for _, qsv := range blockLevel {
		qs.Add("block_level", qsv)
	}

	var blockNetidIR []string
	for _, blockNetidI := range o.BlockNetid {
		blockNetidIS := blockNetidI
		if blockNetidIS != "" {
			blockNetidIR = append(blockNetidIR, blockNetidIS)
		}
	}

	blockNetid := swag.JoinByFormat(blockNetidIR, "multi")

	for _, qsv := range blockNetid {
		qs.Add("block_netid", qsv)
	}

	var blockProtocolIR []string
	for _, blockProtocolI := range o.BlockProtocol {
		blockProtocolIS := blockProtocolI
		if blockProtocolIS != "" {
			blockProtocolIR = append(blockProtocolIR, blockProtocolIS)
		}
	}

	blockProtocol := swag.JoinByFormat(blockProtocolIR, "multi")

	for _, qsv := range blockProtocol {
		qs.Add("block_protocol", qsv)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt64(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var operationDestinationIR []string
	for _, operationDestinationI := range o.OperationDestination {
		operationDestinationIS := operationDestinationI
		if operationDestinationIS != "" {
			operationDestinationIR = append(operationDestinationIR, operationDestinationIS)
		}
	}

	operationDestination := swag.JoinByFormat(operationDestinationIR, "multi")

	for _, qsv := range operationDestination {
		qs.Add("operation_destination", qsv)
	}

	var operationIDIR []string
	for _, operationIDI := range o.OperationID {
		operationIDIS := operationIDI
		if operationIDIS != "" {
			operationIDIR = append(operationIDIR, operationIDIS)
		}
	}

	operationID := swag.JoinByFormat(operationIDIR, "multi")

	for _, qsv := range operationID {
		qs.Add("operation_id", qsv)
	}

	var operationKindIR []string
	for _, operationKindI := range o.OperationKind {
		operationKindIS := operationKindI
		if operationKindIS != "" {
			operationKindIR = append(operationKindIR, operationKindIS)
		}
	}

	operationKind := swag.JoinByFormat(operationKindIR, "multi")

	for _, qsv := range operationKind {
		qs.Add("operation_kind", qsv)
	}

	var operationParticipantIR []string
	for _, operationParticipantI := range o.OperationParticipant {
		operationParticipantIS := operationParticipantI
		if operationParticipantIS != "" {
			operationParticipantIR = append(operationParticipantIR, operationParticipantIS)
		}
	}

	operationParticipant := swag.JoinByFormat(operationParticipantIR, "multi")

	for _, qsv := range operationParticipant {
		qs.Add("operation_participant", qsv)
	}

	var operationSourceIR []string
	for _, operationSourceI := range o.OperationSource {
		operationSourceIS := operationSourceI
		if operationSourceIS != "" {
			operationSourceIR = append(operationSourceIR, operationSourceIS)
		}
	}

	operationSource := swag.JoinByFormat(operationSourceIR, "multi")

	for _, qsv := range operationSource {
		qs.Add("operation_source", qsv)
	}

	var orderQ string
	if o.Order != nil {
		orderQ = *o.Order
	}
	if orderQ != "" {
		qs.Set("order", orderQ)
	}

	var sortByQ string
	if o.SortBy != nil {
		sortByQ = *o.SortBy
	}
	if sortByQ != "" {
		qs.Set("sort_by", sortByQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetOperationGroupsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetOperationGroupsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetOperationGroupsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetOperationGroupsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetOperationGroupsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetOperationGroupsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
