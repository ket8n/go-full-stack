import React, { useState, useEffect } from "react";
import StudentTable from "./components/StudentTable";
import StudentForm from "./components/StudentForm";
import "./App.css";

function App() {
  const [students, setStudents] = useState([]);
  const [editingStudent, setEditingStudent] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchStudents();
  }, []);

  const fetchStudents = async () => {
    try {
      setIsLoading(true);
      const response = await fetch("http://localhost:5555/students", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setStudents(data.data || []);
    } catch (error) {
      console.error("Error fetching students:", error);
      setError("Failed to fetch students. Please try again later.");
    } finally {
      setIsLoading(false);
    }
  };

  const handleCreate = async (newStudent) => {
    try {
      const response = await fetch("http://localhost:5555/students", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newStudent),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      await fetchStudents();
    } catch (error) {
      console.error("Error creating student:", error);
      setError("Failed to create student. Please try again.");
    }
  };

  const handleUpdate = async (updatedStudent) => {
    try {
      const response = await fetch(
        `http://localhost:5555/students/${updatedStudent._id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(updatedStudent),
        }
      );
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      await fetchStudents();
      setEditingStudent(null);
    } catch (error) {
      console.error("Error updating student:", error);
      setError("Failed to update student. Please try again.");
    }
  };

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:5555/students/${id}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      await fetchStudents();
    } catch (error) {
      console.error("Error deleting student:", error);
      setError("Failed to delete student. Please try again.");
    }
  };

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className="App">
      <h1>Student Management System</h1>
      <StudentForm
        onSubmit={editingStudent ? handleUpdate : handleCreate}
        initialData={editingStudent}
      />
      <StudentTable
        students={students}
        onEdit={setEditingStudent}
        onDelete={handleDelete}
      />
    </div>
  );
}

export default App;
