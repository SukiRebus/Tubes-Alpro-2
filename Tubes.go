// Back Up
package main

import "fmt"

const NMAX int = 10000

type Workout struct {
	LATIHAN, JADWAL, JENISLOLAHRAGA string
	LAMA, KALORIBURN                int
}
type WorkHariIni [NMAX]Workout
type SimpanWork [NMAX][NMAX]Workout

type Latihan struct {
	Angkatbeban, Cardio, Yoga string
}
type DLatihan [18]Latihan

func main() {
	var pilih, jumlahRiwayat int
	var hari WorkHariIni
	var data1, data2 DLatihan
	var SQsearch, BNsearch bool
	var SimpWO SimpanWork
	data1[0] = Latihan{"BenchPress", "Running", "BridgePose"}
	data1[1] = Latihan{"ShoulderPress", "Burpees", "Cat-CowPose"}
	data1[2] = Latihan{"LateralRaise", "Cycling", "ChildPose"}
	data1[3] = Latihan{"Dips", "JumpRope", "CobraPose"}
	data1[4] = Latihan{"PushUp", "StairsClimbing", "DownwardDog"}
	data1[5] = Latihan{"PushDown", "Swimming", "Plank"}
	data1[6] = Latihan{"DeadLift", "Elliptical", "SeatedTwist"}
	data1[7] = Latihan{"PullUp", "HighKnees", "TrianglePose"}
	data1[8] = Latihan{"Rows", "MountainClimbers", "TreePose"}
	data1[9] = Latihan{"BarbelCurl", "Rowing", "WarriorPose"}
	data1[10] = Latihan{"HammerCurl", "JumpingJacks", "LotusPose"}
	data1[11] = Latihan{"FacePull", "SkaterJumps", "BoatPose"}
	data1[12] = Latihan{"Squat", "ShadowBoxing", "HappyBabyPose"}
	data1[13] = Latihan{"LegPress", "Kickboxing", "PigeonPose"}
	data1[14] = Latihan{"Lunge", "Zumba", "FishPose"}
	data1[15] = Latihan{"LegCurl", "BoxJumps", "CamelPose"}
	data1[16] = Latihan{"CalfRaise", "SpeedWalking", "CrowPose"}
	data1[17] = Latihan{"RDL", "BearCrawls", "ShoulderStand"}

	data2[0] = Latihan{"BarbelCurl", "BearCrawls", "BoatPose"}
	data2[1] = Latihan{"BenchPress", "BoxJumps", "BridgePose"}
	data2[2] = Latihan{"CalfRaise", "Burpees", "CamelPose"}
	data2[3] = Latihan{"DeadLift", "Cycling", "Cat-CowPose"}
	data2[4] = Latihan{"Dips", "Elliptical", "ChildPose"}
	data2[5] = Latihan{"FacePull", "HighKnees", "CobraPose"}
	data2[6] = Latihan{"HammerCurl", "JumpRope", "CrowPose"}
	data2[7] = Latihan{"LateralRaise", "JumpingJacks", "DownwardDog"}
	data2[8] = Latihan{"LegCurl", "Kickboxing", "FishPose"}
	data2[9] = Latihan{"LegPress", "MountainClimbers", "HappyBabyPose"}
	data2[10] = Latihan{"Lunge", "Rowing", "LotusPose"}
	data2[11] = Latihan{"PullUp", "Running", "PigeonPose"}
	data2[12] = Latihan{"PushDown", "ShadowBoxing", "Plank"}
	data2[13] = Latihan{"PushUp", "SkaterJumps", "SeatedTwist"}
	data2[14] = Latihan{"RDL", "SpeedWalking", "ShoulderStand"}
	data2[15] = Latihan{"Rows", "StairsClimbing", "TreePose"}
	data2[16] = Latihan{"ShoulderPress", "Swimming", "TrianglePose"}
	data2[17] = Latihan{"Squat", "Zumba", "WarriorPose"}

	for pilih != 8 {
		fmt.Printf("---------------------------\n           PILIH\n---------------------------\n")
		fmt.Printf("1.Tambah Riwayat Workout\n2.Ubah Riwayat Workout\n3.Hapus Riwayat Workout\n4.Tampilkan Riwayat Workout\n5.Rekomendasi Workout\n6.Cari Latihan(SQ)\n7.Cari Latihan(BN)\n8.Exit\n")
		fmt.Printf("Pilihan Anda: ")
		fmt.Scan(&pilih)
		fmt.Println()
		switch pilih {
		case 1:
			menambah(&hari, &SimpWO, &jumlahRiwayat)
		case 2:
			Mengubah(&hari)
		case 3:
			Menghapus(&hari, &jumlahRiwayat)
		case 4:
			Tampilkan(hari, SimpWO, jumlahRiwayat)
		case 5:
			Rekomendasi(&hari, &jumlahRiwayat)
		case 6:
			SQsearch = SequentialSearch(&data1)
			if SQsearch == true {
				fmt.Println("Latihan Ditemukan!")
			} else {
				fmt.Println("Latihan tidak Ditemukan..")
			}
			fmt.Println()
		case 7:
			BNsearch = BinarySearch(&data2)
			if BNsearch == true {
				fmt.Println("Latihan ditemukan!")
			} else {
				fmt.Println("Latihan tidak ditemukan..")
			}

		}
	}
}

