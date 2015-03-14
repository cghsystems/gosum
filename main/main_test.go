package main_test

import (
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("main", func() {
	It("starts the http server", func() {
		binPath, err := gexec.Build("github.com/cghsystems/gosum/main/")
		Ω(err).ShouldNot(HaveOccurred())
		cmd := exec.Command(binPath)

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		defer func() {
			session.Terminate()
			session.Wait()
			Eventually(session).Should(gexec.Exit())
		}()
		Ω(err).ShouldNot(HaveOccurred())

		execute := func() error {
			if _, err := http.Get("http://localhost:8080/api"); err != nil {
				return nil
			}
			return err
		}
		Eventually(execute).ShouldNot(HaveOccurred())
	})
})
