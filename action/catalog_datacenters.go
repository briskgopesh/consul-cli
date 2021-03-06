package action

import (
	"flag"
)

type catalogDatacenters struct {
	*config
}

func CatalogDatacentersAction() Action {
	return &catalogDatacenters{
		config: &gConfig,
	}
}

func (c *catalogDatacenters) CommandFlags() *flag.FlagSet {
	return c.newFlagSet(FLAG_OUTPUT)
}

func (c *catalogDatacenters) Run(args []string) error {
	client, err := c.newCatalog()
	if err != nil {
		return err
	}

	config, err := client.Datacenters()
	if err != nil {
		return err
	}

	return c.Output(config)
}