func menambah(A *WorkHariIni, B *SimpanWork, jumlahriwayat *int) {
	var i, lama, kaloriburn, j, totalWaktu int
	var jadwal, MetohdLatihan, latihan, jenisolahraga string
	kaloriburn = 0
	fmt.Println("Olahraga Apa?\n-> Ankatbeban\n-> Cardio\n-> Yoga")
	fmt.Print("Pilih: ")
	fmt.Scan(&jenisolahraga)
	if jenisolahraga == "Angkatbeban" || jenisolahraga == "angkatbeban" {
		fmt.Println("Latihan Angkat Beban Apa?\n-> Push\n-> Pull\n-> Legs")
		fmt.Print("Pilih: ")
		fmt.Scan(&MetohdLatihan)
		if MetohdLatihan == "Push" || MetohdLatihan == "push" {
			for latihan != "." {
				j = 0
				fmt.Println("Pilih Latihan Berikut:\n-> BenchPress\n-> ShoulderPress\n-> LateralRaise\n-> Dips\n-> PushUp\n-> PushDown")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for j = 0; j < lama; j++ {
						switch latihan {
						case "BenchPress":
							kaloriburn += 8
						case "ShoulderPress":
							kaloriburn += 6
						case "LateralRaise":
							kaloriburn += 5
						case "Dips":
							kaloriburn += 6
						case "PushUp":
							kaloriburn += 4
						case "PushDown":
							kaloriburn += 5
						default:
							kaloriburn += 5
						}
					}
					totalWaktu += lama
					B[*jumlahriwayat][i] = Workout{LATIHAN: latihan}
				}

			}

		} else if MetohdLatihan == "Pull" || MetohdLatihan == "pull" {
			for latihan != "." {
				j = 0
				fmt.Println("Pilih Latihan Berikut:\n-> DeadLift\n-> PullUp\n-> Rows\n-> BarbelCurl\n-> HammerCurl\n-> FacePull")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for j = 0; j < lama; j++ {
						switch latihan {
						case "DeadLift":
							kaloriburn += 8
						case "PullUp":
							kaloriburn += 4
						case "Rows":
							kaloriburn += 6
						case "BarbelCurl":
							kaloriburn += 5
						case "HammerCurl":
							kaloriburn += 4
						case "FacePull":
							kaloriburn += 6
						default:
							kaloriburn += 5
						}
					}
					totalWaktu += lama
				}
			}

		} else {
			for latihan != "." {
				j = 0
				fmt.Println("Pilih Latihan Berikut:\n-> Squat\n-> LegPress\n-> Lunge\n-> LegCurl\n-> CalfRaise\n-> RDL")
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&latihan)
				if latihan != "." {
					fmt.Print("Berapa Lama: ")
					fmt.Scan(&lama)
					for j = 0; j < lama; j++ {
						switch latihan {
						case "Squat":
							kaloriburn += 10
						case "LegPress":
							kaloriburn += 9
						case "Lunge":
							kaloriburn += 11
						case "LegCurl":
							kaloriburn += 8
						case "CalfRaise":
							kaloriburn += 6
						case "RDL":
							kaloriburn += 10
						default:
							kaloriburn += 7
						}
					}
					totalWaktu += lama
				}
			}
		}

	} else if jenisolahraga == "Cardio" || jenisolahraga == "cardio" {
		for latihan != "." {
			j = 0
			fmt.Println("Latihan Cardio Apa?\n-> Running\n-> Burpees\n-> Cycling\n-> JumpRope\n-> StairsClimbing\n-> Swimming\n-> Elliptical\n-> HighKnees\n-> MountainClimbers\n-> Rowing\n-> JumpingJacks-> SkaterJumps\n-> ShadowBoxing\n-> Kickboxing\n-> Zumba\n-> BoxJumps\n-> BoxJumps\n-> SpeedWalking\n-> BearCrawls")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&latihan)
			if latihan != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for j = 0; j < lama; j++ {
					switch latihan {
					case "Running":
						kaloriburn += 10
					case "Burpees":
						kaloriburn += 12
					case "Cycling":
						kaloriburn += 8
					case "JumpRope":
						kaloriburn += 13
					case "StairsClimbing":
						kaloriburn += 11
					case "Swimming":
						kaloriburn += 10
					case "Elliptical":
						kaloriburn += 9
					case "HighKnees":
						kaloriburn += 11
					case "MountainClimbers":
						kaloriburn += 12
					case "Rowing":
						kaloriburn += 9
					case "JumpingJacks":
						kaloriburn += 8
					case "SkaterJumps":
						kaloriburn += 10
					case "ShadowBoxing":
						kaloriburn += 7
					case "Kickboxing":
						kaloriburn += 11
					case "Zumba":
						kaloriburn += 9
					case "BoxJumps":
						kaloriburn += 12
					case "SpeedWalking":
						kaloriburn += 6
					case "BearCrawls":
						kaloriburn += 10
					default:
						kaloriburn += 8
					}
				}
				totalWaktu += lama
			}
		}

	} else {
		for latihan != "." {
			j = 0
			fmt.Println("Latihan Yoga Apa?\n-> BridgePose\n-> Cat-CowPose\n-> ChildPose\n-> CobraPose\n-> DownwardDog\n-> Plank\n-> SeatedTwist\n-> TrianglePose\n-> TreePose\n-> WarriorPose\n-> LotusPose\n-> BoatPose\n-> HappyBabyPose\n-> PigeonPose\n-> FishPose\n-> CamelPose\n-> CrowPose\n-> ShoulderStand")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&latihan)
			if latihan != "." {
				fmt.Print("Berapa Lama: ")
				fmt.Scan(&lama)
				for j = 0; j < lama; j++ {
					switch latihan {
					case "BridgePose":
						kaloriburn += 4
					case "Cat-CowPose":
						kaloriburn += 3
					case "ChildPose":
						kaloriburn += 2
					case "CobraPose":
						kaloriburn += 3
					case "DownwardDog":
						kaloriburn += 4
					case "Plank":
						kaloriburn += 5
					case "SeatedTwist":
						kaloriburn += 3
					case "TrianglePose":
						kaloriburn += 4
					case "TreePose":
						kaloriburn += 3
					case "WarriorPose":
						kaloriburn += 5
					case "LotusPose":
						kaloriburn += 2
					case "BoatPose":
						kaloriburn += 5
					case "HappyBabyPose":
						kaloriburn += 3
					case "PigeonPose":
						kaloriburn += 3
					case "FishPose":
						kaloriburn += 2
					case "CamelPose":
						kaloriburn += 4
					case "CrowPose":
						kaloriburn += 6
					case "ShoulderStand":
						kaloriburn += 4
					default:
						kaloriburn += 3
					}
				}
				totalWaktu += lama
			}
		}
	}
	fmt.Print("Jadwal (DD-MM-YY): ")
	fmt.Scan(&jadwal)
	A[i] = Workout{LATIHAN: latihan, JENISLOLAHRAGA: jenisolahraga, JADWAL: jadwal, LAMA: totalWaktu, KALORIBURN: kaloriburn}
	*jumlahriwayat++
	fmt.Println("Riwayat wWorkout Telah Ditambahkan")
}

