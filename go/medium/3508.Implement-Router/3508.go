// https://leetcode.com/problems/implement-router/
// Level: Medium

package leetcode

import "sort"

type Router struct {
	capacity         int
	packetQueue      [][3]int
	destToTimestamps map[int][]int
	seenPackets      map[[3]int]bool
}

func Constructor(memoryLimit int) Router {
	return Router{
		capacity:         memoryLimit,
		packetQueue:      make([][3]int, 0),
		destToTimestamps: make(map[int][]int),
		seenPackets:      make(map[[3]int]bool),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if this.seenPackets[key] {
		return false
	}

	this.packetQueue = append(this.packetQueue, key)
	this.destToTimestamps[destination] = append(this.destToTimestamps[destination], timestamp)
	sort.Ints(this.destToTimestamps[destination])
	this.seenPackets[key] = true

	if len(this.packetQueue) > this.capacity {
		evicted := this.packetQueue[0]
		this.packetQueue = this.packetQueue[1:]
		evictedDest := evicted[1]

		if len(this.destToTimestamps[evictedDest]) > 0 {
			this.destToTimestamps[evictedDest] = this.destToTimestamps[evictedDest][1:]
		}

		delete(this.seenPackets, evicted)
	}

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.packetQueue) == 0 {
		return []int{}
	}

	packet := this.packetQueue[0]
	this.packetQueue = this.packetQueue[1:]

	dest := packet[1]
	if len(this.destToTimestamps[dest]) > 0 {
		this.destToTimestamps[dest] = this.destToTimestamps[dest][1:]
	}

	delete(this.seenPackets, packet)

	return []int{packet[0], packet[1], packet[2]}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	timestamps := this.destToTimestamps[destination]

	left := sort.Search(len(timestamps), func(i int) bool {
		return timestamps[i] >= startTime
	})
	right := sort.Search(len(timestamps), func(i int) bool {
		return timestamps[i] > endTime
	})

	return right - left
}
