import dgram from "dgram";

const PORT = 5500;

// connect with netcat nc -u 127.0.0.1 5500

const socket = dgram.createSocket("udp4");
socket.bind(PORT, "127.0.0.1");
socket.on("message", (msg, info) => {
  console.log(
    `My UDP server get a datagram ${msg} from ${info.address}:${info.port}`,
  );
});

socket.on("listening", () => {
  console.log(`The NodeJs UDP server is running and listening on port ${PORT}`);
});
