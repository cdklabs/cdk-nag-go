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
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/lib/cdkteststack"
import "github.com/cdklabs/cdk-nag-go/cdknag"


app := awscdk.NewApp()
libcdkteststack.NewCdkTestStack(app, jsii.String("CdkNagDemo"))
// Simple rule informational messages using the AWS Solutions Rule pack
awscdk.Aspects_Of(app).Add(cdknag.NewAwsSolutionsChecks())
// Multiple rule packs can be run against the same app
awscdk.Aspects_Of(app).Add(NewNIST80053R5Checks())
```

</details>

## Suppressing a Rule

<details>
  <summary>Example 1) Default Construct</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


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
	cdknag.NagSuppressions_AddResourceSuppressions(test, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-EC23"),
			Reason: jsii.String("lorem ipsum"),
		},
	})
	return this
}
```

</details><details>
  <summary>Example 2) On Multiple Constructs</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	vpc := awscdk.NewVpc(this, jsii.String("vpc"))
	test1 := awscdk.NewSecurityGroup(this, jsii.String("test"), &SecurityGroupProps{
		Vpc: Vpc,
	})
	test1.AddIngressRule(awscdk.Peer_AnyIpv4(), awscdk.Port_AllTraffic())
	test2 := awscdk.NewSecurityGroup(this, jsii.String("test"), &SecurityGroupProps{
		Vpc: Vpc,
	})
	test2.AddIngressRule(awscdk.Peer_AnyIpv4(), awscdk.Port_AllTraffic())
	cdknag.NagSuppressions_AddResourceSuppressions([]interface{}{
		test1,
		test2,
	}, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-EC23"),
			Reason: jsii.String("lorem ipsum"),
		},
	})
	return this
}
```

</details><details>
  <summary>Example 3) Child Constructs</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


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
			jsii.String("s3:PutObject"),
		},
		Resources: []*string{
			jsii.String("arn:aws:s3:::bucket_name/*"),
		},
	}))
	// Enable adding suppressions to child constructs
	cdknag.NagSuppressions_AddResourceSuppressions(user, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-IAM5"),
			Reason: jsii.String("lorem ipsum"),
			AppliesTo: []nagPackSuppressionAppliesTo{
				jsii.String("Resource::arn:aws:s3:::bucket_name/*"),
			},
		},
	}, jsii.Boolean(true))
	return this
}
```

</details><details>
  <summary>Example 4) Stack Level </summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/lib/cdkteststack"
import "github.com/cdklabs/cdk-nag-go/cdknag"


app := awscdk.NewApp()
stack := libcdkteststack.NewCdkTestStack(app, jsii.String("CdkNagDemo"))
awscdk.Aspects_Of(app).Add(cdknag.NewAwsSolutionsChecks())
cdknag.NagSuppressions_AddStackSuppressions(stack, []NagPackSuppression{
	&NagPackSuppression{
		Id: jsii.String("AwsSolutions-EC23"),
		Reason: jsii.String("lorem ipsum"),
	},
})
```

</details><details>
  <summary>Example 5) Construct path</summary>

If you received the following error on synth/deploy

```bash
[Error at /StackName/Custom::CDKBucketDeployment8675309/ServiceRole/Resource] AwsSolutions-IAM4: The IAM user, role, or group uses AWS managed policies
```

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	awscdk.NewBucketDeployment(this, jsii.String("rDeployment"), &BucketDeploymentProps{
		Sources: []ISource{
		},
		DestinationBucket: awscdk.Bucket_FromBucketName(this, jsii.String("rBucket"), jsii.String("foo")),
	})
	cdknag.NagSuppressions_AddResourceSuppressionsByPath(this, jsii.String("/StackName/Custom::CDKBucketDeployment8675309/ServiceRole/Resource"), []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-IAM4"),
			Reason: jsii.String("at least 10 characters"),
		},
	})
	return this
}
```

</details><details>
  <summary>Example 6) Granular Suppressions of findings</summary>

Certain rules support granular suppressions of `findings`. If you received the following errors on synth/deploy

```bash
[Error at /StackName/rFirstUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Action::s3:*]: The IAM entity contains wildcard permissions and does not have a cdk-nag rule suppression with evidence for those permission.
[Error at /StackName/rFirstUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Resource::*]: The IAM entity contains wildcard permissions and does not have a cdk-nag rule suppression with evidence for those permission.
[Error at /StackName/rSecondUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Action::s3:*]: The IAM entity contains wildcard permissions and does not have a cdk-nag rule suppression with evidence for those permission.
[Error at /StackName/rSecondUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Resource::*]: The IAM entity contains wildcard permissions and does not have a cdk-nag rule suppression with evidence for those permission.
```

