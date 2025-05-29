package main

import "fmt"

const NMAX int = 10000
const MaxLatPerHari int = 18

type Workout struct {
	JADWAL, JENISLOLAHRAGA, METODE       string
	LAMA, KALORIBURN, totalLatihan, hari int
	LATIHAN                              [MaxLatPerHari]string
}
type WorkHariIni [NMAX]Workout

type Latihan struct {
	Angkatbeban, Cardio, Yoga string
}
type DLatihan [18]Latihan

func main() {
	var pilih, jumlahRiwayat, printten, n, KaloriB, m, hygdiubah int
	var hari WorkHariIni
	var data DLatihan
	var SQsearch, BNsearch bool

	data[0] = Latihan{"BarbelCurl", "BearCrawls", "BoatPose"}
	data[1] = Latihan{"BenchPress", "BoxJumps", "BridgePose"}
	data[2] = Latihan{"CalfRaise", "Burpees", "CamelPose"}
	data[3] = Latihan{"DeadLift", "Cycling", "Cat-CowPose"}
	data[4] = Latihan{"Dips", "Elliptical", "ChildPose"}
	data[5] = Latihan{"FacePull", "HighKnees", "CobraPose"}
	data[6] = Latihan{"HammerCurl", "JumpRope", "CrowPose"}
	data[7] = Latihan{"LateralRaise", "JumpingJacks", "DownwardDog"}
	data[8] = Latihan{"LegCurl", "Kickboxing", "FishPose"}
	data[9] = Latihan{"LegPress", "MountainClimbers", "HappyBabyPose"}
	data[10] = Latihan{"Lunge", "Rowing", "LotusPose"}
	data[11] = Latihan{"PullUp", "Running", "PigeonPose"}
	data[12] = Latihan{"PushDown", "ShadowBoxing", "Plank"}
	data[13] = Latihan{"PushUp", "SkaterJumps", "SeatedTwist"}
	data[14] = Latihan{"RDL", "SpeedWalking", "ShoulderStand"}
	data[15] = Latihan{"Rows", "StairsClimbing", "TreePose"}
	data[16] = Latihan{"ShoulderPress", "Swimming", "TrianglePose"}
	data[17] = Latihan{"Squat", "Zumba", "WarriorPose"}

	n = 0
	m = 0
	printten = 0
	KaloriB = 0

	for pilih != 10 {
		// Menampilkan 10 Riwayat Terakhir
		if printten == 10+n {
			TampilkanTenWiwayat(hari, jumlahRiwayat, n)
			n += 10
		}

		// Menu
		fmt.Printf("--------------------------------------\n                 MENU\n--------------------------------------\n")
		if KaloriB == 7+m {
			totalKaloriBurn(&hari, jumlahRiwayat, m)
			fmt.Println("--------------------------------------")
			m += 7
		}
		fmt.Printf("1.Tambah Riwayat Workout\n2.Ubah Riwayat Workout\n3.Hapus Riwayat Workout\n4.Tampilkan Riwayat Workout\n5.Rekomendasi Workout\n6.Cari Latihan(SQ)\n7.Cari Latihan(BN)\n8.Selection Sort (Berdasarkan Kalori)\n9.Insertion Sort (Berdasarkan Durasi)\n10.Exit\n")
		// Menampilkan Total Kalori yang Terbakar per 7 hari
		fmt.Printf("Pilihan Anda: ")
		fmt.Scan(&pilih)
		fmt.Println()
		switch pilih {
		case 1:
			menambahjenisWorkout(&hari, &jumlahRiwayat)
			printten++
			KaloriB++
		case 2:
			Mengubah(&hari, hygdiubah, jumlahRiwayat)
		case 3:
			Menghapus(&hari, &jumlahRiwayat)
			printten--
			KaloriB--
		case 4:
			Tampilkan(hari, jumlahRiwayat)
		case 5:
			Rekomendasi(&hari, &jumlahRiwayat)
		case 6:
			SQsearch = SequentialSearch(&data, &hari, jumlahRiwayat)
			if SQsearch == true {
				fmt.Println("Latihan Ditemukan!")
			} else {
				fmt.Println("Latihan tidak ditemukan/Anda Keluar dari Pencarian")
			}
		case 7:
			BNsearch = BinarySearch(&data, &hari, jumlahRiwayat)
			if BNsearch == true {
				fmt.Println("Latihan ditemukan!")
			} else {
				fmt.Println("Latihan tidak ditemukan/Anda Keluar dari Pencarian")
			}
		case 8:
			selectionSort(&hari, jumlahRiwayat)
		case 9:
			insertionSort(&hari, jumlahRiwayat)
		}
		fmt.Println()
	}
}

