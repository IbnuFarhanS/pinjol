package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/IbnuFarhanS/pinjol/model"
)

func ExportTransactionsToCSV(transactions []model.Transactions, filePath string) error {
	// Membuka file CSV untuk penulisan
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Membuat penulis CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Menulis header kolom ke penulis CSV
	header := []string{"ID", "Product ID", "User ID", "Status", "Amount", "Created At", "Due Date"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	// Menulis data transaksi ke penulis CSV
	for _, transaction := range transactions {
		row := []string{
			strconv.FormatInt(transaction.ID, 10),
			strconv.FormatInt(transaction.ProductsID, 10),
			strconv.FormatInt(transaction.UsersID, 10),
			strconv.FormatBool(transaction.Status),
			strconv.FormatFloat(transaction.Amount, 'f', -1, 64),
			transaction.Created_At.String(),
			transaction.Due_Date.String(),
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
