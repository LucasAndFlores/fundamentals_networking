import { createServer } from "net";

const server = createServer((socket) => {
  console.log(
    `TCP handshake successful with ${socket.remoteAddress}:${socket.remotePort}`,
  );

  socket.write("Hello client from node js!");

  socket.on("data", (data) => {
    console.log(`Received data: ${data.toString()}`);
    return;
  });
});

server.listen(8800, "127.0.0.1");