// Menambah Jenis Workout
func menambahjenisWorkout(A *WorkHariIni, jumlahriwayat *int) {
	var jenisolahraga string
	fmt.Println()
	fmt.Println("Olahraga Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
	fmt.Print("Pilih (1 - 3) | (-) untuk Kembali : ")
	fmt.Scan(&jenisolahraga)
	fmt.Println()
	if jenisolahraga == "1" {
		A[*jumlahriwayat].JENISLOLAHRAGA = "AngkatBeban"
		menambahMetode(A, jumlahriwayat)
	} else if jenisolahraga == "2" {
		A[*jumlahriwayat].JENISLOLAHRAGA = "Cardio"
		MenambahCardio(A, jumlahriwayat)
	} else if jenisolahraga == "3" {
		A[*jumlahriwayat].JENISLOLAHRAGA = "Yoga"
		MenambahYoga(A, jumlahriwayat)
	} else if jenisolahraga == "-" {
		fmt.Println()
		A[*jumlahriwayat].JENISLOLAHRAGA = " "
	}
	fmt.Println()
}

// Menambah Metode
func menambahMetode(A *WorkHariIni, jumlahriwayat *int) {
	var MetohdLatihan string
	fmt.Println()
	fmt.Println("Latihan Angkat Beban Apa?\n1. Push\n2. Pull\n3. Legs")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali : ")
	fmt.Scan(&MetohdLatihan)
	if MetohdLatihan == "1" {
		A[*jumlahriwayat].METODE = "Push"
		menambahAngkatBeban(A, jumlahriwayat)
	} else if MetohdLatihan == "2" {
		A[*jumlahriwayat].METODE = "Pull"
		menambahAngkatBeban(A, jumlahriwayat)
	} else if MetohdLatihan == "3" {
		A[*jumlahriwayat].METODE = "Legs"
		menambahAngkatBeban(A, jumlahriwayat)
	} else if MetohdLatihan == "-" {
		A[*jumlahriwayat].METODE = " "
		menambahjenisWorkout(A, jumlahriwayat)
	}
	fmt.Println()
}

// Menambah Latihan Angkat Beban
func menambahAngkatBeban(A *WorkHariIni, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	if A[*jumlahriwayat].JENISLOLAHRAGA == "AngkatBeban" {
		if A[*jumlahriwayat].METODE == "Push" {
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						menambahMetode(A, jumlahriwayat)
					}
					if latihanKe >= 1 {
						menambahAngkatBeban(A, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 8
								A[*jumlahriwayat].LATIHAN[latihanKe] = "BenchPress"
							case "2":
								kaloriburn += 6
								A[*jumlahriwayat].LATIHAN[latihanKe] = "ShoulderPress"
							case "3":
								kaloriburn += 5
								A[*jumlahriwayat].LATIHAN[latihanKe] = "LateralRaise"
							case "4":
								kaloriburn += 6
								A[*jumlahriwayat].LATIHAN[latihanKe] = "Dips"
							case "5":
								kaloriburn += 4
								A[*jumlahriwayat].LATIHAN[latihanKe] = "PushUp"
							case "6":
								kaloriburn += 5
								A[*jumlahriwayat].LATIHAN[latihanKe] = "PushDown"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*jumlahriwayat].LAMA = totalWaktu
						A[*jumlahriwayat].KALORIBURN = kaloriburn
						A[*jumlahriwayat].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {
				menambahJadwal(A, jumlahriwayat)
			}
		} else if A[*jumlahriwayat].METODE == "Pull" {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						menambahMetode(A, jumlahriwayat)
					}
					if latihanKe >= 1 {
						menambahAngkatBeban(A, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 8
								A[*jumlahriwayat].LATIHAN[latihanKe] = "DeadLift"
							case "2":
								kaloriburn += 4
								A[*jumlahriwayat].LATIHAN[latihanKe] = "PullUp"
							case "3":
								kaloriburn += 6
								A[*jumlahriwayat].LATIHAN[latihanKe] = "Rows"
							case "4":
								kaloriburn += 5
								A[*jumlahriwayat].LATIHAN[latihanKe] = "BarbelCurl"
							case "5":
								kaloriburn += 4
								A[*jumlahriwayat].LATIHAN[latihanKe] = "HammerCurl"
							case "6":
								kaloriburn += 6
								A[*jumlahriwayat].LATIHAN[latihanKe] = "FacePull"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*jumlahriwayat].LAMA = totalWaktu
						A[*jumlahriwayat].KALORIBURN = kaloriburn
						A[*jumlahriwayat].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {
				menambahJadwal(A, jumlahriwayat)
			}
		} else {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						menambahMetode(A, jumlahriwayat)
					}
					if latihanKe >= 1 {
						menambahAngkatBeban(A, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 10
								A[*jumlahriwayat].LATIHAN[latihanKe] = "Squat"
							case "2":
								kaloriburn += 9
								A[*jumlahriwayat].LATIHAN[latihanKe] = "LegPress"
							case "3":
								kaloriburn += 11
								A[*jumlahriwayat].LATIHAN[latihanKe] = "Lunge"
							case "4":
								kaloriburn += 8
								A[*jumlahriwayat].LATIHAN[latihanKe] = "LegCurl"
							case "5":
								kaloriburn += 6
								A[*jumlahriwayat].LATIHAN[latihanKe] = "CalfRaise"
							case "6":
								kaloriburn += 10
								A[*jumlahriwayat].LATIHAN[latihanKe] = "RDL"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*jumlahriwayat].LAMA = totalWaktu
						A[*jumlahriwayat].KALORIBURN = kaloriburn
						A[*jumlahriwayat].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {
				menambahJadwal(A, jumlahriwayat)
			}
		}
	}
}

// Menambah Latihan Cardio
func MenambahCardio(A *WorkHariIni, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	for latihanKe < MaxLatPerHari && latihan != "." {
		i = 0
		fmt.Println("Pilih Latihan Cardio Berikut:\n1. Running\n2. Burpees\n3. Cycling\n4. JumpRope\n5. StairsClimbing\n6. Swimming\n7. Elliptical\n8. HighKnees\n9. MountainClimbers\n10. Rowing\n11. JumpingJacks\n12. SkaterJumps\n13. ShadowBoxing\n14. Kickboxing\n15. Zumba\n16. BoxJumps\n17. SpeedWalking\n18. BearCrawls")
		fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
		fmt.Print("Pililah (1 - 18) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
		fmt.Scan(&latihan2)
		if latihan2 == "-" {
			if latihanKe < 1 {
				latihan = "."
				menambahjenisWorkout(A, jumlahriwayat)
			}
			if latihanKe >= 1 {
				MenambahCardio(A, jumlahriwayat)
				latihanKe--
				latihan = "."
			}
		} else {
			if latihan2 != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan2 {
					case "1":
						kaloriburn += 10
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Running"
					case "2":
						kaloriburn += 12
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Burpees"
					case "3":
						kaloriburn += 8
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Cycling"
					case "4":
						kaloriburn += 13
						A[*jumlahriwayat].LATIHAN[latihanKe] = "JumpRope"
					case "5":
						kaloriburn += 11
						A[*jumlahriwayat].LATIHAN[latihanKe] = "StairsClimbing"
					case "6":
						kaloriburn += 10
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Swimming"
					case "7":
						kaloriburn += 9
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Elliptical"
					case "8":
						kaloriburn += 11
						A[*jumlahriwayat].LATIHAN[latihanKe] = "HighKnees"
					case "9":
						kaloriburn += 12
						A[*jumlahriwayat].LATIHAN[latihanKe] = "MountainClimbers"
					case "10":
						kaloriburn += 9
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Rowing"
					case "11":
						kaloriburn += 8
						A[*jumlahriwayat].LATIHAN[latihanKe] = "JumpingJacks"
					case "12":
						kaloriburn += 10
						A[*jumlahriwayat].LATIHAN[latihanKe] = "SkaterJumps"
					case "13":
						kaloriburn += 7
						A[*jumlahriwayat].LATIHAN[latihanKe] = "ShadowBoxing"
					case "14":
						kaloriburn += 11
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Kickboxing"
					case "15":
						kaloriburn += 9
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Zumba"
					case "16":
						kaloriburn += 12
						A[*jumlahriwayat].LATIHAN[latihanKe] = "BoxJumps"
					case "17":
						kaloriburn += 6
						A[*jumlahriwayat].LATIHAN[latihanKe] = "SpeedWalking"
					case "18":
						kaloriburn += 10
						A[*jumlahriwayat].LATIHAN[latihanKe] = "BearCrawls"
					}
				}
				totalWaktu += lama
				latihanKe++
			} else {
				A[*jumlahriwayat].LAMA = totalWaktu
				A[*jumlahriwayat].KALORIBURN = kaloriburn
				A[*jumlahriwayat].totalLatihan = latihanKe
				latihan = "."
			}
		}

	}
	if latihan2 != "-" {
		menambahJadwal(A, jumlahriwayat)
	}
}

