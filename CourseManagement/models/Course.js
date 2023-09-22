const mongoose = require("mongoose");

// Course structure for storing course details.
const CourseSchema = new mongoose.Schema(
  {
    id: {
      type: String,
      required: [true, "Please add an ID"],
    },
    CourseName: {
      type: String,
      required: [true, "Please add a course name"],
    },
    CourseCode: {
      type: String,
      required: [true, "Please add a course code"],
    },
  },
  {
    toJSON: { virtuals: true },
    toObject: { virtuals: true },
  }
);

module.exports = mongoose.model("Course", CourseSchema);
