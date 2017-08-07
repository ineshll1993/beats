ackage htop

import (
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
)

var schema = s.Schema{
	"cpu":          c.Int("cpu"),
	"ram": c.Int("ram")
}