// Menambah Latihan Yoga
func MenambahYoga(A *WorkHariIni, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	for latihanKe < MaxLatPerHari && latihan != "." {
		i = 0
		fmt.Println("Pilih Latihan Yoga Berikut:\n1. BridgePose\n2. Cat-CowPose\n3. ChildPose\n4. CobraPose\n5. DownwardDog\n6. Plank\n7. SeatedTwist\n8. TrianglePose\n9. TreePose\n10. WarriorPose\n11. LotusPose\n12. BoatPose\n13. HappyBabyPose\n14. PigeonPose\n15. FishPose\n16. CamelPose\n17. CrowPose\n18. ShoulderStand")
		fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
		fmt.Print("Pililah (1 - 18) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
		fmt.Scan(&latihan2)
		if latihan2 == "-" {
			if latihanKe < 1 {
				menambahjenisWorkout(A, jumlahriwayat)
			} else {
				latihanKe--
				totalWaktu -= lama
				MenambahYoga(A, jumlahriwayat)
			}
		} else {
			if latihan2 != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan2 {
					case "1":
						kaloriburn += 4
						A[*jumlahriwayat].LATIHAN[latihanKe] = "BridgePose"
					case "2":
						kaloriburn += 3
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Cat-CowPose"
					case "3":
						kaloriburn += 2
						A[*jumlahriwayat].LATIHAN[latihanKe] = "ChildPose"
					case "4":
						kaloriburn += 3
						A[*jumlahriwayat].LATIHAN[latihanKe] = "CobraPose"
					case "5":
						kaloriburn += 4
						A[*jumlahriwayat].LATIHAN[latihanKe] = "DownwardDog"
					case "6":
						kaloriburn += 5
						A[*jumlahriwayat].LATIHAN[latihanKe] = "Plank"
					case "7":
						kaloriburn += 3
						A[*jumlahriwayat].LATIHAN[latihanKe] = "SeatedTwist"
					case "8":
						kaloriburn += 4
						A[*jumlahriwayat].LATIHAN[latihanKe] = "TrianglePose"
					case "9":
						kaloriburn += 3
						A[*jumlahriwayat].LATIHAN[latihanKe] = "TreePose"
					case "10":
						kaloriburn += 5
						A[*jumlahriwayat].LATIHAN[latihanKe] = "WarriorPose"
					case "11":
						kaloriburn += 2
						A[*jumlahriwayat].LATIHAN[latihanKe] = "LotusPose"
					case "12":
						kaloriburn += 5
						A[*jumlahriwayat].LATIHAN[latihanKe] = "BoatPose"
					case "13":
						kaloriburn += 3
						A[*jumlahriwayat].LATIHAN[latihanKe] = "HappyBabyPose"
					case "14":
						kaloriburn += 3
						A[*jumlahriwayat].LATIHAN[latihanKe] = "PigeonPose"
					case "15":
						kaloriburn += 2
						A[*jumlahriwayat].LATIHAN[latihanKe] = "FishPose"
					case "16":
						kaloriburn += 4
						A[*jumlahriwayat].LATIHAN[latihanKe] = "CamelPose"
					case "17":
						kaloriburn += 6
						A[*jumlahriwayat].LATIHAN[latihanKe] = "CrowPose"
					case "18":
						kaloriburn += 4
						A[*jumlahriwayat].LATIHAN[latihanKe] = "ShoulderStand"
					}
				}
				totalWaktu += lama
				latihanKe++
			} else {
				A[*jumlahriwayat].LAMA = totalWaktu
				A[*jumlahriwayat].KALORIBURN = kaloriburn
				A[*jumlahriwayat].totalLatihan = latihanKe
				latihan = "."
			}
		}
	}
	if latihan2 != "-" {
		menambahJadwal(A, jumlahriwayat)
	}
}

// Menambah Jadwal
func menambahJadwal(A *WorkHariIni, jumlahriwayat *int) {
	var jadwal string
	fmt.Print("Jadwal (DD-MM-YYYY): ")
	fmt.Scan(&jadwal)
	A[*jumlahriwayat].hari = *jumlahriwayat + 1
	A[*jumlahriwayat].JADWAL = jadwal
	*jumlahriwayat++
	fmt.Println("Riwayat Berhasil Ditambahkan")
	fmt.Println()
}

// Mengubah Jenis Workout
func mengubahjenisWorkout(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var jenisolahraga string
	fmt.Println()
	fmt.Println("Olahraga Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
	fmt.Print("Pilih (1 - 3) | (-) untuk Kembali : ")
	fmt.Scan(&jenisolahraga)
	fmt.Println()
	if jenisolahraga == "1" {
		A[*hariyangdiUbah].JENISLOLAHRAGA = "AngkatBeban"
		mengubahMetode1(A, hariyangdiUbah, jumlahriwayat)
	} else if jenisolahraga == "2" {
		A[*hariyangdiUbah].JENISLOLAHRAGA = "Cardio"
		mengubahCardio1(A, hariyangdiUbah, jumlahriwayat)
	} else if jenisolahraga == "3" {
		A[*hariyangdiUbah].JENISLOLAHRAGA = "Yoga"
		mengubahYoga1(A, hariyangdiUbah, jumlahriwayat)
	} else if jenisolahraga == "-" {
		Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
	}
	fmt.Println()
}

// Mengubah Metode Opsi 1
func mengubahMetode1(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var MetohdLatihan string
	fmt.Println()
	fmt.Println("Latihan Angkat Beban Apa?\n1. Push\n2. Pull\n3. Legs")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali : ")
	fmt.Scan(&MetohdLatihan)
	if MetohdLatihan == "1" {
		A[*hariyangdiUbah].METODE = "Push"
		mengubahAngkatBeban1(A, hariyangdiUbah, jumlahriwayat)
	} else if MetohdLatihan == "2" {
		A[*hariyangdiUbah].METODE = "Pull"
		mengubahAngkatBeban1(A, hariyangdiUbah, jumlahriwayat)
	} else if MetohdLatihan == "3" {
		A[*hariyangdiUbah].METODE = "Legs"
		mengubahAngkatBeban1(A, hariyangdiUbah, jumlahriwayat)
	} else if MetohdLatihan == "-" {
		A[*hariyangdiUbah].METODE = " "
		mengubahjenisWorkout(A, hariyangdiUbah, jumlahriwayat)
	}
	fmt.Println()
}

// Mengubah Metode Opsi 2
func mengubahMetode2(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var MetohdLatihan string
	fmt.Println()
	fmt.Println("Latihan Angkat Beban Apa?\n1. Push\n2. Pull\n3. Legs")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali : ")
	fmt.Scan(&MetohdLatihan)
	if MetohdLatihan == "1" {
		A[*hariyangdiUbah].METODE = "Push"
		mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
	} else if MetohdLatihan == "2" {
		A[*hariyangdiUbah].METODE = "Pull"
		mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
	} else if MetohdLatihan == "3" {
		A[*hariyangdiUbah].METODE = "Legs"
		mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
	} else if MetohdLatihan == "-" {
		A[*hariyangdiUbah].METODE = " "
		Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
	}
	fmt.Println()
}

// Mengubah Latihan Angkat Beban Opsi 1 (Mengubah Jenis Workout)
func mengubahAngkatBeban1(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	if A[*hariyangdiUbah].JENISLOLAHRAGA == "AngkatBeban" {
		if A[*hariyangdiUbah].METODE == "Push" {
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						mengubahMetode1(A, hariyangdiUbah, jumlahriwayat)
					}
					if latihanKe >= 1 {
						mengubahAngkatBeban1(A, hariyangdiUbah, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 8
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "BenchPress"
							case "2":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShoulderPress"
							case "3":
								kaloriburn += 5
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "LateralRaise"
							case "4":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Dips"
							case "5":
								kaloriburn += 4
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "PushUp"
							case "6":
								kaloriburn += 5
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "PushDown"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*hariyangdiUbah].LAMA = totalWaktu
						A[*hariyangdiUbah].KALORIBURN = kaloriburn
						A[*hariyangdiUbah].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {

			}
		} else if A[*hariyangdiUbah].METODE == "Pull" {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						mengubahMetode1(A, hariyangdiUbah, jumlahriwayat)
					}
					if latihanKe >= 1 {
						mengubahAngkatBeban1(A, hariyangdiUbah, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 8
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "DeadLift"
							case "2":
								kaloriburn += 4
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "PullUp"
							case "3":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Rows"
							case "4":
								kaloriburn += 5
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "BarbelCurl"
							case "5":
								kaloriburn += 4
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "HammerCurl"
							case "6":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "FacePull"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*hariyangdiUbah].LAMA = totalWaktu
						A[*hariyangdiUbah].KALORIBURN = kaloriburn
						A[*hariyangdiUbah].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {

			}
		} else {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						mengubahMetode1(A, hariyangdiUbah, jumlahriwayat)
					}
					if latihanKe >= 1 {
						mengubahAngkatBeban1(A, hariyangdiUbah, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 10
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Squat"
							case "2":
								kaloriburn += 9
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "LegPress"
							case "3":
								kaloriburn += 11
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Lunge"
							case "4":
								kaloriburn += 8
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "LegCurl"
							case "5":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "CalfRaise"
							case "6":
								kaloriburn += 10
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "RDL"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*hariyangdiUbah].LAMA = totalWaktu
						A[*hariyangdiUbah].KALORIBURN = kaloriburn
						A[*hariyangdiUbah].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {

			}
		}
	}
}

// Mengubah Latihan Angkat Beban Opsi 2 (Mengubah Metode)
func mengubahAngkatBeban2(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	if A[*hariyangdiUbah].JENISLOLAHRAGA == "AngkatBeban" {
		if A[*hariyangdiUbah].METODE == "Push" {
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						mengubahMetode2(A, hariyangdiUbah, jumlahriwayat)
					}
					if latihanKe >= 1 {
						mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 8
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "BenchPress"
							case "2":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShoulderPress"
							case "3":
								kaloriburn += 5
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "LateralRaise"
							case "4":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Dips"
							case "5":
								kaloriburn += 4
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "PushUp"
							case "6":
								kaloriburn += 5
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "PushDown"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*hariyangdiUbah].LAMA = totalWaktu
						A[*hariyangdiUbah].KALORIBURN = kaloriburn
						A[*hariyangdiUbah].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {

			}
		} else if A[*hariyangdiUbah].METODE == "Pull" {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						mengubahMetode2(A, hariyangdiUbah, jumlahriwayat)
					}
					if latihanKe >= 1 {
						mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 8
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "DeadLift"
							case "2":
								kaloriburn += 4
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "PullUp"
							case "3":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Rows"
							case "4":
								kaloriburn += 5
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "BarbelCurl"
							case "5":
								kaloriburn += 4
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "HammerCurl"
							case "6":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "FacePull"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*hariyangdiUbah].LAMA = totalWaktu
						A[*hariyangdiUbah].KALORIBURN = kaloriburn
						A[*hariyangdiUbah].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {

			}
		} else {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
				fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
				fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
				fmt.Scan(&latihan2)
				if latihan2 == "-" {
					if latihanKe < 1 {
						latihan = "."
						mengubahMetode2(A, hariyangdiUbah, jumlahriwayat)
					}
					if latihanKe >= 1 {
						mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
						latihanKe--
						latihan = "."
					}
				} else {
					if latihan2 != "." {
						fmt.Print("Berapa Lama : ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan2 {
							case "1":
								kaloriburn += 10
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Squat"
							case "2":
								kaloriburn += 9
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "LegPress"
							case "3":
								kaloriburn += 11
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "Lunge"
							case "4":
								kaloriburn += 8
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "LegCurl"
							case "5":
								kaloriburn += 6
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "CalfRaise"
							case "6":
								kaloriburn += 10
								A[*hariyangdiUbah].LATIHAN[latihanKe] = "RDL"
							}
						}
						totalWaktu += lama
						latihanKe++
					} else {
						A[*hariyangdiUbah].LAMA = totalWaktu
						A[*hariyangdiUbah].KALORIBURN = kaloriburn
						A[*hariyangdiUbah].totalLatihan = latihanKe
						latihan = "."
					}
				}
			}
			if latihan2 != "-" {

			}
		}
	}
}

