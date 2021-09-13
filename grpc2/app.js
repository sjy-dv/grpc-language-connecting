const PROTO_PATH = __dirname + "/idl/hello.proto";

const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const hello_proto = grpc.loadPackageDefinition(packageDefinition).idl;

function sayHello(call, callback) {
  console.log(call);
  console.log(callback);
  callback(null, { body: "Hello from node micro server!" });
}

function main() {
  const server = new grpc.Server();
  server.addService(hello_proto.ChatService.service, { sayHello: sayHello });
  server.bindAsync(
    "localhost:3002",
    grpc.ServerCredentials.createInsecure(),
    () => {
      server.start();
    }
  );
}

main();
