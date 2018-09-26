package openshift

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"github.com/hidevopsio/hiboot/pkg/log"
	authorization_v1 "github.com/openshift/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"github.com/openshift/client-go/authorization/clientset/versioned/fake"
)

func TestRoleBindingCrd(t *testing.T) {
	name := "admin"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	roleBinding, err := NewRoleBinding(clientSet, name, namespace)
	assert.Equal(t, nil, err)
	roleBinding1 := &authorization_v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: name,
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "User",
				Name:      "chen",
			},
			{
				Kind:      "User",
				Name:      "shi",
			},
		},
	}
	role, err := roleBinding.Create(roleBinding1)
	log.Debug(role)
	log.Debug(err)

	// Get
	binding, err := roleBinding.Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, name, binding.Name)

	err = roleBinding.Delete()
	assert.Equal(t, nil, err)
}


func TestCreateImagePullers(t *testing.T)  {
	name := "system:image-pullers"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	binding, err := NewRoleBinding(clientSet, name, namespace)
	assert.Equal(t, nil, err)
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: "system:image-puller",
			Kind: "ClusterRole",
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "Group",
				Name:      "system:serviceaccounts:" + namespace,
				Namespace: namespace,
			},
		},
	}
	role, err := binding.Create(roleBinding)
	log.Debug(role)
	log.Debug(err)

}


func TestCreateImageBuilders(t *testing.T)  {

	name := "system:image-builders"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	binding, err := NewRoleBinding(clientSet, name, namespace)
	assert.Equal(t, nil, err)
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: "system:image-builder",
			Kind: "ClusterRole",
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "ServiceAccount",
				Name:      "builder",
				Namespace: namespace,
			},
		},
	}
	role, err := binding.Create(roleBinding)
	log.Debug(role)
	log.Debug(err)


}



func TestCreateSystemDeployers(t *testing.T) {
	name := "system:deployers"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	binding, err := NewRoleBinding(clientSet, name, namespace)
	assert.Equal(t, nil, err)
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: "system:deployer",
			Kind: "ClusterRole",
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "ServiceAccount",
				Name:      "deployer",
				Namespace: namespace,
			},
		},
	}
	role, err := binding.Create(roleBinding)
	log.Debug(role)
	log.Debug(err)

}

func TestRoleBindingUpdate(t *testing.T) {
	name := "admin"
	namespace := "default"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	binding, err := NewRoleBinding(clientSet, name, namespace)
	assert.Equal(t, nil, err)
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: name,
		},
		Subjects: []corev1.ObjectReference{
			{
				Name: "chen",
				Kind: "User",
			}, {
				Kind: "User",
				Name: "shi",
			},
		},
	}
	role, err := binding.Update(roleBinding)
	log.Debug(role)
	log.Debug(err)
}