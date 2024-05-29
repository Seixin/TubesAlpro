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
	name        string
	creator     *user
	members     [NMAX]*user
	messages    [NMAX]string
	memberCount int
}

type PrivateChat struct {
	sender   user
	receiver user
	content  string
}

const NMAX int = 100

type tabuser [NMAX]user
type tabgroup [NMAX]Group
type tabPrivateChats [NMAX]PrivateChat

// Global array variables
var PrivateChats tabPrivateChats
var groups tabgroup

// Lengths
var nPrivateChat int = 0
var nGroup int = 0
var nUser int = 0

// Simple 'authentication'
var currentUser *user

func main() {
	var users tabuser
	var role string

	// Dummy data
	users[0] = user{username: "a", password: "a", approved: true}
	users[1] = user{username: "s", password: "s", approved: true}
	nUser = 2

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

func adminmenu(users *tabuser) {
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

func viewUsers(users tabuser) {

	fmt.Println("Pengguna yang terdaftar:")
	var u user
	for i := 0; i < NMAX; i++ {
		u = users[i]
		if u.username != "" {
			if u.approved {
				fmt.Println(u.username)
			}
		}
	}
	fmt.Println()
}

func viewUsers2(users tabuser) {

	fmt.Println("Pengguna yang belum diapprove:")
	var u user
	for i := 0; i < NMAX; i++ {
		u = users[i]
		if u.username != "" {
			if !u.approved {
				fmt.Println(u.username)
			}
		}
	}
	fmt.Println()
}
func approveRejectUsers(users *tabuser) {
	var username string
	var action string
	fmt.Println()
	fmt.Print("Masukkan nama pengguna yang ingin disetujui/tolak: ")
	fmt.Scan(&username)

	for i := 0; i < NMAX; i++ {
		if users[i].username == username {
			fmt.Print("Apakah Anda ingin menyetujui atau menolak pengguna ini? (approve/reject): ")
			fmt.Scan(&action)

			switch action {
			case "approve":
				users[i].approved = true
				fmt.Printf("Pengguna %s telah disetujui.\n", username)
				return
			case "reject":
				users[i].approved = false
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

func usermenu(users *tabuser) {
	fmt.Println()
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
			login(*users)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func register(users *tabuser) {
	var newUser user
	fmt.Println("Registrasi Pengguna Baru:")
	fmt.Print("Username: ")
	fmt.Scan(&newUser.username)
	fmt.Print("Password: ")
	fmt.Scan(&newUser.password)
	newUser.approved = false

	for i := 0; i < NMAX; i++ {
		if users[i].username == newUser.username {
			fmt.Println("Username sudah terdaftar. Silakan pilih username lain.")
			fmt.Println()
			return
		}
	}

	for i := 0; i < NMAX; i++ {
		if users[i].username == "" {
			users[i] = newUser
			fmt.Printf("Pengguna %s berhasil terdaftar. Mohon tunggu persetujuan admin.\n", newUser.username)
			fmt.Println()

			nUser++
			return
		}
	}
	fmt.Println("Batas maksimum pengguna telah tercapai.")
}

func login(users tabuser) {
	var username, password string
	fmt.Println("Masukkan informasi login:")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < NMAX; i++ {
		if users[i].username == username && users[i].password == password {
			if users[i].approved {
				fmt.Printf("Selamat datang, %s! Anda berhasil login.\n", username)
				currentUser = &users[i]
				userLoggedInMenu(&users)

				return
			} else {
				fmt.Println("Akun Anda masih menunggu persetujuan admin. Mohon tunggu.")
				return
			}
		}
	}

	fmt.Println("Username atau password salah.")

}

func userLoggedInMenu(users *tabuser) {
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
			sendPrivateMessage(users)
		case 2:
			viewInbox(users)
		case 3:
			ViewSendMessagers(users)
		case 4:
			groupMenu(users)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func sendPrivateMessage(users *tabuser) {
	var receiver, message string
	fmt.Println("Kirim Pesan Pribadi:")
	fmt.Print("Username Penerima: ")
	fmt.Scan(&receiver)

	var userReceiver user

	found := false
	for i := 0; i < NMAX; i++ {
		if users[i].username == receiver {
			if users[i].approved {
				found = true
				userReceiver = users[i]
				break
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

	PrivateChats[nPrivateChat] = PrivateChat{sender: *currentUser, receiver: userReceiver, content: message}
	nPrivateChat++

	return
}

func viewInbox(users *tabuser) {
	fmt.Println()
	fmt.Println("Inbox")

	var inboxCount int

	for i := 0; i < nPrivateChat; i++ {
		message := PrivateChats[i]

		if message.receiver == *currentUser {
			fmt.Println("[", inboxCount+1, "]", "Pesan dari", message.sender.username)
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

	for !(selectedPrivateChat >= 1 && selectedPrivateChat <= nPrivateChat) {

		if selectedPrivateChat == 0 {
			return
		}

		fmt.Print("Pilih inbox (0 untuk exit): ")
		fmt.Scan(&selectedPrivateChat)
	}

	message := PrivateChats[selectedPrivateChat-1]

	fmt.Println()
	fmt.Println("From\t:\t", (message.sender).username)
	fmt.Println("Content\t:\t", message.content)
	fmt.Println("To\t:\t", message.receiver.username)
	fmt.Println()

	return

}

func ViewSendMessagers(users *tabuser) {
	fmt.Println()
	fmt.Println("Pesan yang dikirim:")

	var MessageCount int

	for i := 0; i < nPrivateChat; i++ {
		message := PrivateChats[i]

		if message.sender == *currentUser {
			fmt.Println("[", MessageCount+1, "]", "Pesan untuk", message.receiver.username)
			MessageCount++
		}
	}

	if MessageCount == 0 {
		fmt.Println("Tidak ada pesan.")
		fmt.Println()
	}

	var selectedPrivateChat int
	fmt.Print("Pilih Pesan (0 untuk exit): ")
	fmt.Scan(&selectedPrivateChat)

	for !(selectedPrivateChat >= 1 && selectedPrivateChat <= nPrivateChat) {

		if selectedPrivateChat == 0 {
			return
		}

		fmt.Print("Pilih Pesan (0 untuk exit): ")
		fmt.Scan(&selectedPrivateChat)
	}

	message := PrivateChats[selectedPrivateChat-1]

	fmt.Println()
	fmt.Println("From\t:\t", (message.sender).username)
	fmt.Println("Content\t:\t", message.content)
	fmt.Println("To\t:\t", message.receiver.username)
	fmt.Println()

	return

}
func groupMenu(users *tabuser) {
	for {
		var choice int
		fmt.Println("Group Menu:")
		fmt.Println("1. Buat Group")
		fmt.Println("2. Lihat Group")
		fmt.Println("3. Kirim Pesan ke Group")
		fmt.Println("4. Kembali")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createGroup(users)
		case 2:
			viewGroups()
		case 3:
			sendgroup()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func findUserByUsername(users *tabuser, username string, lengthUsed int) int {
	// Sequential search

	for i := 0; i < lengthUsed; i++ {
		if users[i].username == username {
			return i
		}

	}

	return -1
}

func userIsMemberGroup(u *user, g *Group) bool {
	for _, member := range g.members {
		if member == u {
			return true
		}
	}

	return false
}

func createGroup(users *tabuser) {
	var groupName, memberUsername string
	fmt.Println("Buat Group Baru:")
	fmt.Print("Nama Group: ")
	fmt.Scan(&groupName)

	groups[nGroup].name = groupName
	groups[nGroup].creator = currentUser

	for {
		fmt.Print("Username orang yang ingin diundang (ketik selesai jika sudah selesai): ")
		fmt.Scan(&memberUsername)

		if memberUsername == "selesai" {
			fmt.Println("Group berhasil dibuat dengan nama:", groupName)
			for i := 0; i < groups[nGroup].memberCount; i++ {
				fmt.Println(groups[nGroup].members[i].username, "telah bergabung dalam group.")
			}

			fmt.Println(groups[nGroup].members[0:groups[nGroup].memberCount])
			nGroup++
			return
		}

		alreadyInvited := memberUsername == groups[nGroup].creator.username
		for j := 0; j < groups[nGroup].memberCount && !alreadyInvited; j++ {
			if groups[nGroup].members[j].username == memberUsername {
				alreadyInvited = true
			}
		}

		if alreadyInvited {
			fmt.Println("Anda sudah mengundang pengguna ini sebelumnya.")
		} else {
			// Coba cari id, kalau ga nemu bakal return -1, kalau selain -1 lanjut
			var temp int = findUserByUsername(users, memberUsername, nUser)

			if temp != -1 {
				if users[temp].approved && groups[nGroup].memberCount < NMAX {
					groups[nGroup].members[groups[nGroup].memberCount] = &users[temp]
					groups[nGroup].memberCount++
				}
			} else {
				fmt.Println("Username tidak valid atau belum diapprove. Silakan coba lagi.")
			}
		}
	}
}

func viewGroups() {
	fmt.Println()
	for i := 0; i < nGroup; i++ {
		if groups[i].creator == currentUser {
			fmt.Println(groups[i].name)
		} else {
			for j := 0; j < groups[i].memberCount; j++ {
				// :)))))
				if groups[i].members[j].username == currentUser.username {
					fmt.Println(groups[i].name)
				}
			}
		}
	}
}

func sendgroup() {

}
