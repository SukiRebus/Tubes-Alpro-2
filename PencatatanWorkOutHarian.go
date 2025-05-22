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
	var pilih, jumlahRiwayat, printten, n, KaloriB, m int
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
			menambah(&hari, &jumlahRiwayat)
			printten++
			KaloriB++
		case 2:
			Mengubah(&hari)
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
				fmt.Println("Latihan tidak Ditemukan..")
			}
		case 7:
			BNsearch = BinarySearch(&data, &hari, jumlahRiwayat)
			if BNsearch == true {
				fmt.Println("Latihan ditemukan!")
			} else {
				fmt.Println("Latihan tidak ditemukan..")
			}
		case 8:
			selectionSort(&hari, jumlahRiwayat)
		case 9:
			insertionSort(&hari, jumlahRiwayat)
		}
		fmt.Println()
	}
}

// Prosedure Menambah
func menambah(A *WorkHariIni, jumlahriwayat *int) {
	var lama, kaloriburn, i, totalWaktu, latihanKe int
	var jadwal, MetohdLatihan, latihan, jenisolahraga string
	kaloriburn = 0
	fmt.Println("Olahraga Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
	fmt.Print("Pilih: ")
	fmt.Scan(&jenisolahraga)
	fmt.Println()

	//Angkat Beban
	if jenisolahraga == "1" {
		A[*jumlahriwayat].JENISLOLAHRAGA = "Angkat Beban"
		fmt.Println("Latihan Angkat Beban Apa?\n1. Push\n2. Pull\n3. Legs")
		fmt.Print("Pilih: ")
		fmt.Scan(&MetohdLatihan)

		// Push
		if MetohdLatihan == "1" {
			A[*jumlahriwayat].METODE = "Push"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama : ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
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
				}
			}

			// Pull
		} else if MetohdLatihan == "2" {
			A[*jumlahriwayat].METODE = "Pull"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
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
				}
			}

			// Legs
		} else {
			A[*jumlahriwayat].METODE = "Legs"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
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
				}
			}
		}

		// Cardio
	} else if jenisolahraga == "2" {
		A[*jumlahriwayat].JENISLOLAHRAGA = "Cardio"
		latihanKe = 0
		for latihanKe < MaxLatPerHari && latihan != "." {
			i = 0
			fmt.Println("Pilih Latihan Cardio Berikut:\n1. Running\n2. Burpees\n3. Cycling\n4. JumpRope\n5. StairsClimbing\n6. Swimming\n7. Elliptical\n8. HighKnees\n9. MountainClimbers\n10. Rowing\n11. JumpingJacks\n12. SkaterJumps\n13. ShadowBoxing\n14. Kickboxing\n15. Zumba\n16. BoxJumps\n17. SpeedWalking\n18. BearCrawls")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&latihan)
			if latihan != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan {
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
			}
		}

		// Yoga
	} else if jenisolahraga == "3" {
		A[*jumlahriwayat].JENISLOLAHRAGA = "Yoga"
		latihanKe = 0
		for latihanKe < MaxLatPerHari && latihan != "." {
			i = 0
			fmt.Println("Pilih Latihan Yoga Berikut:\n1. BridgePose\n2. Cat-CowPose\n3. ChildPose\n4. CobraPose\n5. DownwardDog\n6. Plank\n7. SeatedTwist\n8. TrianglePose\n9. TreePose\n10. WarriorPose\n11. LotusPose\n12. BoatPose\n13. HappyBabyPose\n14. PigeonPose\n15. FishPose\n16. CamelPose\n17. CrowPose\n18. ShoulderStand")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&latihan)
			if latihan != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for i = 0; i < lama; i++ {
					switch latihan {
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
			}
		}

	} else {
		fmt.Print("Salah Input Jenis Workout!")
	}
	fmt.Print("Jadwal (DD-MM-YY): ")
	fmt.Scan(&jadwal)
	A[*jumlahriwayat].hari = *jumlahriwayat + 1
	A[*jumlahriwayat].JADWAL = jadwal
	A[*jumlahriwayat].LAMA = totalWaktu
	A[*jumlahriwayat].KALORIBURN = kaloriburn
	A[*jumlahriwayat].totalLatihan = latihanKe
	*jumlahriwayat = *jumlahriwayat + 1
	fmt.Println("Riwayat Workout Telah Ditambahkan")
}

