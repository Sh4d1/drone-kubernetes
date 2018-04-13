package main

import (
	"log"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta2"
)

func applyDeploymentAppsV1beta2(deployment *appsv1beta2.Deployment, deploymentSet v1beta2.DeploymentInterface) error {
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

func applyDaemonSetAppsV1beta2(daemonSet *appsv1beta2.DaemonSet, daemonSetSet v1beta2.DaemonSetInterface) error {
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

func applyReplicaSetAppsV1beta2(replicaSet *appsv1beta2.ReplicaSet, replicaSetSet v1beta2.ReplicaSetInterface) error {
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

func applyStatefulSetAppsV1beta2(statefulSet *appsv1beta2.StatefulSet, statefulSetSet v1beta2.StatefulSetInterface) error {
	statefulSetName := statefulSet.GetObjectMeta().GetName()
	statefulSets, err := statefulSetSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing stateful sets")
		return err
	}

	update := false
	for _, sta := range statefulSets.Items {
		if sta.GetObjectMeta().GetName() == statefulSetName {
			update = true
		}
	}

	if update {
		_, err := statefulSetSet.Get(statefulSetName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old stateful set")
			return err
		}

		_, err = statefulSetSet.Update(statefulSet)
		if err != nil {
			log.Println("Error when updating statefulSet")
			return err
		}
		log.Println("StatefulSet " + statefulSetName + " updated")

		return err
	} else {
		_, err := statefulSetSet.Create(statefulSet)
		if err != nil {
			log.Println("Error when creating statefulSet")
			return err
		}

		log.Println("StatefulSet " + statefulSetName + " created")
		return err
	}
}
