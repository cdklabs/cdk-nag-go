<!--
Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
-->

# cdk-nag

[![PyPI version](https://img.shields.io/pypi/v/cdk-nag)](https://pypi.org/project/cdk-nag/)
[![npm version](https://img.shields.io/npm/v/cdk-nag)](https://www.npmjs.com/package/cdk-nag)
[![Maven version](https://img.shields.io/maven-central/v/io.github.cdklabs/cdknag)](https://search.maven.org/search?q=a:cdknag)
[![NuGet version](https://img.shields.io/nuget/v/Cdklabs.CdkNag)](https://www.nuget.org/packages/Cdklabs.CdkNag)
[![Go version](https://img.shields.io/github/go-mod/go-version/cdklabs/cdk-nag-go?color=blue&filename=cdknag%2Fgo.mod)](https://github.com/cdklabs/cdk-nag-go)

[![View on Construct Hub](https://constructs.dev/badge?package=cdk-nag)](https://constructs.dev/packages/cdk-nag)

Check CDK applications or [CloudFormation templates](#using-on-cloudformation-templates) for best practices using a combination of available rule packs. Inspired by [cfn_nag](https://github.com/stelligent/cfn_nag).

Check out [this blog post](https://aws.amazon.com/blogs/devops/manage-application-security-and-compliance-with-the-aws-cloud-development-kit-and-cdk-nag/) for a guided overview!

![demo](cdk_nag.gif)

## Available Rules and Packs

See [RULES](./RULES.md) for more information on all the available packs.

1. [AWS Solutions](./RULES.md#awssolutions)
2. [HIPAA Security](./RULES.md#hipaa-security)
3. [NIST 800-53 rev 4](./RULES.md#nist-800-53-rev-4)
4. [NIST 800-53 rev 5](./RULES.md#nist-800-53-rev-5)
5. [PCI DSS 3.2.1](./RULES.md#pci-dss-321)
6. [Serverless](./RULES.md#serverless)

[RULES](./RULES.md) also includes a collection of [additional rules](./RULES.md#additional-rules) that are not currently included in any of the pre-built NagPacks, but are still available for inclusion in custom NagPacks.

Read the [NagPack developer docs](./docs/NagPack.md) if you are interested in creating your own pack.

## Usage

For a full list of options See `NagPackProps` in the [API.md](./API.md#struct-nagpackprops)

<details>
<summary>Including in an application</summary>

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"

var cdkTestStack interface{}

app := awscdk.NewApp()
NewCdkTestStack(app, jsii.String("CdkNagDemo"))
// Simple rule informational messages using the AWS Solutions Rule pack
awscdk.Validations_Of(app).AddPlugins(cdknag.NewAwsSolutionsChecks(app))
// Multiple rule packs can be run against the same app
awscdk.Validations_Of(app).AddPlugins(cdknag.NewNIST80053R5Checks(app))
```

</details>

## Acknowledging a Rule

Use CDK's native `Validations.of()` API to acknowledge (suppress) rule violations on specific constructs.

<details>
  <summary>Example 1) Acknowledging a rule on a construct</summary>

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"

type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	test := awscdk.NewSecurityGroup(this, jsii.String("test"), &SecurityGroupProps{
		Vpc: awscdk.NewVpc(this, jsii.String("vpc")),
	})
	test.AddIngressRule(awscdk.Peer_AnyIpv4(), awscdk.Port_AllTraffic())
	awscdk.Validations_Of(test).Acknowledge(&Acknowledgment{
		Id: jsii.String("AwsSolutions-EC23"),
		Reason: jsii.String("This security group is used for internal testing only."),
	})
	return this
}
```

</details><details>
  <summary>Example 2) Acknowledging a rule on a stack</summary>

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"

var cdkTestStack interface{}

app := awscdk.NewApp()
stack := NewCdkTestStack(app, jsii.String("CdkNagDemo"))
awscdk.Validations_Of(app).AddPlugins(cdknag.NewAwsSolutionsChecks(app))
awscdk.Validations_Of(stack).Acknowledge(&Acknowledgment{
	Id: jsii.String("AwsSolutions-EC23"),
	Reason: jsii.String("All security groups in this stack are internal only."),
})
```

</details><details>
  <summary>Example 3) Acknowledging a specific finding</summary>

Certain rules report multiple findings per resource (e.g., IAM wildcard permissions). Each finding has its own ID in the format `RuleId[FindingId]`.

If you received the following errors on synth/deploy:

```bash
[Error at /StackName/rUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Action::s3:*]: The IAM entity contains wildcard permissions.
[Error at /StackName/rUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Resource::*]: The IAM entity contains wildcard permissions.
```

You can acknowledge a specific finding:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"

type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	user := awscdk.NewUser(this, jsii.String("rUser"))
	user.AddToPolicy(
	awscdk.NewPolicyStatement(&PolicyStatementProps{
		Actions: []*string{
			jsii.String("s3:*"),
		},
		Resources: []*string{
			jsii.String("*"),
		},
	}))
	// Only acknowledge the s3:* action — Resource::* still triggers
	awscdk.Validations_Of(user).Acknowledge(&Acknowledgment{
		Id: jsii.String("AwsSolutions-IAM5[Action::s3:*]"),
		Reason: jsii.String("Need s3:* for cross-account replication."),
	})
	return this
}
```

</details>

## Rules and Property Overrides

In some cases L2 Constructs do not have a native option to remediate an issue and must be fixed via [Raw Overrides](https://docs.aws.amazon.com/cdk/latest/guide/cfn_layer.html#cfn_layer_raw). Since raw overrides take place after template synthesis these fixes are not caught by cdk-nag. In this case you should remediate the issue and acknowledge the rule.

<details>
  <summary>Example) Property Overrides</summary>

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"

type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	instance := awscdk.NewInstance(this, jsii.String("rInstance"), &InstanceProps{
		Vpc: awscdk.NewVpc(this, jsii.String("rVpc")),
		InstanceType: awscdk.NewInstanceType(awscdk.InstanceClass_T3),
		MachineImage: awscdk.MachineImage_LatestAmazonLinux(),
	})
	cfnIns := instance.Node.defaultChild.(CfnInstance)
	cfnIns.AddPropertyOverride(jsii.String("DisableApiTermination"), jsii.Boolean(true))
	awscdk.Validations_Of(instance).Acknowledge(&Acknowledgment{
		Id: jsii.String("AwsSolutions-EC29"),
		Reason: jsii.String("Remediated through property override."),
	})
	return this
}
```

</details>

## Audit Trail: CloudFormation Metadata

By default, cdk-nag writes violations to CDK's `policy-validation-report.json` in the cloud assembly. If you need the v2-compatible `cdk_nag` metadata block in your synthesized CloudFormation templates (for existing compliance tooling), enable `writeSuppressionsToCloudFormation`:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"

app := awscdk.NewApp()
// Writes acknowledged rules into CfnResource Metadata as cdk_nag: { rules_to_suppress: [...] }
awscdk.Validations_Of(app).AddPlugins(cdknag.NewAwsSolutionsChecks(app, &NagPackProps{
	WriteSuppressionsToCloudFormation: jsii.Boolean(true),
}))
```

This registers a `WriteNagSuppressionsToCloudFormationAspect` that runs during synthesis and copies `Validations.of().acknowledge()` data into the CloudFormation template Metadata section, preserving the same format as cdk-nag v2.

## Using on CloudFormation templates

You can use cdk-nag on existing CloudFormation templates by using the [cloudformation-include](https://docs.aws.amazon.com/cdk/latest/guide/use-cfn-template.html#use-cfn-template-import) module.

<details>
  <summary>Example) CloudFormation template</summary>

Sample App

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"

var cdkTestStack interface{}

app := awscdk.NewApp()
NewCdkTestStack(app, jsii.String("CdkNagDemo"))
awscdk.Validations_Of(app).AddPlugins(cdknag.NewAwsSolutionsChecks(app))
```

Sample Stack with imported template

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"

type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	template := awscdk.NewCfnInclude(this, jsii.String("Template"), &CfnIncludeProps{
		TemplateFile: jsii.String("my-template.json"),
	})
	// Acknowledge rules on imported resources
	bucket := template.GetResource(jsii.String("rBucket"))
	awscdk.Validations_Of(bucket).Acknowledge(&Acknowledgment{
		Id: jsii.String("AwsSolutions-S1"),
		Reason: jsii.String("Logging not required for this bucket."),
	})
	return this
}
```

</details>

## Migrating from v2

cdk-nag v3 replaces the custom `NagSuppressions` API with CDK's native `Validations.of().acknowledge()` mechanism.

| v2 | v3 |
|---|---|
| `NagSuppressions.addResourceSuppressions(construct, [{ id, reason }])` | `Validations.of(construct).acknowledge({ id, reason })` |
| `NagSuppressions.addStackSuppressions(stack, [{ id, reason }])` | `Validations.of(stack).acknowledge({ id, reason })` |
| `NagSuppressions.addResourceSuppressionsByPath(stack, path, [...])` | `Validations.of(construct).acknowledge({ id, reason })` |
| `appliesTo: ['Action::s3:*']` | `id: 'AwsSolutions-IAM5[Action::s3:*]'` |
| `{ id: 'CdkNagValidationFailure', reason: '...' }` | `Validations.of(construct).acknowledge({ id: 'RuleId', reason: '...' })` |

**Note on bulk suppression:** In v2, suppressing a rule without `appliesTo` would suppress all findings for that rule on the construct. In v3, each finding must be acknowledged individually (e.g., `AwsSolutions-IAM5[Action::s3:*]` and `AwsSolutions-IAM5[Resource::*]` are separate acknowledgments). Prefix matching (acknowledging `AwsSolutions-IAM5` to suppress all findings) is not yet supported — tracked via [issue link].

**Removed APIs:**

* `NagSuppressions` (use `Validations.of().acknowledge()`)
* `INagSuppressionIgnore` and all condition classes
* `NagPackSuppression` interface
* `CdkNagValidationFailure` concept
* `logIgnores` and `suppressionIgnoreCondition` props

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md) for more information.

## License

This project is licensed under the Apache-2.0 License.
