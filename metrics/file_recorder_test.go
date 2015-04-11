package metrics_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	. "github.com/cghsystems/gosum/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Metrics File Recorder", func() {
	Context("metrics are recorded", func() {
		const filePath = "/tmp/FileRecorderTest"
		BeforeEach(func() {
			InitFileRecorder(filePath)
			Ω(filePath).Should(FileExists())
		})

		AfterEach(func() {
			os.Remove(filePath)
		})

		It("records the time against the api call", func() {
			now := time.Now()
			Record("/test", now)

			expectedContents := fmt.Sprintf("/test,%v", now)
			actualContents := fileContents(filePath)

			Ω(expectedContents).Should(Equal(actualContents))
		})
	})
})

func FileExists() types.GomegaMatcher {
	return &FileExistsMatcher{}
}

type FileExistsMatcher struct{}

func (matcher *FileExistsMatcher) Match(actual interface{}) (success bool, err error) {
	filePath := actual.(string)
	_, err = os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return os.IsNotExist(err), nil
}

func (matcher *FileExistsMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to exist")
}

func (matcher *FileExistsMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to exist")
}

func fileContents(filePath string) string {
	bytes, err := ioutil.ReadFile(filePath)
	Ω(err).ShouldNot(HaveOccurred())
	return string(bytes)
}
