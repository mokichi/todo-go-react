import useTodo from './hooks/useTodo'
import TaskForm from './components/TaskForm'
import TodoList from './components/TodoList'

export default function App() {
  const {
    state: { tasks, completedTasks },
    addTask,
    toggleTaskCompleted,
    updateTask,
    deleteTask
  } = useTodo()

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
