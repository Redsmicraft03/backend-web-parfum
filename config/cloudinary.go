package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cld *cloudinary.Cloudinary

func InitCloudinary() {
	var err error

	Cld, err = cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatalf("Gagal menginisialisasi CLOUDINARY: %v", err)
	}

	log.Println("Koneksi Cloudinary berhasil dibuat")
}