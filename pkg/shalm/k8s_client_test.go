package shalm

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("k8s client", func() {

	It("can read kubernetes service", func() {
		client, err := newK8sClient("")
		if err != nil {
			Skip("no connection to k8s")
		}
		namespace := "default"
		obj, err := client.Get().
			Namespace(&namespace).
			Resource("services").
			Name("kubernetes").
			Do().
			Get()

		Expect(err).NotTo(HaveOccurred())
		Expect(obj.Kind).To(Equal("Service"))
	})

})
