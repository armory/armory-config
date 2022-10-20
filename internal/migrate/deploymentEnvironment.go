package migrate

import "strings"

func GetDeploymentEnvironment(KustomizeData Kustomize) string {

	return GetCustomSizing(KustomizeData)
}

func GetCustomSizing(KustomizeData Kustomize) string {
	str := `
	customSizing:
	  # This applies sizings to the clouddriver container as well as any sidecar 
	  # containers running with clouddriver. (Use without spin- to only include the clouddriver container)
	  spin-clouddriver:
	    replicas: 1     # Set the number of replicas (pods) to be created for this service.
	    # limits:
	    #   cpu: 2        # Set the kubernetes CPU limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 4 cpu.
	    #   memory: 4Gi   # Set the kubernetes memory limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 16Gi of memory.
	    # requests:
	    #   cpu: 500m
	    #   memory: 2Gi
	  # This applies sizings to only the service container and not to any sidecar 
	  # containers running with service. (Use spin-<service> to include sidecars)
	  clouddriver:
	    limits:
	      cpu: 2        # Set the kubernetes CPU limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 4 cpu.
	      memory: 4Gi   # Set the kubernetes memory limits for each pod. For reference a m5.xlarge EC2 instance has a capacity of 16Gi of memory.
	    requests:
	      cpu: 500m
	      memory: 2Gi
	  deck:`

	str = strings.Replace(str, "\t", "        ", -1)
	return str
}
