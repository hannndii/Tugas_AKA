

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const NMAX = 100

type tabDesa [NMAX]dataDesa

type dataDesa struct {
    namaDesa   string
    alamatDesa string
    jumlahRt   int
    jumlahRw   int
	pendapatanUMKM int
    penduduk   [NMAX]dataPenduduk // Penambahan slice untuk menyimpan penduduk
}

type dataPenduduk struct {
    namaPenduduk     string
    umurPenduduk     int
    alamatRumah      string
    noRT             int
    noRW             int
    noNIK            int
    statusPerkawinan string
	pendudukUMKM int
}

	type Login struct {
		username string
		email    string
		password string
	}

// ----------KAMUS GLOBAL----------
var dataD tabDesa
var nDesa int
var nPenduduk int
var shouldExit bool
var namaDicari string
var pilih, nomorNIK int

func main() {
	// var input Login
	// login(&input)
	loading()
    menu()
}

// Tampilan Menu Utama
func menu() {
    for !shouldExit {
        fmt.Println("================================")
        fmt.Printf("%18s\n", "SI DESA")
        fmt.Println(" Aplikasi Sistem Informasi Desa")
        fmt.Println("================================")
        fmt.Println("Pilih menu:")
        fmt.Println("1. Input Data")
        fmt.Println("2. Cari Data")
        fmt.Println("3. Hapus Data")
        fmt.Println("4. Edit Data")
        fmt.Println("5. Cetak Data")
        fmt.Println("6. UMKM")
        fmt.Println("7. Exit")
        fmt.Println()
        fmt.Print("Pilih:")
        fmt.Scan(&pilih)
        fmt.Println("================================")
        fmt.Println()

        switch pilih {
        case 1:
            clearScreen()
            inputData(&dataD, &nDesa, &nPenduduk)
        case 2:
            clearScreen()
            cariData(dataD, nDesa, nPenduduk)
        case 3:
            deleteData(&dataD, &nDesa, &nPenduduk)
         case 4:
            clearScreen()
            editData(&dataD, nDesa, nPenduduk)
        case 5:
            clearScreen()
            cetakData(dataD, nDesa ,nPenduduk)
		case 6:
            clearScreen()
            tambahUMKM(&dataD, nDesa)
        case 7:
            shouldExit = true
            fmt.Println("================================")
            fmt.Println("TERIMA KASIH")
            fmt.Println("================================")
        default:
            fmt.Println("Masukan tidak valid!")
            clearScreen()
            menu()
        }
    }
}

// -------LOADING---------
func loading() {
    fmt.Println("Loading...")

    frames := []string{
        "⠋",
        "⠙",
        "⠹",
        "⠸",
        "⠼",
        "⠴",
        "⠦",
        "⠧",
        "⠇",
        "⠏",
    }

    for i := 0; i < 10; i++ {
        fmt.Printf("\r%s Processing... ", frames[i])
        time.Sleep(200 * time.Millisecond)
    }

    fmt.Println("\nDone!")
    time.Sleep(2 * time.Second)
}


// -------FUNGSI UNTUK LOGIN---------
func login(input *Login) {
	durasi := 3
	percobaanMaks := 2
	percobaan := 0
	input.username = "Admin"
	input.email = "Admin"
	input.password = "Admin"

	time.Sleep(time.Duration(durasi) * time.Second)
	fmt.Println("================================")
	fmt.Println("SELAMAT DATANG DI SIDESA!!")
	fmt.Println("================================")

	for percobaan <= percobaanMaks {
		fmt.Print("Username: ")
		fmt.Scan(&input.username)
		fmt.Print("Email: ")
		fmt.Scan(&input.email)
		fmt.Print("Password: ")
		fmt.Scan(&input.password)
		if input.username == "Admin" && input.email == "Admin" && input.password == "Admin" {
			break
		} else {
			fmt.Println("Data tidak sesuai")
			percobaan++
		}
		if percobaan == percobaanMaks {
			fmt.Println("================================")
			fmt.Println("MOHON MAAF DATA KAMU SALAH")
			fmt.Println("================================")
			fmt.Println("SILAHKAN TUNGGU DALAM WAKTU 1 MENIT")
			fmt.Println("================================")
			time.Sleep(time.Duration(durasi) * time.Second)
			login(input)
			return
		}
	}
}

