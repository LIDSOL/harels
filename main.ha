use bufio;
use io;
use fs;
use fmt;
use os;
use strings;
use encoding::json;
use log;
use regex;
use strconv;

fn initialize(id: int) void = {
	const capabilities = json::newobject();
	json::put(&capabilities, "hoverProvider", true)!;

	const obj = json::newobject();
	json::put(&obj, "jsonrpc", "2.0")!;
	json::put(&obj, "result", capabilities)!;
	json::put(&obj, "id", id: f64)!;

	const s = json::dumpstr(obj);

	fmt::printf("Content-Length: {}\r\n\r\n", len(s))!;
	fmt::print(s)!;
	fmt::print("\r\n")!;
	log::println(s);
};

export fn main() void = {

	const re = regex::compile(`Content-Length: [0-9]+[^0-9]+$`)!;
	const re2 = regex::compile(`Content-Length: `)!;
	// Crear un buffer (o cadena) vacío
	let container: str = "";
	for (true) {
		for (!regex::test(&re, container)!) {
			const byte = bufio::read_byte(os::stdin)! as u8;
			const bufferlol: [1]u8 = [byte];
			const nomanches = strings::fromutf8(bufferlol)!;
			log::println(nomanches);
			const nuevo = strings::concat(container, nomanches)!;
			container = nuevo;
		};

		for (bufio::read_byte(os::stdin)! as u8 != '{') void;

		log::println(container);
		const num_str = strings::trim(regex::replace(&re2, container, ``)!);
		container = "";
		const num = strconv::stoi(num_str)!;
		log::println(num);
		const buf: []u8 = alloc([0...], (num-1): size)!;
		defer free(buf);

		const siz = match(io::read(os::stdin, buf)) {
		case let s: size =>
			log::println("YA LEI EL MENSAJE");
			log::println(s);
			yield s;
		case io::EOF =>
			fmt::fatalf("EOF Reached when receiving message");
		case let err: io::error =>
			fmt::fatalf("ME VALE");
		};

		const msg_str = strings::concat("{", strings::fromutf8(buf)!)!;
		log::println(msg_str);

		const val = match (json::loadstr(msg_str)) {
		case let v: json::value =>
			yield v;
		case nomem =>
			fmt::fatalf("No memory");
		case let err: json::error =>
			fmt::fatalf("Error when reading JSON object");
		};

		match (val) {
		case let obj: json::object =>
			const method = *(json::get(&obj, "method") as *json::value) as str;


			switch (method) {
			case "initialize" =>
				const id = *(json::get(&obj, "id") as *json::value) as f64: int;
				initialize(id);
			case "initialized" =>
				log::println("I INITIALIZED");
			case =>
				fmt::fatalf("I wanna die 2");
			};
		case =>
			fmt::fatalf("I wanna die");

		};
		log::println("QUE ONDA, WUUUU");
	};

};
