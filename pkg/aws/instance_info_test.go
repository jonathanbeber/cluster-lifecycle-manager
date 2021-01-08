package aws

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/stretchr/testify/require"
)

type mockEC2 struct {
	ec2iface.EC2API
}

func (mock *mockEC2) DescribeInstanceTypesPages(_ *ec2.DescribeInstanceTypesInput, callback func(*ec2.DescribeInstanceTypesOutput, bool) bool) error {
	if !callback(&ec2.DescribeInstanceTypesOutput{
		InstanceTypes: []*ec2.InstanceTypeInfo{
			{
				InstanceType: aws.String("m4.xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(4),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(16384),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{},
			},
			{
				InstanceType: aws.String("i3.4xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(16),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(124928),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{
					Disks: []*ec2.DiskInfo{
						{
							Count:    aws.Int64(2),
							SizeInGB: aws.Int64(1900),
						},
					},
				},
			},
			{
				InstanceType: aws.String("m5.xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(4),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(16384),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{},
			},
			{
				InstanceType: aws.String("m5d.4xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(16),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(65536),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{
					Disks: []*ec2.DiskInfo{
						{
							Count:    aws.Int64(2),
							SizeInGB: aws.Int64(300),
						},
					},
				},
			},
		},
	}, false) {
		return nil
	}
	callback(&ec2.DescribeInstanceTypesOutput{
		InstanceTypes: []*ec2.InstanceTypeInfo{
			{
				InstanceType: aws.String("c5d.xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(4),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(8192),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{
					Disks: []*ec2.DiskInfo{
						{
							Count:    aws.Int64(1),
							SizeInGB: aws.Int64(100),
						},
					},
				},
			},
			{
				InstanceType: aws.String("r5d.xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(4),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(32768),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{
					Disks: []*ec2.DiskInfo{
						{
							Count:    aws.Int64(1),
							SizeInGB: aws.Int64(150),
						},
					},
				},
			},
			{
				InstanceType: aws.String("m5d.xlarge"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(4),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(16384),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{
					Disks: []*ec2.DiskInfo{
						{
							Count:    aws.Int64(1),
							SizeInGB: aws.Int64(150),
						},
					},
				},
			},
			{
				InstanceType: aws.String("r5d.large"),
				VCpuInfo: &ec2.VCpuInfo{
					DefaultVCpus: aws.Int64(2),
				},
				MemoryInfo: &ec2.MemoryInfo{
					SizeInMiB: aws.Int64(16384),
				},
				InstanceStorageInfo: &ec2.InstanceStorageInfo{
					Disks: []*ec2.DiskInfo{
						{
							Count:    aws.Int64(1),
							SizeInGB: aws.Int64(75),
						},
					},
				},
			},
		},
	}, false)
	return nil
}

func TestMain(m *testing.M) {
	err := InitInstanceTypes(&mockEC2{})
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestInstanceInfo(t *testing.T) {
	for _, tc := range []Instance{
		{
			InstanceType: "m4.xlarge",
			VCPU:         4,
			Memory:       17179869184,
		},
		{
			InstanceType:              "i3.4xlarge",
			VCPU:                      16,
			Memory:                    130996502528,
			InstanceStorageDevices:    2,
			InstanceStorageDeviceSize: 1900 * gigabyte,
		},
	} {
		t.Run(tc.InstanceType, func(t *testing.T) {
			info, err := InstanceInfo(tc.InstanceType)
			require.NoError(t, err)
			require.Equal(t, tc.InstanceType, info.InstanceType)
			require.Equal(t, tc.VCPU, info.VCPU)
			require.Equal(t, tc.Memory, info.Memory)
			require.Equal(t, tc.InstanceStorageDevices, info.InstanceStorageDevices)
		})
	}
}

func TestInstanceInfoError(t *testing.T) {
	_, err := InstanceInfo("invalid.type")
	require.Error(t, err)
}

func TestSyntheticInstanceInfo(t *testing.T) {
	for _, tc := range []struct {
		name             string
		instanceTypes    []string
		expectedError    bool
		expectedInstance Instance
	}{
		{
			name:          "one type",
			instanceTypes: []string{"m4.xlarge"},
			expectedInstance: Instance{
				InstanceType: "m4.xlarge",
				VCPU:         4,
				Memory:       17179869184,
			},
		},
		{
			name:          "no types",
			instanceTypes: []string{},
			expectedError: true,
		},
		{
			name:          "invalid type (only one)",
			instanceTypes: []string{"invalid.type"},
			expectedError: true,
		},
		{
			name:          "invalid type (first)",
			instanceTypes: []string{"invalid.type", "m4.xlarge"},
			expectedError: true,
		},
		{
			name:          "invalid type (second)",
			instanceTypes: []string{"m4.xlarge", "invalid.type"},
			expectedError: true,
		},
		{
			name:          "multiple types, different storage device size",
			instanceTypes: []string{"c5d.xlarge", "r5d.large"},
			expectedInstance: Instance{
				InstanceType:              "<multiple>",
				VCPU:                      2,
				Memory:                    8589934592,
				InstanceStorageDevices:    1,
				InstanceStorageDeviceSize: 75 * 1024 * megabyte,
			},
		},
		{
			name:          "multiple types, different storage device count",
			instanceTypes: []string{"m5d.xlarge", "m5d.4xlarge"},
			expectedInstance: Instance{
				InstanceType:              "<multiple>",
				VCPU:                      4,
				Memory:                    17179869184,
				InstanceStorageDevices:    1,
				InstanceStorageDeviceSize: 150 * 1024 * megabyte,
			},
		},
		{
			name:          "multiple types, some with storage, some without",
			instanceTypes: []string{"m5d.xlarge", "m5.xlarge"},
			expectedInstance: Instance{
				InstanceType: "<multiple>",
				VCPU:         4,
				Memory:       17179869184,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			info, err := SyntheticInstanceInfo(tc.instanceTypes)
			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedInstance.InstanceType, info.InstanceType)
				require.Equal(t, tc.expectedInstance.VCPU, info.VCPU)
				require.Equal(t, tc.expectedInstance.Memory, info.Memory)
				require.Equal(t, tc.expectedInstance.InstanceStorageDevices, info.InstanceStorageDevices)
				require.Equal(t, tc.expectedInstance.InstanceStorageDeviceSize, info.InstanceStorageDeviceSize)
			}
		})
	}
}
