package apostle

import "os"

type config struct {
	DeliveryHost string
	DomainKey    string
	Version      string
}

func loadConfig() (c *config) {
	c = &config{
		Version: "1.0.0",
	}
	if deliveryHost := os.Getenv("APOSTLE_DELIVERY_HOST"); len(deliveryHost) != 0 {
		c.DeliveryHost = deliveryHost
	} else {
		c.DeliveryHost = "https://deliver.apostle.io"
	}

	if domainKey := os.Getenv("APOSTLE_DOMAIN_KEY"); len(domainKey) != 0 {
		c.DomainKey = domainKey
	}
	return
}

func SetDomainKey(dk string) {
	conf.DomainKey = dk
}
func SetDeliveryHost(dh string) {
	conf.DeliveryHost = dh
}