By applying the following suppressions

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	firstUser := awscdk.NewUser(this, jsii.String("rFirstUser"))
	firstUser.AddToPolicy(
	NewPolicyStatement(map[string][]*string{
		"actions": []*string{
			jsii.String("s3:*"),
		},
		"resources": []*string{
			jsii.String("*"),
		},
	}))
	secondUser := awscdk.NewUser(this, jsii.String("rSecondUser"))
	secondUser.AddToPolicy(
	NewPolicyStatement(map[string][]*string{
		"actions": []*string{
			jsii.String("s3:*"),
		},
		"resources": []*string{
			jsii.String("*"),
		},
	}))
	thirdUser := awscdk.NewUser(this, jsii.String("rSecondUser"))
	thirdUser.AddToPolicy(
	NewPolicyStatement(map[string][]*string{
		"actions": []*string{
			jsii.String("sqs:CreateQueue"),
		},
		"resources": []*string{
			fmt.Sprintf("arn:aws:sqs:%v:%v:*", this.region, this.account),
		},
	}))
	cdknag.NagSuppressions_AddResourceSuppressions(firstUser, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-IAM5"),
			Reason: jsii.String("Only suppress AwsSolutions-IAM5 's3:*' finding on First User."),
			AppliesTo: []nagPackSuppressionAppliesTo{
				jsii.String("Action::s3:*"),
			},
		},
	}, jsii.Boolean(true))
	cdknag.NagSuppressions_AddResourceSuppressions(secondUser, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-IAM5"),
			Reason: jsii.String("Suppress all AwsSolutions-IAM5 findings on Second User."),
		},
	}, jsii.Boolean(true))
	cdknag.NagSuppressions_AddResourceSuppressions(thirdUser, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-IAM5"),
			Reason: jsii.String("Suppress AwsSolutions-IAM5 on the SQS resource."),
			AppliesTo: []*nagPackSuppressionAppliesTo{
				&RegexAppliesTo{
					Regex: jsii.String("/^Resource::arn:aws:sqs:(.*):\\*$/g"),
				},
			},
		},
	}, jsii.Boolean(true))
	return this
}
```

You would see the following error on synth/deploy

```bash
[Error at /StackName/rFirstUser/DefaultPolicy/Resource] AwsSolutions-IAM5[Resource::*]: The IAM entity contains wildcard permissions and does not have a cdk-nag rule suppression with evidence for those permission.
```

</details>

## Suppressing Rule Validation Failures

When a rule validation fails it is handled similarly to a rule violation, and can be suppressed in the same manner. The `ID` for a rule failure is `CdkNagValidationFailure`.

If a rule is suppressed in a non-granular manner (i.e. `appliesTo` is not set, see example 1 above) then validation failures on that rule are also suppressed.

Validation failure suppression respects any applied [Suppression Ignore Conditions](#conditionally-ignoring-suppressions)

<details>
  <summary>Example 1) Suppress all Validation Failures on a Resource</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


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
	cdknag.NagSuppressions_AddResourceSuppressions(test, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("CdkNagValidationFailure"),
			Reason: jsii.String("lorem ipsum"),
		},
	})
	return this
}
```

</details><details>
  <summary>Example 2) Granular Suppression of Validation Failures</summary>
Validation failures can be suppressed for individual rules by using `appliesTo` to list the desired rules

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


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
	cdknag.NagSuppressions_AddResourceSuppressions(test, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("CdkNagValidationFailure"),
			Reason: jsii.String("lorem ipsum"),
			AppliesTo: []nagPackSuppressionAppliesTo{
				jsii.String("AwsSolutions-L1"),
			},
		},
	})
	return this
}
```

</details>

## Suppressing `aws-cdk-lib/pipelines` Violations

The [aws-cdk-lib/pipelines.CodePipeline](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.pipelines.CodePipeline.html) construct and its child constructs are not guaranteed to be "Visited" by `Aspects`, as they are not added during the "Construction" phase of the [cdk lifecycle](https://docs.aws.amazon.com/cdk/v2/guide/apps.html#lifecycle). Because of this behavior, you may experience problems such as rule violations not appearing or the inability to suppress violations on these constructs.

You can remediate these rule violation and suppression problems by forcing the pipeline construct creation forward by calling `.buildPipeline()` on your `CodePipeline` object. Otherwise you may see errors such as:

```
Error: Suppression path "/this/construct/path" did not match any resource. This can occur when a resource does not exist or if a suppression is applied before a resource is created.
```

See [this issue](https://github.com/aws/aws-cdk/issues/18440) for more information.

<details>
  <summary>Example) Suppressing Violations in Pipelines</summary>

`example-app.ts`

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"
import "github.com/aws-samples/dummy/lib/examplepipeline"


app := awscdk.NewApp()
libexamplepipeline.NewExamplePipeline(app, jsii.String("example-cdk-pipeline"))
awscdk.Aspects_Of(app).Add(cdknag.NewAwsSolutionsChecks(&NagPackProps{
	Verbose: jsii.Boolean(true),
}))
app.Synth()
```

