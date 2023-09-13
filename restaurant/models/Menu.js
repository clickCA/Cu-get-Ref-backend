const mongoose = require("mongoose");

const MenuSchema = new mongoose.Schema(
  {
    id: {
      type: String,
      required: [true, "Please add an ID"],
    },
    name: {
      type: String,
      required: [true, "Please add a name"],
    },
    price: {
      type: Number,
      required: [true, "Please add a price"],
    },
  },
  {
    toJSON: { virtuals: true },
    toObject: { virtuals: true },
  }
);

module.exports = mongoose.model("Menu", MenuSchema);
