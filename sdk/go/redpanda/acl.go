// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redpanda

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-redpanda/sdk/go/redpanda/internal"
)

// Creates an Access Control List (ACL) in a Redpanda cluster.
//
// ## Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/pulumiverse/pulumi-redpanda/sdk/go/redpanda"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testNamespace, err := redpanda.NewNamespace(ctx, "testNamespace", nil)
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			region := "us-east-1"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			cloudProvider := "aws"
//			if param := cfg.Get("cloudProvider"); param != "" {
//				cloudProvider = param
//			}
//			testNetwork, err := redpanda.NewNetwork(ctx, "testNetwork", &redpanda.NetworkArgs{
//				NamespaceId:   testNamespace.ID(),
//				CloudProvider: pulumi.String(cloudProvider),
//				Region:        pulumi.String(region),
//				ClusterType:   pulumi.String("dedicated"),
//				CidrBlock:     pulumi.String("10.0.0.0/20"),
//			})
//			if err != nil {
//				return err
//			}
//			zones := []string{
//				"use1-az2",
//				"use1-az4",
//				"use1-az6",
//			}
//			if param := cfg.GetObject("zones"); param != nil {
//				zones = param
//			}
//			throughputTier := "tier-1-aws-v2-arm"
//			if param := cfg.Get("throughputTier"); param != "" {
//				throughputTier = param
//			}
//			testCluster, err := redpanda.NewCluster(ctx, "testCluster", &redpanda.ClusterArgs{
//				NamespaceId:    testNamespace.ID(),
//				NetworkId:      testNetwork.ID(),
//				CloudProvider:  pulumi.String(cloudProvider),
//				Region:         pulumi.String(region),
//				ClusterType:    pulumi.String("dedicated"),
//				ConnectionType: pulumi.String("public"),
//				ThroughputTier: pulumi.String(throughputTier),
//				Zones:          pulumi.Any(zones),
//				AllowDeletion:  pulumi.Bool(true),
//				Tags: pulumi.StringMap{
//					"key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			userPw := "password"
//			if param := cfg.Get("userPw"); param != "" {
//				userPw = param
//			}
//			mechanism := "scram-sha-256"
//			if param := cfg.Get("mechanism"); param != "" {
//				mechanism = param
//			}
//			testUser, err := redpanda.NewUser(ctx, "testUser", &redpanda.UserArgs{
//				Password:      pulumi.String(userPw),
//				Mechanism:     pulumi.String(mechanism),
//				ClusterApiUrl: testCluster.ClusterApiUrl,
//			})
//			if err != nil {
//				return err
//			}
//			partitionCount := float64(3)
//			if param := cfg.GetFloat64("partitionCount"); param != 0 {
//				partitionCount = param
//			}
//			replicationFactor := float64(3)
//			if param := cfg.GetFloat64("replicationFactor"); param != 0 {
//				replicationFactor = param
//			}
//			testTopic, err := redpanda.NewTopic(ctx, "testTopic", &redpanda.TopicArgs{
//				PartitionCount:    pulumi.Float64(partitionCount),
//				ReplicationFactor: pulumi.Float64(replicationFactor),
//				ClusterApiUrl:     testCluster.ClusterApiUrl,
//				AllowDeletion:     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = redpanda.NewAcl(ctx, "testAcl", &redpanda.AclArgs{
//				ResourceType:        pulumi.String("TOPIC"),
//				ResourceName:        testTopic.Name,
//				ResourcePatternType: pulumi.String("LITERAL"),
//				Principal: testUser.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("User:%v", name), nil
//				}).(pulumi.StringOutput),
//				Host:           pulumi.String("*"),
//				Operation:      pulumi.String("READ"),
//				PermissionType: pulumi.String("ALLOW"),
//				ClusterApiUrl:  testCluster.ClusterApiUrl,
//			})
//			if err != nil {
//				return err
//			}
//			namespaceName := "testname"
//			if param := cfg.Get("namespaceName"); param != "" {
//				namespaceName = param
//			}
//			networkName := "testname"
//			if param := cfg.Get("networkName"); param != "" {
//				networkName = param
//			}
//			clusterName := "testname"
//			if param := cfg.Get("clusterName"); param != "" {
//				clusterName = param
//			}
//			userName := "test-username"
//			if param := cfg.Get("userName"); param != "" {
//				userName = param
//			}
//			topicName := "test-topic"
//			if param := cfg.Get("topicName"); param != "" {
//				topicName = param
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Limitations
//
// We are not currently able to support ACL creation in self hosted clusters. This is an area of active development so expect that to change soon.
//
// ## Import
//
// We do not support the import of ACLs into the Redpanda provider at this time.
type Acl struct {
	pulumi.CustomResourceState

	// The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
	ClusterApiUrl pulumi.StringOutput `pulumi:"clusterApiUrl"`
	// The host address to use for this ACL
	Host pulumi.StringOutput `pulumi:"host"`
	// The operation type that shall be allowed or denied (e.g READ)
	Operation pulumi.StringOutput `pulumi:"operation"`
	// The permission type. It determines whether the operation should be ALLOWED or DENIED
	PermissionType pulumi.StringOutput `pulumi:"permissionType"`
	// The principal to apply this ACL for
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The name of the resource this ACL entry will be on
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
	ResourcePatternType pulumi.StringOutput `pulumi:"resourcePatternType"`
	// The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterApiUrl == nil {
		return nil, errors.New("invalid value for required argument 'ClusterApiUrl'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Operation == nil {
		return nil, errors.New("invalid value for required argument 'Operation'")
	}
	if args.PermissionType == nil {
		return nil, errors.New("invalid value for required argument 'PermissionType'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourcePatternType == nil {
		return nil, errors.New("invalid value for required argument 'ResourcePatternType'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Acl
	err := ctx.RegisterResource("redpanda:index/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("redpanda:index/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
	ClusterApiUrl *string `pulumi:"clusterApiUrl"`
	// The host address to use for this ACL
	Host *string `pulumi:"host"`
	// The operation type that shall be allowed or denied (e.g READ)
	Operation *string `pulumi:"operation"`
	// The permission type. It determines whether the operation should be ALLOWED or DENIED
	PermissionType *string `pulumi:"permissionType"`
	// The principal to apply this ACL for
	Principal *string `pulumi:"principal"`
	// The name of the resource this ACL entry will be on
	ResourceName *string `pulumi:"resourceName"`
	// The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
	ResourcePatternType *string `pulumi:"resourcePatternType"`
	// The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
	ResourceType *string `pulumi:"resourceType"`
}

type AclState struct {
	// The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
	ClusterApiUrl pulumi.StringPtrInput
	// The host address to use for this ACL
	Host pulumi.StringPtrInput
	// The operation type that shall be allowed or denied (e.g READ)
	Operation pulumi.StringPtrInput
	// The permission type. It determines whether the operation should be ALLOWED or DENIED
	PermissionType pulumi.StringPtrInput
	// The principal to apply this ACL for
	Principal pulumi.StringPtrInput
	// The name of the resource this ACL entry will be on
	ResourceName pulumi.StringPtrInput
	// The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
	ResourcePatternType pulumi.StringPtrInput
	// The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
	ResourceType pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
	ClusterApiUrl string `pulumi:"clusterApiUrl"`
	// The host address to use for this ACL
	Host string `pulumi:"host"`
	// The operation type that shall be allowed or denied (e.g READ)
	Operation string `pulumi:"operation"`
	// The permission type. It determines whether the operation should be ALLOWED or DENIED
	PermissionType string `pulumi:"permissionType"`
	// The principal to apply this ACL for
	Principal string `pulumi:"principal"`
	// The name of the resource this ACL entry will be on
	ResourceName string `pulumi:"resourceName"`
	// The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
	ResourcePatternType string `pulumi:"resourcePatternType"`
	// The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
	ResourceType string `pulumi:"resourceType"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
	ClusterApiUrl pulumi.StringInput
	// The host address to use for this ACL
	Host pulumi.StringInput
	// The operation type that shall be allowed or denied (e.g READ)
	Operation pulumi.StringInput
	// The permission type. It determines whether the operation should be ALLOWED or DENIED
	PermissionType pulumi.StringInput
	// The principal to apply this ACL for
	Principal pulumi.StringInput
	// The name of the resource this ACL entry will be on
	ResourceName pulumi.StringInput
	// The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
	ResourcePatternType pulumi.StringInput
	// The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
	ResourceType pulumi.StringInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

// AclArrayInput is an input type that accepts AclArray and AclArrayOutput values.
// You can construct a concrete instance of `AclArrayInput` via:
//
//	AclArray{ AclArgs{...} }
type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

// AclMapInput is an input type that accepts AclMap and AclMapOutput values.
// You can construct a concrete instance of `AclMapInput` via:
//
//	AclMap{ "key": AclArgs{...} }
type AclMapInput interface {
	pulumi.Input

	ToAclMapOutput() AclMapOutput
	ToAclMapOutputWithContext(context.Context) AclMapOutput
}

type AclMap map[string]AclInput

func (AclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (i AclMap) ToAclMapOutput() AclMapOutput {
	return i.ToAclMapOutputWithContext(context.Background())
}

func (i AclMap) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclMapOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

// The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
func (o AclOutput) ClusterApiUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ClusterApiUrl }).(pulumi.StringOutput)
}

// The host address to use for this ACL
func (o AclOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The operation type that shall be allowed or denied (e.g READ)
func (o AclOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.Operation }).(pulumi.StringOutput)
}

// The permission type. It determines whether the operation should be ALLOWED or DENIED
func (o AclOutput) PermissionType() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.PermissionType }).(pulumi.StringOutput)
}

// The principal to apply this ACL for
func (o AclOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The name of the resource this ACL entry will be on
func (o AclOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
func (o AclOutput) ResourcePatternType() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ResourcePatternType }).(pulumi.StringOutput)
}

// The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
func (o AclOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].([]*Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclMapOutput struct{ *pulumi.OutputState }

func (AclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (o AclMapOutput) ToAclMapOutput() AclMapOutput {
	return o
}

func (o AclMapOutput) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return o
}

func (o AclMapOutput) MapIndex(k pulumi.StringInput) AclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].(map[string]*Acl)[vs[1].(string)]
	}).(AclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclInput)(nil)).Elem(), &Acl{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclArrayInput)(nil)).Elem(), AclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclMapInput)(nil)).Elem(), AclMap{})
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclMapOutput{})
}