`example-pipeline.ts`

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"
import "github.com/aws/constructs-go/constructs"


type ExamplePipeline struct {
	Stack
}

func NewExamplePipeline(scope Construct, id *string, props StackProps) *ExamplePipeline {
	this := &ExamplePipeline{}
	newStack_Override(this, scope, id, props)

	exampleSynth := awscdk.NewShellStep(jsii.String("ExampleSynth"), &ShellStepProps{
		Commands: []*string{
			jsii.String("yarn build --frozen-lockfile"),
		},
		Input: awscdk.CodePipelineSource_CodeCommit(
		awscdk.NewRepository(this, jsii.String("ExampleRepo"), &RepositoryProps{
			RepositoryName: jsii.String("ExampleRepo"),
		}), jsii.String("main")),
	})

	examplePipeline := awscdk.NewCodePipeline(this, jsii.String("ExamplePipeline"), &CodePipelineProps{
		Synth: exampleSynth,
	})

	// Force the pipeline construct creation forward before applying suppressions.
	// @See https://github.com/aws/aws-cdk/issues/18440
	examplePipeline.BuildPipeline()

	// The path suppression will error if you comment out "ExamplePipeline.buildPipeline();""
	cdknag.NagSuppressions_AddResourceSuppressionsByPath(this, jsii.String("/example-cdk-pipeline/ExamplePipeline/Pipeline/ArtifactsBucket/Resource"), []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-S1"),
			Reason: jsii.String("Because I said so"),
		},
	})
	return this
}
```

</details>

## Rules and Property Overrides

In some cases L2 Constructs do not have a native option to remediate an issue and must be fixed via [Raw Overrides](https://docs.aws.amazon.com/cdk/latest/guide/cfn_layer.html#cfn_layer_raw). Since raw overrides take place after template synthesis these fixes are not caught by cdk-nag. In this case you should remediate the issue and suppress the issue like in the following example.

<details>
  <summary>Example) Property Overrides</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"
import "github.com/cdklabs/cdk-nag-go/cdknag"


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
	cdknag.NagSuppressions_AddResourceSuppressions(instance, []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-EC29"),
			Reason: jsii.String("Remediated through property override."),
		},
	})
	return this
}
```

</details>

## Conditionally Ignoring Suppressions

You can optionally create a condition that prevents certain rules from being suppressed. You can create conditions for any variety of reasons. Examples include a condition that always ignores a suppression, a condition that ignores a suppression based on the date, a condition that ignores a suppression based on the reason. You can read [the developer docs](./docs/IgnoreSuppressionConditions.md) for more information on creating your own conditions.

<details>
  <summary>Example) Using the pre-built `SuppressionIgnoreErrors` class to ignore suppressions on any `Error` level rules.</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/lib/cdkteststack"
import "github.com/cdklabs/cdk-nag-go/cdknag"


app := awscdk.NewApp()
libcdkteststack.NewCdkTestStack(app, jsii.String("CdkNagDemo"))
// Ignore Suppressions on any errors
awscdk.Aspects_Of(app).Add(
cdknag.NewAwsSolutionsChecks(&NagPackProps{
	SuppressionIgnoreCondition: *cdknag.NewSuppressionIgnoreErrors(),
}))
```

</details>

## Customizing Logging

`NagLogger`s give `NagPack` authors and users the ability to create their own custom reporting mechanisms. All pre-built `NagPacks`come with the `AnnotationsLogger`and the `NagReportLogger` (with CSV reports) enabled by default.

See the [NagLogger](./docs/NagLogger.md) developer docs for more information.

<details>
  <summary>Example) Adding the `ExtremelyHelpfulConsoleLogger` example from the NagLogger docs</summary>

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/lib/cdkteststack"
import "github.com/aws-samples/dummy/docs/nagLogger"
import "github.com/cdklabs/cdk-nag-go/cdknag"


app := awscdk.NewApp()
libcdkteststack.NewCdkTestStack(app, jsii.String("CdkNagDemo"))
awscdk.Aspects_Of(app).Add(
cdknag.NewAwsSolutionsChecks(&NagPackProps{
	AdditionalLoggers: []INagLogger{
		docsNagLogger.NewExtremelyHelpfulConsoleLogger(),
	},
}))
```