// -------FUNGSI UNTUK MASUKKAN DATA DESA & PENDUDUK---------
func inputData(K *tabDesa, nDesa, nPendudukDesa *int) {
    /* {I.S. ___ 
        F.S. ___}
    */
	fmt.Println("================================")
	fmt.Println("SILAHKAN MASUKKAN DATA PENDUDUK")
	fmt.Println("================================")
	fmt.Print("Masukkan jumlah Desa: ")
	fmt.Scan(nDesa)

	for i := 0; i < *nDesa; i++ {
		fmt.Printf("Data Desa Ke-%d\n", i+1)
		fmt.Print("Nama Desa: ")
		fmt.Scan(&(*K)[i].namaDesa)
		fmt.Print("Alamat Desa: ")
		fmt.Scan(&(*K)[i].alamatDesa)
		fmt.Print("Jumlah RT: ")
		fmt.Scan(&(*K)[i].jumlahRt)
		fmt.Print("Jumlah RW: ")
		fmt.Scan(&(*K)[i].jumlahRw)

		fmt.Print("Masukkan jumlah penduduk: ")
		fmt.Scan(&(*nPendudukDesa))

		for j := 0; j < *nPendudukDesa; j++ {
			fmt.Printf("Data Penduduk Ke-%d Desa %s\n", j+1, (*K)[i].namaDesa)
			fmt.Print("Nama Penduduk: ")
			fmt.Scan(&(*K)[i].penduduk[j].namaPenduduk)
			fmt.Print("Alamat Penduduk: ")
			fmt.Scan(&(*K)[i].penduduk[j].alamatRumah)
			fmt.Print("Umur: ")
			fmt.Scan(&(*K)[i].penduduk[j].umurPenduduk)
			fmt.Print("RT: ")
			fmt.Scan(&(*K)[i].penduduk[j].noRT)
			fmt.Print("RW: ")
			fmt.Scan(&(*K)[i].penduduk[j].noRW)
			fmt.Print("Masukkan Nomor NIK: ")
			fmt.Scan(&(*K)[i].penduduk[j].noNIK)
			fmt.Print("Status Kawin?: ")
			fmt.Scan(&(*K)[i].penduduk[j].statusPerkawinan)
		}
	}
}

// -------FUNGSI UNTUK MENCETAK DATA DESA & PENDUDUK---------
func cetakData(K tabDesa, nDesa, nPenduduk int) {
    /* {I.S. ___ 
        F.S. ___}
    */
	fmt.Println("================================")
	loading()
    fmt.Println("================================")
    fmt.Println("CETAK DATA PENDUDUK DESA")
    fmt.Println("================================")

    for i := 0; i < nDesa; i++ {
        fmt.Printf("Data Desa ke-%d\n", i+1)
        fmt.Printf("Nama Desa: %s\n", K[i].namaDesa)
        fmt.Printf("Alamat Desa: %s\n", K[i].alamatDesa)
        fmt.Printf("Jumlah RT: %d\n", K[i].jumlahRt)
        fmt.Printf("Jumlah RW: %d\n", K[i].jumlahRw)
        fmt.Printf("Pendapatan UMKM Desa: %d\n", K[i].pendapatanUMKM)
        fmt.Println("================================")

        fmt.Printf("%-20s%-10s%-20s%-5s%-5s%-15s%-20s\n", "Nama Penduduk", "Umur", "Alamat", "RT", "RW", "NIK", "Status Perkawinan")
        for j := 0; j < nPenduduk; j++ {
            fmt.Printf("%-20s%-10d%-20s%-5d%-5d%-15d%-20s\n", 
            K[i].penduduk[j].namaPenduduk, K[i].penduduk[j].umurPenduduk, 
            K[i].penduduk[j].alamatRumah, K[i].penduduk[j].noRT, K[i].penduduk[j].noRW, 
            K[i].penduduk[j].noNIK, K[i].penduduk[j].statusPerkawinan)
        }
        fmt.Println("================================")
    }
}


