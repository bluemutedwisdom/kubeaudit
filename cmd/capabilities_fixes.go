package cmd

import k8sRuntime "k8s.io/apimachinery/pkg/runtime"

func fixCapabilitiesNIL(resource k8sRuntime.Object) k8sRuntime.Object {
	containers := getContainers(resource)
	for _, container := range containers {
		if container.SecurityContext.Capabilities == nil {
			container.SecurityContext.Capabilities = &Capabilities{}
		}
		if container.SecurityContext.Capabilities.Drop == nil {
			container.SecurityContext.Capabilities.Drop = []Capability{}
		}
		if container.SecurityContext.Capabilities.Add == nil {
			container.SecurityContext.Capabilities.Add = []Capability{}
		}
	}
	return setContainers(resource, containers)
}

func fixCapabilityNotDropped(resource k8sRuntime.Object, occurrence Occurrence) k8sRuntime.Object {
	containers := getContainers(resource)
	for _, container := range containers {
		if occurrence.container == container.Name {
			container.SecurityContext.Capabilities.Drop = append(container.SecurityContext.Capabilities.Drop, Capability(occurrence.metadata["CapName"]))
		}
	}
	return setContainers(resource, containers)
}

func fixCapabilityAdded(resource k8sRuntime.Object, occurrence Occurrence) k8sRuntime.Object {
	containers := getContainers(resource)
	for _, container := range containers {
		if occurrence.container == container.Name {
			add := []Capability{}
			for _, cap := range container.SecurityContext.Capabilities.Add {
				if string(cap) != occurrence.metadata["CapName"] {
					add = append(add, cap)
				}
			}
			container.SecurityContext.Capabilities.Add = add
		}
	}
	return setContainers(resource, containers)
}
