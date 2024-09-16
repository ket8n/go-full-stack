import React, { useState, useEffect } from "react";

function StudentForm({ onSubmit, initialData }) {
  const [formData, setFormData] = useState({
    name: "",
    age: "",
    grade: "",
  });

  useEffect(() => {
    if (initialData) {
      setFormData(initialData);
    } else {
      setFormData({ name: "", age: "", grade: "" });
    }
  }, [initialData]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
    setFormData({ name: "", age: "", grade: "" });
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        name="name"
        value={formData.name}
        onChange={handleChange}
        placeholder="Name"
        required
      />
      <input
        type="text"
        name="age"
        value={formData.age}
        onChange={handleChange}
        placeholder="Age"
        required
      />
      <input
        type="text"
        name="grade"
        value={formData.grade}
        onChange={handleChange}
        placeholder="Grade"
        required
      />
      <button
        type="submit"
        style={{
          backgroundColor: "red",
          color: "white",
          borderRadius: "20px",
          padding: "10px 20px",
          border: "none",
          cursor: "pointer",
        }}
      >
        {initialData ? "Update" : "Create"} Student
      </button>
    </form>
  );
}

export default StudentForm;
