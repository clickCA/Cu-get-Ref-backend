const client = require("./client");

const path = require("path");
const express = require("express");
const bodyParser = require("body-parser");

const app = express();

app.set("views", path.join(__dirname, "views"));
app.set("view engine", "hbs");

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

app.get("/", (req, res) => {
  client.getAllCourse(null, (err, data) => {
    if (!err) {
      res.render("course", {
        results: data.course,
      });
    }
  });
});

app.post("/save", (req, res) => {
  let newCourseItem = {
    courseName: req.body.courseName,
    courseCode: req.body.courseCode,
    courseDescription: req.body.courseDescription,
  };

  client.insert(newCourseItem, (err, data) => {
    if (err) throw err;

    console.log("New Course created successfully", data);
    res.redirect("/");
  });
});

app.post("/update", (req, res) => {
  const updateCourseItem = {
    id: req.body.id,
    courseName: req.body.courseName,
    courseCode: req.body.courseCode,
    courseDescription: req.body.courseDescription,
  };
  console.log(
    "update Item %s %s %d",
    updateCourseItem.id,
    req.body.name,
    req.body.price
  );

  client.update(updateCourseItem, (err, data) => {
    if (err) throw err;

    console.log("Course Item updated successfully", data);
    res.redirect("/");
  });
});

app.post("/remove", (req, res) => {
  client.remove({ id: req.body.id }, (err, _) => {
    if (err) throw err;
    console.log("Course Item removed successfully");
    res.redirect("/");
  });
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log("Server running at port %d", PORT);
});
