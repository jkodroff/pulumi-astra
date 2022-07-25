// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-index/sdk/go/index"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := astra.NewDatabase(ctx, "example", &astra.DatabaseArgs{
// 			CloudProvider: pulumi.String("gcp"),
// 			Keyspace:      pulumi.String("keyspace"),
// 			Regions: pulumi.StringArray{
// 				pulumi.String("us-east1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import astra:index/database:Database example 48bfc13b-c1a5-48db-b70f-b6ef9709872b
// ```
type Database struct {
	pulumi.CustomResourceState

	// Additional keyspaces
	AdditionalKeyspaces pulumi.StringArrayOutput `pulumi:"additionalKeyspaces"`
	// The cloud provider to launch the database. (Currently supported: aws, azure, gcp)
	CloudProvider pulumi.StringOutput `pulumi:"cloudProvider"`
	// The cqlsh_url
	CqlshUrl pulumi.StringOutput `pulumi:"cqlshUrl"`
	// The data_endpoint_url
	DataEndpointUrl pulumi.StringOutput `pulumi:"dataEndpointUrl"`
	// Map of Datacenter IDs. The map key is "cloud_provider.region". Example: "GCP.us-east4".
	Datacenters pulumi.StringMapOutput `pulumi:"datacenters"`
	// The grafana_url
	GrafanaUrl pulumi.StringOutput `pulumi:"grafanaUrl"`
	// The graphql_url
	GraphqlUrl pulumi.StringOutput `pulumi:"graphqlUrl"`
	// Initial keyspace name. For additional keyspaces, use the astra_keyspace resource.
	Keyspace pulumi.StringOutput `pulumi:"keyspace"`
	// Astra database name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The node_count
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// The org id.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The owner id.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Cloud regions to launch the database. (see https://docs.datastax.com/en/astra/docs/database-regions.html for supported
	// regions)
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// The replication_factor
	ReplicationFactor pulumi.IntOutput `pulumi:"replicationFactor"`
	// The status
	Status pulumi.StringOutput `pulumi:"status"`
	// The total_storage
	TotalStorage pulumi.IntOutput `pulumi:"totalStorage"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudProvider == nil {
		return nil, errors.New("invalid value for required argument 'CloudProvider'")
	}
	if args.Keyspace == nil {
		return nil, errors.New("invalid value for required argument 'Keyspace'")
	}
	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("astra:index/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("astra:index/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Additional keyspaces
	AdditionalKeyspaces []string `pulumi:"additionalKeyspaces"`
	// The cloud provider to launch the database. (Currently supported: aws, azure, gcp)
	CloudProvider *string `pulumi:"cloudProvider"`
	// The cqlsh_url
	CqlshUrl *string `pulumi:"cqlshUrl"`
	// The data_endpoint_url
	DataEndpointUrl *string `pulumi:"dataEndpointUrl"`
	// Map of Datacenter IDs. The map key is "cloud_provider.region". Example: "GCP.us-east4".
	Datacenters map[string]string `pulumi:"datacenters"`
	// The grafana_url
	GrafanaUrl *string `pulumi:"grafanaUrl"`
	// The graphql_url
	GraphqlUrl *string `pulumi:"graphqlUrl"`
	// Initial keyspace name. For additional keyspaces, use the astra_keyspace resource.
	Keyspace *string `pulumi:"keyspace"`
	// Astra database name.
	Name *string `pulumi:"name"`
	// The node_count
	NodeCount *int `pulumi:"nodeCount"`
	// The org id.
	OrganizationId *string `pulumi:"organizationId"`
	// The owner id.
	OwnerId *string `pulumi:"ownerId"`
	// Cloud regions to launch the database. (see https://docs.datastax.com/en/astra/docs/database-regions.html for supported
	// regions)
	Regions []string `pulumi:"regions"`
	// The replication_factor
	ReplicationFactor *int `pulumi:"replicationFactor"`
	// The status
	Status *string `pulumi:"status"`
	// The total_storage
	TotalStorage *int `pulumi:"totalStorage"`
}

type DatabaseState struct {
	// Additional keyspaces
	AdditionalKeyspaces pulumi.StringArrayInput
	// The cloud provider to launch the database. (Currently supported: aws, azure, gcp)
	CloudProvider pulumi.StringPtrInput
	// The cqlsh_url
	CqlshUrl pulumi.StringPtrInput
	// The data_endpoint_url
	DataEndpointUrl pulumi.StringPtrInput
	// Map of Datacenter IDs. The map key is "cloud_provider.region". Example: "GCP.us-east4".
	Datacenters pulumi.StringMapInput
	// The grafana_url
	GrafanaUrl pulumi.StringPtrInput
	// The graphql_url
	GraphqlUrl pulumi.StringPtrInput
	// Initial keyspace name. For additional keyspaces, use the astra_keyspace resource.
	Keyspace pulumi.StringPtrInput
	// Astra database name.
	Name pulumi.StringPtrInput
	// The node_count
	NodeCount pulumi.IntPtrInput
	// The org id.
	OrganizationId pulumi.StringPtrInput
	// The owner id.
	OwnerId pulumi.StringPtrInput
	// Cloud regions to launch the database. (see https://docs.datastax.com/en/astra/docs/database-regions.html for supported
	// regions)
	Regions pulumi.StringArrayInput
	// The replication_factor
	ReplicationFactor pulumi.IntPtrInput
	// The status
	Status pulumi.StringPtrInput
	// The total_storage
	TotalStorage pulumi.IntPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The cloud provider to launch the database. (Currently supported: aws, azure, gcp)
	CloudProvider string `pulumi:"cloudProvider"`
	// Initial keyspace name. For additional keyspaces, use the astra_keyspace resource.
	Keyspace string `pulumi:"keyspace"`
	// Astra database name.
	Name *string `pulumi:"name"`
	// Cloud regions to launch the database. (see https://docs.datastax.com/en/astra/docs/database-regions.html for supported
	// regions)
	Regions []string `pulumi:"regions"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The cloud provider to launch the database. (Currently supported: aws, azure, gcp)
	CloudProvider pulumi.StringInput
	// Initial keyspace name. For additional keyspaces, use the astra_keyspace resource.
	Keyspace pulumi.StringInput
	// Astra database name.
	Name pulumi.StringPtrInput
	// Cloud regions to launch the database. (see https://docs.datastax.com/en/astra/docs/database-regions.html for supported
	// regions)
	Regions pulumi.StringArrayInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//          DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//          DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// Additional keyspaces
func (o DatabaseOutput) AdditionalKeyspaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Database) pulumi.StringArrayOutput { return v.AdditionalKeyspaces }).(pulumi.StringArrayOutput)
}

