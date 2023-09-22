const PROTO_PATH = `${__dirname}../../../../course.proto`;

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

const getCourses = (req, res) => {
  client.getAllCourse({}, (error, result) => {
    if (!error) {
      res.status(200).json(result);
    } else {
      res.status(400).json(error);
    }
  });
};

const getCourse = (req, res) => {
  const { id } = req.params;
  client.get({ id }, (error, result) => {
    if (!error) {
      res.status(200).json(result);
    } else {
      res.status(400).json(error);
    }
  });
};

const insertCourse = (req, res) => {
  const { courseName, courseCode, courseDescription } = req.body;
  const id = 1;
  const newCourseItem = {
    id,
    courseName,
    courseCode,
    courseDescription,
  };
  client.insert(newCourseItem, (error, courseItem) => {
    if (!error) {
      res.status(201).json(courseItem);
    } else {
      res.status(400).json(error);
    }
  });
};

const updateCourse = (req, res) => {
  const { id } = req.params;
  const { courseName, courseCode, courseDescription } = req.body;
  const courseItem = {
    id,
    courseName,
    courseCode,
    courseDescription,
  };
  client.update(courseItem, (error, result) => {
    if (!error) {
      res.status(200).json(result);
    } else {
      res.status(400).json(error);
    }
  });
};

const removeCourse = (req, res) => {
  const { id } = req.params;
  client.remove({ id }, (error, result) => {
    if (!error) {
      res.status(200).json(result);
    } else {
      res.status(400).json(error);
    }
  });
};
module.exports = {
  getCourses,
  getCourse,
  removeCourse,
  updateCourse,
  insertCourse,
};
