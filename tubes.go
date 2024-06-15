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

type tabuser [NMAX]user
type tabgroup [NMAX]Group
type tabPrivateChats [NMAX]PrivateChat

var PrivateChats tabPrivateChats
var groups tabgroup

var nPrivateChat int = 0
var nGroup int = 0
var nUser int = 0

var currentUser *user

func main() {
	var users tabuser
	var role string

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
	for i := 0; i < NMAX; i++ {
		if users[i].username != "" && users[i].approved {
			fmt.Println(users[i].username)
		}
	}
	fmt.Println()
}

func viewUsers2(users tabuser) {
	fmt.Println("Pengguna yang belum diapprove:")
	for i := 0; i < NMAX; i++ {
		if users[i].username != "" && !users[i].approved {
			fmt.Println(users[i].username)
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
	fmt.Println("From:", message.sender.username)
	fmt.Println("To:", message.receiver.username)
	fmt.Println("Message:", message.content)

	fmt.Println()
}

func ViewSendMessagers(users *tabuser) {
	fmt.Println()
	fmt.Println("Pesan yang Dikirim")

	var sendCount int

	for i := 0; i < nPrivateChat; i++ {
		message := PrivateChats[i]

		if message.sender == *currentUser {
			fmt.Println("[", sendCount+1, "]", "Pesan ke", message.receiver.username)
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

	for !(selectedPrivateChat >= 1 && selectedPrivateChat <= nPrivateChat) {

		if selectedPrivateChat == 0 {
			return
		}

		fmt.Print("Pilih pesan yang dikirim (0 untuk exit): ")
		fmt.Scan(&selectedPrivateChat)
	}

	message := PrivateChats[selectedPrivateChat-1]

	fmt.Println()
	fmt.Println("From:", message.sender.username)
	fmt.Println("To:", message.receiver.username)
	fmt.Println("Message:", message.content)

	fmt.Println()
}

func groupMenu(users *tabuser) {
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
			createGroup(users)
		case 2:
			viewJoinedGroups()
		case 3:
			sendGroupMessage()
		case 4:
			groupMessage(users)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang valid.")
		}
	}
}

func viewJoinedGroups() {
	fmt.Println("Grup yang Anda ikuti:")
	for i := 0; i < nGroup; i++ {
		isMember := false
		for j := 0; j < groups[i].memberCount; j++ {
			if groups[i].members[j].username == currentUser.username {
				isMember = true

			}
		}
		if isMember {
			fmt.Printf("[%d] %s\n", i+1, groups[i].name)
		}
	}
}

func createGroup(users *tabuser) {
	var groupName string
	fmt.Print("Masukkan nama grup baru: ")
	fmt.Scan(&groupName)

	for i := 0; i < nGroup; i++ {
		if groups[i].name == groupName {
			fmt.Println("Nama grup sudah ada. Silakan pilih nama lain.")
			return
		}
	}

	groups[nGroup] = Group{
		name:         groupName,
		creator:      currentUser,
		memberCount:  1,
		messageCount: 0,
	}

	groups[nGroup].members[0] = currentUser

	fmt.Println("Daftar pengguna yang terdaftar:")
	for i := 0; i < NMAX; i++ {
		if users[i].username != "" && users[i].approved && users[i].username != currentUser.username {
			fmt.Println("-", users[i].username)
		}
	}

	for {
		var addUser string
		fmt.Print("Masukkan nama pengguna untuk diundang ke grup (0 untuk berhenti): ")
		fmt.Scan(&addUser)

		if addUser == "0" {
			break
		}

		userFound := false
		for i := 0; i < NMAX; i++ {
			if users[i].username == addUser && users[i].approved {
				userFound = true
				alreadyMember := false
				for j := 0; j < groups[nGroup].memberCount; j++ {
					if groups[nGroup].members[j].username == addUser {
						alreadyMember = true
						break
					}
				}
				if alreadyMember {
					fmt.Println("Pengguna sudah menjadi anggota grup.")
				} else {
					groups[nGroup].members[groups[nGroup].memberCount] = &users[i]
					groups[nGroup].memberCount++
					fmt.Printf("Pengguna %s telah ditambahkan ke grup.\n", addUser)
				}
				break
			}
		}

		if !userFound {
			fmt.Println("Pengguna tidak ditemukan atau belum diapprove.")
		}
	}

	fmt.Printf("Grup %s berhasil dibuat.\n", groupName)
	nGroup++
}

func viewGroups() {
	fmt.Println("Daftar grup yang Anda ikuti:")
	for i := 0; i < nGroup; i++ {
		isMember := false
		for j := 0; j < groups[i].memberCount; j++ {
			if groups[i].members[j].username == currentUser.username {
				isMember = true
				break
			}
		}
		if isMember {
			fmt.Printf("[%d] %s\n", i+1, groups[i].name)
		}
	}

	var selectedGroup int
	fmt.Print("Pilih nomor grup untuk melihat anggota (0 untuk kembali): ")
	fmt.Scan(&selectedGroup)

	if selectedGroup == 0 {
		return
	}

	if selectedGroup < 1 || selectedGroup > nGroup {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	groupIndex := selectedGroup - 1
	fmt.Printf("Anggota grup %s:\n", groups[groupIndex].name)
	for j := 0; j < groups[groupIndex].memberCount; j++ {
		if groups[groupIndex].members[j] != nil {
			fmt.Printf("- %s\n", groups[groupIndex].members[j].username)
		}
	}
}

func sendGroupMessage() {
	fmt.Println("Daftar grup:")
	for i := 0; i < nGroup; i++ {
		fmt.Printf("[%d] %s\n", i+1, groups[i].name)
	}

	var selectedGroup int
	fmt.Print("Pilih nomor grup untuk mengirim pesan (0 untuk kembali): ")
	fmt.Scan(&selectedGroup)

	if selectedGroup == 0 {
		return
	}

	if selectedGroup < 1 || selectedGroup > nGroup {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	groupIndex := selectedGroup - 1

	isMember := false
	for i := 0; i < groups[groupIndex].memberCount; i++ {
		if groups[groupIndex].members[i].username == currentUser.username {
			isMember = true
			break
		}
	}

	if !isMember {
		fmt.Println("Anda bukan anggota grup ini.")
		return
	}

	var message string
	fmt.Print("Masukkan pesan (akhiri dengan ';'): ")
	var temp byte
	fmt.Scanf("%c", &temp) // Membaca karakter newline sebelum pesan
	for temp != ';' {
		message += string(temp)
		fmt.Scanf("%c", &temp)
	}

	groups[groupIndex].messages[groups[groupIndex].messageCount] = message
	groups[groupIndex].messageCount++
	fmt.Println("Pesan berhasil dikirim ke grup.")
}

func groupMessage(users *tabuser) {
	fmt.Println("Daftar grup yang Anda ikuti:")
	for i := 0; i < nGroup; i++ {
		isMember := false
		for j := 0; j < groups[i].memberCount; j++ {
			if groups[i].members[j].username == currentUser.username {
				isMember = true
				break
			}
		}
		if isMember {
			fmt.Printf("[%d] %s\n", i+1, groups[i].name)
		}
	}

	var selectedGroup int
	fmt.Print("Pilih nomor grup untuk melihat pesan (0 untuk kembali): ")
	fmt.Scan(&selectedGroup)

	if selectedGroup == 0 {
		return
	}

	if selectedGroup < 1 || selectedGroup > nGroup {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	groupIndex := selectedGroup - 1

	isMember := false
	for i := 0; i < groups[groupIndex].memberCount; i++ {
		if groups[groupIndex].members[i].username == currentUser.username {
			isMember = true
			break
		}
	}

	if !isMember {
		fmt.Println("Anda bukan anggota grup ini.")
		return
	}

	fmt.Printf("Pesan dalam grup %s:\n", groups[groupIndex].name)
	for i := 0; i < groups[groupIndex].messageCount; i++ {
		fmt.Printf("[%d] %s\n", i+1, groups[groupIndex].messages[i])
	}
}
