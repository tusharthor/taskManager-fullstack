import React, { useState } from 'react';
import './CreateTaskForm.css';

function CreateTaskForm({setTasks}) {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8080/tasks',{
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({title, description, completed: false}),
        });
        if (response.ok){
            const newTask = await response.json();
            setTasks((prevTasks) => [...prevTasks, newTask]);
            setTitle('');
            setDescription('');
        }
    };

    return (
        <div className="task-card">
            <h3 className="task-title">Add New Task</h3>
            <form onSubmit={handleSubmit} className="task-form">
                <div className="input-group">
                    <label>Title: </label>
                    <input 
                    type="text"
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                    />
                </div>
                <div className="input-group">
                    <label>Description:</label>
                    <textarea 
                        placeholder="Description"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                </div>
                <button type="submit" className="submit-button">Add Task</button>
            </form>
        </div>
    );
}

export default CreateTaskForm;