// Mengubah Latihan Angkat Beban Opsi 3 (Mengubah Latihan)
func mengubahAngkatBeban3(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	if A[*hariyangdiUbah].METODE == "Push" {
		for latihanKe < MaxLatPerHari && latihan != "." {
			i = 0
			fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
			fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
			fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
			fmt.Scan(&latihan2)
			if latihan2 == "-" {
				if latihanKe < 1 {
					latihan = "."
					Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
				}
				if latihanKe >= 1 {
					mengubahAngkatBeban3(A, hariyangdiUbah, jumlahriwayat)
					latihanKe--
					latihan = "."
				}
			} else {
				if latihan2 != "." {
					fmt.Print("Berapa Lama : ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan2 {
						case "1":
							kaloriburn += 8
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "BenchPress"
						case "2":
							kaloriburn += 6
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShoulderPress"
						case "3":
							kaloriburn += 5
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "LateralRaise"
						case "4":
							kaloriburn += 6
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "Dips"
						case "5":
							kaloriburn += 4
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "PushUp"
						case "6":
							kaloriburn += 5
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "PushDown"
						}
					}
					totalWaktu += lama
					latihanKe++
				} else {
					A[*hariyangdiUbah].LAMA = totalWaktu
					A[*hariyangdiUbah].KALORIBURN = kaloriburn
					A[*hariyangdiUbah].totalLatihan = latihanKe
					latihan = "."
				}
			}
		}
		if latihan2 != "-" {

		}
	} else if A[*hariyangdiUbah].METODE == "Pull" {
		latihanKe = 0
		for latihanKe < MaxLatPerHari && latihan != "." {
			i = 0
			fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
			fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
			fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
			fmt.Scan(&latihan2)
			if latihan2 == "-" {
				if latihanKe < 1 {
					latihan = "."
					Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
				}
				if latihanKe >= 1 {
					mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
					latihanKe--
					latihan = "."
				}
			} else {
				if latihan2 != "." {
					fmt.Print("Berapa Lama : ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan2 {
						case "1":
							kaloriburn += 8
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "DeadLift"
						case "2":
							kaloriburn += 4
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "PullUp"
						case "3":
							kaloriburn += 6
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "Rows"
						case "4":
							kaloriburn += 5
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "BarbelCurl"
						case "5":
							kaloriburn += 4
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "HammerCurl"
						case "6":
							kaloriburn += 6
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "FacePull"
						}
					}
					totalWaktu += lama
					latihanKe++
				} else {
					A[*hariyangdiUbah].LAMA = totalWaktu
					A[*hariyangdiUbah].KALORIBURN = kaloriburn
					A[*hariyangdiUbah].totalLatihan = latihanKe
					latihan = "."
				}
			}
		}
		if latihan2 != "-" {

		}
	} else {
		latihanKe = 0
		for latihanKe < MaxLatPerHari && latihan != "." {
			i = 0
			fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
			fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
			fmt.Print("Pililah (1 - 6) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
			fmt.Scan(&latihan2)
			if latihan2 == "-" {
				if latihanKe < 1 {
					latihan = "."
					Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
				}
				if latihanKe >= 1 {
					mengubahAngkatBeban2(A, hariyangdiUbah, jumlahriwayat)
					latihanKe--
					latihan = "."
				}
			} else {
				if latihan2 != "." {
					fmt.Print("Berapa Lama : ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan2 {
						case "1":
							kaloriburn += 10
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "Squat"
						case "2":
							kaloriburn += 9
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "LegPress"
						case "3":
							kaloriburn += 11
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "Lunge"
						case "4":
							kaloriburn += 8
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "LegCurl"
						case "5":
							kaloriburn += 6
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "CalfRaise"
						case "6":
							kaloriburn += 10
							A[*hariyangdiUbah].LATIHAN[latihanKe] = "RDL"
						}
					}
					totalWaktu += lama
					latihanKe++
				} else {
					A[*hariyangdiUbah].LAMA = totalWaktu
					A[*hariyangdiUbah].KALORIBURN = kaloriburn
					A[*hariyangdiUbah].totalLatihan = latihanKe
					latihan = "."
				}
			}
		}
		if latihan2 != "-" {

		}
	}
}

