package main

import (
	"fmt"
)

type user struct {
	username string
	password string
	approved bool
}

type group struct {
	name   string
	member string
}

type Chat struct {
	sender   user
	receiver user
	content  string
}

const NMAX int = 100

type tabuser [NMAX]user
type tabgroup [NMAX]group
type tabChats [NMAX]Chat

var chats tabChats
var nChat int = 0

var currentUser user

func main() {
	var users tabuser
	var role string
	users[0] = user{username: "a", password: "a", approved: true}
	users[1] = user{username: "s", password: "s", approved: true}
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
	const adminpassword = "jojo"
	var password string
	fmt.Print("Masukkan password admin: ")
	fmt.Scan(&password)
	if password != adminpassword {
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
			return
		}
	}
	fmt.Println("Batas maksimum pengguna telah tercapai.")
}

func login(users tabuser) {
	var username, password string
	var groups tabgroup
	fmt.Println("Masukkan informasi login:")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < NMAX; i++ {
		if users[i].username == username && users[i].password == password {
			if users[i].approved {
				fmt.Printf("Selamat datang, %s! Anda berhasil login.\n", username)
				currentUser = users[i]
				userLoggedInMenu(&users, &groups)

				return
			} else {
				fmt.Println("Akun Anda masih menunggu persetujuan admin. Mohon tunggu.")
				return
			}
		}
	}

	fmt.Println("Username atau password salah.")

}

func userLoggedInMenu(users *tabuser, groups *tabgroup) {
	for {
		var choice int
		fmt.Println("User Logged In Menu:")
		fmt.Println("1. Kirim Pesan Pribadi")
		fmt.Println("2. Inbox")
		fmt.Println("3. Group")
		fmt.Println("4. Kembali")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			sendPrivateMessage(users)
		case 2:
			viewInbox(users)
		case 3:
			groupMenu(groups)
		case 4:
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

	fmt.Print("Isi Pesan: ")
	fmt.Scan(&message)

	chats[nChat] = Chat{sender: currentUser, receiver: userReceiver, content: message}
	nChat++

	return
}

func viewInbox(users *tabuser) {
	fmt.Println()
	fmt.Println("Inbox")

	var inboxCount int

	for i := 0; i < nChat; i++ {
		message := chats[i]

		if message.receiver == currentUser {
			fmt.Println("[", i+1, "]", "Pesan dari", message.sender.username)
			inboxCount++
		}
	}

	if inboxCount == 0 {
		fmt.Println("Tidak ada pesan.")
		fmt.Println()
	}

	var selectedChat int
	fmt.Print("Pilih inbox (0 untuk exit): ")
	fmt.Scan(&selectedChat)

	for !(selectedChat >= 1 && selectedChat <= nChat) {

		if selectedChat == 0 {
			return
		}

		fmt.Print("Pilih inbox (0 untuk exit): ")
		fmt.Scan(&selectedChat)
	}

	message := chats[selectedChat-1]

	fmt.Println()
	fmt.Println("From\t:\t", message.sender.username)
	fmt.Println("Content\t:\t", message.content)
	fmt.Println("To\t:\t", message.receiver.username)
	fmt.Println()

	return

}

func groupMenu(groups *tabgroup) {
	for {
		var choice int
		fmt.Println("Group Menu:")
		fmt.Println("1. Buat Group")
		fmt.Println("2. Lihat Group")
		fmt.Println("3. Kembali")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createGroup(groups)
		case 2:
			viewGroups(groups)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func createGroup(groups *tabgroup) {

}

func viewGroups(groups *tabgroup) {

}
