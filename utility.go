package util

import (
	"github.com/tealeg/xlsx"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var header = []string {"Title", "Business type", "Rating", "Mobile Number", "Phone Number", "Address", "Website"}
const fileNamePrefix = "data"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const outputdir = "output"

type Item struct {
	Title string
	Rating string
	Mapcoord string
	Mobileno string
	Whatsappno string
	Phonenumber string
	Address string
	Businesstype string
	Email string
	Website string
	Area string
}

func ToFloat(s string) float64  {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalln("Error in util.ToFloat : when converting string to float", err)
		f = 0
	}
	return f
}

func Ftoa(f float64) string {
	return strconv.FormatFloat(f, 64, 2, 2)
}

func ToInt (s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("Error in util.ToInt : when converting string to int", err)
		i = 0
	}
	return i
}

func ToString (i int) string {
	s := strconv.Itoa(i)
	return s
}

func WriteToFile(items []Item, sheetName string)  {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		log.Fatalln("Error in WriteToFile : when adding sheet", err)
	}
	WriteHeader(sheet)
	for _, item := range items {
		WriteRow(sheet, item)
	}
	SaveFile(file)
}

func WriteHeader(sheet *xlsx.Sheet){
	row := sheet.AddRow()
	for _, v := range header {
		addCell(row, v)
	}
}

func addCell (row *xlsx.Row, value string) {
	cell := row.AddCell()
	cell.Value = value
}
func WriteRow(sheet *xlsx.Sheet, item Item) {
	row := sheet.AddRow()
	//use same order as declared in var header
	//"Title", "Business type", "Rating", "Mobile Number", "Phone Number", "Address", "Website"
	addCell(row, item.Title)
	addCell(row, item.Businesstype)
	addCell(row, item.Rating)
	addCell(row, item.Mobileno)
	addCell(row, item.Phonenumber)
	addCell(row, item.Address)
	addCell(row, item.Website)
}

func SaveFile ( file *xlsx.File){
	os.MkdirAll(outputdir, os.ModePerm)
	err := file.Save(getFileName())
	if err != nil {
		log.Fatalln("Error in util.WriteHeader : when saving file", err)
	}

}

func getFileName()string {
	fname := fileNamePrefix+"_"+time.Now().Format("20060102150405")+".xlsx"
	return filepath.Join(outputdir, fname)
}

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}