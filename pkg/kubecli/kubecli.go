/*
 * This file is part of the kubevirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2017 Red Hat, Inc.
 *
 */

package kubecli

import (
	"flag"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"kubevirt.io/kubevirt/pkg/api/v1"
)

var (
	kubeconfig string
	master     string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.StringVar(&master, "master", "", "master url")
}

func GetFromFlags(master string, kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}

	config.GroupVersion = &v1.GroupVersion
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON

	return kubernetes.NewForConfig(config)
}

func Get() (*kubernetes.Clientset, error) {
	return GetFromFlags(master, kubeconfig)
}

func GetRESTClient() (*rest.RESTClient, error) {
	return GetRESTClientFromFlags(master, kubeconfig)
}

func GetKubevirtClientFromFlags(master string, kubeconfig string) (KubevirtClient, error) {
	restClient, err := GetRESTClientFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}
	return &kubevirt{restClient}, nil
}

func GetKubevirtClient() (KubevirtClient, error) {
	return GetKubevirtClientFromFlags(master, kubeconfig)
}

func GetRESTClientFromFlags(master string, kubeconfig string) (*rest.RESTClient, error) {

	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}

	config.GroupVersion = &v1.GroupVersion
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON

	return rest.RESTClientFor(config)
}
