// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type KernelModuleSpecSpec -type KernelParamSpecSpec -type KernelParamStatusSpec -type MachineStatusSpec -type MountStatusSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package runtime

// DeepCopy generates a deep copy of KernelModuleSpecSpec.
func (o KernelModuleSpecSpec) DeepCopy() KernelModuleSpecSpec {
	var cp KernelModuleSpecSpec = o
	return cp
}

// DeepCopy generates a deep copy of KernelParamSpecSpec.
func (o KernelParamSpecSpec) DeepCopy() KernelParamSpecSpec {
	var cp KernelParamSpecSpec = o
	return cp
}

// DeepCopy generates a deep copy of KernelParamStatusSpec.
func (o KernelParamStatusSpec) DeepCopy() KernelParamStatusSpec {
	var cp KernelParamStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of MachineStatusSpec.
func (o MachineStatusSpec) DeepCopy() MachineStatusSpec {
	var cp MachineStatusSpec = o
	if o.Status.UnmetConditions != nil {
		cp.Status.UnmetConditions = make([]UnmetCondition, len(o.Status.UnmetConditions))
		copy(cp.Status.UnmetConditions, o.Status.UnmetConditions)
	}
	return cp
}

// DeepCopy generates a deep copy of MountStatusSpec.
func (o MountStatusSpec) DeepCopy() MountStatusSpec {
	var cp MountStatusSpec = o
	if o.Options != nil {
		cp.Options = make([]string, len(o.Options))
		copy(cp.Options, o.Options)
	}
	return cp
}