// Mengubah Latihan Cardio Opsi 1 (Mengubah Jenis Workout)
func mengubahCardio1(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	for latihanKe < MaxLatPerHari && latihan != "." {
		i = 0
		fmt.Println("Pilih Latihan Cardio Berikut:\n1. Running\n2. Burpees\n3. Cycling\n4. JumpRope\n5. StairsClimbing\n6. Swimming\n7. Elliptical\n8. HighKnees\n9. MountainClimbers\n10. Rowing\n11. JumpingJacks\n12. SkaterJumps\n13. ShadowBoxing\n14. Kickboxing\n15. Zumba\n16. BoxJumps\n17. SpeedWalking\n18. BearCrawls")
		fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
		fmt.Print("Pililah (1 - 18) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
		fmt.Scan(&latihan2)
		if latihan2 == "-" {
			if latihanKe < 1 {
				latihan = "."
				mengubahjenisWorkout(A, hariyangdiUbah, jumlahriwayat)
			}
			if latihanKe >= 1 {
				mengubahCardio1(A, hariyangdiUbah, jumlahriwayat)
				latihanKe--
				latihan = "."
			}
		} else {
			if latihan2 != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan2 {
					case "1":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Running"
					case "2":
						kaloriburn += 12
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Burpees"
					case "3":
						kaloriburn += 8
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Cycling"
					case "4":
						kaloriburn += 13
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "JumpRope"
					case "5":
						kaloriburn += 11
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "StairsClimbing"
					case "6":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Swimming"
					case "7":
						kaloriburn += 9
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Elliptical"
					case "8":
						kaloriburn += 11
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "HighKnees"
					case "9":
						kaloriburn += 12
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "MountainClimbers"
					case "10":
						kaloriburn += 9
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Rowing"
					case "11":
						kaloriburn += 8
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "JumpingJacks"
					case "12":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "SkaterJumps"
					case "13":
						kaloriburn += 7
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShadowBoxing"
					case "14":
						kaloriburn += 11
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Kickboxing"
					case "15":
						kaloriburn += 9
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Zumba"
					case "16":
						kaloriburn += 12
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BoxJumps"
					case "17":
						kaloriburn += 6
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "SpeedWalking"
					case "18":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BearCrawls"
					}
				}
				totalWaktu += lama
				latihanKe++
			} else {
				A[*hariyangdiUbah].LAMA = totalWaktu
				A[*hariyangdiUbah].KALORIBURN = kaloriburn
				A[*hariyangdiUbah].totalLatihan = latihanKe
				latihan = "."
			}
		}

	}
	if latihan2 != "-" {

	}
}

// Mengubah Latihan Cardio Opsi 1 (Mengubah Latihan)
func mengubahCardio2(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	for latihanKe < MaxLatPerHari && latihan != "." {
		i = 0
		fmt.Println("Pilih Latihan Cardio Berikut:\n1. Running\n2. Burpees\n3. Cycling\n4. JumpRope\n5. StairsClimbing\n6. Swimming\n7. Elliptical\n8. HighKnees\n9. MountainClimbers\n10. Rowing\n11. JumpingJacks\n12. SkaterJumps\n13. ShadowBoxing\n14. Kickboxing\n15. Zumba\n16. BoxJumps\n17. SpeedWalking\n18. BearCrawls")
		fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
		fmt.Print("Pililah (1 - 18) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
		fmt.Scan(&latihan2)
		if latihan2 == "-" {
			if latihanKe < 1 {
				latihan = "."
				Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
			}
			if latihanKe >= 1 {
				mengubahCardio2(A, hariyangdiUbah, jumlahriwayat)
				latihanKe--
				latihan = "."
			}
		} else {
			if latihan2 != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan2 {
					case "1":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Running"
					case "2":
						kaloriburn += 12
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Burpees"
					case "3":
						kaloriburn += 8
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Cycling"
					case "4":
						kaloriburn += 13
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "JumpRope"
					case "5":
						kaloriburn += 11
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "StairsClimbing"
					case "6":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Swimming"
					case "7":
						kaloriburn += 9
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Elliptical"
					case "8":
						kaloriburn += 11
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "HighKnees"
					case "9":
						kaloriburn += 12
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "MountainClimbers"
					case "10":
						kaloriburn += 9
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Rowing"
					case "11":
						kaloriburn += 8
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "JumpingJacks"
					case "12":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "SkaterJumps"
					case "13":
						kaloriburn += 7
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShadowBoxing"
					case "14":
						kaloriburn += 11
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Kickboxing"
					case "15":
						kaloriburn += 9
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Zumba"
					case "16":
						kaloriburn += 12
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BoxJumps"
					case "17":
						kaloriburn += 6
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "SpeedWalking"
					case "18":
						kaloriburn += 10
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BearCrawls"
					}
				}
				totalWaktu += lama
				latihanKe++
			} else {
				A[*hariyangdiUbah].LAMA = totalWaktu
				A[*hariyangdiUbah].KALORIBURN = kaloriburn
				A[*hariyangdiUbah].totalLatihan = latihanKe
				latihan = "."
			}
		}

	}
	if latihan2 != "-" {

	}
}

// Mengubah Latihan Yoga Opsi 1 (Mengubah Jenis Workout)
func mengubahYoga1(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	for latihanKe < MaxLatPerHari && latihan != "." {
		i = 0
		fmt.Println("Pilih Latihan Yoga Berikut:\n1. BridgePose\n2. Cat-CowPose\n3. ChildPose\n4. CobraPose\n5. DownwardDog\n6. Plank\n7. SeatedTwist\n8. TrianglePose\n9. TreePose\n10. WarriorPose\n11. LotusPose\n12. BoatPose\n13. HappyBabyPose\n14. PigeonPose\n15. FishPose\n16. CamelPose\n17. CrowPose\n18. ShoulderStand")
		fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
		fmt.Print("Pililah (1 - 18) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
		fmt.Scan(&latihan2)
		if latihan2 == "-" {
			if latihanKe < 1 {
				mengubahjenisWorkout(A, hariyangdiUbah, jumlahriwayat)
			} else {
				latihanKe--
				totalWaktu -= lama
				mengubahYoga1(A, hariyangdiUbah, jumlahriwayat)
			}
		} else {
			if latihan2 != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan2 {
					case "1":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BridgePose"
					case "2":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Cat-CowPose"
					case "3":
						kaloriburn += 2
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "ChildPose"
					case "4":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "CobraPose"
					case "5":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "DownwardDog"
					case "6":
						kaloriburn += 5
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Plank"
					case "7":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "SeatedTwist"
					case "8":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "TrianglePose"
					case "9":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "TreePose"
					case "10":
						kaloriburn += 5
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "WarriorPose"
					case "11":
						kaloriburn += 2
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "LotusPose"
					case "12":
						kaloriburn += 5
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BoatPose"
					case "13":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "HappyBabyPose"
					case "14":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "PigeonPose"
					case "15":
						kaloriburn += 2
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "FishPose"
					case "16":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "CamelPose"
					case "17":
						kaloriburn += 6
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "CrowPose"
					case "18":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShoulderStand"
					}
				}
				totalWaktu += lama
				latihanKe++
			} else {
				A[*hariyangdiUbah].LAMA = totalWaktu
				A[*hariyangdiUbah].KALORIBURN = kaloriburn
				A[*hariyangdiUbah].totalLatihan = latihanKe
				latihan = "."
			}
		}
	}
	if latihan2 != "-" {

	}
}

