package main

import "fmt"

type Warga struct {
	ID          int
	Nama        string
	Alamat      string
	JenisSampah string
	BeratSampah float64
}

type Setoran struct {
	ID          int
	WargaID     int
	NamaWarga   string
	JenisSampah string
	Berat       float64
	Hari        int
	Bulan       int
	Tahun       int
	MingguKe    int
	Tanggal     string
}
func main() {
	var warga [100]Warga
	var jumlahWarga int
	var setoran [100]Setoran
	var jumlahSetoran int

	for {
		fmt.Println("\n=== Waste-Track ===")
		fmt.Println("1. Tambah data warga")
		fmt.Println("2. Ubah data warga")
		fmt.Println("3. Hapus data warga")
		fmt.Println("4. Tambah log setoran sampah")
		fmt.Println("5. Cari data warga")
		fmt.Println("6. Urutkan data warga")
		fmt.Println("7. Tampilkan data warga")
		fmt.Println("8. Tampilkan statistik sampah seminggu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)
		if pilihan == 0 {
			fmt.Println("Input tidak valid, silakan coba lagi.")
			return
		}

		switch pilihan {
		case 1:
			tambahWarga(&warga, &jumlahWarga)
		case 2:
			ubahWarga(&warga, jumlahWarga)
		case 3:
			hapusWarga(&warga, &jumlahWarga)
		case 4:
			tambahSetoran(&warga, jumlahWarga, &setoran, &jumlahSetoran)
		case 5:
			fmt.Println("\nPilih cara mencari data warga:")
			fmt.Println("1. Berdasarkan ID (linear search)")
			fmt.Println("2. Berdasarkan nama (binary search)")
			fmt.Print("Pilihan: ")
			var mode int
			fmt.Scanln(&mode)
			switch mode {
			case 1:
				cariSequential(warga, jumlahWarga)
			case 2:
				cariBinary(warga, jumlahWarga)
			default:
				fmt.Println("Pilihan tidak tersedia.")
			}
		case 6:
			fmt.Println("\nPilih cara mengurutkan data warga:")
			fmt.Println("1. Berdasarkan berat (selection sort)")
			fmt.Println("2. Berdasarkan ID (insertion sort)")
			fmt.Print("Pilihan: ")
			var mode int
			fmt.Scanln(&mode)
			switch mode {
			case 1:
				urutkanSelection(warga, jumlahWarga)
			case 2:
				urutkanInsertion(warga, jumlahWarga)
			default:
				fmt.Println("Pilihan tidak tersedia.")
			}
		case 7:
			tampilkanWarga(warga, jumlahWarga)
		case 8:
			tampilkanStatistik(setoran, jumlahSetoran)
		case 0:
			fmt.Println("Terima kasih telah menggunakan Waste-Track.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
func nextID(warga [100]Warga, count int) int {
	maxID := 0
	for i := 0; i < count; i++ {
		if warga[i].ID > maxID {
			maxID = warga[i].ID
		}
	}
	return maxID + 1
}
func nextSetoranID(setoran [100]Setoran, count int) int {
	maxID := 0
	for i := 0; i < count; i++ {
		if setoran[i].ID > maxID {
			maxID = setoran[i].ID
		}
	}
	return maxID + 1
}
func tambahWarga(warga *[100]Warga, count *int) {
	if *count >= 100 {
		fmt.Println("Kapasitas data warga sudah penuh.")
		return
	}
	var nama, alamat, jenis string
	var berat float64
	fmt.Print("Nama warga: ")
	fmt.Scanln(&nama)
	fmt.Print("Alamat: ")
	fmt.Scanln(&alamat)
	fmt.Print("Jenis sampah utama: ")
	fmt.Scanln(&jenis)
	fmt.Print("Berat sampah (kg): ")
	fmt.Scanln(&berat)
	(*warga)[*count] = Warga{ID: nextID(*warga, *count), Nama: nama, Alamat: alamat, JenisSampah: jenis, BeratSampah: berat}
	*count++
	fmt.Println("Data warga berhasil ditambahkan.")
}
func ubahWarga(warga *[100]Warga, count int) {
	if count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}
	var id int
	fmt.Print("Masukkan ID warga yang ingin diubah: ")
	fmt.Scanln(&id)
	for i := 0; i < count; i++ {
		if (*warga)[i].ID == id {
			var nama, alamat, jenis string
			var berat float64
			fmt.Print("Nama baru: ")
			fmt.Scanln(&nama)
			fmt.Print("Alamat baru: ")
			fmt.Scanln(&alamat)
			fmt.Print("Jenis sampah utama baru: ")
			fmt.Scanln(&jenis)
			fmt.Print("Berat sampah baru (kg): ")
			fmt.Scanln(&berat)
			(*warga)[i].Nama = nama
			(*warga)[i].Alamat = alamat
			(*warga)[i].JenisSampah = jenis
			(*warga)[i].BeratSampah = berat
			fmt.Println("Data warga berhasil diubah.")
			return
		}
	}
	fmt.Println("ID warga tidak ditemukan.")
}
func hapusWarga(warga *[100]Warga, count *int) {
	if *count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}
	var id int
	fmt.Print("Masukkan ID warga yang ingin dihapus: ")
	fmt.Scanln(&id)
	for i := 0; i < *count; i++ {
		if (*warga)[i].ID == id {
			for j := i; j < *count-1; j++ {
				(*warga)[j] = (*warga)[j+1]
			}
			*count--
			fmt.Println("Data warga berhasil dihapus.")
			return
		}
	}
	fmt.Println("ID warga tidak ditemukan.")
}
func tambahSetoran(warga *[100]Warga, countWarga int, setoran *[100]Setoran, count *int) {
	if countWarga == 0 {
		fmt.Println("Belum ada warga yang bisa dijadikan referensi setoran.")
		return
	}
	if *count >= 100 {
		fmt.Println("Kapasitas log setoran sudah penuh.")
		return
	}
	var wargaID int
	var berat float64
	var hari, bulan, tahun int
	fmt.Print("Masukkan ID warga: ")
	fmt.Scanln(&wargaID)
	idx := -1
	for i := 0; i < countWarga; i++ {
		if (*warga)[i].ID == wargaID {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("ID warga tidak ditemukan.")
		return
	}

	fmt.Println("\nData Setoran")
	fmt.Print("Hari: ")
	fmt.Scanln(&hari)
	fmt.Print("Bulan: ")
	fmt.Scanln(&bulan)
	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	var minggu int

	fmt.Print("Minggu ke (1-52): ")
	fmt.Scanln(&minggu)

	if minggu < 1 || minggu > 52 {
		fmt.Println("Minggu tidak valid.")
		return
	}

	if hari < 1 || bulan < 1 || bulan > 12 || tahun < 1 {
		fmt.Println("Input tanggal tidak valid. Setoran dibatalkan.")
		return
	}

	Tanggal := fmt.Sprintf("%02d/%02d/%04d", hari, bulan, tahun)

	fmt.Print("Berat setoran (kg): ")
	fmt.Scanln(&berat)

	(*setoran)[*count] = Setoran{
		ID:          nextSetoranID(*setoran, *count),
		WargaID:     wargaID,
		NamaWarga:   (*warga)[idx].Nama,
		JenisSampah: (*warga)[idx].JenisSampah,
		Berat:       berat,
		Hari:        hari,
		Bulan:       bulan,
		Tahun:       tahun,
		MingguKe:    minggu,
		Tanggal:     Tanggal,
	}
	*count++
	fmt.Printf("Log setoran berhasil dicatat pada tanggal %s.\n", Tanggal)
}
func cariSequential(warga [100]Warga, count int) {
	if count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}

	var id int
	fmt.Print("Masukkan ID: ")
	fmt.Scanln(&id)

	for i := 0; i < count; i++ {
		if warga[i].ID == id {
			fmt.Printf("Ditemukan (Linear search berdasarkan ID): ID=%d, Nama=%s, Berat=%.2f kg\n", warga[i].ID, warga[i].Nama, warga[i].BeratSampah)
			return
		}
	}
	fmt.Println("Data warga tidak ditemukan.")
}
func cariBinary(warga [100]Warga, count int) {
	if count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}

	var nama string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	copyData := sortByNama(warga, count)
	if found, ok := binarySearchByNama(copyData, count, nama); ok {
		fmt.Printf("Ditemukan (Binary search berdasarkan nama): ID=%d, Nama=%s, Berat=%.2f kg\n", found.ID, found.Nama, found.BeratSampah)
	} else {
		fmt.Println("Data warga tidak ditemukan.")
	}
}
func sortByID(data [100]Warga, count int) [100]Warga {
	result := data
	for i := 1; i < count; i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j].ID > key.ID {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}
func sortByNama(data [100]Warga, count int) [100]Warga {
	result := data
	for i := 1; i < count; i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j].Nama > key.Nama {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}
func binarySearchByID(data [100]Warga, count int, target int) (Warga, bool) {
	low, high := 0, count-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case data[mid].ID == target:
			return data[mid], true
		case data[mid].ID < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return Warga{}, false
}
func binarySearchByNama(data [100]Warga, count int, target string) (Warga, bool) {
	low, high := 0, count-1
	for low <= high {
		mid := (low + high) / 2
		midName := data[mid].Nama
		switch {
		case midName == target:
			return data[mid], true
		case midName < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return Warga{}, false
}
func urutkanSelection(warga [100]Warga, count int) {
	if count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}
	result := selectionSortByBerat(warga, count)
	fmt.Println("Hasil urutan (Pengurutan seleksi) berdasarkan berat sampah terbanyak:")
	for i := 0; i < count; i++ {
		fmt.Printf("ID=%d | Nama=%s | Berat=%.2f kg\n", result[i].ID, result[i].Nama, result[i].BeratSampah)
	}
}
func urutkanInsertion(warga [100]Warga, count int) {
	if count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}
	result := insertionSortByID(warga, count)
	fmt.Println("Hasil urutan (Pengurutan sisipan) berdasarkan ID:")
	for i := 0; i < count; i++ {
		fmt.Printf("ID=%d | Nama=%s | Berat=%.2f kg\n", result[i].ID, result[i].Nama, result[i].BeratSampah)
	}
}
func selectionSortByBerat(data [100]Warga, count int) [100]Warga {
	result := data
	for i := 0; i < count-1; i++ {
		maxIdx := i
		for j := i + 1; j < count; j++ {
			if result[j].BeratSampah > result[maxIdx].BeratSampah {
				maxIdx = j
			}
		}
		result[i], result[maxIdx] = result[maxIdx], result[i]
	}
	return result
}
func insertionSortByBerat(data [100]Warga, count int) [100]Warga {
	result := data
	for i := 1; i < count; i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j].BeratSampah < key.BeratSampah {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}
func insertionSortByID(data [100]Warga, count int) [100]Warga {
	result := data
	for i := 1; i < count; i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j].ID > key.ID {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}
func tampilkanWarga(warga [100]Warga, count int) {
	if count == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}
	fmt.Println("Daftar warga:")
	for i := 0; i < count; i++ {
		fmt.Printf("ID=%d | Nama=%s | Alamat=%s | Jenis=%s | Berat=%.2f kg\n", warga[i].ID, warga[i].Nama, warga[i].Alamat, warga[i].JenisSampah, warga[i].BeratSampah)
	}
}
func tampilkanStatistik(setoran [100]Setoran, count int) {
	if count == 0 {
		fmt.Println("Belum ada log setoran sampah.")
		return
	}

	var minggu int
	var tahun int

	fmt.Print("Masukkan minggu yang ingin dilihat: ")
	fmt.Scanln(&minggu)

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	var total float64
	var ditemukan bool

	fmt.Println("\nData Setoran Mingguan")

	for i := 0; i < count; i++ {
		if setoran[i].MingguKe == minggu &&
			setoran[i].Tahun == tahun {

			ditemukan = true
			total += setoran[i].Berat

			fmt.Printf(
				"ID=%d | Nama=%s | Tanggal=%s | Berat=%.2f kg\n",
				setoran[i].WargaID,
				setoran[i].NamaWarga,
				setoran[i].Tanggal,
				setoran[i].Berat,
			)
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada data pada minggu tersebut.")
		return
	}

	fmt.Printf("\nTotal akumulasi sampah minggu ke-%d tahun %d = %.2f kg\n",
		minggu, tahun, total)
}
