// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates an Access Control List (ACL) in a Redpanda cluster.
 *
 * ## Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as redpanda from "@pulumi/redpanda";
 *
 * const testNamespace = new redpanda.Namespace("testNamespace", {});
 * const config = new pulumi.Config();
 * const region = config.get("region") || "us-east-1";
 * const cloudProvider = config.get("cloudProvider") || "aws";
 * const testNetwork = new redpanda.Network("testNetwork", {
 *     namespaceId: testNamespace.id,
 *     cloudProvider: cloudProvider,
 *     region: region,
 *     clusterType: "dedicated",
 *     cidrBlock: "10.0.0.0/20",
 * });
 * const zones = config.getObject("zones") || [
 *     "use1-az2",
 *     "use1-az4",
 *     "use1-az6",
 * ];
 * const throughputTier = config.get("throughputTier") || "tier-1-aws-v2-arm";
 * const testCluster = new redpanda.Cluster("testCluster", {
 *     namespaceId: testNamespace.id,
 *     networkId: testNetwork.id,
 *     cloudProvider: cloudProvider,
 *     region: region,
 *     clusterType: "dedicated",
 *     connectionType: "public",
 *     throughputTier: throughputTier,
 *     zones: zones,
 *     allowDeletion: true,
 *     tags: {
 *         key: "value",
 *     },
 * });
 * const userPw = config.get("userPw") || "password";
 * const mechanism = config.get("mechanism") || "scram-sha-256";
 * const testUser = new redpanda.User("testUser", {
 *     password: userPw,
 *     mechanism: mechanism,
 *     clusterApiUrl: testCluster.clusterApiUrl,
 * });
 * const partitionCount = config.getNumber("partitionCount") || 3;
 * const replicationFactor = config.getNumber("replicationFactor") || 3;
 * const testTopic = new redpanda.Topic("testTopic", {
 *     partitionCount: partitionCount,
 *     replicationFactor: replicationFactor,
 *     clusterApiUrl: testCluster.clusterApiUrl,
 *     allowDeletion: true,
 * });
 * const testAcl = new redpanda.Acl("testAcl", {
 *     resourceType: "TOPIC",
 *     resourceName: testTopic.name,
 *     resourcePatternType: "LITERAL",
 *     principal: pulumi.interpolate`User:${testUser.name}`,
 *     host: "*",
 *     operation: "READ",
 *     permissionType: "ALLOW",
 *     clusterApiUrl: testCluster.clusterApiUrl,
 * });
 * const namespaceName = config.get("namespaceName") || "testname";
 * const networkName = config.get("networkName") || "testname";
 * const clusterName = config.get("clusterName") || "testname";
 * const userName = config.get("userName") || "test-username";
 * const topicName = config.get("topicName") || "test-topic";
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Limitations
 *
 * We are not currently able to support ACL creation in self hosted clusters. This is an area of active development so expect that to change soon.
 *
 * ## Import
 *
 * We do not support the import of ACLs into the Redpanda provider at this time.
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'redpanda:index/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
     */
    public readonly clusterApiUrl!: pulumi.Output<string>;
    /**
     * The host address to use for this ACL
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The operation type that shall be allowed or denied (e.g READ)
     */
    public readonly operation!: pulumi.Output<string>;
    /**
     * The permission type. It determines whether the operation should be ALLOWED or DENIED
     */
    public readonly permissionType!: pulumi.Output<string>;
    /**
     * The principal to apply this ACL for
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * The name of the resource this ACL entry will be on
     */
    public readonly resourceName!: pulumi.Output<string>;
    /**
     * The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
     */
    public readonly resourcePatternType!: pulumi.Output<string>;
    /**
     * The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
     */
    public readonly resourceType!: pulumi.Output<string>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            resourceInputs["clusterApiUrl"] = state ? state.clusterApiUrl : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["operation"] = state ? state.operation : undefined;
            resourceInputs["permissionType"] = state ? state.permissionType : undefined;
            resourceInputs["principal"] = state ? state.principal : undefined;
            resourceInputs["resourceName"] = state ? state.resourceName : undefined;
            resourceInputs["resourcePatternType"] = state ? state.resourcePatternType : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.clusterApiUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterApiUrl'");
            }
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.operation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operation'");
            }
            if ((!args || args.permissionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissionType'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            if ((!args || args.resourceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceName'");
            }
            if ((!args || args.resourcePatternType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourcePatternType'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["clusterApiUrl"] = args ? args.clusterApiUrl : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["operation"] = args ? args.operation : undefined;
            resourceInputs["permissionType"] = args ? args.permissionType : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["resourceName"] = args ? args.resourceName : undefined;
            resourceInputs["resourcePatternType"] = args ? args.resourcePatternType : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Acl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
     */
    clusterApiUrl?: pulumi.Input<string>;
    /**
     * The host address to use for this ACL
     */
    host?: pulumi.Input<string>;
    /**
     * The operation type that shall be allowed or denied (e.g READ)
     */
    operation?: pulumi.Input<string>;
    /**
     * The permission type. It determines whether the operation should be ALLOWED or DENIED
     */
    permissionType?: pulumi.Input<string>;
    /**
     * The principal to apply this ACL for
     */
    principal?: pulumi.Input<string>;
    /**
     * The name of the resource this ACL entry will be on
     */
    resourceName?: pulumi.Input<string>;
    /**
     * The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
     */
    resourcePatternType?: pulumi.Input<string>;
    /**
     * The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
     */
    resourceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * The cluster API URL. Changing this will prevent deletion of the resource on the existing cluster. It is generally a better idea to delete an existing resource and create a new one than to change this value unless you are planning to do state imports
     */
    clusterApiUrl: pulumi.Input<string>;
    /**
     * The host address to use for this ACL
     */
    host: pulumi.Input<string>;
    /**
     * The operation type that shall be allowed or denied (e.g READ)
     */
    operation: pulumi.Input<string>;
    /**
     * The permission type. It determines whether the operation should be ALLOWED or DENIED
     */
    permissionType: pulumi.Input<string>;
    /**
     * The principal to apply this ACL for
     */
    principal: pulumi.Input<string>;
    /**
     * The name of the resource this ACL entry will be on
     */
    resourceName: pulumi.Input<string>;
    /**
     * The pattern type of the resource. It determines the strategy how the provided resource name is matched (LITERAL, MATCH, PREFIXED, etc ...) against the actual resource names
     */
    resourcePatternType: pulumi.Input<string>;
    /**
     * The type of the resource (TOPIC, GROUP, etc...) this ACL shall target
     */
    resourceType: pulumi.Input<string>;
}