// -------FUNGSI UNTUK MENCARI PENDUDUK---------
func cariData(K tabDesa, nDesa, nPenduduk int) {
    /* {I.S. ___ 
        F.S. ___}
    */
    var data dataPenduduk
    fmt.Println("================================")
    fmt.Println("PENCARIAN DATA PENDUDUK DESA")
    fmt.Println("================================")
    fmt.Println("MENU PILIHAN PENCARIAN")
    fmt.Println("1. Nama")
    fmt.Println("2. NIK")
    fmt.Println("3. EXIT")
    fmt.Print("Pilih: ")
    fmt.Scan(&pilih)
    var searchField string
    switch pilih {
    case 1:
        searchField = "Nama"
    case 2:
        searchField = "NIK"
    case 3:
        return
    }

    fmt.Println("================================")
    fmt.Printf("Masukkan %s: ", searchField)
    var searchQuery string
    fmt.Scan(&searchQuery)

    found := false
    for i := 0; i < nDesa; i++ { //jika i kurang dari nDesa dia jalan
        for j := 0; j < nPenduduk; j++ { // jika j kurang dari jumlah RT yang ada di desa ke-i dia jalan
                 if pilih == 1 {
                     data = K[i].penduduk[j]
                 } else {
                     data = K[i].penduduk[j]
                 }

                var match bool
                if pilih == 1 {
                    match = data.namaPenduduk == searchQuery
                } else {
                    match = data.noNIK == nomorNIK
                }

                if match {
					fmt.Println("================================")
				    loading()
                    fmt.Printf("Data Ditemukan: %+v\n", data)
					
                    found = true
                }
        }
    }
    if !found {
        fmt.Println("Data Tidak Ada")
    }

}

