package cln

import "time"

type queueNode struct {
  t    time.Time
  next *queueNode
}

// A TimeQueue is a queue that stores elements of type time.Time
type TimeQueue struct {
  head *queueNode
  tail *queueNode

  len int
}

func newQueueNode(t time.Time) *queueNode {
  qn := new(queueNode)
  qn.t = t
  qn.next = nil
  return qn
}

// Len Returns the length of the tq
func (tq *TimeQueue) Len() int {
  return tq.len
}

// Empty returns true if  tq has a length of 0
func (tq *TimeQueue) Empty() bool {
  return tq.len == 0
}

// Head returns the head (front) of tq, or the zero value of time.Time if Empty
func (tq *TimeQueue) Head() time.Time {
  var t time.Time
  if !tq.Empty() {
    t = tq.head.t
  }
  return t
}

// Tail returns the tail (back) of tq, or the zero value of time.Time if Empty
func (tq *TimeQueue) Tail() time.Time {
  var t time.Time
  if !tq.Empty() {
    t = tq.tail.t
  }
  return t
}

// Add adds the time.Time t to the back of tq
func (tq *TimeQueue) Add(t time.Time) {
  if tq.Empty() {
    tq.head = newQueueNode(t)
    tq.tail = tq.head
  } else {
    tq.tail.next = newQueueNode(t)
    tq.tail = tq.tail.next
  }
  tq.len++
}

// Poll removes and returns the head of tq (first time.Time)
func (tq *TimeQueue) Poll() time.Time {
  var t time.Time
  if !tq.Empty() {
    t = tq.Head()
    tq.head = tq.head.next
    tq.len--
  }
  return t
}

// RemoveTimesBefore removes all time.Time elements that are before t, assuming that the time.Time elements in
// tq are in sorted order (earliest to latest)
func (tq *TimeQueue) RemoveTimesBefore(t time.Time) {
  for !tq.Empty() && tq.Head().Before(t) {
    tq.Poll()
  }
}