</details>

## Using on CloudFormation templates

You can use cdk-nag on existing CloudFormation templates by using the [cloudformation-include](https://docs.aws.amazon.com/cdk/latest/guide/use-cfn-template.html#use-cfn-template-import) module.

<details>
  <summary>Example 1) CloudFormation template with suppression</summary>

Sample CloudFormation template with suppression

```json
{
  "Resources": {
    "rBucket": {
      "Type": "AWS::S3::Bucket",
      "Properties": {
        "BucketName": "some-bucket-name"
      },
      "Metadata": {
        "cdk_nag": {
          "rules_to_suppress": [
            {
              "id": "AwsSolutions-S1",
              "reason": "at least 10 characters"
            }
          ]
        }
      }
    }
  }
}
```

Sample App

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/lib/cdkteststack"
import "github.com/cdklabs/cdk-nag-go/cdknag"


app := awscdk.NewApp()
libcdkteststack.NewCdkTestStack(app, jsii.String("CdkNagDemo"))
awscdk.Aspects_Of(app).Add(cdknag.NewAwsSolutionsChecks())
```

Sample Stack with imported template

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"


type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	awscdk.NewCfnInclude(this, jsii.String("Template"), &CfnIncludeProps{
		TemplateFile: jsii.String("my-template.json"),
	})
	// Add any additional suppressions
	cdknag.NagSuppressions_AddResourceSuppressionsByPath(this, jsii.String("/CdkNagDemo/Template/rBucket"), []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-S2"),
			Reason: jsii.String("at least 10 characters"),
		},
	})
	return this
}
```

</details><details>
  <summary>Example 2) CloudFormation template with granular suppressions</summary>

Sample CloudFormation template with suppression

```json
{
  "Resources": {
    "myPolicy": {
      "Type": "AWS::IAM::Policy",
      "Properties": {
        "PolicyDocument": {
          "Statement": [
            {
              "Action": [
                "kms:Decrypt",
                "kms:DescribeKey",
                "kms:Encrypt",
                "kms:ReEncrypt*",
                "kms:GenerateDataKey*"
              ],
              "Effect": "Allow",
              "Resource": ["some-key-arn"]
            }
          ],
          "Version": "2012-10-17"
        }
      },
      "Metadata": {
        "cdk_nag": {
          "rules_to_suppress": [
            {
              "id": "AwsSolutions-IAM5",
              "reason": "Allow key data access",
              "applies_to": [
                "Action::kms:ReEncrypt*",
                "Action::kms:GenerateDataKey*"
              ]
            }
          ]
        }
      }
    }
  }
}
```

Sample App

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/lib/cdkteststack"
import "github.com/cdklabs/cdk-nag-go/cdknag"


app := awscdk.NewApp()
libcdkteststack.NewCdkTestStack(app, jsii.String("CdkNagDemo"))
awscdk.Aspects_Of(app).Add(cdknag.NewAwsSolutionsChecks())
```

Sample Stack with imported template

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-nag-go/cdknag"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/constructs-go/constructs"


type CdkTestStack struct {
	Stack
}

func NewCdkTestStack(scope Construct, id *string, props StackProps) *CdkTestStack {
	this := &CdkTestStack{}
	newStack_Override(this, scope, id, props)
	awscdk.NewCfnInclude(this, jsii.String("Template"), &CfnIncludeProps{
		TemplateFile: jsii.String("my-template.json"),
	})
	// Add any additional suppressions
	cdknag.NagSuppressions_AddResourceSuppressionsByPath(this, jsii.String("/CdkNagDemo/Template/myPolicy"), []NagPackSuppression{
		&NagPackSuppression{
			Id: jsii.String("AwsSolutions-IAM5"),
			Reason: jsii.String("Allow key data access"),
			AppliesTo: []nagPackSuppressionAppliesTo{
				jsii.String("Action::kms:ReEncrypt*"),
				jsii.String("Action::kms:GenerateDataKey*"),
			},
		},
	})
	return this
}
```

</details>

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md) for more information.

## License

This project is licensed under the Apache-2.0 License.
