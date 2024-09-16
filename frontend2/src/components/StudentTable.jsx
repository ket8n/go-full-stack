import React from "react";

function StudentTable({ students, onEdit, onDelete }) {
  if (!students || students.length === 0) {
    return <p>No students found.</p>;
  }

  return (
    <table>
      <thead style={{ backgroundColor: "yellow" }}>
        <tr>
          <th>Name</th>
          <th>Age</th>
          <th>Grade</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {students.map((student) => (
          <tr key={student._id}>
            <td>{student.name}</td>
            <td>{student.age}</td>
            <td>{student.grade}</td>
            <td>
              <button
                onClick={() => onEdit(student)}
                style={{
                  backgroundColor: "green",
                  color: "white",
                  marginRight: "5px",
                }}
              >
                Edit
              </button>
              <button
                onClick={() => onDelete(student._id)}
                style={{ backgroundColor: "red", color: "white" }}
              >
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default StudentTable;