func editData(T *tabDesa, nDesa, nPenduduk int) {
    /* {I.S. ___ 
        F.S. ___}
    */
    var (
        //Variabel untuk edit desa
        nama_desa string
        alamat_desa string
        jumlah_rt int
        jumlah_rw int
        //variabel untuk edit penduduk
        nama string
        umur int
        NIK int
        alamat string
        rt int
        rw int
    )
	fmt.Println("================================")
    fmt.Println("Menu Edit Data")
	fmt.Println("1. Edit Data Desa")
	fmt.Println("2. Edit Data Penduduk")
	fmt.Println("3. Edit Status Perkawinan")
	fmt.Println("4. Exit")
	fmt.Println("================================")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilih)

	switch pilih {
    case 1:
        fmt.Print("Masukan Nama Desa: ")
        fmt.Scan(&nama_desa)
    case 2:
        fmt.Println("Masukan Nama Penduduk: ")
        fmt.Scan(&nama)
    case 3:
        ubahStatusPerkawinan(T, nDesa, nPenduduk)
    case 4:
        return
    }

    found := false
    for i := 0; i < nDesa; i++ {
        for j := 0; j < nPenduduk; j++ {
                var desaMatch, pendudukMatch bool
                if pilih == 1 {
                    desaMatch = T[i].namaDesa == nama_desa
                } else {
                    //cariData(*T, nDesa, nPenduduk)
                    pendudukMatch = T[i].penduduk[j].namaPenduduk == nama
                }

                if desaMatch {
                    fmt.Printf("Data Desa Ditemukan:\n")
				    fmt.Printf("%-20s%-20s%-20s%-20s\n", "Nama Desa", "Alamat Desa", "Jumlah RT", "Jumlah RW")
				    fmt.Printf("%-20s%-20s%-20d%-20d\n", T[i].namaDesa, T[i].alamatDesa, T[i].jumlahRt, T[i].jumlahRw )
                    fmt.Println("")
                    fmt.Println("Masukan data desa yang ingin diubah: ")
                    fmt.Print("Nama Desa: ")
		            fmt.Scan(&nama_desa)
		            fmt.Print("Alamat Desa: ")
		            fmt.Scan(&alamat_desa)
		            fmt.Print("Jumlah RT: ")
		            fmt.Scan(&jumlah_rt)
		            fmt.Print("Jumlah RW: ")
		            fmt.Scan(&jumlah_rw)
                    fmt.Println("Data Berhasil Diubah!")
                    T[i].namaDesa = nama_desa
                    T[i].alamatDesa = alamat_desa
                    T[i].jumlahRt = jumlah_rt
                    T[i].jumlahRw = jumlah_rw
                    found = true
                } else if pendudukMatch {
                    fmt.Printf("Data Penduduk Ditemukan:\n")
				    fmt.Printf("%-20s%-20s%-10s%-5s%-5s%-15s%-20s\n", "Nama Penduduk", "Umur", "Alamat", "RT", "RW", "NIK", "Status Perkawinan")
				    fmt.Printf("%-20s%-20s%-10d%-5d%-5d%-15d%-20s\n", T[i].penduduk[j].namaPenduduk, T[i].penduduk[j].alamatRumah, T[i].penduduk[j].umurPenduduk, T[i].penduduk[j].noRT, T[i].penduduk[j].noRW, T[i].penduduk[j].noNIK, T[i].penduduk[j].statusPerkawinan)
                    fmt.Println("")
                    fmt.Println("Masukan data penduduk yang ingin diubah: ")
                    fmt.Print("Nama: ")
                    fmt.Scan(&nama)
                    fmt.Print("Alamat: ")
                    fmt.Scan(&alamat)
                    fmt.Print("Umur: ")
                    fmt.Scan(&umur)
                    fmt.Print("RT: ")
                    fmt.Scan(&rt)
                    fmt.Print("RW: ")
                    fmt.Scan(&rw)
                    fmt.Print("NIK: ")
                    fmt.Scan(&NIK)
					fmt.Println("================================")

				    loading()
                    fmt.Println("Data Berhasil Diubah!")
                    T[i].penduduk[j].namaPenduduk = nama
                    T[i].penduduk[j].umurPenduduk = umur
                    T[i].penduduk[j].alamatRumah = alamat
                    T[i].penduduk[j].noRT = rt
                    T[i].penduduk[j].noRW = rw
                    T[i].penduduk[j].noNIK = NIK
                    found = true
                }
        }
    }
    if !found {
        fmt.Println("Data Tidak Ada")
    }
}

// -------FUNGSI UBAH STATUS PERKAWINAN PENDUDUK--------
func ubahStatusPerkawinan(K *tabDesa, nDesa, nPenduduk int) {
    fmt.Println("================================")
    fmt.Println("PERUBAHAN STATUS PERKAWINAN")
    fmt.Println("================================")
    fmt.Print("Masukkan Nomor NIK penduduk: ")
    fmt.Scan(&nomorNIK)

    found := false
    for i := 0; i < nDesa; i++ {
        for j := 0; j < nPenduduk; j++ {
            if K[i].penduduk[j].noNIK == nomorNIK {
				fmt.Printf("Data Penduduk Ditemukan:\n")
				fmt.Printf("%-20s%-20s%-10s%-5s%-5s%-15s%-20s\n", "Nama Penduduk", "Umur", "Alamat", "RT", "RW", "NIK", "Status Perkawinan")
				fmt.Printf("%-20s%-20s%-10d%-5d%-5d%-15d%-20s\n", (*K)[i].penduduk[j].namaPenduduk, 
                (*K)[i].penduduk[j].alamatRumah, (*K)[i].penduduk[j].umurPenduduk, (*K)[i].penduduk[j].noRT, 
                (*K)[i].penduduk[j].noRW, (*K)[i].penduduk[j].noNIK, (*K)[i].penduduk[j].statusPerkawinan)

				fmt.Print("Masukkan status perkawinan baru: ")
				var newStatus string
				fmt.Scan(&newStatus)
				fmt.Println("================================")

				loading()
				(*K)[i].penduduk[j].statusPerkawinan = newStatus
				fmt.Println("Status perkawinan berhasil diubah.")
				found = true
			}
        }
    }
    if !found {
        fmt.Println("Data Penduduk dengan NIK tersebut tidak ditemukan.")
    }
}

