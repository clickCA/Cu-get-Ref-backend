const express = require("express");
const grpcController = require("../controllers/grpc");

const router = express.Router();

router.get("/courses", grpcController.getCourses);
router.post("/courses", grpcController.insertCourse);
router.get("/courses/:id", grpcController.getCourse);
router.delete("/courses/:id", grpcController.removeCourse);
router.put("/courses/:id", grpcController.updateCourse);

module.exports = router;
