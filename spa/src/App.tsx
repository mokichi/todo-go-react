import { useState, useEffect } from 'react'
import TaskForm from './components/TaskForm'
import TodoList from './components/TodoList'

export default function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const addTask = (title: string) => {
    const newTask = {
      id: tasks[tasks.length - 1].id + 1,
      title: title,
      completed: false,
      createdAt: new Date(),
      updatedAt: new Date()
    }
    setTasks(tasks.concat(newTask))
  }
  const toggleTaskCompleted = (task: Task) => {
    const newTasks = tasks.map(t => {
      if (t.id === task.id) {
        t.completed = !t.completed
      }
      return t;
    })
    setTasks(newTasks)
  }
  
  useEffect(() => {
    setTasks([
      {
        id: 1,
        title: 'Task 1',
        completed: false,
        createdAt: new Date(),
        updatedAt: new Date()
      },
      {
        id: 2,
        title: 'Task 2',
        completed: false,
        createdAt: new Date(),
        updatedAt: new Date()
      }
    ])
  }, [])

  return (
    <div>
      <TaskForm addTask={addTask} />
      <TodoList
        tasks={tasks} 
        toggleTaskCompleted={toggleTaskCompleted} />
    </div>
  )
}
