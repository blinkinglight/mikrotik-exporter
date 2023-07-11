package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/routeros.v2/proto"
)

type psuCollector struct {
	props        []string
	descriptions map[string]*prometheus.Desc
}

func newpsuCollector() routerOSCollector {
	c := &psuCollector{}
	c.init()
	return c
}

func (c *psuCollector) init() {
	c.props = []string{"psu1-state", "psu2-state"}

	labelNames := []string{"name", "address"}
	helpText := []string{
		"State of PSU1",
		"State of PSU2",
	}
	c.descriptions = make(map[string]*prometheus.Desc)
	for i, p := range c.props {
		c.descriptions[p] = descriptionForPropertyNameHelpText("psu", p, labelNames, helpText[i])
	}
}

func (c *psuCollector) describe(ch chan<- *prometheus.Desc) {
	for _, d := range c.descriptions {
		ch <- d
	}
}

func (c *psuCollector) collect(ctx *collectorContext) error {
	stats, err := c.fetch(ctx)
	if err != nil {
		return err
	}

	for _, re := range stats {
		if metric, ok := re.Map["name"]; ok {
			c.collectMetricForProperty(metric, re, ctx)
		} else {
			c.collectForStat(re, ctx)
		}
	}

	return nil
}

func (c *psuCollector) fetch(ctx *collectorContext) ([]*proto.Sentence, error) {
	reply, err := ctx.client.Run("/system/health/print")
	if err != nil {
		log.WithFields(log.Fields{
			"device": ctx.device.Name,
			"error":  err,
		}).Error("error fetching system health metrics")
		return nil, err
	}

	return reply.Re, nil
}

func (c *psuCollector) collectForStat(re *proto.Sentence, ctx *collectorContext) {
	for _, p := range c.props[:2] {
		c.collectMetricForProperty(p, re, ctx)
	}
}

func (c *psuCollector) collectMetricForProperty(property string, re *proto.Sentence, ctx *collectorContext) {
	name := property
	value, ok1 := re.Map["value"]
	if value == "" || !ok1 {
		if value, ok1 = re.Map["value"]; !ok1 {
			return
		}
	}

	var v float64
	if value == "ok" {
		v = 1
	} else if value == "fail" {
		v = 0
	} else {
		return
	}
	// log.Printf("%v", re.Map)
	desc := c.descriptions[name]
	// log.Printf("%v", desc)
	ctx.ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, v, ctx.device.Name, ctx.device.Address)
}
