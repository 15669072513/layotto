package apollo

import (
	"mosn.io/layotto/components/configstores"
	"sync"
)

// Holding subscribers' chan and ctx.
type subscriberHolder struct {
	sync.RWMutex
	chanMap map[subscriberKey][]*subscriber
}

func (h *subscriberHolder) findByTopic(namespace string, keyWithLabel string) []*subscriber {
	h.RLock()
	defer h.RUnlock()
	load, ok := h.chanMap[subscriberKey{
		group:        namespace,
		keyWithLabel: keyWithLabel,
	}]
	if !ok {
		return nil
	}
	return load
}

func (h *subscriberHolder) addByTopic(namespace string, keyWithLabel string, respChan chan *configstores.SubscribeResp) error {
	if respChan == nil {
		return errParamsMissingField("respChan")
	}
	key := subscriberKey{
		group:        namespace,
		keyWithLabel: keyWithLabel,
	}
	h.Lock()
	defer h.Unlock()
	_, ok := h.chanMap[key]
	if !ok {
		h.chanMap[key] = make([]*subscriber, 0, 1)
	}
	s := &subscriber{
		respChan:      respChan,
		group:         namespace,
		subscriberKey: &key,
	}
	h.chanMap[key] = append(h.chanMap[key], s)
	return nil
}

func (h *subscriberHolder) remove(s *subscriber) {
	// check
	if s == nil || s.subscriberKey == nil {
		return
	}
	key := *s.subscriberKey
	h.Lock()
	defer h.Unlock()
	// find related slice
	slice, ok := h.chanMap[key]
	if !ok {
		return
	}
	// find and remove the subscriber
	n := len(slice)
	for i := 0; i < n; i++ {
		if slice[i] == s {
			//	remove
			h.chanMap[key] = append(h.chanMap[key][:i], h.chanMap[key][i+1:]...)
			return
		}
	}
}

func (h *subscriberHolder) reset() {
	h.Lock()
	defer h.Unlock()
	h.chanMap = make(map[subscriberKey][]*subscriber)
}

type subscriberKey struct {
	//appId        string
	group        string
	keyWithLabel string
}

type subscriber struct {
	respChan      chan *configstores.SubscribeResp
	group         string
	subscriberKey *subscriberKey
	//	TODO add "context" field for canceling
}

func newSubscriberHolder() *subscriberHolder {
	return &subscriberHolder{
		chanMap: make(map[subscriberKey][]*subscriber),
	}
}