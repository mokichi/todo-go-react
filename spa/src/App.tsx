import { useState, useEffect } from 'react'
import TaskForm from './components/TaskForm'
import TodoList from './components/TodoList'
import axios from 'axios'

export default function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [completedTasks, setCompletedTasks] = useState<Task[]>([])

  const compareTasks = (a: Task, b: Task) => {
    return Date.parse(b.updatedAt) - Date.parse(a.updatedAt)
  }
  const addTask = async (title: string) => {
    if (!title) {
      return
    }
    const { data } = await axios.post<Task>('/tasks', { title })
    setTasks([data].concat(tasks))
  }
  const updateTask = async (task: Task, params: TaskUpdateParams) => {
    if (!params.title) {
      return
    }
    const { data } = await axios.patch<Task>(`/tasks/${task.id}`, params)
    setTasks(tasks.map(t => {
      if (t.id === data.id) {
        return data
      }
      return t
    }).sort(compareTasks))
  }
  const toggleTaskCompleted = async (task: Task) => {
    const { data } = await axios.patch<Task>(`/tasks/${task.id}`, {
      completed: !task.completed
    })
    if (data.completed) {
      setTasks(tasks.filter(t => t.id !== data.id))
      setCompletedTasks(completedTasks.concat(data).sort(compareTasks))
    } else {
      setCompletedTasks(completedTasks.filter(t => t.id !== data.id))
      setTasks(tasks.concat(data).sort(compareTasks))
    }
  }
  const deleteTask = async (task: Task) => {
    await axios.delete(`/tasks/${task.id}`)
    setTasks(tasks.filter(t => t.id !== task.id))
    setCompletedTasks(completedTasks.filter(t => t.id !== task.id))
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
        toggleTaskCompleted={toggleTaskCompleted}
        updateTask={updateTask}
        deleteTask={deleteTask} />
      <TodoList
        tasks={completedTasks} 
        toggleTaskCompleted={toggleTaskCompleted}
        updateTask={updateTask}
        deleteTask={deleteTask} />
    </div>
  )
}
