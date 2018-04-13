package main

import (
	"log"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

func applyDeploymentAppsV1beta1(deployment *appsv1beta1.Deployment, deploymentSet v1beta1.DeploymentInterface) error {
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

func applyStatefulSetAppsV1beta1(statefulSet *appsv1beta1.StatefulSet, statefulSetSet v1beta1.StatefulSetInterface) error {
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
