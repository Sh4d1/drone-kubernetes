package main

import (
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func applyConfigMap(configMap *corev1.ConfigMap, configMapSet v1.ConfigMapInterface) error {
	configMapName := configMap.GetObjectMeta().GetName()
	configMaps, err := configMapSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing configMaps")
		return err
	}

	update := false
	for _, cm := range configMaps.Items {
		if cm.GetObjectMeta().GetName() == configMapName {
			update = true
		}
	}

	if update {
		_, err := configMapSet.Get(configMapName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old configMap")
			return err
		}

		_, err = configMapSet.Update(configMap)
		if err != nil {
			log.Println("Error when updating configMap")
			return err
		}
		log.Println("ConfigMap " + configMapName + " updated")

		return err
	} else {
		_, err := configMapSet.Create(configMap)
		if err != nil {
			log.Println("Error when creating configMap")
			return err
		}

		log.Println("ConfigMap " + configMapName + " created")
		return err
	}
}

func applyPersistentVolume(persistentVolume *corev1.PersistentVolume, persistentVolumeSet v1.PersistentVolumeInterface) error {
	persistentVolumeName := persistentVolume.GetObjectMeta().GetName()
	persistentVolumes, err := persistentVolumeSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing persistentVolumes")
		return err
	}

	update := false
	for _, pv := range persistentVolumes.Items {
		if pv.GetObjectMeta().GetName() == persistentVolumeName {
			update = true
		}
	}

	if update {
		_, err := persistentVolumeSet.Get(persistentVolumeName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old persistentVolume")
			return err
		}

		_, err = persistentVolumeSet.Update(persistentVolume)
		if err != nil {
			log.Println("Error when updating persistentVolume")
			return err
		}
		log.Println("PersistentVolume " + persistentVolumeName + " updated")

		return err
	} else {
		_, err := persistentVolumeSet.Create(persistentVolume)
		if err != nil {
			log.Println("Error when creating persistentVolume")
			return err
		}

		log.Println("PersistentVolume " + persistentVolumeName + " created")
		return err
	}
}

func applyPersistentVolumeClaim(persistentVolumeClaim *corev1.PersistentVolumeClaim, persistentVolumeClaimSet v1.PersistentVolumeClaimInterface) error {
	persistentVolumeClaimName := persistentVolumeClaim.GetObjectMeta().GetName()
	persistentVolumeClaims, err := persistentVolumeClaimSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing persistentVolumeClaims")
		return err
	}

	update := false
	for _, pvc := range persistentVolumeClaims.Items {
		if pvc.GetObjectMeta().GetName() == persistentVolumeClaimName {
			update = true
		}
	}

	if update {
		_, err := persistentVolumeClaimSet.Get(persistentVolumeClaimName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old persistentVolumeClaim")
			return err
		}

		_, err = persistentVolumeClaimSet.Update(persistentVolumeClaim)
		if err != nil {
			log.Println("Error when updating persistentVolumeClaim")
			return err
		}
		log.Println("PersistentVolumeClaim " + persistentVolumeClaimName + " updated")

		return err
	} else {
		_, err := persistentVolumeClaimSet.Create(persistentVolumeClaim)
		if err != nil {
			log.Println("Error when creating persistentVolumeClaim")
			return err
		}

		log.Println("PersistentVolumeClaim " + persistentVolumeClaimName + " created")
		return err
	}
}

func applyPod(pod *corev1.Pod, podSet v1.PodInterface) error {
	podName := pod.GetObjectMeta().GetName()
	pods, err := podSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing pods")
		return err
	}

	update := false
	for _, p := range pods.Items {
		if p.GetObjectMeta().GetName() == podName {
			update = true
		}
	}

	if update {
		_, err := podSet.Get(podName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old pod")
			return err
		}

		_, err = podSet.Update(pod)
		if err != nil {
			log.Println("Error when updating pod")
			return err
		}
		log.Println("Pod " + podName + " updated")

		return err
	} else {
		_, err := podSet.Create(pod)
		if err != nil {
			log.Println("Error when creating pod")
			return err
		}

		log.Println("Pod " + podName + " created")
		return err
	}
}

func applyReplicationController(replicationController *corev1.ReplicationController, replicationControllerSet v1.ReplicationControllerInterface) error {
	replicationControllerName := replicationController.GetObjectMeta().GetName()
	replicationControllers, err := replicationControllerSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing replicationControllers")
		return err
	}

	update := false
	for _, rc := range replicationControllers.Items {
		if rc.GetObjectMeta().GetName() == replicationControllerName {
			update = true
		}
	}

	if update {
		_, err := replicationControllerSet.Get(replicationControllerName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old replicationController")
			return err
		}

		_, err = replicationControllerSet.Update(replicationController)
		if err != nil {
			log.Println("Error when updating replicationController")
			return err
		}
		log.Println("ReplicationController " + replicationControllerName + " updated")

		return err
	} else {
		_, err := replicationControllerSet.Create(replicationController)
		if err != nil {
			log.Println("Error when creating replicationController")
			return err
		}

		log.Println("ReplicationController " + replicationControllerName + " created")
		return err
	}
}

func applyService(service *corev1.Service, serviceSet v1.ServiceInterface) error {
	serviceName := service.GetObjectMeta().GetName()
	services, err := serviceSet.List(metav1.ListOptions{})
	if err != nil {
		log.Println("Error when listing services")
		return err
	}

	update := false
	for _, svc := range services.Items {
		if svc.GetObjectMeta().GetName() == serviceName {
			update = true
		}
	}

	if update {
		_, err := serviceSet.Get(serviceName, metav1.GetOptions{})
		if err != nil {
			log.Println("Error when getting old service")
			return err
		}

		_, err = serviceSet.Update(service)
		if err != nil {
			log.Println("Error when updating service")
			return err
		}
		log.Println("Service " + serviceName + " updated")

		return err
	} else {
		_, err := serviceSet.Create(service)
		if err != nil {
			log.Println("Error when creating service")
			return err
		}

		log.Println("Service " + serviceName + " created")
		return err
	}
}
