// Code generated by go-swagger; DO NOT EDIT.

package node_pool_config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// NewAddOrUpdateNodePoolConfigItemParams creates a new AddOrUpdateNodePoolConfigItemParams object
// with the default values initialized.
func NewAddOrUpdateNodePoolConfigItemParams() *AddOrUpdateNodePoolConfigItemParams {
	var ()
	return &AddOrUpdateNodePoolConfigItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrUpdateNodePoolConfigItemParamsWithTimeout creates a new AddOrUpdateNodePoolConfigItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOrUpdateNodePoolConfigItemParamsWithTimeout(timeout time.Duration) *AddOrUpdateNodePoolConfigItemParams {
	var ()
	return &AddOrUpdateNodePoolConfigItemParams{

		timeout: timeout,
	}
}

// NewAddOrUpdateNodePoolConfigItemParamsWithContext creates a new AddOrUpdateNodePoolConfigItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOrUpdateNodePoolConfigItemParamsWithContext(ctx context.Context) *AddOrUpdateNodePoolConfigItemParams {
	var ()
	return &AddOrUpdateNodePoolConfigItemParams{

		Context: ctx,
	}
}

// NewAddOrUpdateNodePoolConfigItemParamsWithHTTPClient creates a new AddOrUpdateNodePoolConfigItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOrUpdateNodePoolConfigItemParamsWithHTTPClient(client *http.Client) *AddOrUpdateNodePoolConfigItemParams {
	var ()
	return &AddOrUpdateNodePoolConfigItemParams{
		HTTPClient: client,
	}
}

/*AddOrUpdateNodePoolConfigItemParams contains all the parameters to send to the API endpoint
for the add or update node pool config item operation typically these are written to a http.Request
*/
type AddOrUpdateNodePoolConfigItemParams struct {

	/*ClusterID
	  ID of the cluster.

	*/
	ClusterID string
	/*ConfigKey
	  Key for the config value.

	*/
	ConfigKey string
	/*NodePoolName
	  Name of the node pool.

	*/
	NodePoolName string
	/*Value
	  Config value.

	*/
	Value *models.ConfigValue

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithTimeout(timeout time.Duration) *AddOrUpdateNodePoolConfigItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithContext(ctx context.Context) *AddOrUpdateNodePoolConfigItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithHTTPClient(client *http.Client) *AddOrUpdateNodePoolConfigItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithClusterID(clusterID string) *AddOrUpdateNodePoolConfigItemParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithConfigKey adds the configKey to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithConfigKey(configKey string) *AddOrUpdateNodePoolConfigItemParams {
	o.SetConfigKey(configKey)
	return o
}

// SetConfigKey adds the configKey to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetConfigKey(configKey string) {
	o.ConfigKey = configKey
}

// WithNodePoolName adds the nodePoolName to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithNodePoolName(nodePoolName string) *AddOrUpdateNodePoolConfigItemParams {
	o.SetNodePoolName(nodePoolName)
	return o
}

// SetNodePoolName adds the nodePoolName to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetNodePoolName(nodePoolName string) {
	o.NodePoolName = nodePoolName
}

// WithValue adds the value to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) WithValue(value *models.ConfigValue) *AddOrUpdateNodePoolConfigItemParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the add or update node pool config item params
func (o *AddOrUpdateNodePoolConfigItemParams) SetValue(value *models.ConfigValue) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrUpdateNodePoolConfigItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param config_key
	if err := r.SetPathParam("config_key", o.ConfigKey); err != nil {
		return err
	}

	// path param node_pool_name
	if err := r.SetPathParam("node_pool_name", o.NodePoolName); err != nil {
		return err
	}

	if o.Value != nil {
		if err := r.SetBodyParam(o.Value); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
