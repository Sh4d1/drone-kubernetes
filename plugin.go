package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
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

	// connect to Kubernetes
	clientset, err := p.createKubeClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// parse the template file and do substitutions
	//txt, err := openAndSub(p.Config.Template, p)
	//if err != nil {
	//	return err
	//}

	txt, err := ioutil.ReadFile(p.Config.Template)
	if err != nil {
		panic(err)
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode
	//obj, groupVersionKind, err := decode([]byte(txt), nil, nil)

	for _, s := range strings.Split(string(txt), "---") {
		obj, _, err := decode([]byte(s), nil, nil)
		if err != nil {
			log.Fatal(err.Error())
		}

		switch o := obj.(type) {
		case *v1.Pod:
			fmt.Printf("pod")
		case *appsv1.Deployment:
			result, err := clientset.AppsV1().Deployments("default").Create(o)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

		case *v1beta1.Ingress:
			fmt.Printf("ing")
		default:
			fmt.Printf("other")
		}
	}

	return err
}

// open up the template and then sub variables in. Handlebar stuff.
func openAndSub(templateFile string, p Plugin) (string, error) {
	u, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return "", err
	}
	//potty humor!  Render trim toilet paper!  Ha ha, so funny.
	return RenderTrim(string(u), p)
}

// create the connection to kubernetes based on parameters passed in.
// the kubernetes/client-go project is really hard to understand.
func (p Plugin) createKubeClient() (*kubernetes.Clientset, error) {

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
