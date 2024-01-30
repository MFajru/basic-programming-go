package solution

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type product struct {
	barcode string
	name    string
	price   int
}

type purchasedProduct struct {
	product
	qty      int
	subtotal int
}

func calculatePurchase(barcodes []string) (map[string]purchasedProduct, int) {
	products := map[string]product{
		"00001": {
			barcode: "00001",
			name:    "Coca-cola",
			price:   3000,
		},
		"00002": {
			barcode: "00002",
			name:    "Sprite",
			price:   2500,
		},
		"00003": {
			barcode: "00003",
			name:    "Fanta",
			price:   2500,
		},
		"00004": {
			barcode: "00004",
			name:    "Instant Noodles",
			price:   3500,
		},
		"00005": {
			barcode: "00005",
			name:    "Coffee",
			price:   5000,
		},
	}

	buyedProducts := map[string]purchasedProduct{}
	var grandTotal int

	for _, barcode := range barcodes {
		buyed, isBuyed := buyedProducts[barcode]
		_, prodExist := products[barcode]
		if !isBuyed && prodExist {
			buyed.product = products[barcode]
			buyed.qty = 1
			buyed.subtotal = products[barcode].price
			buyedProducts[barcode] = buyed
			grandTotal += products[barcode].price
		} else if isBuyed && prodExist {
			buyed.qty++
			buyed.subtotal += products[barcode].price
			buyedProducts[barcode] = buyed
			grandTotal += products[barcode].price
		}
	}

	return buyedProducts, grandTotal
}

func generateReceiptText(buyedProducts map[string]purchasedProduct, grandTotal int) {
	printOut := tabwriter.NewWriter(os.Stdout, 17, 1, 1, ' ', 0)
	fmt.Println("Purchased Summary")
	for _, product := range buyedProducts {
		fmt.Fprintf(printOut, "%s\t%d  %d\n", product.name, product.qty, product.subtotal)
		printOut.Flush()
	}
	fmt.Printf("Total: %d\n", grandTotal)
}

func Pos() {
	var barcodes []string

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)

	// scanner.Scan()

	// barcodes = append(barcodes, scanner.Text())
	var input string
	fmt.Print("Welcome to Juju Mart\n\n")
	fmt.Println("Please input dash symbol (-) to end the input")
	fmt.Println("Enter the barcode: ")
	for {
		fmt.Scan(&input)
		if input == "-" {
			break
		}
		barcodes = append(barcodes, input)
	}

	buyedProduct, grandTotal := calculatePurchase(barcodes)
	generateReceiptText(buyedProduct, grandTotal)

}
