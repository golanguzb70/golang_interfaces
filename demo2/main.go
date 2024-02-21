package main

import "fmt"

type Printer struct {
}

type Printer2 struct {
}

type Scanner struct {
}

type printer interface {
	Print()
}

type scanner interface {
	Scan()
}

type printerScanner interface {
	printer
	scanner
}

func (p *Printer) Print() {
	fmt.Println("Print something by Printer 1")
}

func (p *Printer2) Print() {
	fmt.Println("Print something by Printer 2")
}

func (s *Scanner) Scan() {
	fmt.Println("Scanned something")
}

type PrinterScanner struct {
}

func (p *PrinterScanner) Print() {
	fmt.Println("Print something by new Printer")
}

func (s *PrinterScanner) Scan() {
	fmt.Println("Scanned something by new Scanner")
}

func main() {
	var printerSc printerScanner

	printerSc = &PrinterScanner{}

	printerSc.Print()
}
