// +build collision

package collision

import (
	"github.com/autovelop/playthos"
)

func init() {
	engine.NewSystem(&Collision{})
}

type Collision struct {
	engine.System
	colliders []*Collider
}

func (c *Collision) InitSystem() {}

func (c *Collision) NewIntegrant(integrant engine.IntegrantRoutine) {}

func (c *Collision) NewComponent(component engine.ComponentRoutine) {
	switch component := component.(type) {
	case *Collider:
		c.colliders = append(c.colliders, component)
		break
	}
}

func (c *Collision) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Collider{}}
}

func (c *Collision) Update() {
	var prev_collider *Collider
	for _, collider := range c.colliders {
		if prev_collider == nil {
			prev_collider = collider
		} else {
			c1 := collider.Entity()
			c2 := prev_collider.Entity()
			if c1 != nil && c2 != nil {
				// p1 := c1.GetComponent(&render.Transform{}).(*render.Transform)
				// p2 := c2.GetComponent(&render.Transform{}).(*render.Transform)

				if CheckCollisionAABB(collider, prev_collider) {
					// collider.Hit(c2)
					prev_collider.Hit(c1)
				}
				// if Distance3(p1.GetPosition(), p2.GetPosition()) < 80 {
				// 	collider.Hit()
				// 	prev_collider.Hit()
				// }
			}
		}
	}
}
