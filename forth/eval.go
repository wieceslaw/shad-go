//go:build !solution

package main

import (
	"errors"
	"strconv"
	"strings"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Peek() (int, bool) {
	if s == nil {
		return 0, false
	}
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack) Pop() (int, bool) {
	if s == nil {
		return 0, false
	}
	if len(s.elements) == 0 {
		return 0, false
	}

	n := len(s.elements) - 1
	value := s.elements[n]
	s.elements[n] = 0
	if n == 0 {
		s.elements = []int{}
	} else {
		s.elements = s.elements[:n]
	}
	return value, true
}

func (s *Stack) Size() int {
	if s == nil {
		return 0
	}
	return len(s.elements)
}

func (s *Stack) Append(other *Stack) {
	s.elements = append(s.elements, other.elements...)
}

type Evaluator struct {
	stack                Stack
	definitions          map[string][]string
	processingDefinition bool
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{
		stack:                Stack{},
		definitions:          map[string][]string{},
		processingDefinition: false,
	}
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.

func (e *Evaluator) Process(row string) ([]int, error) {
	tokens := strings.Fields(strings.ToLower(row))
	if len(tokens) == 0 {
		return []int{}, nil
	}

	index := 0
	for index < len(tokens) {
		token := tokens[index]
		if token == ":" {
			err := e.processDefinition(tokens, &index)
			if err != nil {
				return []int{}, err
			}
		} else {
			err := processSingle(&e.stack, e.definitions, tokens, &index)
			if err != nil {
				return []int{}, err
			}
		}
	}

	return e.stack.elements, nil
}

func (e *Evaluator) processDefinition(tokens []string, index *int) error {
	if e.processingDefinition {
		return errors.New("already processing definition")
	}
	e.processingDefinition = true

	*index += 1
	word, err := atIndex(tokens, *index)
	if err != nil {
		return err
	}
	if isNumber(word) {
		return errors.New("unexpected number")
	}

	*index += 1
	stack := Stack{}
	for tokens[*index] != ";" {
		err = processSingle(&stack, e.definitions, tokens, index)
		if err != nil {
			return err
		}
	}
	e.definitions[word] = stack

	*index += 1
	e.processingDefinition = false

	return nil
}

func atIndex[T any](arr []T, index int) (T, error) {
	if index < 0 || index >= len(arr) {
		var zero T
		return zero, errors.New("")
	}
	return arr[index], nil
}

func processSingle(
	stack *Stack,
	defs map[string]Stack,
	tokens []string,
	index *int,
) error {
	token := tokens[*index]
	switch {
	case token == "+":
		err := handlePlusOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "-":
		err := handleMinusOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "*":
		err := handleMultiplyOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "/":
		err := handleDivisionOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "dup":
		err := handleDupOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "over":
		err := handleOverOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "drop":
		err := handleDropOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	case token == "swap":
		err := handleSwapOperation(stack, tokens, index)
		if err != nil {
			return err
		}
	default:
		if isNumber(token) {
			err := processNumber(stack, tokens, index)
			if err != nil {
				return err
			}
		} else if _, ok := defs[token]; ok {
			err := handleWordOperation(stack, defs, tokens, index)
			if err != nil {
				return err
			}
		} else {
			return errors.New("unknown operation: '" + token + "'")
		}

	}
	return nil
}

func handlePlusOperation(stack *Stack, _ []string, index *int) error {
	if stack.Size() < 2 {
		return errors.New("not enough elements")
	}
	first, _ := stack.Pop()
	second, _ := stack.Pop()
	stack.Push(first + second)
	*index += 1
	return nil
}

func handleMinusOperation(stack *Stack, _ []string, index *int) error {
	if stack.Size() < 2 {
		return errors.New("not enough elements")
	}
	first, _ := stack.Pop()
	second, _ := stack.Pop()
	stack.Push(second - first)
	*index += 1
	return nil
}

func handleMultiplyOperation(stack *Stack, _ []string, index *int) error {
	if stack.Size() < 2 {
		return errors.New("not enough elements")
	}
	first, _ := stack.Pop()
	second, _ := stack.Pop()
	stack.Push(first * second)
	*index += 1
	return nil
}

func handleDivisionOperation(stack *Stack, _ []string, index *int) error {
	if stack.Size() < 2 {
		return errors.New("not enough elements")
	}
	first, _ := stack.Pop()
	if first == 0 {
		return errors.New("zero division")
	}
	second, _ := stack.Pop()
	stack.Push(second / first)
	*index += 1
	return nil
}

func handleDupOperation(stack *Stack, _ []string, index *int) error {
	top, ok := stack.Peek()
	if !ok {
		return errors.New("expected at least one element")
	}

	stack.Push(top)

	*index += 1
	return nil
}

func handleOverOperation(stack *Stack, _ []string, index *int) error {
	first, ok := stack.Pop()
	if !ok {
		return errors.New("expected at least one element")
	}
	second, ok := stack.Peek()
	if !ok {
		return errors.New("expected at least one element")
	}

	stack.Push(first)
	stack.Push(second)

	*index += 1
	return nil
}

func handleDropOperation(stack *Stack, _ []string, index *int) error {
	_, ok := stack.Pop()
	if !ok {
		return errors.New("expected at least one element")
	}

	*index += 1
	return nil
}

func handleSwapOperation(stack *Stack, _ []string, index *int) error {
	first, ok := stack.Pop()
	if !ok {
		return errors.New("expected at least one element")
	}
	second, ok := stack.Pop()
	if !ok {
		return errors.New("expected at least one element")
	}

	stack.Push(first)
	stack.Push(second)

	*index += 1
	return nil
}

func handleWordOperation(
	stack *Stack,
	defs map[string][]string,
	tokens []string,
	index *int,
) error {
	word := tokens[*index]

	def, ok := defs[word]
	if !ok {
		return errors.New("unknown symbol: " + word)
	}

	// TODO: execute def tokens
	stack.Append(&def)

	*index += 1
	return nil
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func processNumber(stack *Stack, tokens []string, index *int) error {
	token := tokens[*index]

	number, err := strconv.Atoi(token)
	if err != nil {
		return err
	}
	stack.Push(number)

	*index += 1
	return nil
}
