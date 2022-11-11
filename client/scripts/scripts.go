// TODO move to v1 to make it consistent
package scripts

import (
	"context"
	"github.com/kubeshop/testkube-operator/pkg/utils"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	scriptv1 "github.com/kubeshop/testkube-operator/apis/script/v1"
)

func NewClient(client client.Client, namespace string) *ScriptsClient {
	return &ScriptsClient{
		Client:    client,
		Namespace: namespace,
	}
}

type ScriptsClient struct {
	Client    client.Client
	Namespace string
}

func (s ScriptsClient) List(tags []string) (*scriptv1.ScriptList, error) {
	list := &scriptv1.ScriptList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: s.Namespace})
	if len(tags) == 0 {
		return list, err
	}

	toReturn := &scriptv1.ScriptList{}
	for _, script := range list.Items {
		hasTags := false
		for _, tag := range tags {
			if utils.ContainsTag(script.Spec.Tags, tag) {
				hasTags = true
			} else {
				hasTags = false
			}
		}
		if hasTags {
			toReturn.Items = append(toReturn.Items, script)

		}
	}

	return toReturn, nil
}

func (s ScriptsClient) ListTags() ([]string, error) {
	tags := []string{}
	list := &scriptv1.ScriptList{}
	err := s.Client.List(context.Background(), list, &client.ListOptions{Namespace: s.Namespace})
	if err != nil {
		return tags, err
	}

	for _, test := range list.Items {
		tags = append(tags, test.Spec.Tags...)
	}

	tags = utils.RemoveDuplicates(tags)

	return tags, nil
}

func (s ScriptsClient) Get(name string) (*scriptv1.Script, error) {
	script := &scriptv1.Script{}
	err := s.Client.Get(context.Background(), client.ObjectKey{Namespace: s.Namespace, Name: name}, script)
	return script, err
}

func (s ScriptsClient) Create(script *scriptv1.Script) (*scriptv1.Script, error) {
	err := s.Client.Create(context.Background(), script)
	return script, err
}

func (s ScriptsClient) Update(script *scriptv1.Script) (*scriptv1.Script, error) {
	err := s.Client.Update(context.Background(), script)
	return script, err
}

func (s ScriptsClient) Delete(name string) error {
	script, err := s.Get(name)
	if err != nil {
		return err
	}

	err = s.Client.Delete(context.Background(), script)
	return err
}

func (s ScriptsClient) DeleteAll() error {

	u := &unstructured.Unstructured{}
	u.SetKind("script")
	u.SetAPIVersion("tests.testkube.io/v1")
	err := s.Client.DeleteAllOf(context.Background(), u, client.InNamespace(s.Namespace))
	return err
}