// The cloud provider to launch the database. (Currently supported: aws, azure, gcp)
func (o DatabaseOutput) CloudProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CloudProvider }).(pulumi.StringOutput)
}

// The cqlsh_url
func (o DatabaseOutput) CqlshUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CqlshUrl }).(pulumi.StringOutput)
}

// The data_endpoint_url
func (o DatabaseOutput) DataEndpointUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DataEndpointUrl }).(pulumi.StringOutput)
}

// Map of Datacenter IDs. The map key is "cloud_provider.region". Example: "GCP.us-east4".
func (o DatabaseOutput) Datacenters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.Datacenters }).(pulumi.StringMapOutput)
}

// The grafana_url
func (o DatabaseOutput) GrafanaUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.GrafanaUrl }).(pulumi.StringOutput)
}

// The graphql_url
func (o DatabaseOutput) GraphqlUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.GraphqlUrl }).(pulumi.StringOutput)
}

// Initial keyspace name. For additional keyspaces, use the astra_keyspace resource.
func (o DatabaseOutput) Keyspace() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Keyspace }).(pulumi.StringOutput)
}

// Astra database name.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The node_count
func (o DatabaseOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

// The org id.
func (o DatabaseOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The owner id.
func (o DatabaseOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Cloud regions to launch the database. (see https://docs.datastax.com/en/astra/docs/database-regions.html for supported
// regions)
func (o DatabaseOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Database) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// The replication_factor
func (o DatabaseOutput) ReplicationFactor() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.ReplicationFactor }).(pulumi.IntOutput)
}

// The status
func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The total_storage
func (o DatabaseOutput) TotalStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.TotalStorage }).(pulumi.IntOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
