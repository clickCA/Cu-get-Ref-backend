const mongoose = require("mongoose");
const Menu = require("../models/Menu");

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
const getAllMenus = () =>
  new Promise(async (resolve, reject) => {
    try {
      const data = Menu.find().sort({ id: "asc" }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

const getMenu = (key = "") =>
  new Promise(async (resolve, reject) => {
    try {
      const data = await Menu.findOne({ id: key }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

const insertMenu = (data) =>
  new Promise(async (resolve, reject) => {
    try {
      const menu = new Menu(data);
      const newMenu = await menu.save();
      resolve(newMenu);
    } catch (e) {
      reject(e);
    }
  });

const removeMenu = (key = "") =>
  new Promise(async (resolve, reject) => {
    try {
      const data = await Menu.findOneAndDelete({ id: key }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

const updateMenu = (key = "", arg) =>
  new Promise(async (resolve, reject) => {
    try {
      const data = await Menu.findOneAndUpdate({ id: key }, arg, {
        new: true,
      }).exec();
      resolve(data);
    } catch (e) {
      reject(e);
    }
  });

module.exports = {
  connectDB,
  getAllMenus,
  getMenu,
  insertMenu,
  removeMenu,
  updateMenu,
};
