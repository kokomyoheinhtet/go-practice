package main

import "fmt"

type Size int

const (
	small Size = iota
	medium
	large
)

type Color int

const (
	red Color = iota
	green
	blue
)

type Product struct {
	name  string
	size  Size
	color Color
}

func PrintProducts(name string, products []Product) {
	fmt.Println(name)
	for _, v := range products {
		fmt.Println(v)
	}
}

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []Product {
	result := make([]Product, 0)

	for _, v := range products {
		if v.color == color {
			result = append(result, v)
		}
	}

	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []Product {
	result := make([]Product, 0)

	for _, v := range products {
		if v.size == size {
			result = append(result, v)
		}
	}

	return result
}

// obeying OCP

type Specification interface {
	IsSatisfied(p Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []Product {
	result := make([]Product, 0)
	for _, v := range products {
		if spec.IsSatisfied(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	apple := Product{"apple", small, green}
	tree := Product{"tree", large, green}
	house := Product{"house", large, red}
	products := []Product{apple, tree, house}

	fmt.Println("Old Filter")
	f := Filter{}
	PrintProducts("Green Products", f.FilterByColor(products, green))
	PrintProducts("Large Products", f.FilterBySize(products, large))
	fmt.Println()

	fmt.Println("New Filter")
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	bf := BetterFilter{}
	PrintProducts("Green Products", bf.Filter(products, greenSpec))
	PrintProducts("Large Products", bf.Filter(products, largeSpec))
	PrintProducts("Large Green Products", bf.Filter(products, largeGreenSpec))
}