func Mengubah(A *WorkHariIni) {
	var HyangdiUbah, lama, kaloriburn int
	var jadwal, latihan, jenisolahraga string
	fmt.Print("Hari yang Ingin Diubah: ")
	fmt.Scan(&HyangdiUbah)
	HyangdiUbah -= 1
	fmt.Print("Jenis Olahraga: ")
	fmt.Scan(&jenisolahraga)
	fmt.Print("Latihan: ")
	fmt.Scan(&latihan)
	fmt.Print("Jadwal (DD-MM-YY): ")
	fmt.Scan(&jadwal)
	fmt.Print("Lama: ")
	fmt.Scan(&lama)
	fmt.Print("Kalori Terbakar: ")
	fmt.Scan(&kaloriburn)
	A[HyangdiUbah] = Workout{LATIHAN: latihan, JENISLOLAHRAGA: jenisolahraga, JADWAL: jadwal, LAMA: lama, KALORIBURN: kaloriburn}
}

func Menghapus(A *WorkHariIni, jumlahriwayat *int) {
	var HyangdiHapus, i int
	fmt.Print("Hari yang Ingin Dihapus: ")
	fmt.Scan(&HyangdiHapus)
	HyangdiHapus -= 1
	for i = HyangdiHapus; i < *jumlahriwayat; i++ {
		A[i] = A[i+1]
	}
}

