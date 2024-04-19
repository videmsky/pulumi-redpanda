// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redpanda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/videmsky/pulumi-redpanda/sdk/go/redpanda/internal"
)

// Data source for a Redpanda Cloud network
//
// ## Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/videmsky/pulumi-redpanda/sdk/go/redpanda"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redpanda.LookupNetwork(ctx, &redpanda.LookupNetworkArgs{
//				Id: "network_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkResult
	err := ctx.Invoke("redpanda:index/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// UUID of the network
	Id string `pulumi:"id"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// The cidrBlock to create the network in
	CidrBlock string `pulumi:"cidrBlock"`
	// The cloud provider to create the network in. Can also be set at the provider level
	CloudProvider string `pulumi:"cloudProvider"`
	// The type of cluster this network is associated with, can be one of dedicated or cloud
	ClusterType string `pulumi:"clusterType"`
	// UUID of the network
	Id string `pulumi:"id"`
	// Name of the network
	Name string `pulumi:"name"`
	// The id of the namespace in which to create the network
	NamespaceId string `pulumi:"namespaceId"`
	// The region to create the network in. Can also be set at the provider level
	Region string `pulumi:"region"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			var s LookupNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	// UUID of the network
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// The cidrBlock to create the network in
func (o LookupNetworkResultOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.CidrBlock }).(pulumi.StringOutput)
}

// The cloud provider to create the network in. Can also be set at the provider level
func (o LookupNetworkResultOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.CloudProvider }).(pulumi.StringOutput)
}

// The type of cluster this network is associated with, can be one of dedicated or cloud
func (o LookupNetworkResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

// UUID of the network
func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the network
func (o LookupNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the namespace in which to create the network
func (o LookupNetworkResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

// The region to create the network in. Can also be set at the provider level
func (o LookupNetworkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}