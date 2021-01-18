import { useState, useEffect } from 'react'
import axios from 'axios'

function compareTasks(a: Task, b: Task): number {
  return Date.parse(b.updatedAt) - Date.parse(a.updatedAt)
}

type State = {
  tasks: Task[],
  completedTasks: Task[]
}

export default function useTodo() {
  const [state, setState] = useState<State>({
    tasks: [],
    completedTasks: []
  })
  
  useEffect(() => {
    (async () => {
      const [tasks, completedTasks] = await Promise.all([
        new Promise<Task[]>(resolve => {
          axios.get<Task[]>('/tasks')
            .then(({ data }) => {
              resolve(data)
            })
        }),
        new Promise<Task[]>(resolve => {
          axios.get<Task[]>('/tasks/completed')
            .then(({ data }) => {
              resolve(data)
            })
        })
      ])
      setState({
        ...state,
        tasks,
        completedTasks
      })
    })()
    // eslint-disable-next-line
  }, [])

  const addTask = async (title: string) => {
    if (!title) {
      return
    }
    const { data } = await axios.post<Task>('/tasks', { title })
    setState({
      ...state,
      tasks: [data].concat(state.tasks)
    })
  }

  const updateTask = async (task: Task, params: TaskUpdateParams) => {
    if (!params.title) {
      return
    }
    const { data } = await axios.patch<Task>(`/tasks/${task.id}`, params)
    setState({
      ...state,
      tasks: state.tasks.map(t => {
        if (t.id === data.id) {
          return data
        }
        return t
      }).sort(compareTasks)
    })
  }

  const toggleTaskCompleted = async (task: Task) => {
    const { data } = await axios.patch<Task>(`/tasks/${task.id}`, {
      completed: !task.completed
    })
    if (data.completed) {
      setState({
        ...state,
        tasks: state.tasks.filter(t => t.id !== data.id),
        completedTasks: state.completedTasks.concat(data).sort(compareTasks)
      })
    } else {
      setState({
        ...state,
        completedTasks: state.completedTasks.filter(t => t.id !== data.id),
        tasks: state.tasks.concat(data).sort(compareTasks)
      })
    }
  }

  const deleteTask = async (task: Task) => {
    await axios.delete(`/tasks/${task.id}`)
    setState({
      ...state,
      tasks: state.tasks.filter(t => t.id !== task.id),
      completedTasks: state.completedTasks.filter(t => t.id !== task.id)
    })
  }

  return {
    state,
    addTask,
    updateTask,
    toggleTaskCompleted,
    deleteTask
  }
}
