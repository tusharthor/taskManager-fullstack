import React, { useState, useEffect } from "react";
import "./TaskList.css"; // Import external CSS

const TaskList = () => {
  // State to store tasks, loading state, and errors
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch tasks from the API
  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await fetch("http://localhost:8080/tasks");
        if (!response.ok) {
          throw new Error("Failed to fetch tasks");
        }
        const data = await response.json();
        setTasks(data);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    fetchTasks();
  }, []); // Empty dependency array to run only once when the component mounts

  // Delete task handler
  const onDelete = (id) => {
    setTasks(tasks.filter((task) => task.id !== id));
  };

  // Toggle task completion handler
  const onToggleComplete = (id) => {
    setTasks(
      tasks.map((task) =>
        task.id === id ? { ...task, completed: !task.completed } : task
      )
    );
  };

  // Render loading or error message
  if (loading) return <p>Loading tasks...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div className="task-list-container">
      <h2 className="task-list-title">Your Tasks</h2>

      {/* Display message if no tasks available */}
      {tasks.length === 0 ? (
        <p className="no-tasks">No tasks available. Add some!</p>
      ) : (
        <table className="task-table">
          <thead>
            <tr>
              <th>Title</th>
              <th>Description</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {/* Map through tasks and display each task in a table row */}
            {tasks.map((task) => (
              <tr key={task.id}>
                <td>{task.title}</td>
                <td>{task.description}</td>
                <td>
                  <span
                    className={`task-status ${task.completed ? "completed" : ""}`}
                    onClick={() => onToggleComplete(task.id)}
                  >
                    {task.completed ? "Completed" : "Pending"}
                  </span>
                </td>
                <td>
                  <button
                    className="delete-button"
                    onClick={() => onDelete(task.id)}
                  >
                    ✖
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default TaskList;
