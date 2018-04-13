package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  string
		Status  string
		Link    string
		Started int64
		Created int64
	}

	Job struct {
		Started int64
	}

	Config struct {
		Cert      string
		Server    string
		Token     string
		Namespace string
		Template  string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {

	if p.Config.Server == "" {
		log.Fatal("KUBERNETES_SERVER is not defined")
	}
	if p.Config.Token == "" {
		log.Fatal("KUBERNETES_TOKEN is not defined")
	}
	if p.Config.Cert == "" {
		log.Fatal("KUBERNETES_CERT is not defined")
	}
	if p.Config.Namespace == "" {
		p.Config.Namespace = "default"
	}
	if p.Config.Template == "" {
		log.Fatal("KUBERNETES_TEMPLATE is not defined")
	}

	clientset, err := p.getClient()
	if err != nil {
		return err
	}

	template, err := p.getTemplate()
	if err != nil {
		return err
	}
	decode := scheme.Codecs.UniversalDeserializer().Decode

	// iterate if several yalm files separated by ---
	for _, s := range strings.Split(template, "---") {
		obj, _, err := decode([]byte(s), nil, nil)
		if err != nil {
			log.Println("Error when decoding template YAML")
			return err
		}

		switch o := obj.(type) {
		// appsv1
		case *appsv1.DaemonSet:
			daemonSetSet := clientset.AppsV1().DaemonSets(p.Config.Namespace)
			err := applyDaemonSetAppsV1(o, daemonSetSet)
			if err != nil {
				return err
			}

		case *appsv1.Deployment:
			deploymentSet := clientset.AppsV1().Deployments(p.Config.Namespace)
			err := applyDeploymentAppsV1(o, deploymentSet)
			if err != nil {
				return err
			}

		case *appsv1.ReplicaSet:
			replicatSetSet := clientset.AppsV1().ReplicaSets(p.Config.Namespace)
			err := applyReplicaSetAppsV1(o, replicatSetSet)
			if err != nil {
				return err
			}

		case *appsv1.StatefulSet:
			statefulSetSet := clientset.AppsV1().StatefulSets(p.Config.Namespace)
			err := applyStatefulSetAppsV1(o, statefulSetSet)
			if err != nil {
				return err
			}

		// appsv1beta1
		case *appsv1beta1.Deployment:
			deploymentSet := clientset.AppsV1beta1().Deployments(p.Config.Namespace)
			err := applyDeploymentAppsV1beta1(o, deploymentSet)
			if err != nil {
				return err
			}

		case *appsv1beta1.StatefulSet:
			statefulSetSet := clientset.AppsV1beta1().StatefulSets(p.Config.Namespace)
			err := applyStatefulSetAppsV1beta1(o, statefulSetSet)
			if err != nil {
				return err
			}

		// appsv1beta2
		case *appsv1beta2.DaemonSet:
			daemonSetSet := clientset.AppsV1beta2().DaemonSets(p.Config.Namespace)
			err := applyDaemonSetAppsV1beta2(o, daemonSetSet)
			if err != nil {
				return err
			}

		case *appsv1beta2.Deployment:
			deploymentSet := clientset.AppsV1beta2().Deployments(p.Config.Namespace)
			err := applyDeploymentAppsV1beta2(o, deploymentSet)
			if err != nil {
				return err
			}

		case *appsv1beta2.ReplicaSet:
			replicatSetSet := clientset.AppsV1beta2().ReplicaSets(p.Config.Namespace)
			err := applyReplicaSetAppsV1beta2(o, replicatSetSet)
			if err != nil {
				return err
			}

		case *appsv1beta2.StatefulSet:
			statefulSetSet := clientset.AppsV1beta2().StatefulSets(p.Config.Namespace)
			err := applyStatefulSetAppsV1beta2(o, statefulSetSet)
			if err != nil {
				return err
			}

		// corev1
		case *corev1.ConfigMap:
			configMapSet := clientset.CoreV1().ConfigMaps(p.Config.Namespace)
			err := applyConfigMap(o, configMapSet)

			if err != nil {
				return err
			}

		case *corev1.PersistentVolume:
			persistentVolumeSet := clientset.CoreV1().PersistentVolumes()
			err := applyPersistentVolume(o, persistentVolumeSet)

			if err != nil {
				return err
			}

		case *corev1.PersistentVolumeClaim:
			persistentVolumeClaimSet := clientset.CoreV1().PersistentVolumeClaims(p.Config.Namespace)
			err := applyPersistentVolumeClaim(o, persistentVolumeClaimSet)

			if err != nil {
				return err
			}

		case *corev1.Pod:
			podSet := clientset.CoreV1().Pods(p.Config.Namespace)
			err := applyPod(o, podSet)

			if err != nil {
				return err
			}

		case *corev1.ReplicationController:
			replicationControllerSet := clientset.CoreV1().ReplicationControllers(p.Config.Namespace)
			err := applyReplicationController(o, replicationControllerSet)

			if err != nil {
				return err
			}

		case *corev1.Service:
			serviceSet := clientset.CoreV1().Services(p.Config.Namespace)
			err := applyService(o, serviceSet)

			if err != nil {
				return err
			}

		// extensionsv1beta1
		case *extensionsv1beta1.DaemonSet:
			daemonSetSet := clientset.ExtensionsV1beta1().DaemonSets(p.Config.Namespace)
			err := applyDaemonSetExtensionsV1beta1(o, daemonSetSet)
			if err != nil {
				return err
			}

		case *extensionsv1beta1.Deployment:
			deploymentSet := clientset.ExtensionsV1beta1().Deployments(p.Config.Namespace)
			err := applyDeploymentExtensionsV1beta1(o, deploymentSet)
			if err != nil {
				return err
			}

		case *extensionsv1beta1.Ingress:
			ingressSet := clientset.ExtensionsV1beta1().Ingresses(p.Config.Namespace)
			err := applyIngressExtensionsV1beta1(o, ingressSet)

			if err != nil {
				return err
			}

		case *extensionsv1beta1.ReplicaSet:
			replicatSetSet := clientset.ExtensionsV1beta1().ReplicaSets(p.Config.Namespace)
			err := applyReplicaSetExtensionsV1beta1(o, replicatSetSet)
			if err != nil {
				return err
			}

		default:
			fmt.Printf("other")
		}
	}

	return nil
}

func (p Plugin) getClient() (*kubernetes.Clientset, error) {

	cert, err := base64.StdEncoding.DecodeString(p.Config.Cert)
	config := clientcmdapi.NewConfig()
	config.Clusters["drone"] = &clientcmdapi.Cluster{
		Server: p.Config.Server,
		CertificateAuthorityData: cert,
	}
	config.AuthInfos["drone"] = &clientcmdapi.AuthInfo{
		Token: p.Config.Token,
	}

	config.Contexts["drone"] = &clientcmdapi.Context{
		Cluster:  "drone",
		AuthInfo: "drone",
	}

	config.CurrentContext = "drone"

	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, "drone", &clientcmd.ConfigOverrides{}, nil)
	actualCfg, err := clientBuilder.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	return kubernetes.NewForConfig(actualCfg)
}

