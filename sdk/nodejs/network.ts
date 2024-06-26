// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a network in the Redpanda Cloud.
 *
 * ## Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as redpanda from "@pulumiverse/redpanda";
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
 * const namespaceName = config.get("namespaceName") || "testname";
 * const networkName = config.get("networkName") || "testname";
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import redpanda:index/network:Network example networkId
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'redpanda:index/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * The cidrBlock to create the network in
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The cloud provider to create the network in. Can also be set at the provider level
     */
    public readonly cloudProvider!: pulumi.Output<string | undefined>;
    /**
     * The type of cluster this network is associated with, can be one of dedicated or cloud
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * Name of the network
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the namespace in which to create the network
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * The region to create the network in. Can also be set at the provider level
     */
    public readonly region!: pulumi.Output<string | undefined>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["cloudProvider"] = state ? state.cloudProvider : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            if ((!args || args.cidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["cloudProvider"] = args ? args.cloudProvider : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Network.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The cidrBlock to create the network in
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The cloud provider to create the network in. Can also be set at the provider level
     */
    cloudProvider?: pulumi.Input<string>;
    /**
     * The type of cluster this network is associated with, can be one of dedicated or cloud
     */
    clusterType?: pulumi.Input<string>;
    /**
     * Name of the network
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the namespace in which to create the network
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The region to create the network in. Can also be set at the provider level
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The cidrBlock to create the network in
     */
    cidrBlock: pulumi.Input<string>;
    /**
     * The cloud provider to create the network in. Can also be set at the provider level
     */
    cloudProvider?: pulumi.Input<string>;
    /**
     * The type of cluster this network is associated with, can be one of dedicated or cloud
     */
    clusterType: pulumi.Input<string>;
    /**
     * Name of the network
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the namespace in which to create the network
     */
    namespaceId: pulumi.Input<string>;
    /**
     * The region to create the network in. Can also be set at the provider level
     */
    region?: pulumi.Input<string>;
}
