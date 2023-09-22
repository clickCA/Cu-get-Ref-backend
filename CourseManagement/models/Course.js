const mongoose = require("mongoose");

const CourseSchema = new mongoose.Schema(
  {
    id: {
      type: String,
      required: [true, "Please add an ID"],
    },
    CourseName: {
      type: String,
      required: [true, "Please add a name"],
    },
    CourseCode: {
      type: String,
      required: [true, "Please add a price"],
    },
  },
  {
    toJSON: { virtuals: true },
    toObject: { virtuals: true },
  }
);

module.exports = mongoose.model("Course", COurseSchema);
