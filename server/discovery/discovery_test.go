package discovery

import (
	"os"
	"testing"

	"github.com/choria-io/go-protocol/protocol"

	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-config"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	os.Setenv("MCOLLECTIVE_CERTNAME", "rip.mcollective")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server/Discovery")
}

var _ = Describe("Server/Discovery", func() {
	var fw *choria.Framework
	var log *logrus.Entry
	var err error
	var mgr *Manager
	var req protocol.Request
	var filter *protocol.Filter
	var agents []string

	BeforeSuite(func() {
		log = logrus.WithFields(logrus.Fields{"test": true})
		cfg := config.NewConfigForTests()
		cfg.DisableTLS = true

		fw, err = choria.NewWithConfig(cfg)
		Expect(err).ToNot(HaveOccurred())

		fw.Config.ClassesFile = "testdata/classes.txt"
		fw.Config.FactSourceFile = "testdata/facts.yaml"
		fw.Config.Identity = "test.example.net"
	})

	BeforeEach(func() {
		mgr = New(fw, log)
		rid, err := fw.NewRequestID()
		Expect(err).ToNot(HaveOccurred())

		req, err = fw.NewRequest(protocol.RequestV1, "test", "testid", "callerid", 60, rid, "mcollective")
		Expect(err).ToNot(HaveOccurred())

		filter = req.NewFilter()
		req.SetFilter(filter)

		agents = []string{"apache", "rpcutil"}
	})

	It("Should match on empty filters", func() {
		Expect(mgr.ShouldProcess(req, []string{})).To(BeTrue())
	})

	It("Should match if all filters matched", func() {
		filter.AddAgentFilter("apache")
		filter.AddClassFilter("role::testing")
		filter.AddClassFilter("/test/")
		filter.AddFactFilter("nested.string", "=~", "/hello/")
		filter.AddIdentityFilter("/test/")

		Expect(mgr.ShouldProcess(req, agents)).To(BeTrue())
	})

	It("Should fail if some filters matched", func() {
		filter.AddAgentFilter("apache")
		filter.AddClassFilter("role::test")
		filter.AddFactFilter("nested.string", "=~", "/meh/")

		Expect(mgr.ShouldProcess(req, agents)).To(BeFalse())
	})
})
