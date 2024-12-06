# Latihan Golang E-Commerce API

Proyek ini adalah API backend sederhana untuk mengelola produk dan kategori menggunakan **Golang**. Dibangun dengan **GORM** untuk ORM dan **Fiber** sebagai framework web, proyek ini cocok untuk aplikasi e-commerce atau sistem inventarisasi.

## 🛠️ Fitur

- **Manajemen Produk**: Tambah, baca, ubah, dan hapus data produk.
- **Manajemen Kategori**: Tambah, baca, ubah, dan hapus data kategori.
- **Relasi Produk dan Kategori**: Produk terkait dengan kategori melalui `CategoryID`.
- **Validasi Input**: Validasi data menggunakan tag bawaan struct.
- **RESTful API**: Endpoint untuk operasi CRUD.

## 🚀 Teknologi yang Digunakan

- **Golang** (Backend)
- **Fiber** (Web Framework)
- **GORM** (ORM)
- **PostgreSQL** (Database)

## 📦 Struktur Proyek

```plaintext
latihan-golang/
├── database/          # Koneksi dan migrasi database
├── models/            # Definisi model (Product & Category)
├── routes/            # Routing API
├── main.go            # Entry point aplikasi
└── README.md          # Dokumentasi proyek