// Mengubah Latihan Yoga Opsi 2 (Mengubah Latihan)
func mengubahYoga2(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var totalWaktu, latihanKe, lama, kaloriburn, i int
	var latihan, latihan2 string
	kaloriburn = 0
	for latihanKe < MaxLatPerHari && latihan != "." {
		i = 0
		fmt.Println("Pilih Latihan Yoga Berikut:\n1. BridgePose\n2. Cat-CowPose\n3. ChildPose\n4. CobraPose\n5. DownwardDog\n6. Plank\n7. SeatedTwist\n8. TrianglePose\n9. TreePose\n10. WarriorPose\n11. LotusPose\n12. BoatPose\n13. HappyBabyPose\n14. PigeonPose\n15. FishPose\n16. CamelPose\n17. CrowPose\n18. ShoulderStand")
		fmt.Printf("[INFO] Input Ke-%d\n", latihanKe+1)
		fmt.Print("Pililah (1 - 18) | (.) untuk selesai & (-) untuk Input ulang(Reset) / (-) pada input pertama untuk kembali : ")
		fmt.Scan(&latihan2)
		if latihan2 == "-" {
			if latihanKe < 1 {
				Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
			} else {
				latihanKe--
				totalWaktu -= lama
				mengubahYoga2(A, hariyangdiUbah, jumlahriwayat)
			}
		} else {
			if latihan2 != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan2 {
					case "1":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BridgePose"
					case "2":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Cat-CowPose"
					case "3":
						kaloriburn += 2
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "ChildPose"
					case "4":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "CobraPose"
					case "5":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "DownwardDog"
					case "6":
						kaloriburn += 5
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "Plank"
					case "7":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "SeatedTwist"
					case "8":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "TrianglePose"
					case "9":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "TreePose"
					case "10":
						kaloriburn += 5
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "WarriorPose"
					case "11":
						kaloriburn += 2
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "LotusPose"
					case "12":
						kaloriburn += 5
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "BoatPose"
					case "13":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "HappyBabyPose"
					case "14":
						kaloriburn += 3
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "PigeonPose"
					case "15":
						kaloriburn += 2
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "FishPose"
					case "16":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "CamelPose"
					case "17":
						kaloriburn += 6
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "CrowPose"
					case "18":
						kaloriburn += 4
						A[*hariyangdiUbah].LATIHAN[latihanKe] = "ShoulderStand"
					}
				}
				totalWaktu += lama
				latihanKe++
			} else {
				A[*hariyangdiUbah].LAMA = totalWaktu
				A[*hariyangdiUbah].KALORIBURN = kaloriburn
				A[*hariyangdiUbah].totalLatihan = latihanKe
				latihan = "."
			}
		}
	}
	if latihan2 != "-" {

	}
}

// Mengubah Jadwal
func mengubahJadwal(A *WorkHariIni, hariyangdiUbah, jumlahriwayat *int) {
	var jadwal string
	fmt.Print("Jadwal (DD-MM-YYYY) | (-) untuk kembali: ")
	fmt.Scan(&jadwal)
	if jadwal == "-" {
		Mengubah(A, *hariyangdiUbah, *jumlahriwayat)
	} else {
		A[*hariyangdiUbah].JADWAL = jadwal
		fmt.Println("Riwayat Berhasil Diubah")
		fmt.Println()
	}

}

// Prosedure Mengubah
func Mengubah(A *WorkHariIni, hariyangdiUbah, jumlahriwayat int) {
	var ubahapa string
	Tampilkan(*A, jumlahriwayat)
	fmt.Print("Hari yang Ingin Diubah | (-1) untuk kembali: ")
	fmt.Scan(&hariyangdiUbah)
	if hariyangdiUbah != -1 {
		hariyangdiUbah -= 1
		fmt.Println("Ubah Apa?\n1. Jenisolahraga\n2. Metode\n3. Jadwal\n4. Latihan")
		fmt.Print("Pilih: ")
		fmt.Scan(&ubahapa)
		fmt.Print("Pilih: ")

		//Ubah Jenis Olahraga
		if ubahapa == "1" {
			mengubahjenisWorkout(A, &hariyangdiUbah, &jumlahriwayat)
			//Mengubah Metode
		} else if ubahapa == "2" && A[hariyangdiUbah].JENISLOLAHRAGA == "AngkatBeban" {
			mengubahMetode2(A, &hariyangdiUbah, &jumlahriwayat)
			// Mengubah Jadwal
		} else if ubahapa == "3" {
			mengubahJadwal(A, &hariyangdiUbah, &jumlahriwayat)
			// Mengubah Latihan
		} else if ubahapa == "4" {
			if A[hariyangdiUbah].JENISLOLAHRAGA == "AngkatBeban" {
				mengubahAngkatBeban3(A, &hariyangdiUbah, &jumlahriwayat)
			} else if A[hariyangdiUbah].JENISLOLAHRAGA == "Cardio" {
				mengubahCardio2(A, &hariyangdiUbah, &jumlahriwayat)
			} else if A[hariyangdiUbah].JENISLOLAHRAGA == "Yoga" {
				mengubahYoga2(A, &hariyangdiUbah, &jumlahriwayat)
			}
		} else if ubahapa == "-" {
			fmt.Println()
		}
	} else {

	}
}

// Prosedure Menghapus
func Menghapus(A *WorkHariIni, jumlahriwayat *int) {
	var HyangdiHapus, i int
	fmt.Print("Hari yang Ingin Dihapus: ")
	fmt.Scan(&HyangdiHapus)
	HyangdiHapus -= 1
	for i = HyangdiHapus; i < *jumlahriwayat; i++ {
		A[i] = A[i+1]
	}
	*jumlahriwayat--
}

// Prosedure Menampilkan
func Tampilkan(A WorkHariIni, jumlahriwayat int) {
	var i, j int
	fmt.Println("Riwayat Workout Anda:")
	if jumlahriwayat <= 0 {
		fmt.Println("Anda Belum Memiliki Riwayat Workout")
	} else {
		for i = 0; i < jumlahriwayat; i++ {
			if A[i].JENISLOLAHRAGA == "AngkatBeban" {
				fmt.Printf("Hari Ke: %d\nJenis Olahraga: %s\nMetode Latihan: %s\nJadwal: %s\n", A[i].hari, A[i].JENISLOLAHRAGA, A[i].METODE, A[i].JADWAL)
				for j = 0; j < A[i].totalLatihan; j++ {
					fmt.Printf("-> Latihan Ke-%d: %s\n", j+1, A[i].LATIHAN[j])
				}
				fmt.Printf("Lama Workout: %d menit\nJumlah Kalori yang Terbakar: %d kkal\n", A[i].LAMA, A[i].KALORIBURN)
				fmt.Println()

			} else if A[i].JENISLOLAHRAGA == "Cardio" {
				fmt.Printf("Hari Ke: %d\nJenis Olahraga: %s\nJadwal: %s\n", A[i].hari, A[i].JENISLOLAHRAGA, A[i].JADWAL)
				for j = 0; j < A[i].totalLatihan; j++ {
					fmt.Printf("-> Latihan Ke-%d: %s\n", j+1, A[i].LATIHAN[j])
				}
				fmt.Printf("Lama Workout: %d menit\nJumlah Kalori yang Terbakar: %d kkal\n", A[i].LAMA, A[i].KALORIBURN)
				fmt.Println()

			} else {
				fmt.Printf("Hari Ke: %d\nJenis Olahraga: %s\nJadwal: %s\n", A[i].hari, A[i].JENISLOLAHRAGA, A[i].JADWAL)
				for j = 0; j < A[i].totalLatihan; j++ {
					fmt.Printf("-> Latihan Ke-%d: %s\n", j+1, A[i].LATIHAN[j])
				}
				fmt.Printf("Lama Workout: %d menit\nJumlah Kalori yang Terbakar: %d kkal\n", A[i].LAMA, A[i].KALORIBURN)
				fmt.Println()
			}
		}
	}
}

