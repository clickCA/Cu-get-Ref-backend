const mongoose = require("mongoose");
const Course = require("../models/Course");

//Connect to database
const connectDB = async () => {
  mongoose.set("strictQuery", true);
  const conn = await mongoose.connect(process.env.MONGO_URI, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  });

  console.log(`MongoDB Connected: ${conn.connection.host}`);
};

//Define the functions that will interact with the database
const getAllCourses = () =>
  new Promise(async (resolve, reject) => {
    try {
      const data = Course.find().sort({ courseCode: "asc" }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

const getCourse = (key = "") =>
  new Promise(async (resolve, reject) => {
    try {
      const data = await Menu.findOne({ courseCode: key }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

const insertCourse = (data) =>
  new Promise(async (resolve, reject) => {
    try {
      const course = new Course(data);
      const newCourse = await course.save();
      resolve(newCourse);
    } catch (e) {
      reject(e);
    }
  });

const removeCourse = (key = "") =>
  new Promise(async (resolve, reject) => {
    try {
      const data = await Menu.findOneAndDelete({ courseCode: key }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

const updateCourse = (key = "", arg) =>
  new Promise(async (resolve, reject) => {
    try {
      const data = await Menu.findOneAndUpdate({ courseCode: key }, arg, {
        new: true,
      }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

module.exports = {
  connectDB,
  getAllCourses,
  getCourse,
  insertCourse,
  removeCourse,
  updateCourse,
};
