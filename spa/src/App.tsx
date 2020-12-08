import { useState, useEffect } from 'react'
import TaskForm from './components/TaskForm'
import TodoList from './components/TodoList'
import axios from 'axios'

export default function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [completedTasks, setCompletedTasks] = useState<Task[]>([])
  const addTask = async (title: string) => {
    if (!title) {
      return
    }
    const { data } = await axios.post<Task>('/tasks', { title })
    setTasks(tasks.concat(data))
  }
  const toggleTaskCompleted = async (task: Task) => {
    const { data } = await axios.patch<Task>(`/tasks/${task.id}`, {
      completed: !task.completed
    })
    if (data.completed) {
      setTasks(tasks.filter(t => t.id !== data.id))
      setCompletedTasks(completedTasks.concat(data).sort((a, b) => {
        return Date.parse(b.updatedAt) - Date.parse(a.updatedAt)
      }))
    } else {
      setCompletedTasks(completedTasks.filter(t => t.id !== data.id))
      setTasks(tasks.concat(data).sort((a, b) => {
        return Date.parse(b.updatedAt) - Date.parse(a.updatedAt)
      }))
    }
  }
  
  useEffect(() => {
    axios.get<Task[]>('/tasks')
      .then(({ data }) => {
        setTasks(data)
      })
    axios.get<Task[]>('/tasks/completed')
      .then(({ data }) => {
        setCompletedTasks(data)
      })
  }, [])

  return (
    <div>
      <TaskForm addTask={addTask} />
      <TodoList
        tasks={tasks} 
        toggleTaskCompleted={toggleTaskCompleted} />
      <TodoList
        tasks={completedTasks} 
        toggleTaskCompleted={toggleTaskCompleted} />
    </div>
  )
}