// Prosedure Menampilkan 10 Riwayat Terakhir
func TampilkanTenWiwayat(A WorkHariIni, jumlahriwayat int, n int) {
	var i, j int
	fmt.Println("Riwayat Workout Anda:")
	if jumlahriwayat <= 0 {
		fmt.Println("Anda Belum Memiliki Riwayat Workout")
	} else {
		for i = 0 + n; i < jumlahriwayat; i++ {
			if A[i].JENISLOLAHRAGA == "AngkatBeban" {
				fmt.Printf("Hari Ke: %d\nJenis Olahraga: %s\nMetode Latihan: %s\nJadwal: %s\n", A[i].hari, A[i].JENISLOLAHRAGA, A[i].METODE, A[i].JADWAL)
				for j = 0; j < A[i].totalLatihan; j++ {
					fmt.Printf("-> Latihan Ke-%d: %s\n", j+1, A[i].LATIHAN[j])
				}
				fmt.Printf("Lama Workout: %d menit\nJumlah Kalori yang Terbakar: %d kkal\n", A[i].LAMA, A[i].KALORIBURN)
				fmt.Println()

			} else if A[i].JENISLOLAHRAGA == "Cardio" {
				fmt.Printf("Hari Ke: %d\nJenis Olahraga: %s\nJadwal: %s\n", A[i].hari, A[i].JENISLOLAHRAGA, A[i].JADWAL)
				for j = 0; j < A[i].totalLatihan; j++ {
					fmt.Printf("-> Latihan Ke-%d: %s\n", j+1, A[i].LATIHAN[j])
				}
				fmt.Printf("Lama Workout: %d menit\nJumlah Kalori yang Terbakar: %d kkal\n", A[i].LAMA, A[i].KALORIBURN)
				fmt.Println()

			} else {
				fmt.Printf("Hari Ke: %d\nJenis Olahraga: %s\nJadwal: %s\n", A[i].hari, A[i].JENISLOLAHRAGA, A[i].JADWAL)
				for j = 0; j < A[i].totalLatihan; j++ {
					fmt.Printf("-> Latihan Ke-%d: %s\n", j+1, A[i].LATIHAN[j])
				}
				fmt.Printf("Lama Workout: %d menit\nJumlah Kalori yang Terbakar: %d kkal\n", A[i].LAMA, A[i].KALORIBURN)
				fmt.Println()
			}
		}
	}
}

// Prosedure Rekomendasi
func Rekomendasi(A *WorkHariIni, jumlahriwayat *int) {
	var workTerakhir Workout
	var want string

	if *jumlahriwayat == 0 {
		fmt.Println("Anda Belum Memiliki Riwayat Workout.")
	} else {
		workTerakhir = A[*jumlahriwayat-1]

		// Jika workout terakhir adalah Angkat Beban
		if workTerakhir.JENISLOLAHRAGA == "AngkatBeban" {
			fmt.Println("Mau latihan apa?\n1. Angkatbeban\n2. Cardio\n3. Yoga")
			fmt.Print("Pilih (1 -3) | (-) untuk kembali: ")
			fmt.Scan(&want)
			fmt.Println()
			if want != "-" {
				if want == "1" || want == "Angkatbeban" || want == "angkatbeban" {
					if workTerakhir.METODE == "Push" {
						fmt.Println("Rekomendasi: Coba metode Pull atau Legs.")
						fmt.Println("Pull: DeadLift, PullUp, Rows, BarbelCurl")
						fmt.Println("Legs: Squat, Lunge, RDL, LegPress")

					} else if workTerakhir.METODE == "Pull" {
						fmt.Println("Rekomendasi: Coba metode Push atau Legs.")
						fmt.Println("Push: BenchPress, ShoulderPress, PushUp, Dips")
						fmt.Println("Legs: LegCurl, CalfRaise, RDL")

					} else {
						fmt.Println("Rekomendasi: Coba metode Push atau Pull.")
						fmt.Println("Push: PushUp, PushDown, LateralRaise")
						fmt.Println("Pull: FacePull, Rows, HammerCurl")
					}

				} else if want == "2" || want == "Cardio" || want == "cardio" {
					fmt.Println("Rekomendasi Cardio:")
					fmt.Println("Burpees, JumpRope, MountainClimbers, Swimming, Kickboxing")

				} else {
					fmt.Println("Rekomendasi Yoga:")
					fmt.Println("Plank, BoatPose, TreePose, WarriorPose, CobraPose")
				}
			} else {

			}

			// Jika workout terakhir adalah Cardio
		} else if workTerakhir.JENISLOLAHRAGA == "Cardio" {
			fmt.Println("Mau latihan apa?\n1. Angkatbeban\n2. Cardio\n3. Yoga")
			fmt.Print("Pilih (1 -3) | (-) untuk kembali: ")
			fmt.Scan(&want)
			fmt.Println()
			if want != "-" {
				if want == "1" || want == "Angkatbeban" || want == "angkatbeban" {
					fmt.Println("Rekomendasi Angkatbeban:")
					fmt.Println("Push: BenchPress, PushUp")
					fmt.Println("Pull: Rows, DeadLift")
					fmt.Println("Legs: Squat, CalfRaise")

				} else if want == "2" || want == "Cardio" || want == "cardio" {
					fmt.Println("Rekomendasi variasi Cardio:")
					fmt.Println("Cycling, BoxJumps, Zumba, SkaterJumps, BearCrawls")

				} else {
					fmt.Println("Rekomendasi Yoga:")
					fmt.Println("SeatedTwist, CamelPose, CrowPose, ChildPose")
				}
			} else {

			}

			// Jika workout terakhir adalah Yoga
		} else if workTerakhir.JENISLOLAHRAGA == "Yoga" {
			fmt.Println("Mau latihan apa?\n1. Angkatbeban\n2. Cardio\n3. Yoga")
			fmt.Print("Pilih (1 -3) | (-) untuk kembali: ")
			fmt.Scan(&want)
			fmt.Println()
			if want != "-" {
				if want == "1" || want == "Angkatbeban" || want == "angkatbeban" {
					fmt.Println("Rekomendasi Angkatbeban:")
					fmt.Println("Push: ShoulderPress, Dips")
					fmt.Println("Pull: BarbelCurl, FacePull")
					fmt.Println("Legs: LegPress, RDL")

				} else if want == "2" || want == "Cardio" || want == "cardio" {
					fmt.Println("Rekomendasi Cardio:")
					fmt.Println("ShadowBoxing, Running, Elliptical, SpeedWalking")

				} else {
					fmt.Println("Rekomendasi Yoga:")
					fmt.Println("BridgePose, TrianglePose, FishPose, ShoulderStand")
				}
			} else {

			}

		}
	}
}

