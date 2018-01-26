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
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
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
		case *appsv1.Deployment:
			deploymentSet := clientset.AppsV1().Deployments(p.Config.Namespace)
			err := applyDeployment(o, deploymentSet)
			if err != nil {
				return err
			}
		case *appsv1.DaemonSet:
			daemonSet := clientset.AppsV1().DaemonSets(p.Config.Namespace)
			err := applyDaemonSet(o, daemonSet)
			if err != nil {
				return err
			}
			//	case *apiv1.Service:
			//		serviceSet := clientset.CoreV1().Services(p.Config.Namespace)
			//		_, err := applyService(o, serviceSet)
			//		if err != nil {
			//			return err
			//		}
		case *v1beta1.Ingress:
			fmt.Printf("ing")
		default:
			fmt.Printf("other")
		}
	}

	return nil
}

func applyDeployment(deployment *appsv1.Deployment, deploymentSet typedappsv1.DeploymentInterface) error {
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

func applyDaemonSet(daemonSet *appsv1.DaemonSet, daemonSetSet typedappsv1.DaemonSetInterface) error {
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

	return template, nil
}