func deleteData(T *tabDesa, nDesa, nPenduduk *int){
    var nama_desa, nama string
    fmt.Println("================================")
    fmt.Println("Menu Edit Data")
	fmt.Println("1. Hapus Desa")
	fmt.Println("2. Hapus Data Penduduk")
	fmt.Println("3. Exit")
	fmt.Println("================================")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilih)
    
    found := false
	switch pilih {
    case 1:
        fmt.Print("Masukan Nama Desa: ")
        fmt.Scan(&nama_desa)
        for i := 0; i < *nDesa; i++{
            if T[i].namaDesa == nama_desa {
                for k := i; k < *nDesa - 1; k++{
                    T[k].namaDesa = T[k+1].namaDesa
                    T[k].alamatDesa = T[k+1].alamatDesa
                    T[k].jumlahRt = T[k+1].jumlahRt
                    T[k].jumlahRw = T[k+1].jumlahRw
                }
                T[*nDesa] = T[*nDesa - 1]
                *nDesa--
                found = true
            }
    }
    case 2:
        fmt.Println("Masukan Nama Penduduk: ")
        fmt.Scan(&nama)
        for i := 0; i < *nDesa; i++ {
            for j := 0; j < len(T[i].penduduk); j++ {
                if T[i].penduduk[j].namaPenduduk == nama {
                    for k := j; k < len(T[i].penduduk) - 1; k++{
                        T[i].penduduk[k].namaPenduduk = T[i].penduduk[k+1].namaPenduduk
                        T[i].penduduk[k].alamatRumah = T[i].penduduk[k+1].alamatRumah
                        T[i].penduduk[k].umurPenduduk = T[i].penduduk[k+1].umurPenduduk
                        T[i].penduduk[k].noRT = T[i].penduduk[k+1].noRT
                        T[i].penduduk[k].noRW = T[i].penduduk[k+1].noRW
                        T[i].penduduk[k].noNIK = T[i].penduduk[k+1].noNIK
                        T[i].penduduk[k].statusPerkawinan = T[i].penduduk[k+1].statusPerkawinan
                    }
                    
                    *nPenduduk--
                    found = true
                }
            }
        }
    case 3:
        return
    }
    
    if !found {
        fmt.Println("Data Tidak Ada")
    }
}

// -------FUNGSI UNTUK MENAMBAH DATA UMKM---------
func tambahUMKM(K *tabDesa, nDesa int) {
	var namaDesa, namaUMKM string
	var found bool
	var pendapatanUMKM int

	fmt.Println("================================")
	fmt.Println("Tambah Data UMKM")
	fmt.Println("================================")
	fmt.Print("Masukkan Nama Desa: ")
	fmt.Scan(&namaDesa)

	for i := 0; i < nDesa; i++ {
		if K[i].namaDesa == namaDesa {
			fmt.Printf("Data Desa Ditemukan: %+v\n", K[i].namaDesa)
			fmt.Print("Nama UMKM: ")
			fmt.Scan(&namaUMKM)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			namaUMKM = scanner.Text()
			fmt.Print("Pendapatan UMKM: ")
			fmt.Scan(&K[i].pendapatanUMKM)
			K[i].pendapatanUMKM += pendapatanUMKM
			fmt.Println("================================")

			loading()
			fmt.Println("Data UMKM berhasil ditambahkan.")
			found = true
		}
	}

	if !found {
		fmt.Println("Data Desa tidak ditemukan.")
	} 
}

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func waitForEnter() {
	fmt.Println("\nTekan Enter untuk lanjutkan...")
	scanner := bufio.NewReader(os.Stdin)
	scanner.ReadBytes('\n')
}