// Prosedure Menampilkan Total Kalori yang Terbakar per 7 hari
func totalKaloriBurn(A *WorkHariIni, jumlahRiwayat int, m int) {
	var i, totalkalori int
	for i = 0 + m; i < jumlahRiwayat; i++ {
		totalkalori += A[i].KALORIBURN
	}
	fmt.Printf("[INFO] Total Kalori yang Terbakar dalam 7 Hari: %d kkal\n", totalkalori)
}

// function SequentialSearch
func SequentialSearch(A *DLatihan, B *WorkHariIni, jumlahriwayat int) bool {
	var want, namaL string
	var i int
	fmt.Println("Mau Latihan Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali: ")
	fmt.Scan(&want)
	if want != "-" {
		fmt.Print("Nama Latihan: ")
		fmt.Scan(&namaL)
		fmt.Println()
		// Menampilkan Latihan yang tersedia
		for i = 0; i < 18; i++ {
			if (want == "1" || want == "Angkatbeban" || want == "angkatbeban") && namaL == A[i].Angkatbeban {
				return true
			} else if (want == "2" || want == "Cardio" || want == "cardio") && namaL == A[i].Cardio {
				return true
			} else if (want == "3" || want == "Yoga" || want == "yoga") && namaL == A[i].Yoga {
				return true
			}
		}
	} else {

	}
	return false
}

// function BinarySearch
func BinarySearch(A *DLatihan, B *WorkHariIni, jumlahriwayat int) bool {
	var want, namaL, sort string
	var left, mid, right int
	left = 0
	right = 18
	fmt.Println("Mau Latihan Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali: ")
	fmt.Scan(&want)
	if want != "-" {
		fmt.Print("Nama Latihan: ")
		fmt.Scan(&namaL)
		fmt.Println()

		for left <= right {
			mid = (left + right) / 2
			if want == "1" || want == "Angkatbeban" || want == "angkatbeban" {
				sort = A[mid].Angkatbeban
			} else if want == "2" || want == "Cardio" || want == "cardio" {
				sort = A[mid].Cardio
			} else if want == "3" || want == "Yoga" || want == "yoga" {
				sort = A[mid].Yoga
			} else {
				return false
			}

			if sort == namaL {
				return true
			} else if sort < namaL {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	} else {

	}
	return false
}

// function InsertionSort
func insertionSort(A *WorkHariIni, jumlahRiwayat int) {
	var i, j int
	var temp Workout
	var pilihan string
	fmt.Println("Urutkan Secara:\n1. Ascending\n2. Descending\n3. Reset")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali: ")
	fmt.Scan(&pilihan)
	fmt.Println()
	if pilihan != "-" {
		// Ascending
		if pilihan == "1" || pilihan == "Ascending" || pilihan == "ascending" {
			if jumlahRiwayat == 0 {
				fmt.Println("Anda Belum Memiliki Riwayat Workout.")
			} else {
				for i = 1; i < jumlahRiwayat; i++ {
					temp = A[i]
					j = i - 1
					for j >= 0 && A[j].LAMA > temp.LAMA {
						A[j+1] = A[j]
						j--
					}
					A[j+1] = temp
				}
				Tampilkan(*A, jumlahRiwayat)
				fmt.Println("Data berhasil diurutkan Menaik berdasarkan Waktu!")
			}

			// Descending
		} else if pilihan == "2" || pilihan == "Descending" || pilihan == "descending" {
			if jumlahRiwayat == 0 {
				fmt.Println("Anda Belum Memiliki Riwayat Workout.")
			} else {
				for i = 1; i < jumlahRiwayat; i++ {
					temp = A[i]
					j = i - 1
					for j >= 0 && A[j].LAMA < temp.LAMA {
						A[j+1] = A[j]
						j--
					}
					A[j+1] = temp
				}
				Tampilkan(*A, jumlahRiwayat)
				fmt.Println("Data berhasil diurutkan Menurun berdasarkan Waktu!")
			}
			// Reset
		} else {
			if jumlahRiwayat == 0 {
				fmt.Println("Anda Belum Memiliki Riwayat Workout.")
			} else {
				for i = 1; i < jumlahRiwayat; i++ {
					temp = A[i]
					j = i - 1
					for j >= 0 && A[j].hari > temp.hari {
						A[j+1] = A[j]
						j--
					}
					A[j+1] = temp
				}
				Tampilkan(*A, jumlahRiwayat)
				fmt.Println("Data berhasil direset!")
			}
		}
	} else {

	}
}

// function SelectionSort
func selectionSort(A *WorkHariIni, jumlahRiwayat int) {
	var pass, idx, i int
	var temp Workout
	var pilihan string
	fmt.Println("Urutkan Secara:\n1. Ascending\n2. Descending\n3. Reset")
	fmt.Print("Pilih (1 - 3) | (-) untuk kembali: ")
	fmt.Scan(&pilihan)
	fmt.Println()
	if pilihan != "-" {
		// Ascending
		if pilihan == "1" || pilihan == "Ascending" || pilihan == "ascending" {
			if jumlahRiwayat == 0 {
				fmt.Println("Anda Belum Memiliki Riwayat Workout.")
			} else {
				pass = 0
				for pass < jumlahRiwayat {
					idx = pass
					i = pass + 1
					for i < jumlahRiwayat {
						if A[idx].KALORIBURN > A[i].KALORIBURN {
							idx = i
						}
						i++
					}
					temp = A[pass]
					A[pass] = A[idx]
					A[idx] = temp
					pass++
				}
				Tampilkan(*A, jumlahRiwayat)
				fmt.Println("Data berhasil diurutkan Menaik berdasarkan Kalori!")
			}

			// Descending
		} else if pilihan == "2" || pilihan == "Descending" || pilihan == "descending" {
			if jumlahRiwayat == 0 {
				fmt.Println("Anda Belum Memiliki Riwayat Workout.")
			} else {
				pass = 0
				for pass < jumlahRiwayat {
					idx = pass
					i = pass + 1
					for i < jumlahRiwayat {
						if A[idx].KALORIBURN < A[i].KALORIBURN {
							idx = i
						}
						i++
					}
					temp = A[pass]
					A[pass] = A[idx]
					A[idx] = temp
					pass++
				}
				Tampilkan(*A, jumlahRiwayat)
				fmt.Println("Data berhasil diurutkan Menurun berdasarkan Kalori!")
			}

			// Reset
		} else {
			if jumlahRiwayat == 0 {
				fmt.Println("Anda Belum Memiliki Riwayat Workout.")
			} else {
				pass = 0
				for pass < jumlahRiwayat {
					idx = pass
					i = pass + 1
					for i < jumlahRiwayat {
						if A[idx].hari > A[i].hari {
							idx = i
						}
						i++
					}
					temp = A[pass]
					A[pass] = A[idx]
					A[idx] = temp
					pass++
				}
				Tampilkan(*A, jumlahRiwayat)
				fmt.Println("Data berhasil direset!")
			}
		}
	} else {

	}
}
