const PROTO_PATH = "../course.proto";

const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  arrays: true,
});

var courseService = grpc.loadPackageDefinition(packageDefinition).CourseService;

const client = new courseService(
  "localhost:30043",
  grpc.credentials.createInsecure()
);

module.exports = client;
