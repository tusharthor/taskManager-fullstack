import React, {useState, useEffect} from 'react';
import CreateTaskForm from './components/CreateTaskForm';
import TaskList from './components/TaskList';
import './components/CreateTaskForm.css';

function App() {

  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
    .then((res) => res.json())
    .then((data) => setTasks(data))
    .catch((err) => console.error(err));
  }, []);

  return (
    <div>
      <h1 className="task-title">Task Manager</h1>
      <CreateTaskForm setTasks={setTasks} />
      <TaskList tasks={tasks} />
    </div>
  );
}

export default App;
