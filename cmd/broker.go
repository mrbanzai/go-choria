package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/choria-io/go-choria/broker/adapter"
	"github.com/choria-io/go-choria/broker/federation"
	"github.com/choria-io/go-choria/broker/network"
	log "github.com/sirupsen/logrus"
)

type brokerCommand struct {
	command
}

type brokerRunCommand struct {
	command

	disableTLS       bool
	disableTLSVerify bool
	pidFile          string

	server     *network.Server
	federation *federation.FederationBroker
}

// broker
func (b *brokerCommand) Setup() (err error) {
	b.cmd = cli.app.Command("broker", "Choria Network Broker")

	return
}

func (b *brokerCommand) Run(wg *sync.WaitGroup) (err error) {
	return
}

// broker run
func (r *brokerRunCommand) Setup() (err error) {
	if broker, ok := cmdWithFullCommand("broker"); ok {
		r.cmd = broker.Cmd().Command("run", "Runs a Choria Network Broker instance").Default()
		r.cmd.Flag("disable-tls", "Disables TLS").Hidden().Default("false").BoolVar(&r.disableTLS)
		r.cmd.Flag("disable-ssl-verification", "Disables SSL Verification").Hidden().Default("false").BoolVar(&r.disableTLSVerify)
		r.cmd.Flag("pid", "Write running PID to a file").StringVar(&r.pidFile)
	}

	return
}

func (r *brokerRunCommand) Run(wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	net := config.Choria.BrokerNetwork
	discovery := config.Choria.BrokerDiscovery
	federation := config.Choria.BrokerFederation
	adapters := config.Choria.Adapters

	if !net && !discovery && !federation && len(adapters) == 0 {
		return fmt.Errorf("All broker features are disabled")
	}

	if r.pidFile != "" {
		err := ioutil.WriteFile(r.pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0644)
		if err != nil {
			return fmt.Errorf("Could not write PID: %s", err.Error())
		}
	}

	if r.disableTLS {
		c.Config.DisableTLS = true
		log.Warn("Running with TLS disabled, not compatible with production use Choria.")
	}

	if r.disableTLSVerify {
		c.Config.DisableTLSVerify = true
		log.Warn("Running with TLS Verification disabled, not compatible with production use Choria.")
	}

	if len(adapters) > 0 {
		log.Info("Starting Protocol Adapters")

		wg.Add(1)
		go r.runAdapters(ctx, wg)
	}

	if net {
		log.Info("Starting Network Broker")
		if err = r.runBroker(ctx, wg); err != nil {
			return fmt.Errorf("Starting the network broker failed: %s", err.Error())
		}
	}

	if federation {
		log.Infof("Starting Federation Broker on cluster %s", c.Config.Choria.FederationCluster)
		if err = r.runFederation(ctx, wg); err != nil {
			return fmt.Errorf("Starting the federation broker failed: %s", err.Error())
		}
	}

	if discovery {
		log.Warn("The Broker is configured to support Discovery but it's not been implemented yet.")
	}

	return
}

func (r *brokerRunCommand) runAdapters(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	err := adapter.RunAdapters(ctx, c, wg)
	if err != nil {
		log.Errorf("Failed to run Protocol Adapters: %s", err.Error())
		cancel()
	}
}

func (r *brokerRunCommand) runFederation(ctx context.Context, wg *sync.WaitGroup) (err error) {
	r.federation, err = federation.NewFederationBroker(c.Config.Choria.FederationCluster, c)
	if err != nil {
		return fmt.Errorf("Could not set up Choria Federation Broker: %s", err.Error())
	}

	wg.Add(1)
	r.federation.Start(ctx, wg)

	return
}

func (r *brokerRunCommand) runBroker(ctx context.Context, wg *sync.WaitGroup) (err error) {
	r.server, err = network.NewServer(c, debug)
	if err != nil {
		return fmt.Errorf("Could not set up Choria Network Broker: %s", err.Error())
	}

	wg.Add(1)
	go r.server.Start(ctx, wg)

	return
}

func init() {
	cli.commands = append(cli.commands, &brokerCommand{})
	cli.commands = append(cli.commands, &brokerRunCommand{})
}