func Tampilkan(A WorkHariIni, B SimpanWork, jumlahriwayat int) {
	var i, j int
	fmt.Println("Riwayat Workout Anda:")
	if jumlahriwayat <= 0 {
		fmt.Println("Anda Belum Memiliki Riwayat Workout")
	} else {
		for i = 0; i < jumlahriwayat; i++ {
			fmt.Printf("Hari Ke: %d\nJenis Olahraga %s\nLatihan: %s\n", i+1, A[i].JENISLOLAHRAGA, A[i].LATIHAN)
			for j = 0; j < jumlahriwayat; j++ {
				fmt.Printf("Latihan: %s\n", B[jumlahriwayat][i])
			}
			fmt.Printf("Jadwal: %s\nLama Workout: %d menit\nJumlah Kalori yang Terbakar: %d\n", A[i].JADWAL, A[i].LAMA, A[i].KALORIBURN)
		}
		fmt.Println()
	}
}

func Rekomendasi(A *WorkHariIni, jumlahriwayat *int) {
	var workTerakhir Workout
	var want string
	workTerakhir = A[*jumlahriwayat-1]
	if workTerakhir.JENISLOLAHRAGA == "Angkatbeban" || workTerakhir.JENISLOLAHRAGA == "angkatbeban" {
		fmt.Println("mau Latihan apa?\n-> Ankatbeban\n-> Cardio\n-> Yoga")
		fmt.Print("Pilih: ")
		fmt.Scan(&want)
		fmt.Println()
		if want == "Angkatbeban" || want == "angkatbeban" {
			if workTerakhir.LATIHAN == "Push" || workTerakhir.LATIHAN == "push" {
				fmt.Println(" ")
			} else if workTerakhir.LATIHAN == "Push" || workTerakhir.LATIHAN == "push" {
				fmt.Println(" ")
			} else {
				fmt.Println(" ")
			}
		} else if want == "Cardio" || want == "cardio" {
			fmt.Println(" ")
		} else {
			fmt.Println(" ")
		}

	} else if workTerakhir.JENISLOLAHRAGA == "Cardio" || workTerakhir.JENISLOLAHRAGA == "cardio" {
		fmt.Println("mau Latihan apa?\n-> Ankatbeban\n-> Cardio\n-> Yoga")
		fmt.Print("Pilih: ")
		fmt.Scan(&want)
		fmt.Println()
		if want == "Angkatbeban" || want == "angkatbeban" {
			fmt.Println(" ")
		} else if want == "Cardio" || want == "cardio" {
			fmt.Println(" ")
		} else {
			fmt.Println(" ")
		}

	} else if workTerakhir.JENISLOLAHRAGA == "Yoga" || workTerakhir.JENISLOLAHRAGA == "yoga" {
		fmt.Println("mau Latihan apa?\n-> Ankatbeban\n-> Cardio\n-> Yoga")
		fmt.Print("Pilih: ")
		fmt.Scan(&want)
		fmt.Println()
		if want == "Angkatbeban" || want == "angkatbeban" {
			fmt.Println(" ")
		} else if want == "Cardio" || want == "cardio" {
			fmt.Println(" ")
		} else {
			fmt.Println(" ")
		}
	}
}

func SequentialSearch(A *DLatihan) bool {
	var want, namaL string
	var i int
	fmt.Println("Mau Latihan Apa?\n-> Ankatbeban\n-> Cardio\n-> Yoga")
	fmt.Print("Pilih: ")
	fmt.Scan(&want)
	fmt.Print("Nama Latihan: ")
	fmt.Scan(&namaL)
	fmt.Println()
	for i = 0; i < 10; i++ {
		if (want == "Angkatbeban" || want == "angkatbeban") && namaL == A[i].Angkatbeban {
			return true
		} else if (want == "Cardio" || want == "cardio") && namaL == A[i].Cardio {
			return true
		} else if (want == "Yoga" || want == "yoga") && namaL == A[i].Cardio {
			return true
		}
	}
	return false
}

func BinarySearch(A *DLatihan) bool {
	var want, namaL, sort string
	var left, mid, right int
	left = 0
	right = 9
	fmt.Println("Mau Latihan Apa?\n-> Ankatbeban\n-> Cardio\n-> Yoga")
	fmt.Print("Pilih: ")
	fmt.Scan(&want)
	fmt.Print("Nama Latihan: ")
	fmt.Scan(&namaL)
	fmt.Println()

	for left <= right {
		mid = (left + right) / 2
		if want == "Angkatbeban" || want == "angkatbeban" {
			sort = A[mid].Angkatbeban
		} else if want == "Cardio" || want == "cardio" {
			sort = A[mid].Cardio
		} else if want == "Yoga" || want == "yoga" {
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
