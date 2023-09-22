const PROTO_PATH = "./course.proto";
const dotenv = require("dotenv");
const db = require("../config/db");

var grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");

//Load env vars
dotenv.config({ path: "./config/config.env" });

//Connect to database
db.connectDB();

var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  arrays: true,
});

var courseProto = grpc.loadPackageDefinition(packageDefinition);

const { v4: uuidv4 } = require("uuid");

const server = new grpc.Server();

server.addService(courseProto.CourseService.service, {
  getAllCourse: (_, callback) => {
    db.getAllCourses()
      .then((result) => {
        if (!result) {
          callback(null, { course: [] });
        }
        callback(null, { course: result });
      })
      .catch((e) => {
        callback(e);
      });
  },
  get: (call, callback) => {
    const { id } = call.request;
    db.getCourse(id)
      .then((result) => {
        if (!result) {
          callback({
            code: grpc.status.NOT_FOUND,
            details: "Not found",
          });
        }
        callback(null, result);
      })
      .catch((e) => {
        callback(e);
      });
  },
  insert: (call, callback) => {
    let courseItem = call.request;
    courseItem.id = uuidv4();

    db.insertCourse(courseItem)
      .then(() => {
        callback(null, courseItem);
      })
      .catch((e) => {
        callback(e);
      });
  },
  update: (call, callback) => {
    const course = call.request;
    const { id } = course;
    db.updateCourse(id, course)
      .then(() => {
        callback(null, course);
      })
      .catch((e) => {
        callback({
          code: grpc.status.NOT_FOUND,
          details: "Not Found",
        });
      });
  },
  remove: (call, callback) => {
    const { id } = call.request;
    db.removeCourse(id)
      .then(() => {
        callback(null, {});
      })
      .catch((e) => {
        callback({
          code: grpc.status.NOT_FOUND,
          details: "NOT Found",
        });
      });
  },
});

server.bindAsync(
  "127.0.0.1:30043",
  grpc.ServerCredentials.createInsecure(),
  () => {
    server.start();
  }
);
console.log("Server running at http://127.0.0.1:30043");