// Prosedure Mengubah
func Mengubah(A *WorkHariIni) {
	var HyangdiUbah, lama, i, kaloriburn, latihanKe, totalWaktu int
	var jadwal, latihan, jenisolahraga, MetohdLatihan, ubahapa string
	fmt.Print("Hari yang Ingin Diubah: ")
	fmt.Scan(&HyangdiUbah)
	HyangdiUbah -= 1
	fmt.Println("Ubah Apa?\n1. Jenisolahraga\n2. Metode\n3. Jadwal\n4. Latihan")
	fmt.Print("Pilih: ")
	fmt.Scan(&ubahapa)
	fmt.Print("Pilih: ")

	//Ubah Jenis Olahraga
	if ubahapa == "1" {
		fmt.Println("Olahraga Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
		fmt.Print("Pilih: ")
		fmt.Scan(&jenisolahraga)

		// Angkat Beban
		if jenisolahraga == "1" {
			A[HyangdiUbah].JENISLOLAHRAGA = "Angkat Beban"
			fmt.Println("Latihan Angkat Beban Apa?\n1. Push\n2. Pull\n3. Legs")
			fmt.Print("Pilih: ")
			fmt.Scan(&MetohdLatihan)

			// Push
			if MetohdLatihan == "1" {
				A[HyangdiUbah].METODE = "Push"
				latihanKe = 0
				for latihanKe < MaxLatPerHari && latihan != "." {
					i = 0
					fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&latihan)
					if latihan != "." {
						fmt.Print("Berapa Lama: ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan {
							case "1":
								kaloriburn += 8
								A[HyangdiUbah].LATIHAN[latihanKe] = "BenchPress"
							case "2":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "ShoulderPress"
							case "3":
								kaloriburn += 5
								A[HyangdiUbah].LATIHAN[latihanKe] = "LateralRaise"
							case "4":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "Dips"
							case "5":
								kaloriburn += 4
								A[HyangdiUbah].LATIHAN[latihanKe] = "PushUp"
							case "6":
								kaloriburn += 5
								A[HyangdiUbah].LATIHAN[latihanKe] = "PushDown"
							}
						}
						totalWaktu += lama
						latihanKe++
					}
				}

				// Pull
			} else if MetohdLatihan == "2" {
				A[HyangdiUbah].METODE = "Pull"
				latihanKe = 0
				for latihanKe < MaxLatPerHari && latihan != "." {
					i = 0
					fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&latihan)
					if latihan != "." {
						fmt.Print("Berapa Lama: ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan {
							case "1":
								kaloriburn += 8
								A[HyangdiUbah].LATIHAN[latihanKe] = "DeadLift"
							case "2":
								kaloriburn += 4
								A[HyangdiUbah].LATIHAN[latihanKe] = "PullUp"
							case "3":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "Rows"
							case "4":
								kaloriburn += 5
								A[HyangdiUbah].LATIHAN[latihanKe] = "BarbelCurl"
							case "5":
								kaloriburn += 4
								A[HyangdiUbah].LATIHAN[latihanKe] = "HammerCurl"
							case "6":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "FacePull"
							}
						}
						totalWaktu += lama
						latihanKe++
					}
				}

				// Legs
			} else {
				A[HyangdiUbah].METODE = "Legs"
				latihanKe = 0
				for latihanKe < MaxLatPerHari && latihan != "." {
					i = 0
					fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&latihan)
					if latihan != "." {
						fmt.Print("Berapa Lama: ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan {
							case "1":
								kaloriburn += 10
								A[HyangdiUbah].LATIHAN[latihanKe] = "Squat"
							case "2":
								kaloriburn += 9
								A[HyangdiUbah].LATIHAN[latihanKe] = "LegPress"
							case "3":
								kaloriburn += 11
								A[HyangdiUbah].LATIHAN[latihanKe] = "Lunge"
							case "4":
								kaloriburn += 8
								A[HyangdiUbah].LATIHAN[latihanKe] = "LegCurl"
							case "5":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "CalfRaise"
							case "6":
								kaloriburn += 10
								A[HyangdiUbah].LATIHAN[latihanKe] = "RDL"
							}
						}
						totalWaktu += lama
						latihanKe++
					}
				}

			}

			// Cardio
		} else if jenisolahraga == "2" {
			A[HyangdiUbah].JENISLOLAHRAGA = "Cardio"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Cardio Berikut:\n1. Running\n2. Burpees\n3. Cycling\n4. JumpRope\n5. StairsClimbing\n6. Swimming\n7. Elliptical\n8. HighKnees\n9. MountainClimbers\n10. Rowing\n11. JumpingJacks\n12. SkaterJumps\n13. ShadowBoxing\n14. Kickboxing\n15. Zumba\n16. BoxJumps\n17. SpeedWalking\n18. BearCrawls")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "Running"
						case "2":
							kaloriburn += 12
							A[HyangdiUbah].LATIHAN[latihanKe] = "Burpees"
						case "3":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "Cycling"
						case "4":
							kaloriburn += 13
							A[HyangdiUbah].LATIHAN[latihanKe] = "JumpRope"
						case "5":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "StairsClimbing"
						case "6":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "Swimming"
						case "7":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "Elliptical"
						case "8":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "HighKnees"
						case "9":
							kaloriburn += 12
							A[HyangdiUbah].LATIHAN[latihanKe] = "MountainClimbers"
						case "10":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "Rowing"
						case "11":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "JumpingJacks"
						case "12":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "SkaterJumps"
						case "13":
							kaloriburn += 7
							A[HyangdiUbah].LATIHAN[latihanKe] = "ShadowBoxing"
						case "14":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "Kickboxing"
						case "15":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "Zumba"
						case "16":
							kaloriburn += 12
							A[HyangdiUbah].LATIHAN[latihanKe] = "BoxJumps"
						case "17":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "SpeedWalking"
						case "18":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "BearCrawls"
						}
					}
					totalWaktu += lama
					latihanKe++
				}
			}

			// Yoga
		} else if jenisolahraga == "3" {
			A[HyangdiUbah].JENISLOLAHRAGA = "Yoga"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Yoga Berikut:\n1. BridgePose\n2. Cat-CowPose\n3. ChildPose\n4. CobraPose\n5. DownwardDog\n6. Plank\n7. SeatedTwist\n8. TrianglePose\n9. TreePose\n10. WarriorPose\n11. LotusPose\n12. BoatPose\n13. HappyBabyPose\n14. PigeonPose\n15. FishPose\n16. CamelPose\n17. CrowPose\n18. ShoulderStand")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "BridgePose"
						case "2":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "Cat-CowPose"
						case "3":
							kaloriburn += 2
							A[HyangdiUbah].LATIHAN[latihanKe] = "ChildPose"
						case "4":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "CobraPose"
						case "5":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "DownwardDog"
						case "6":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "Plank"
						case "7":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "SeatedTwist"
						case "8":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "TrianglePose"
						case "9":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "TreePose"
						case "10":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "WarriorPose"
						case "11":
							kaloriburn += 2
							A[HyangdiUbah].LATIHAN[latihanKe] = "LotusPose"
						case "12":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "BoatPose"
						case "13":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "HappyBabyPose"
						case "14":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "PigeonPose"
						case "15":
							kaloriburn += 2
							A[HyangdiUbah].LATIHAN[latihanKe] = "FishPose"
						case "16":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "CamelPose"
						case "17":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "CrowPose"
						case "18":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "ShoulderStand"
						}
					}
					totalWaktu += lama
					latihanKe++
				}
			}

		} else {
			fmt.Print("Salah Input Jenis Workout!")
		}
		A[HyangdiUbah].LAMA = totalWaktu
		A[HyangdiUbah].KALORIBURN = kaloriburn
		A[HyangdiUbah].totalLatihan = latihanKe

		//Mengubah Metode
	} else if ubahapa == "2" && A[HyangdiUbah].JENISLOLAHRAGA == "Angkat Beban" {
		fmt.Println("Latihan Angkat Beban Apa?\n1. Push\n2. Pull\n3. Legs")
		fmt.Print("Pilih: ")
		fmt.Scan(&MetohdLatihan)

		// Push
		if MetohdLatihan == "1" {
			A[HyangdiUbah].METODE = "Push"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)

				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "BenchPress"
						case "2":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "ShoulderPress"
						case "3":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "LateralRaise"
						case "4":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "Dips"
						case "5":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "PushUp"
						case "6":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "PushDown"
						}
					}
					totalWaktu += lama
					latihanKe++
				}

			}

			// Pull
		} else if MetohdLatihan == "2" {
			A[HyangdiUbah].METODE = "Pull"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)

				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "DeadLift"
						case "2":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "PullUp"
						case "3":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "Rows"
						case "4":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "BarbelCurl"
						case "5":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "HammerCurl"
						case "6":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "FacePull"
						}
					}
					totalWaktu += lama
					latihanKe++
				}
			}

			// Legs
		} else {
			A[HyangdiUbah].METODE = "Legs"
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)

				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "Squat"
						case "2":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "LegPress"
						case "3":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "Lunge"
						case "4":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "LegCurl"
						case "5":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "CalfRaise"
						case "6":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "RDL"
						}
					}
					totalWaktu += lama
					latihanKe++
				}
			}
		}
		A[HyangdiUbah].LAMA = totalWaktu
		A[HyangdiUbah].KALORIBURN = kaloriburn
		A[HyangdiUbah].totalLatihan = latihanKe

		// Mengubah Jadwal
	} else if ubahapa == "3" {
		fmt.Print("Jadwal (DD-MM-YY): ")
		fmt.Scan(&jadwal)
		A[HyangdiUbah].JADWAL = jadwal

		// Mengubah Latihan
	} else if ubahapa == "4" {

		// Jika jenis workoutnya adalah Angkat Beban
		if A[HyangdiUbah].JENISLOLAHRAGA == "Angkat Beban" {

			// Push
			if A[HyangdiUbah].METODE == "Push" {
				latihanKe = 0
				for latihanKe < MaxLatPerHari && latihan != "." {
					i = 0
					fmt.Println("Pilih Latihan Berikut:\n1. BenchPress\n2. ShoulderPress\n3. LateralRaise\n4. Dips\n5. PushUp\n6. PushDown")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&latihan)

					if latihan != "." {
						fmt.Print("Berapa Lama: ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan {
							case "1":
								kaloriburn += 8
								A[HyangdiUbah].LATIHAN[latihanKe] = "BenchPress"
							case "2":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "ShoulderPress"
							case "3":
								kaloriburn += 5
								A[HyangdiUbah].LATIHAN[latihanKe] = "LateralRaise"
							case "4":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "Dips"
							case "5":
								kaloriburn += 4
								A[HyangdiUbah].LATIHAN[latihanKe] = "PushUp"
							case "6":
								kaloriburn += 5
								A[HyangdiUbah].LATIHAN[latihanKe] = "PushDown"
							}
						}
						totalWaktu += lama
						latihanKe++
					}
				}
				A[HyangdiUbah].LAMA = totalWaktu
				A[HyangdiUbah].KALORIBURN = kaloriburn
				A[HyangdiUbah].totalLatihan = latihanKe

				// Pull
			} else if A[HyangdiUbah].METODE == "Pull" {
				latihanKe = 0
				for latihanKe < MaxLatPerHari && latihan != "." {
					i = 0
					fmt.Println("Pilih Latihan Berikut:\n1. DeadLift\n2. PullUp\n3. Rows\n4. BarbelCurl\n5. HammerCurl\n6. FacePull")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&latihan)

					if latihan != "." {
						fmt.Print("Berapa Lama: ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan {
							case "1":
								kaloriburn += 8
								A[HyangdiUbah].LATIHAN[latihanKe] = "DeadLift"
							case "2":
								kaloriburn += 4
								A[HyangdiUbah].LATIHAN[latihanKe] = "PullUp"
							case "3":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "Rows"
							case "4":
								kaloriburn += 5
								A[HyangdiUbah].LATIHAN[latihanKe] = "BarbelCurl"
							case "5":
								kaloriburn += 4
								A[HyangdiUbah].LATIHAN[latihanKe] = "HammerCurl"
							case "6":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "FacePull"
							}
						}
						totalWaktu += lama
						latihanKe++
					}
				}
				A[HyangdiUbah].LAMA = totalWaktu
				A[HyangdiUbah].KALORIBURN = kaloriburn
				A[HyangdiUbah].totalLatihan = latihanKe

				// Legs
			} else {
				latihanKe = 0
				for latihanKe < MaxLatPerHari && latihan != "." {
					i = 0
					fmt.Println("Pilih Latihan Berikut:\n1. Squat\n2. LegPress\n3. Lunge\n4. LegCurl\n5. CalfRaise\n6. RDL")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&latihan)
					if latihan != "." {
						fmt.Print("Berapa Lama: ")
						fmt.Scan(&lama)
						for i = 0; i < lama; i++ {
							switch latihan {
							case "1":
								kaloriburn += 10
								A[HyangdiUbah].LATIHAN[latihanKe] = "Squat"
							case "2":
								kaloriburn += 9
								A[HyangdiUbah].LATIHAN[latihanKe] = "LegPress"
							case "3":
								kaloriburn += 11
								A[HyangdiUbah].LATIHAN[latihanKe] = "Lunge"
							case "4":
								kaloriburn += 8
								A[HyangdiUbah].LATIHAN[latihanKe] = "LegCurl"
							case "5":
								kaloriburn += 6
								A[HyangdiUbah].LATIHAN[latihanKe] = "CalfRaise"
							case "6":
								kaloriburn += 10
								A[HyangdiUbah].LATIHAN[latihanKe] = "RDL"
							}
						}
						totalWaktu += lama
						latihanKe++
					}
				}
			}
			A[HyangdiUbah].LAMA = totalWaktu
			A[HyangdiUbah].KALORIBURN = kaloriburn
			A[HyangdiUbah].totalLatihan = latihanKe

			// Jika jenis workoutnya adalah Cardio
		} else if A[HyangdiUbah].JENISLOLAHRAGA == "Cardio" {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Cardio Berikut:\n1. Running\n2. Burpees\n3. Cycling\n4. JumpRope\n5. StairsClimbing\n6. Swimming\n7. Elliptical\n8. HighKnees\n9. MountainClimbers\n10. Rowing\n11. JumpingJacks\n12. SkaterJumps\n13. ShadowBoxing\n14. Kickboxing\n15. Zumba\n16. BoxJumps\n17. SpeedWalking\n18. BearCrawls")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "Running"
						case "2":
							kaloriburn += 12
							A[HyangdiUbah].LATIHAN[latihanKe] = "Burpees"
						case "3":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "Cycling"
						case "4":
							kaloriburn += 13
							A[HyangdiUbah].LATIHAN[latihanKe] = "JumpRope"
						case "5":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "StairsClimbing"
						case "6":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "Swimming"
						case "7":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "Elliptical"
						case "8":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "HighKnees"
						case "9":
							kaloriburn += 12
							A[HyangdiUbah].LATIHAN[latihanKe] = "MountainClimbers"
						case "10":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "Rowing"
						case "11":
							kaloriburn += 8
							A[HyangdiUbah].LATIHAN[latihanKe] = "JumpingJacks"
						case "12":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "SkaterJumps"
						case "13":
							kaloriburn += 7
							A[HyangdiUbah].LATIHAN[latihanKe] = "ShadowBoxing"
						case "14":
							kaloriburn += 11
							A[HyangdiUbah].LATIHAN[latihanKe] = "Kickboxing"
						case "15":
							kaloriburn += 9
							A[HyangdiUbah].LATIHAN[latihanKe] = "Zumba"
						case "16":
							kaloriburn += 12
							A[HyangdiUbah].LATIHAN[latihanKe] = "BoxJumps"
						case "17":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "SpeedWalking"
						case "18":
							kaloriburn += 10
							A[HyangdiUbah].LATIHAN[latihanKe] = "BearCrawls"
						}
					}
					totalWaktu += lama
					latihanKe++
				}
			}
			A[HyangdiUbah].LAMA = totalWaktu
			A[HyangdiUbah].KALORIBURN = kaloriburn
			A[HyangdiUbah].totalLatihan = latihanKe

			// Jika jenis workoutnya adalah Yoga
		} else {
			latihanKe = 0
			for latihanKe < MaxLatPerHari && latihan != "." {
				i = 0
				fmt.Println("Pilih Latihan Yoga Berikut:\n1. BridgePose\n2. Cat-CowPose\n3. ChildPose\n4. CobraPose\n5. DownwardDog\n6. Plank\n7. SeatedTwist\n8. TrianglePose\n9. TreePose\n10. WarriorPose\n11. LotusPose\n12. BoatPose\n13. HappyBabyPose\n14. PigeonPose\n15. FishPose\n16. CamelPose\n17. CrowPose\n18. ShoulderStand")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for i = 0; i < lama; i++ {
						switch latihan {
						case "1":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "BridgePose"
						case "2":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "Cat-CowPose"
						case "3":
							kaloriburn += 2
							A[HyangdiUbah].LATIHAN[latihanKe] = "ChildPose"
						case "4":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "CobraPose"
						case "5":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "DownwardDog"
						case "6":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "Plank"
						case "7":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "SeatedTwist"
						case "8":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "TrianglePose"
						case "9":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "TreePose"
						case "10":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "WarriorPose"
						case "11":
							kaloriburn += 2
							A[HyangdiUbah].LATIHAN[latihanKe] = "LotusPose"
						case "12":
							kaloriburn += 5
							A[HyangdiUbah].LATIHAN[latihanKe] = "BoatPose"
						case "13":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "HappyBabyPose"
						case "14":
							kaloriburn += 3
							A[HyangdiUbah].LATIHAN[latihanKe] = "PigeonPose"
						case "15":
							kaloriburn += 2
							A[HyangdiUbah].LATIHAN[latihanKe] = "FishPose"
						case "16":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "CamelPose"
						case "17":
							kaloriburn += 6
							A[HyangdiUbah].LATIHAN[latihanKe] = "CrowPose"
						case "18":
							kaloriburn += 4
							A[HyangdiUbah].LATIHAN[latihanKe] = "ShoulderStand"
						}
					}
					totalWaktu += lama
					latihanKe++
				}
			}
		}
		A[HyangdiUbah].LAMA = totalWaktu
		A[HyangdiUbah].KALORIBURN = kaloriburn
		A[HyangdiUbah].totalLatihan = latihanKe

	} else {
		fmt.Println("Salah Input!")
	}
	fmt.Println("Riwayat Workout Telah Diubah")
	fmt.Println()
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
			if A[i].JENISLOLAHRAGA == "Angkat Beban" {
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
			if A[i].JENISLOLAHRAGA == "Angkat Beban" {
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
		if workTerakhir.JENISLOLAHRAGA == "Angkat Beban" {
			fmt.Println("Mau latihan apa?\n1. Angkatbeban\n2. Cardio\n3. Yoga")
			fmt.Print("Pilih: ")
			fmt.Scan(&want)
			fmt.Println()

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

			// Jika workout terakhir adalah Cardio
		} else if workTerakhir.JENISLOLAHRAGA == "Cardio" {
			fmt.Println("Mau latihan apa?\n1. Angkatbeban\n2. Cardio\n3. Yoga")
			fmt.Print("Pilih: ")
			fmt.Scan(&want)
			fmt.Println()

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

			// Jika workout terakhir adalah Yoga
		} else if workTerakhir.JENISLOLAHRAGA == "Yoga" {
			fmt.Println("Mau latihan apa?\n1. Angkatbeban\n2. Cardio\n3. Yoga")
			fmt.Print("Pilih: ")
			fmt.Scan(&want)
			fmt.Println()

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
	fmt.Print("Pilih: ")
	fmt.Scan(&want)
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

	return false
}

// function BinarySearch
func BinarySearch(A *DLatihan, B *WorkHariIni, jumlahriwayat int) bool {
	var want, namaL, sort string
	var left, mid, right int
	left = 0
	right = 18
	fmt.Println("Mau Latihan Apa?\n1. Ankatbeban\n2. Cardio\n3. Yoga")
	fmt.Print("Pilih: ")
	fmt.Scan(&want)
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
	return false
}

// function InsertionSort
func insertionSort(A *WorkHariIni, jumlahRiwayat int) {
	var i, j int
	var temp Workout
	var pilihan string
	fmt.Println("Urutkan Secara:\n1. Ascending\n2. Descending\n3. Reset")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	fmt.Println()
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
}

// function SelectionSort
func selectionSort(A *WorkHariIni, jumlahRiwayat int) {
	var pass, idx, i int
	var temp Workout
	var pilihan string
	fmt.Println("Urutkan Secara:\n1. Ascending\n2. Descending\n3. Reset")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	fmt.Println()

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
}
