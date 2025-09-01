package math

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Parser парсер математических выражжений
type Parser struct {
	expression string
	pos        int
}

// NewParser создает новый парсер
func NewParser(expression string) *Parser {
	return &Parser{
		expression: strings.ReplaceAll(expression, " ", ""),
		pos:        0,
	}
}

// Parse разбирает и вычисляет выражение
func (p *Parser) Parse() (float64, error) {
	result, err := p.parseExpression()
	if err != nil {
		return 0, err
	}

	if p.pos < len(p.expression) {
		return 0, fmt.Errorf("неожиданый символ: %c", p.expression[p.pos])
	}
	return result, nil

}

// parseExpression разбирает выражение (самый низкий приоритет)
func (p *Parser) parseExpression() (float64, error) {
	result, err := p.parseTerm()
	if err != nil {
		return 0, err
	}
	for p.pos < len(p.expression) {
		op := p.expression[p.pos]
		if op != '+' && op != '-' {
			break
		}
		p.pos++

		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}

		switch op {
		case '+':
			result += right
		case '-':
			result -= right
		}
	}
	return result, nil
}

// parseTerm разбирает слагаемые (*, /)
func (p *Parser) parseTerm() (float64, error) {
	result, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for p.pos < len(p.expression) {
		op := p.expression[p.pos]
		if op != '*' && op != '/' {
			break
		}
		p.pos++

		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		switch op {
		case '*':
			result *= right
		case '/':
			result /= right
			if right == 0 {
				return 0, fmt.Errorf("деление на ноль")
			}
			result /= right
		}

	}
	return result, nil

}

// parseFactor разбирает множетели (^, !)
func (p *Parser) parseFactor() (float64, error) {
	result, err := p.parsePower()
	if err != nil {
		return 0, err
	}

	//Обработка факториала
	if p.pos < len(p.expression) && p.expression[p.pos] == '!' {
		p.pos++
		if result != float64(int(result)) || result < 0 {
			return 0, fmt.Errorf("факториал только для целых неотрицательных чисел")
		}
		result = float64(factorial(int(result)))
	}

	return result, nil

}

// parsePower разбирает степени (^)
func (p *Parser) parsePower() (float64, error) {
	result, err := p.parsePrimary()
	if err != nil {
		return 0, err
	}

	if p.pos < len(p.expression) && p.expression[p.pos] == '^' {
		p.pos++
		right, err := p.parsePrimary()
		if err != nil {
			return 0, err
		}

		result = Pow(result, right)
	}
	return result, nil

}

func (p *Parser) parsePrimary() (float64, error) {
	if p.pos >= len(p.expression) {
		return 0, fmt.Errorf("неожиданный конец выражения")
	}

	//Обработка скобок
	if p.expression[p.pos] == '(' {
		p.pos++
		result, err := p.parseExpression()
		if err != nil {
			return 0, err
		}
		if p.pos >= len(p.expression) || p.expression[p.pos] != ')' {
			return 0, fmt.Errorf("ожидается закрывающая скобка")
		}
		p.pos++
		return result, nil
	}

	//Обработка унарного минуса
	if p.expression[p.pos] == '-' {
		p.pos++
		result, err := p.parsePrimary()
		if err != nil {
			return 0, err
		}
		return -result, nil
	}

	//Обработка чисел
	start := p.pos
	for p.pos < len(p.expression) && (unicode.IsDigit(rune(p.expression[p.pos])) || p.expression[p.pos] == '.') {
		p.pos++
	}

	if start == p.pos {
		return 0, fmt.Errorf("ожидается число или выражение")
	}

	numberStr := p.expression[start:p.pos]
	result, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, fmt.Errorf("неверное число: %s", numberStr)
	}
	return result, nil

}

// Pow возведение в степень
func Pow(base, exponent float64) float64 {
	result := 1.0
	for i := 0; i < int(exponent); i++ {
		result *= base
	}
	return result
}
