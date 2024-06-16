package main

import (
	"fmt"
)

type user struct {
	username string
	password string
	approved bool
}

type Group struct {
	name         string
	creator      *user
	members      [NMAX]*user
	messages     [NMAX]string
	memberCount  int
	messageCount int
}

type PrivateChat struct {
	sender   user
	receiver user
	content  string
}

const NMAX int = 100

type tabUser struct {
	data   [NMAX]user
	length int
}

type tabGroup struct {
	data   [NMAX]Group
	length int
}

type tabPrivateChats struct {
	data   [NMAX]PrivateChat
	length int
}

var users tabUser
var PrivateChats tabPrivateChats
var groups tabGroup

func main() {
	var role string

	users.data[0] = user{username: "a", password: "a", approved: true}
	users.data[1] = user{username: "s", password: "s", approved: true}
	users.length = 2

	for {
		fmt.Println("Menu Utama:")
		fmt.Println("1. Admin")
		fmt.Println("2. User")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&role)

		switch role {
		case "1":
			adminmenu(&users)
		case "2":
			usermenu(&users)
		case "3":
			fmt.Println("Keluar dari program")
			return
		default:
			fmt.Println("Silahkan isi pilihan yang valid")
		}
	}
}

func adminmenu(users *tabUser) {
	var password string
	fmt.Print("Masukkan password admin: ")
	fmt.Scan(&password)
	if password != "jojo" {
		fmt.Println("Password salah, akses ditolak")
		fmt.Println()
		return
	}
	fmt.Println()
	for {
		var choice int
		fmt.Println("Admin Menu:")
		fmt.Println("1. Lihat Pengguna Yang Sudah Terdaftar")
		fmt.Println("2. Lihat Pengguna Yang Menunggu Diapprove")
		fmt.Println("3. Setujui/Tolak Pengguna")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewUsers(*users)
		case 2:
			viewUsers2(*users)
		case 3:
			approveRejectUsers(users)
		case 4:
			fmt.Println()
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func viewUsers(users tabUser) {
	var pilih int
	fmt.Println("Anda Ingin Menampilka Data Secara Ascending Atau Descending?")
	fmt.Println("Ketik (1) Ascending, (2) Descending")
	fmt.Scan(&pilih)
	if pilih == 1 {
		urutAscending(&users)
	} else if pilih == 2 {
		urutDescending(&users)
	} else {
		fmt.Println("Pilihan tidak valid")
		fmt.Println()
		return
	}
	fmt.Println("Pengguna yang terdaftar:")
	for i := 0; i < users.length; i++ {
		if users.data[i].username != "" && users.data[i].approved {
			fmt.Println(users.data[i].username)
		}
	}
	var cari string
	fmt.Print("Ingin mencari user? (Y/N): ")
	fmt.Scan(&cari)
	if pilih == 1 && cari == "Y" {
		fmt.Print("Masukkan nama user yang ingin dicari: ")
		fmt.Scan(&cari)
		idx := binarySearchAscending(&users, cari)
		if idx != -1 {
			fmt.Printf("User %s ditemukan pada indeks %d\n", cari, idx)
		} else {
			fmt.Printf("User %s tidak ditemukan\n", cari)
		}
	} else if pilih == 2 && cari == "Y" {
		fmt.Print("Masukkan nama user yang ingin dicari: ")
		fmt.Scan(&cari)
		idx := binarySearchDescending(&users, cari)
		if idx != -1 {
			fmt.Printf("User %s ditemukan pada indeks %d\n", cari, idx)
		} else {
			fmt.Printf("User %s tidak ditemukan\n", cari)
		}
	}
	fmt.Println()
}

func urutAscending(users *tabUser) {
	var pass, idx, i int
	var temp user

	pass = 1
	for pass <= users.length-1 {
		idx = pass - 1
		i = pass
		for i < users.length {
			if users.data[idx].username > users.data[i].username {
				idx = i
			}
			i = i + 1
		}
		temp = users.data[pass-1]
		users.data[pass-1] = users.data[idx]
		users.data[idx] = temp
		pass = pass + 1
	}
}
func urutDescending(users *tabUser) {
	var pass, i int
	var temp user
	pass = 1
	for pass <= users.length-1 {
		i = pass
		temp.username = users.data[i].username
		for i > 0 && temp.username > users.data[i-1].username {
			users.data[i].username = users.data[i-1].username
			i = i - 1
		}
		users.data[i].username = temp.username
		pass = pass + 1
	}
}

func viewUsers2(users tabUser) {
	fmt.Println("Pengguna yang belum diapprove:")
	for i := 0; i < users.length; i++ {
		if users.data[i].username != "" && !users.data[i].approved {
			fmt.Println(users.data[i].username)
		}
	}
	fmt.Println()
}

func binarySearchAscending(users *tabUser, target string) int {
	low := 0
	high := users.length - 1

	for low <= high {
		mid := (low + high) / 2
		if users.data[mid].username == target {
			return mid
		} else if users.data[mid].username < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func binarySearchDescending(users *tabUser, target string) int {
	low := 0
	high := users.length - 1

	for low <= high {
		mid := (low + high) / 2
		if users.data[mid].username == target {
			return mid
		} else if users.data[mid].username > target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func approveRejectUsers(users *tabUser) {
	var username string
	var action string
	fmt.Println()
	fmt.Print("Masukkan nama pengguna yang ingin disetujui/tolak: ")
	fmt.Scan(&username)

	for i := 0; i < users.length; i++ {
		if users.data[i].username == username {
			fmt.Print("Apakah Anda ingin menyetujui atau menolak pengguna ini? (approve/reject): ")
			fmt.Scan(&action)

			switch action {
			case "approve":
				users.data[i].approved = true
				fmt.Printf("Pengguna %s telah disetujui.\n", username)
				return
			case "reject":
				users.data[i].approved = false
				fmt.Printf("Pengguna %s telah ditolak.\n", username)
				return
			default:
				fmt.Println("Aksi tidak valid. Silakan pilih 'approve' atau 'reject'.")
				return
			}
		}
	}
	fmt.Println("Pengguna tidak ditemukan.")
	fmt.Println()
}

func usermenu(users *tabUser) {
	fmt.Println()

	var currentUser user
	for {
		var choice int
		fmt.Println("User Menu:")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih Opsi:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			register(users)
		case 2:
			login(*users, &currentUser)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func register(users *tabUser) {
	var newUser user
	fmt.Println("Registrasi Pengguna Baru:")
	fmt.Print("Username: ")
	fmt.Scan(&newUser.username)
	fmt.Print("Password: ")
	fmt.Scan(&newUser.password)
	newUser.approved = false

	for i := 0; i < users.length; i++ {
		if users.data[i].username == newUser.username {
			fmt.Println("Username sudah terdaftar. Silakan pilih username lain.")
			fmt.Println()
			return
		}
	}

	for i := 0; i < NMAX; i++ {
		if users.data[i].username == "" {
			users.data[i] = newUser
			fmt.Printf("Pengguna %s berhasil terdaftar. Mohon tunggu persetujuan admin.\n", newUser.username)
			fmt.Println()

			users.length++
			return
		}
	}
	fmt.Println("Batas maksimum pengguna telah tercapai.")
}

func login(users tabUser, currentUser *user) {
	var username, password string
	fmt.Println("Masukkan informasi login:")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < users.length; i++ {
		if users.data[i].username == username && users.data[i].password == password {
			if users.data[i].approved {
				fmt.Printf("Selamat datang, %s! Anda berhasil login.\n", username)
				currentUser = &users.data[i]
				userLoggedInMenu(&users, currentUser)

				return
			} else {
				fmt.Println("Akun Anda masih menunggu persetujuan admin. Mohon tunggu.")
				return
			}
		}
	}

	fmt.Println("Username atau password salah.")

}

func userLoggedInMenu(users *tabUser, currentUser *user) {
	for {
		var choice int
		fmt.Println("User Logged In Menu:")
		fmt.Println("1. Kirim Pesan Pribadi")
		fmt.Println("2. Inbox")
		fmt.Println("3. Pesan yang Dikirim")
		fmt.Println("4. Group")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih Opsi:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			sendPrivateMessage(users, currentUser)
		case 2:
			viewInbox(currentUser)
		case 3:
			ViewSendMessagers(users, currentUser)
		case 4:
			groupMenu(currentUser)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func sendPrivateMessage(users *tabUser, currentUser *user) {
	var receiver, message string
	fmt.Println("Kirim Pesan Pribadi:")
	fmt.Print("Username Penerima: ")
	fmt.Scan(&receiver)

	var userReceiver user

	found := false
	for i := 0; i < users.length; i++ {
		if users.data[i].username == receiver {
			if users.data[i].approved {
				found = true
				userReceiver = users.data[i]

			} else {
				fmt.Println("Penerima pesan belum diapprove oleh admin.")
				return
			}
		}
	}
	if !found {
		fmt.Println("Penerima pesan tidak ditemukan.")
		return
	}

	fmt.Println("Isi Pesan: ")
	var temp byte
	fmt.Scanf("%c", temp)

	for temp != ';' {
		message += string(temp)
		fmt.Scanf("%c", &temp)
	}

	PrivateChats.data[PrivateChats.length] = PrivateChat{sender: *currentUser, receiver: userReceiver, content: message}
	PrivateChats.length++

	return
}

func viewInbox(currentUser *user) {
	fmt.Println()
	fmt.Println("Inbox")

	var inboxCount int

	for i := 0; i < PrivateChats.length; i++ {
		message := PrivateChats.data[i]

		if message.receiver == *currentUser {
			fmt.Println("[", i+1, "]", "Pesan dari", message.sender.username)
			inboxCount++
		}
	}

	if inboxCount == 0 {
		fmt.Println("Tidak ada pesan.")
		fmt.Println()
	}

	var selectedPrivateChat int
	fmt.Print("Pilih inbox (0 untuk exit): ")
	fmt.Scan(&selectedPrivateChat)

	if selectedPrivateChat == 0 {
		return
	}

	for !(selectedPrivateChat >= 1 && selectedPrivateChat <= PrivateChats.length) && PrivateChats.data[selectedPrivateChat-1].receiver != *currentUser {
		fmt.Println("No inbox tidak valid. Coba lagi.")
		fmt.Print("Pilih inbox (0 untuk exit): ")
		fmt.Scan(&selectedPrivateChat)
	}

	message := PrivateChats.data[selectedPrivateChat-1]

	fmt.Println()
	fmt.Println("From:", message.sender.username)
	fmt.Println("To:", message.receiver.username)
	fmt.Println("Message:", message.content)

	fmt.Println()
}

func ViewSendMessagers(users *tabUser, currentUser *user) {
	fmt.Println()
	fmt.Println("Pesan yang Dikirim")

	var sendCount int

	for i := 0; i < PrivateChats.length; i++ {
		message := PrivateChats.data[i]

		if message.sender == *currentUser {
			fmt.Println("[", i+1, "]", "Pesan ke", message.receiver.username)
			sendCount++
		}
	}

	if sendCount == 0 {
		fmt.Println("Tidak ada pesan.")
		fmt.Println()
	}

	var selectedPrivateChat int
	fmt.Print("Pilih pesan yang dikirim (0 untuk exit): ")
	fmt.Scan(&selectedPrivateChat)

	if selectedPrivateChat == 0 {
		return
	}

	for !(selectedPrivateChat >= 1 && selectedPrivateChat <= PrivateChats.length) && PrivateChats.data[selectedPrivateChat-1].sender != *currentUser {
		fmt.Println("No inbox tidak valid. Coba lagi.")
		fmt.Print("Pilih inbox (0 untuk exit): ")
		fmt.Scan(&selectedPrivateChat)
	}

	message := PrivateChats.data[selectedPrivateChat-1]

	fmt.Println()
	fmt.Println("From:", message.sender.username)
	fmt.Println("To:", message.receiver.username)
	fmt.Println("Message:", message.content)

	fmt.Println()
}

func groupMenu(currentUser *user) {
	for {
		var choice int
		fmt.Println("Group Menu:")
		fmt.Println("1. Buat Group")
		fmt.Println("2. Lihat Group")
		fmt.Println("3. Kirim Pesan ke Group")
		fmt.Println("4. Lihat Pesan dari Group")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createGroup(currentUser)
		case 2:
			viewJoinedGroups(currentUser)
		case 3:
			sendGroupMessage(currentUser)
		case 4:
			groupMessage(currentUser)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func viewJoinedGroups(currentUser *user) {
	fmt.Println("Grup yang Anda ikuti:")
	for i := 0; i < groups.length; i++ {
		if inGroup(*currentUser, groups.data[i]) {
			fmt.Printf("[%d] %s\n", i+1, groups.data[i].name)
			fmt.Println("Members:")

			for j := 0; j < groups.data[i].memberCount; j++ {
				fmt.Printf("- %s\n", groups.data[i].members[j].username)
			}
		}
	}

}

func inGroup(u user, g Group) bool {
	for i := 0; i < g.memberCount; i++ {
		if g.members[i].username == u.username {
			return true
		}
	}

	return false
}

func createGroup(currentUser *user) {
	var groupName string
	fmt.Print("Masukkan nama grup baru: ")
	fmt.Scan(&groupName)

	for i := 0; i < groups.length; i++ {
		if groups.data[i].name == groupName {
			fmt.Println("Nama grup sudah ada. Silakan pilih nama lain.")
			return
		}
	}

	groups.data[groups.length] = Group{
		name:         groupName,
		creator:      currentUser,
		memberCount:  1,
		messageCount: 0,
	}

	groups.data[groups.length].members[0] = currentUser

	fmt.Println("Daftar pengguna yang terdaftar di grup ini:")
	for i := 0; i < groups.data[groups.length].memberCount; i++ {
		fmt.Println("-", groups.data[groups.length].members[i].username)
	}

	var addUser string
	for addUser != "0" {
		fmt.Print("Masukkan nama pengguna untuk diundang ke grup (0 untuk berhenti): ")
		fmt.Scan(&addUser)

		if addUser != "0" {
			userFound := false
			for i := 0; i < NMAX && !userFound; i++ {
				if users.data[i].username == addUser && users.data[i].approved {
					userFound = true
					alreadyMember := false
					for j := 0; j < groups.data[groups.length].memberCount && !alreadyMember; j++ {
						if groups.data[groups.length].members[j].username == addUser {
							alreadyMember = true
						}
					}
					if alreadyMember {
						fmt.Println("Pengguna sudah menjadi anggota grup.")
					} else {
						groups.data[groups.length].members[groups.data[groups.length].memberCount] = &users.data[i]
						groups.data[groups.length].memberCount++
						fmt.Printf("Pengguna %s telah ditambahkan ke grup.\n", addUser)
					}
				}
			}

			if !userFound {
				fmt.Println("Pengguna tidak ditemukan atau belum diapprove.")
			}
		}

	}

	fmt.Printf("Grup %s berhasil dibuat.\n", groupName)
	groups.length++
}

func viewGroups(currentUser *user) {
	fmt.Println("Daftar grup yang Anda ikuti:")
	for i := 0; i < groups.length; i++ {
		if inGroup(*currentUser, groups.data[i]) {
			fmt.Printf("[%d] %s\n", i+1, groups.data[i].name)
		}
	}

	var selectedGroup int
	fmt.Print("Pilih nomor grup untuk melihat anggota (0 untuk kembali): ")
	fmt.Scan(&selectedGroup)

	if selectedGroup == 0 {
		return
	}

	if selectedGroup < 1 || selectedGroup > groups.length {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	groupIndex := selectedGroup - 1
	fmt.Printf("Anggota grup %s:\n", groups.data[groupIndex].name)
	for j := 0; j < groups.data[groupIndex].memberCount; j++ {
		if groups.data[groupIndex].members[j] != nil {
			fmt.Printf("- %s\n", groups.data[groupIndex].members[j].username)
		}
	}
}

func sendGroupMessage(currentUser *user) {
	fmt.Println("Daftar grup:")
	for i := 0; i < groups.length; i++ {
		if inGroup(*currentUser, groups.data[i]) {
			fmt.Printf("[%d] %s\n", i+1, groups.data[i].name)
		}
	}

	var selectedGroup int
	fmt.Print("Pilih nomor grup untuk mengirim pesan (0 untuk kembali): ")
	fmt.Scan(&selectedGroup)

	if selectedGroup == 0 {
		return
	}

	for !inGroup(*currentUser, groups.data[selectedGroup-1]) {
		fmt.Println("Grup tidak ditemukan. Coba lagi.")

		fmt.Print("Pilih nomor grup untuk mengirim pesan (0 untuk kembali): ")
		fmt.Scan(&selectedGroup)
	}

	groupIndex := selectedGroup - 1

	var message string
	fmt.Print("Masukkan pesan (akhiri dengan ';'): ")
	var temp byte
	fmt.Scanf("%c", &temp)
	for temp != ';' {
		message += string(temp)
		fmt.Scanf("%c", &temp)
	}

	groups.data[groupIndex].messages[groups.data[groupIndex].messageCount] = message
	groups.data[groupIndex].messageCount++
	fmt.Println("Pesan berhasil dikirim ke grup.")
}

func groupMessage(currentUser *user) {
	fmt.Println("Daftar grup yang Anda ikuti:")
	for i := 0; i < groups.length; i++ {
		if inGroup(*currentUser, groups.data[i]) {
			fmt.Printf("[%d] %s\n", i+1, groups.data[i].name)
		}
	}

	var selectedGroup int
	fmt.Print("Pilih nomor grup untuk melihat pesan (0 untuk kembali): ")
	fmt.Scan(&selectedGroup)

	if selectedGroup == 0 {
		return
	}

	for !inGroup(*currentUser, groups.data[selectedGroup-1]) {
		fmt.Println("Grup tidak ditemukan. Coba lagi.")

		fmt.Print("Pilih nomor grup untuk melihat pesan (0 untuk kembali): ")
		fmt.Scan(&selectedGroup)
	}

	groupIndex := selectedGroup - 1

	fmt.Printf("Pesan dalam grup %s:\n", groups.data[groupIndex].name)
	for i := 0; i < groups.data[groupIndex].messageCount; i++ {
		fmt.Printf("[%d] %s\n", i+1, groups.data[groupIndex].messages[i])
	}
}
