package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ListTestSuite struct {
	suite.Suite
	list List
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

func (s *ListTestSuite) SetupTest() {
	s.list = NewList()
}

func (s *ListTestSuite) TestCreateInstance() {
	s.IsType(&list{}, s.list)
}

func (s *ListTestSuite) TestPushFrontAndLen() {
	s.Equal(0, s.list.Len())
	s.Nil(s.list.Front())
	s.Nil(s.list.Back())

	s.list.PushFront(1)
	s.Equal(1, s.list.Len())

	s.list.PushFront(2)
	s.list.PushFront(40)
	s.Equal(3, s.list.Len())
}

func (s *ListTestSuite) TestPushBackAndLen() {
	s.Equal(0, s.list.Len())
	s.Nil(s.list.Front())
	s.Nil(s.list.Back())

	s.list.PushBack(1)
	s.Equal(1, s.list.Len())

	s.list.PushBack(2)
	s.list.PushBack(3)
	s.list.PushBack(4)
	s.Equal(4, s.list.Len())
}

func (s *ListTestSuite) TestGetFront() {
	s.list.PushFront(10)
	s.Equal(10, s.list.Front().Value)

	s.list.PushFront("test")
	s.Equal("test", s.list.Front().Value)
	s.Equal(10, s.list.Front().Prev.Value)
}

func (s *ListTestSuite) TestGetBack() {
	s.list.PushBack(10)
	s.Equal(10, s.list.Back().Value)

	s.list.PushBack(20)
	s.Equal(20, s.list.Back().Value)
	s.Equal(10, s.list.Back().Next.Value)
}

func (s *ListTestSuite) TestPushFrontAndGetBack() {

	s.list.PushFront(3)
	s.list.PushFront(2)
	s.list.PushFront(1)

	s.Equal(3, s.list.Len())

	s.Equal(1, s.list.Front().Value)
	s.Equal(2, s.list.Front().Prev.Value)
	s.Equal(3, s.list.Front().Prev.Prev.Value)
	s.Equal(3, s.list.Back().Value)
	s.Equal(2, s.list.Back().Next.Value)
	s.Equal(1, s.list.Back().Next.Next.Value)

	s.Nil(s.list.Front().Prev.Prev.Prev)
	s.Nil(s.list.Back().Next.Next.Next)
}

func (s *ListTestSuite) TestPushBackAndGetFront() {
	s.list.PushBack(1)
	s.list.PushBack(2)
	s.list.PushBack(3)

	s.Equal(3, s.list.Len())

	s.Equal(3, s.list.Back().Value)
	s.Equal(2, s.list.Back().Next.Value)
	s.Equal(1, s.list.Back().Next.Next.Value)
	s.Equal(1, s.list.Front().Value)
	s.Equal(2, s.list.Front().Prev.Value)
	s.Equal(3, s.list.Front().Prev.Prev.Value)

	s.Nil(s.list.Front().Prev.Prev.Prev)
	s.Nil(s.list.Back().Next.Next.Next)
}

func (s *ListTestSuite) TestRemoveOne() {
	i := s.list.PushFront(1)
	s.Equal(1, s.list.Len())
	s.list.Remove(i)
	s.Equal(0, s.list.Len())
	s.Nil(s.list.Front())
	s.Nil(s.list.Back())
}

func (s *ListTestSuite) TestRemoveFirst() {
	s.list.PushFront(3)
	s.list.PushFront(2)
	s.list.PushFront(1)

	i := s.list.Front()
	s.Equal(1, i.Value)

	s.list.Remove(i)
	s.Equal(2, s.list.Len())
	s.Equal(2, s.list.Front().Value)
	s.Equal(3, s.list.Back().Value)
}

func (s *ListTestSuite) TestRemoveLast() {
	s.list.PushFront(3)
	s.list.PushFront(2)
	s.list.PushFront(1)

	i := s.list.Back()
	s.Equal(3, i.Value)

	s.list.Remove(i)
	s.Equal(2, s.list.Len())
	s.Equal(1, s.list.Front().Value)
	s.Equal(2, s.list.Back().Value)
}

func (s *ListTestSuite) TestRemoveMiddle() {
	s.list.PushFront(3)
	s.list.PushFront(2)
	s.list.PushFront(1)

	i := s.list.Front().Prev
	s.Equal(2, i.Value)

	s.list.Remove(i)
	s.Equal(2, s.list.Len())
	s.Equal(1, s.list.Front().Value)
	s.Equal(3, s.list.Back().Value)
}

func TestList(t *testing.T) {
	//	t.Run("complex", func(t *testing.T) {
	//		l := NewList()
	//
	//		l.PushFront(10) // [10]
	//		l.PushBack(20)  // [10, 20]
	//		l.PushBack(20)  // [10, 20]
	//		l.PushBack(30)  // [10, 20, 30]
	//		l.PushBack(30)  // [10, 20, 30]
	//		require.Equal(t, 3, l.Len())
	//		require.Equal(t, 3, l.Len())
	//		require.Equal(t, 3, l.Len())
	//
	//		middle := l.Front().Next // 20
	//		l.Remove(middle)         // [10, 30]
	//		require.Equal(t, 2, l.Len())
	//
	//		for i, v := range [...]int{40, 50, 60, 70, 80} {
	//			if i%2 == 0 {
	//				l.PushFront(v)
	//			} else {
	//				l.PushBack(v)
	//			}
	//		} // [80, 60, 40, 10, 30, 50, 70]
	//
	//		require.Equal(t, 7, l.Len())
	//		require.Equal(t, 80, l.Front().Value)
	//		require.Equal(t, 70, l.Back().Value)
	//
	//		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
	//		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]
	//
	//		elems := make([]int, 0, l.Len())
	//		for i := l.Front(); i != nil; i = i.Next {
	//			elems = append(elems, i.Value.(int))
	//		}
	//		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	//	})
}
