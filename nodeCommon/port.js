function getPort() {

	const _port = parseInt(process.argv[2])

	let port;
	port = _port
	if (!port) {
	  port = 9000
	  console.log('Using default port 9000')
	}

	return port

}

module.exports = getPort()
