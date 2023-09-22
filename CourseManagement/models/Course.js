const mongoose = require("mongoose");

// Course structure for storing course details.
const CourseSchema = new mongoose.Schema(
  {
    id: {
      type: String,
      required: [true, "Please add an ID"],
    },
    courseName: {
      type: String,
      required: [true, "Please add a course name"],
    },
    courseCode: {
      type: String,
      required: [true, "Please add a course code"],
    },
    courseDescription: {
      type: String,
      required: [true, "Please add a course description"],
    },
  },
  {
    toJSON: { virtuals: true },
    toObject: { virtuals: true },
  }
);

module.exports = mongoose.model("Course", CourseSchema);