func (p Plugin) getTemplate() (string, error) {

	var template string
	u, err := url.ParseRequestURI(p.Config.Template)
	if err == nil {
		switch u.Scheme {
		case "http", "https":
			defaultTransport := http.DefaultTransport.(*http.Transport)
			cli := &http.Transport{
				Proxy:                 defaultTransport.Proxy,
				DialContext:           defaultTransport.DialContext,
				MaxIdleConns:          defaultTransport.MaxIdleConns,
				IdleConnTimeout:       defaultTransport.IdleConnTimeout,
				ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
				TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			}

			client := &http.Client{Transport: cli}
			res, err := client.Get(p.Config.Template)
			if err != nil {
				log.Println("Error when getting template URL")
				return template, err
			}
			defer res.Body.Close()
			out, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println("Error when reading template URL")
				return template, err
			}
			template = string(out)
		case "file":
			fmt.Println(u.Path)
			out, err := ioutil.ReadFile(u.Path)
			if err != nil {
				log.Println("Error when reading template file")
				return template, err
			}
			template = string(out)
		}
	} else {
		fmt.Println("file")
		file, err := filepath.Abs(p.Config.Template)
		if err != nil {
			log.Println("Error when getting template path")
			return template, err
		}
		out, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("Error when reading template file")
			return template, err
		}
		template = string(out)
	}

	return RenderTrim(template, p)
}
