// module cyber_shield

import io
import net
import time
import x.json2 as j

pub struct LicenseID {
	pub mut: 
		auth_status			bool
		id					string
		username			string // username == discord tag
		discord_id			string
		rank				string
}

const app_version = "APPLICATION_NAME_AND_VERSION"

// const cs_dns_rotation = ['1.1.1.1', '1.1.1.1'] 
const cs_backend = "15.204.58.133" // DO NOT CHANGE

fn main()
{
	mut lic := LicenseID{id: "4IqF7qAxP9JObVr5R16XOxpKePKPMAFpaZ1gft7Ez75IVp1H"}
	mut cs := connect_to_backend(mut &lic)
	for {
		time.sleep(1*time.second)
	}
}

pub fn connect_to_backend(mut lic LicenseID) LicenseID {
	mut socket := net.dial_tcp("${cs_backend}:17234") or {
		print("[!] Error, Unable to connect to Cyber Shield Server. Cannot authenticate license ID!....\n")
		exit(0)
	}

	mut reader := io.new_buffered_reader(reader: socket)
	
	socket.write_string("${app_version}\n") or {
		print("[!] Error, Unable to interact with Cyber Shield server...!\n")
		exit(0)
	}
	
	socket.write_string("${lic.id}\n") or { exit(0) }
	data := reader.read_line() or { exit(0) }
	socket.set_read_timeout(time.infinite)

	if data.starts_with("{") && data.ends_with("}") { //validating json format
		user_info := (j.raw_decode(data) or { map[string]j.Any{} }).as_map()
		lic.auth_status = (user_info['status'] or { panic("[!] Error, Status")}).bool()
		lic.username = (user_info['discord_tag'] or { panic("[!] Error, User Discord Tag")}).str()
		lic.discord_id = (user_info['discord_id'] or { panic("[!] Error, User Discord ID") }).str()
		lic.rank = (user_info['rank'] or { panic("[!] Error, No rank for user") }).str()
	}
	
	print("Successfully authed! ${lic.rank}}\n")
	go connection(mut &lic, mut socket, mut reader)
	return lic
}

pub fn connection(mut lid LicenseID, mut socket net.TcpConn, mut reader io.BufferedReader) {
	for {
		data := reader.read_line() or { "" }
		if data.len > 0 {
			if data == "close" {
				// utils.clear_screen()
				print("[!] Error, The owner has disconnected you from using Cyber Shield....\n")
				exit(0)
			} else if data == "banned" {
				// utils.clear_screen()
				print("[!] Error, You have been banned from using Cyber Shield....")
				exit(0)
				// write code here to delete this app from the user's system
			}
		}
	}
}