package ec2

var cloudFormationASEnableAtRestEncryptionGoodExamples = []string{
	`---
Resources:
  GoodExample:
    Properties:
      BlockDeviceMappings:
        - DeviceName: root
          Ebs:
            Encrypted: true
      ImageId: ami-123456
      InstanceType: t2.small
    Type: AWS::AutoScaling::LaunchConfiguration
`,
}

var cloudFormationASEnableAtRestEncryptionBadExamples = []string{
	`---
Resources:
  BadExample:
    Properties:
      BlockDeviceMappings:
        - DeviceName: root
          Ebs:
            Encrypted: true
        - DeviceName: data
          Ebs:
            Encrypted: false
      ImageId: ami-123456
      InstanceType: t2.small
    Type: AWS::AutoScaling::LaunchConfiguration
`,
}

var cloudFormationASEnableAtRestEncryptionLinks = []string{}

var cloudFormationASEnableAtRestEncryptionRemediationMarkdown = ``
