package main

import (
	"log"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

func applyDeploymentExtensionsV1beta1(deployment *extensionsv1beta1.Deployment, deploymentSet v1beta1.DeploymentInterface) error {
	deploymentName := deployment.GetObjectMeta().GetName()
	deployments, err := deploymentSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing deployments")
		return err
	}

	update := false
	for _, dep := range deployments.Items {
		if dep.GetObjectMeta().GetName() == deploymentName {
			update = true
		}
	}

	if update {
		_, err := deploymentSet.Get(deploymentName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old deployment")
			return err
		}

		_, err = deploymentSet.Update(deployment)
		if err != nil {
			log.Println("Error when updating deployment")
			return err
		}
		log.Println("Deployment " + deploymentName + " updated")

		return err
	} else {
		_, err := deploymentSet.Create(deployment)
		if err != nil {
			log.Println("Error when creating deployment")
			return err
		}

		log.Println("Deployment " + deploymentName + " created")
		return err
	}
}

func applyDaemonSetExtensionsV1beta1(daemonSet *extensionsv1beta1.DaemonSet, daemonSetSet v1beta1.DaemonSetInterface) error {
	daemonSetName := daemonSet.GetObjectMeta().GetName()
	daemonSets, err := daemonSetSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing daemon sets")
		return err
	}

	update := false
	for _, dep := range daemonSets.Items {
		if dep.GetObjectMeta().GetName() == daemonSetName {
			update = true
		}
	}

	if update {
		_, err := daemonSetSet.Get(daemonSetName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old daemon set")
			return err
		}

		_, err = daemonSetSet.Update(daemonSet)
		if err != nil {
			log.Println("Error when updating daemonSet")
			return err
		}
		log.Println("Deployment " + daemonSetName + " updated")

		return err
	} else {
		_, err := daemonSetSet.Create(daemonSet)
		if err != nil {
			log.Println("Error when creating daemonSet")
			return err
		}

		log.Println("Deployment " + daemonSetName + " created")
		return err
	}
}

func applyReplicaSetExtensionsV1beta1(replicaSet *extensionsv1beta1.ReplicaSet, replicaSetSet v1beta1.ReplicaSetInterface) error {
	replicaSetName := replicaSet.GetObjectMeta().GetName()
	replicaSets, err := replicaSetSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing replica sets")
		return err
	}

	update := false
	for _, rep := range replicaSets.Items {
		if rep.GetObjectMeta().GetName() == replicaSetName {
			update = true
		}
	}

	if update {
		_, err := replicaSetSet.Get(replicaSetName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old replica set")
			return err
		}

		_, err = replicaSetSet.Update(replicaSet)
		if err != nil {
			log.Println("Error when updating replicaSet")
			return err
		}
		log.Println("ReplicaSet " + replicaSetName + " updated")

		return err
	} else {
		_, err := replicaSetSet.Create(replicaSet)
		if err != nil {
			log.Println("Error when creating replicaSet")
			return err
		}

		log.Println("ReplicaSet " + replicaSetName + " created")
		return err
	}
}

func applyIngressExtensionsV1beta1(ingress *extensionsv1beta1.Ingress, ingressSet v1beta1.IngressInterface) error {
	ingressName := ingress.GetObjectMeta().GetName()
	ingresss, err := ingressSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing ingress")
		return err
	}

	update := false
	for _, ing := range ingresss.Items {
		if ing.GetObjectMeta().GetName() == ingressName {
			update = true
		}
	}

	if update {
		_, err := ingressSet.Get(ingressName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old ingress")
			return err
		}

		_, err = ingressSet.Update(ingress)
		if err != nil {
			log.Println("Error when updating ingress")
			return err
		}
		log.Println("Ingress " + ingressName + " updated")

		return err
	} else {
		_, err := ingressSet.Create(ingress)
		if err != nil {
			log.Println("Error when creating ingress")
			return err
		}

		log.Println("Ingress " + ingressName + " created")
		return err
	}
}
