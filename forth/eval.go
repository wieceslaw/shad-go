//go:build !solution

package main

import (
	"errors"
	"strings"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Peek() int {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Pop() (int, bool) {
	if s == nil {
		return 0, false
	}

	n := len(s.elements) - 1
	value := s.elements[n]
	s.elements[n] = 0
	if n == 0 {
		s.elements = nil
	} else {
		s.elements = s.elements[:n]
	}
	return value, true
}

func (s *Stack) Size() int {
	return len(s.elements)
}

type Evaluator struct {
	stack       Stack
	definitions map[string]string
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.

func (e *Evaluator) apply(word string) error {
	switch word {
	case "+":
		{
			if e.stack.Size() < 2 {
				return errors.New("Not enough elements")
			}
			first, _ := e.stack.Pop()
			second, _ := e.stack.Pop()
			e.stack.Push(first + second)
		}
	case "-":
		{

		}
	}
	return nil
}

func (e *Evaluator) Process(row string) ([]int, error) {
	if len(row) == 0 {
		return nil, errors.New("Empty string")
	}

	words := strings.Fields(row)
	if words[0] == ":" {
		// TODO: definition
		return e.stack.elements, nil
	}

	for word := range words {

	}

	return nil, nil
}
